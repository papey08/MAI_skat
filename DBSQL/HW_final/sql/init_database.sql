alter table if exists "user"
drop constraint if exists fk_avatar_id;

alter table if exists message
drop constraint if exists fk_chat_id;
alter table if exists message
drop constraint if exists fk_user_id;
alter table if exists message
drop constraint if exists fk_file_id;
alter table if exists message
drop constraint if exists fk_reply_id;
alter table if exists message
drop constraint if exists fk_forward_id;

alter table if exists private_chat
drop constraint if exists fk_user_id01;
alter table if exists private_chat
drop constraint if exists fk_user_id02;

alter table if exists group_chat
drop constraint if exists fk_avatar_id;
alter table if exists group_chat
drop constraint if exists fk_creator_id;

alter table if exists chat
drop constraint if exists fk_group_chat;
alter table if exists chat
drop constraint if exists fk_private_chat;

alter table if exists user_group_chat
drop constraint if exists fk_group_chat_id;
alter table if exists user_group_chat
drop constraint if exists fk_user_id;

alter table if exists reaction
drop constraint if exists fk_message_id;
alter table if exists reaction
drop constraint if exists fk_user_id;

drop table if exists "user";
drop table if exists message;
drop table if exists private_chat;
drop table if exists group_chat;
drop table if exists chat;
drop table if exists user_group_chat;
drop table if exists reaction;
drop table if exists file;

drop type if exists admin_rights;
drop type if exists emoji;

create table "user" (
    id serial primary key,
    username varchar(50) unique not null,
    first_name varchar(100) not null,
    second_name varchar(100),
    email varchar(100) unique not null,
    avatar_id bigint,
    created_at timestamp not null default current_timestamp,

    constraint check_columns_not_empty check (
        trim(username) <> '' and trim(first_name) <> ''
    ),
    constraint validate_email check (
        email like '%_@%_.%_'
    )
);

-- функция для проверки того, что пользователь, который пишет сообщение в чат, есть в этом чате
create or replace function check_if_user_in_chat(p_user_id bigint, p_chat_id bigint) returns boolean as $$
    begin
        if (select private_chat_id from chat where id = p_chat_id) is null then
            return p_user_id in (
                select user_id from chat
                inner join group_chat gc on chat.group_chat_id = gc.id
                inner join user_group_chat ugc on gc.id = ugc.group_chat_id
                where chat.id = p_chat_id
            );
        else
            return p_user_id in (
                select user_id01 from chat
                inner join private_chat pc on chat.private_chat_id = pc.id
                where chat.id = p_chat_id
            ) or p_user_id in (
                select user_id02 from chat
                inner join private_chat pc on chat.private_chat_id = pc.id
                where chat.id = p_chat_id
            );
        end if;
    end;
$$ language plpgsql;

create table message (
    id serial primary key,
    chat_id bigint not null,
    user_id bigint not null,
    message_text text,
    message_file bigint,
    reply_message_id bigint,
    forward_message_id bigint,
    created_at timestamp not null default current_timestamp,

--  проверка на то, что присутствует хотя бы одно из обязательных полей
    constraint check_columns_not_nil check (
        message_text is not null
        or message_file is not null
        or reply_message_id is not null
        or forward_message_id is not null
    ),

--  проверка на то, что пользователь есть в чате, в который он пишет сообщение
    constraint check_user_in_chat check (
        check_if_user_in_chat(user_id, chat_id)
    )
);

create table private_chat (
    id serial primary key,
    user_id01 bigint not null,
    user_id02 bigint not null,

--  проверка на уникальность чата двух пользователей
    constraint unique_users unique (user_id01, user_id02)
);
-- функция сортирует user_id01 и user_id02 для корректной работы unique_users
create or replace function private_chat_order_user_ids()
returns trigger as $$
    declare temp_buf bigint;
    begin
        if new.user_id01 > new.user_id02 then
            -- меняем значения user_id01 и user_id02 местами
            temp_buf := new.user_id01;
            new.user_id01 = new.user_id02;
            new.user_id02 = temp_buf;
        end if;
        return new;
    end;
$$ language plpgsql;
-- триггер, который выполняет сортировку user_id01 и user_id02 при вставке в private_chat
create trigger private_chat_trigger_order_user_ids
before insert on private_chat
for each row
execute function private_chat_order_user_ids();

create table group_chat (
    id serial primary key,
    title varchar(200) not null,
    description text,
    avatar_id bigint,
    creator_id bigint not null,
    created_at timestamp not null default current_timestamp
);
-- функция, которая выполняет вставку создателя группового чата в созданный им чат
create or replace function add_creator_to_group_chat()
returns trigger as $$
    begin
        insert into user_group_chat (user_id, group_chat_id, rights) values
            (new.creator_id, new.id, 'creator');
        return new;
    end;
$$ language plpgsql;
-- триггер, который выполняет вставку создателя группового чата в созданный им чат
create trigger group_chat_trigger_add_creator_to_group_chat
after insert on group_chat
for each row
execute function add_creator_to_group_chat();

create table chat (
    id serial primary key,
    private_chat_id bigint unique,
    group_chat_id bigint unique,
    messages_amount bigint not null default 0,

--  проверка на то, что запись ссылается только на один чат
    constraint check_chat_id_is_null check (
    --  реализуем оператор xor
        (private_chat_id is null and group_chat_id is not null)
        or (private_chat_id is not null and group_chat_id is null)
    )
);

create type admin_rights as enum (
    'creator',
    'admin'
);
create table user_group_chat (
    user_id bigint not null,
    group_chat_id bigint not null,
    rights admin_rights,

--  проверка на то, что один пользователь состоит в одном чате не больше 1 раза
    constraint unique_user_group unique (user_id, group_chat_id)
);

create type emoji as enum (
    'thumb_up',
    'thumb_down',
    'red_heart',
    'rolling_on_the_floor_laughing',
    'enraged_face',
    'smiling_face_with_horns',
    'loudly_crying_face',
    'exploding_head',
    'face_screaming_in_fear',
    'clown_face',
    'pile_of_poo'
);
create table reaction (
    message_id bigint not null,
    user_id bigint not null,
    react emoji not null,

--  проверка на то, что один пользователь поставил на одно сообщение не больше одной реакции
    constraint unique_message_user unique (message_id, user_id)
);

create table file (
    id serial primary key,
    name varchar(200) not null,
    mime_type varchar(200) not null,
    hash text not null unique,
    body bytea not null,
    created_at timestamp not null default current_timestamp
);


-- добавляем внешние ключи после того, как все таблицы созданы
alter table "user"
add constraint fk_avatar_id
foreign key (avatar_id) references file(id);

alter table message
add constraint fk_chat_id
foreign key (chat_id) references chat(id);
alter table message
add constraint fk_user_id
foreign key (user_id) references "user"(id);
alter table message
add constraint fk_file_id
foreign key (message_file) references file(id);
alter table message
add constraint fk_reply_id
foreign key (reply_message_id) references message(id);
alter table message
add constraint fk_forward_id
foreign key (forward_message_id) references message(id);

alter table private_chat
add constraint fk_user_id01
foreign key (user_id01) references "user"(id);
alter table private_chat
add constraint fk_user_id02
foreign key (user_id02) references "user"(id);

alter table group_chat
add constraint fk_avatar_id
foreign key (avatar_id) references file(id);
alter table group_chat
add constraint fk_creator_id
foreign key (creator_id) references "user"(id);

alter table chat
add constraint fk_private_chat
foreign key (private_chat_id) references private_chat(id);
alter table chat
add constraint fk_group_chat
foreign key (group_chat_id) references group_chat(id);

alter table user_group_chat
add constraint fk_user_id
foreign key (user_id) references "user"(id);
alter table user_group_chat
add constraint fk_group_chat_id
foreign key (group_chat_id) references group_chat(id);

alter table reaction
add constraint fk_message_id
foreign key (message_id) references message(id);
alter table reaction
add constraint fk_user_id
foreign key (user_id) references "user"(id);

-- триггер, который будет увеличивать количество сообщений в чате при добавлении новых сообщений
create or replace function increment_messages_amount()
returns trigger as $$
    begin
        update chat
        set messages_amount = messages_amount + 1
        where id = new.chat_id;
        return new;
    end;
$$ language plpgsql;
create trigger chat_trigger_increment_messages_count
after insert on message
for each row
execute function increment_messages_amount();

-- триггер, который при создании нового личного чата добавляет его в таблицу chat
create or replace function init_chat_with_private_chat()
returns trigger as $$
    begin
        insert into chat (private_chat_id, group_chat_id)
        values (new.id, null);
        return new;
    end;
$$ language plpgsql;
create trigger trigger_init_chat_with_private_chat
after insert on private_chat
for each row
execute function init_chat_with_private_chat();

-- триггер, который при создании нового группового чата добавляет его в таблицу chat
create or replace function init_chat_with_group_chat()
returns trigger as $$
    begin
        insert into chat (private_chat_id, group_chat_id)
        values (null, new.id);
        return new;
    end;
$$ language plpgsql;
create trigger trigger_init_chat_with_group_chat
after insert on group_chat
for each row
execute function init_chat_with_group_chat();
