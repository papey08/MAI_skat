insert into "user" (username, first_name, second_name, email, avatar_id) values
    ('papey08', 'Matvey', 'Popov', 'matveypopov@test.com', NULL),
    ('Farayld', 'Willie', 'Sandoval', 'williesandoval@test.com', NULL),
    ('Alistel', 'Bella', 'Norman', 'bellanorman@test.com', NULL),
    ('Erielie', 'Phoebe', 'Le', 'phoebele@test.com', NULL),
    ('Oronanc', 'Yasmin', 'Vargas', 'yasminvargas@test.com', NULL),
    ('Miziala', 'Adele', 'Byrne', 'adelebyrne@test.com', NULL),
    ('Zievetr', 'Lillie', 'Hicks', 'lilliehicks@test.com', NULL),
    ('Iliamaw', 'Zainab', 'Steele', 'zainabsteele@test.com', NULL);

insert into group_chat (title, description, avatar_id, creator_id) values
    ('chat01', NULL, NULL, 1),
    ('chat02', NULL, NULL, 2);

insert into private_chat (user_id01, user_id02) values
    (1, 2),
    (3, 4),
    (5, 6),
    (8, 7);

insert into user_group_chat (user_id, group_chat_id, rights) values
    (2, 1, NULL),
    (3, 1, NULL),
    (1, 2, NULL),
    (3, 2, NULL);

insert into message (chat_id, user_id, message_text, message_file, reply_message_id, forward_message_id) values
    (1, 1, 'Hello world 01', NULL, NULL, NULL),
    (1, 1, 'Hello world 02', NULL, NULL, NULL),
    (1, 2, 'Hello world 03', NULL, NULL, NULL),
    (1, 3, 'Hello world 04', NULL, NULL, NULL),
    (2, 1, 'Hello world 05', NULL, NULL, NULL),
    (2, 2, 'Hello world 06', NULL, NULL, NULL),
    (2, 3, 'Hello world 07', NULL, NULL, NULL),
    (2, 1, 'Hello world 08', NULL, NULL, NULL),
    (3, 1, 'Hello world 09', NULL, NULL, NULL),
    (3, 2, 'Hello world 10', NULL, NULL, NULL),
    (4, 3, 'Hello world 11', NULL, NULL, NULL),
    (4, 4, 'Hello world 12', NULL, NULL, NULL),
    (5, 5, 'Hello world 13', NULL, NULL, NULL),
    (5, 6, 'Hello world 14', NULL, NULL, NULL);
