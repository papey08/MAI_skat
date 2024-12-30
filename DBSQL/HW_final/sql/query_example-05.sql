-- window function
-- пронумеруем участников чата с id = 1 по алфавитному порядку фамилии имени
select second_name, first_name,
    row_number() over (order by second_name, first_name) as numeration
from "user"
inner join user_group_chat ugc on "user".id = ugc.user_id
inner join group_chat gc on ugc.group_chat_id = gc.id
inner join chat c on gc.id = c.group_chat_id
where c.id = 1;
