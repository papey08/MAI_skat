-- CTE
-- выберем все групповые чаты, в которых написано больше 2 сообщений
with group_chats as (
    select chat.id, messages_amount, group_chat.title from chat
    inner join group_chat on group_chat_id = group_chat.id
    where private_chat_id is null
)
select id, title from group_chats
where messages_amount > 2;
