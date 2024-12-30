-- subquery
-- выберем все сообщения, которые написал пользователь по имени Матвей
select first_name, message_text, chat_id from (
    select * from message
    inner join "user" u on u.id = message.user_id
) as t1
where first_name = 'Matvey';
