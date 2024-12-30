-- пишем сообщение в чат с id = 2
insert into message (chat_id, user_id, message_text, message_file, reply_message_id, forward_message_id) values
    (2, 1, 'Hello world 15', null, null, null);
