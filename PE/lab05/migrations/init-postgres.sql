create table "user" (
    id serial primary key,
    first_name text not null,
    last_name text not null,
    "login" text not null unique,
    email text not null unique,
    "password" text not null,
    is_admin boolean default false,
    created_at timestamp default now()
);

create index user_login_idx on "user" ("login");

insert into "user" (first_name, last_name, "login", email, "password", is_admin) 
values 
    ('Илья', 'Ирбитский', 'admin', 'myhamster@mail.ru', '4d4384524b0f5b3a4750c60df2e70b7d4592966b757c698ed4a592ba928d9517', true), -- password: secret
    ('Матвей', 'Попов', 'papey08', 'papey08@mail.ru', 'fbcdb98979cef27dc4258157243dbf30fc91fd804b82e131fed12a51d3ec4b2d', false), -- password: aaaaAAAA
    ('Дмитрий', 'Дзюба', 'DVDemon', 'dvdemon@mail.ru', 'dce2cd1399900510a077280f6b3cef57599b46e4a1cd03d204f31e5c637eea1c', false); -- password: bbbbBBBB
