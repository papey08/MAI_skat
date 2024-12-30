--
-- PostgreSQL database dump
--

-- Dumped from database version 15.5 (Debian 15.5-1.pgdg120+1)
-- Dumped by pg_dump version 15.5 (Debian 15.5-1.pgdg120+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: admin_rights; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.admin_rights AS ENUM (
    'creator',
    'admin'
);


ALTER TYPE public.admin_rights OWNER TO postgres;

--
-- Name: emoji; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.emoji AS ENUM (
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


ALTER TYPE public.emoji OWNER TO postgres;

--
-- Name: add_creator_to_group_chat(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.add_creator_to_group_chat() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    begin
        insert into user_group_chat (user_id, group_chat_id, rights) values
            (new.creator_id, new.id, 'creator');
        return new;
    end;
$$;


ALTER FUNCTION public.add_creator_to_group_chat() OWNER TO postgres;

--
-- Name: check_if_group_chat_creator_exists(bigint); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.check_if_group_chat_creator_exists(p_creator_id bigint) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
    begin
        return p_creator_id in (select id from "user");
    end;
$$;


ALTER FUNCTION public.check_if_group_chat_creator_exists(p_creator_id bigint) OWNER TO postgres;

--
-- Name: check_if_user_in_chat(bigint, bigint); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.check_if_user_in_chat(p_user_id bigint, p_chat_id bigint) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
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
$$;


ALTER FUNCTION public.check_if_user_in_chat(p_user_id bigint, p_chat_id bigint) OWNER TO postgres;

--
-- Name: increment_messages_amount(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.increment_messages_amount() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    begin
        update chat
        set messages_amount = messages_amount + 1
        where id = new.chat_id;
        return new;
    end;
$$;


ALTER FUNCTION public.increment_messages_amount() OWNER TO postgres;

--
-- Name: init_chat_with_group_chat(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.init_chat_with_group_chat() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    begin
        insert into chat (private_chat_id, group_chat_id)
        values (null, new.id);
        return new;
    end;
$$;


ALTER FUNCTION public.init_chat_with_group_chat() OWNER TO postgres;

--
-- Name: init_chat_with_private_chat(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.init_chat_with_private_chat() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    begin
        insert into chat (private_chat_id, group_chat_id)
        values (new.id, null);
        return new;
    end;
$$;


ALTER FUNCTION public.init_chat_with_private_chat() OWNER TO postgres;

--
-- Name: private_chat_order_user_ids(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.private_chat_order_user_ids() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
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
$$;


ALTER FUNCTION public.private_chat_order_user_ids() OWNER TO postgres;

--
-- Name: select_and_mark_as_read_messages(bigint, bigint, bigint); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.select_and_mark_as_read_messages(p_chat_id bigint, p_limit bigint, p_offset bigint) RETURNS TABLE(id bigint, chat_id bigint, user_id bigint, message_text text, message_file bigint, reply_message_id bigint, forward_message_id bigint, is_read boolean, created_at timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
    begin
        with selected_rows as (
            select * from message
            where chat_id = p_chat_id
            order by created_at desc
            limit p_limit offset p_offset
        )
        update message
        set is_read = true
        from selected_rows
        where chat_id = p_chat_id and is_read = false;

        return query
        select id, chat_id, user_id, message_text, message_file, reply_message_id, forward_message_id, is_read, created_at from message where chat_id = p_chat_id
        order by created_at desc
        limit p_limit offset p_offset;
    end;
$$;


ALTER FUNCTION public.select_and_mark_as_read_messages(p_chat_id bigint, p_limit bigint, p_offset bigint) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: chat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.chat (
    id integer NOT NULL,
    private_chat_id bigint,
    group_chat_id bigint,
    messages_amount bigint DEFAULT 0 NOT NULL,
    CONSTRAINT check_chat_id_is_null CHECK ((((private_chat_id IS NULL) AND (group_chat_id IS NOT NULL)) OR ((private_chat_id IS NOT NULL) AND (group_chat_id IS NULL))))
);


ALTER TABLE public.chat OWNER TO postgres;

--
-- Name: chat_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.chat_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.chat_id_seq OWNER TO postgres;

--
-- Name: chat_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.chat_id_seq OWNED BY public.chat.id;


--
-- Name: file; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.file (
    id integer NOT NULL,
    name character varying(200) NOT NULL,
    mime_type character varying(200) NOT NULL,
    hash text NOT NULL,
    body bytea NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.file OWNER TO postgres;

--
-- Name: file_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.file_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.file_id_seq OWNER TO postgres;

--
-- Name: file_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.file_id_seq OWNED BY public.file.id;


--
-- Name: group_chat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.group_chat (
    id integer NOT NULL,
    title character varying(200) NOT NULL,
    description text,
    avatar_id bigint,
    creator_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.group_chat OWNER TO postgres;

--
-- Name: group_chat_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.group_chat_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.group_chat_id_seq OWNER TO postgres;

--
-- Name: group_chat_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.group_chat_id_seq OWNED BY public.group_chat.id;


--
-- Name: message; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.message (
    id integer NOT NULL,
    chat_id bigint NOT NULL,
    user_id bigint NOT NULL,
    message_text text,
    message_file bigint,
    reply_message_id bigint,
    forward_message_id bigint,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT check_columns_not_nil CHECK (((message_text IS NOT NULL) OR (message_file IS NOT NULL) OR (reply_message_id IS NOT NULL) OR (forward_message_id IS NOT NULL))),
    CONSTRAINT check_user_in_chat CHECK (public.check_if_user_in_chat(user_id, chat_id))
);


ALTER TABLE public.message OWNER TO postgres;

--
-- Name: message_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.message_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.message_id_seq OWNER TO postgres;

--
-- Name: message_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.message_id_seq OWNED BY public.message.id;


--
-- Name: private_chat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.private_chat (
    id integer NOT NULL,
    user_id01 bigint NOT NULL,
    user_id02 bigint NOT NULL
);


ALTER TABLE public.private_chat OWNER TO postgres;

--
-- Name: private_chat_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.private_chat_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.private_chat_id_seq OWNER TO postgres;

--
-- Name: private_chat_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.private_chat_id_seq OWNED BY public.private_chat.id;


--
-- Name: reaction; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reaction (
    message_id bigint NOT NULL,
    user_id bigint NOT NULL,
    react public.emoji NOT NULL
);


ALTER TABLE public.reaction OWNER TO postgres;

--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    first_name character varying(100) NOT NULL,
    second_name character varying(100),
    email character varying(100) NOT NULL,
    avatar_id bigint,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT check_columns_not_empty CHECK (((TRIM(BOTH FROM username) <> ''::text) AND (TRIM(BOTH FROM first_name) <> ''::text))),
    CONSTRAINT validate_email CHECK (((email)::text ~~ '%_@%_.%_'::text))
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Name: user_group_chat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_group_chat (
    user_id bigint NOT NULL,
    group_chat_id bigint NOT NULL,
    rights public.admin_rights
);


ALTER TABLE public.user_group_chat OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- Name: chat id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.chat ALTER COLUMN id SET DEFAULT nextval('public.chat_id_seq'::regclass);


--
-- Name: file id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.file ALTER COLUMN id SET DEFAULT nextval('public.file_id_seq'::regclass);


--
-- Name: group_chat id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.group_chat ALTER COLUMN id SET DEFAULT nextval('public.group_chat_id_seq'::regclass);


--
-- Name: message id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message ALTER COLUMN id SET DEFAULT nextval('public.message_id_seq'::regclass);


--
-- Name: private_chat id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.private_chat ALTER COLUMN id SET DEFAULT nextval('public.private_chat_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Data for Name: chat; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.chat (id, private_chat_id, group_chat_id, messages_amount) FROM stdin;
6	4	\N	0
1	\N	1	4
3	1	\N	2
4	2	\N	2
5	3	\N	2
2	\N	2	5
\.


--
-- Data for Name: file; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.file (id, name, mime_type, hash, body, created_at) FROM stdin;
\.


--
-- Data for Name: group_chat; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.group_chat (id, title, description, avatar_id, creator_id, created_at) FROM stdin;
1	chat01	\N	\N	1	2024-10-27 04:18:24.076978
2	chat02	\N	\N	2	2024-10-27 04:18:24.076978
\.


--
-- Data for Name: message; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.message (id, chat_id, user_id, message_text, message_file, reply_message_id, forward_message_id, created_at) FROM stdin;
1	1	1	Hello world 01	\N	\N	\N	2024-10-27 04:18:24.107746
2	1	1	Hello world 02	\N	\N	\N	2024-10-27 04:18:24.107746
3	1	2	Hello world 03	\N	\N	\N	2024-10-27 04:18:24.107746
4	1	3	Hello world 04	\N	\N	\N	2024-10-27 04:18:24.107746
5	2	1	Hello world 05	\N	\N	\N	2024-10-27 04:18:24.107746
6	2	2	Hello world 06	\N	\N	\N	2024-10-27 04:18:24.107746
7	2	3	Hello world 07	\N	\N	\N	2024-10-27 04:18:24.107746
8	2	1	Hello world 08	\N	\N	\N	2024-10-27 04:18:24.107746
9	3	1	Hello world 09	\N	\N	\N	2024-10-27 04:18:24.107746
10	3	2	Hello world 10	\N	\N	\N	2024-10-27 04:18:24.107746
11	4	3	Hello world 11	\N	\N	\N	2024-10-27 04:18:24.107746
12	4	4	Hello world 12	\N	\N	\N	2024-10-27 04:18:24.107746
13	5	5	Hello world 13	\N	\N	\N	2024-10-27 04:18:24.107746
14	5	6	Hello world 14	\N	\N	\N	2024-10-27 04:18:24.107746
15	2	1	Hello world 15	\N	\N	\N	2024-10-27 04:19:54.282841
\.


--
-- Data for Name: private_chat; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.private_chat (id, user_id01, user_id02) FROM stdin;
1	1	2
2	3	4
3	5	6
4	7	8
\.


--
-- Data for Name: reaction; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.reaction (message_id, user_id, react) FROM stdin;
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" (id, username, first_name, second_name, email, avatar_id, created_at) FROM stdin;
1	papey08	Matvey	Popov	matveypopov@test.com	\N	2024-10-27 04:18:24.065912
2	Farayld	Willie	Sandoval	williesandoval@test.com	\N	2024-10-27 04:18:24.065912
3	Alistel	Bella	Norman	bellanorman@test.com	\N	2024-10-27 04:18:24.065912
4	Erielie	Phoebe	Le	phoebele@test.com	\N	2024-10-27 04:18:24.065912
5	Oronanc	Yasmin	Vargas	yasminvargas@test.com	\N	2024-10-27 04:18:24.065912
6	Miziala	Adele	Byrne	adelebyrne@test.com	\N	2024-10-27 04:18:24.065912
7	Zievetr	Lillie	Hicks	lilliehicks@test.com	\N	2024-10-27 04:18:24.065912
8	Iliamaw	Zainab	Steele	zainabsteele@test.com	\N	2024-10-27 04:18:24.065912
\.


--
-- Data for Name: user_group_chat; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_group_chat (user_id, group_chat_id, rights) FROM stdin;
1	1	creator
2	2	creator
2	1	\N
3	1	\N
1	2	\N
3	2	\N
\.


--
-- Name: chat_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.chat_id_seq', 6, true);


--
-- Name: file_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.file_id_seq', 1, false);


--
-- Name: group_chat_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.group_chat_id_seq', 2, true);


--
-- Name: message_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.message_id_seq', 15, true);


--
-- Name: private_chat_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.private_chat_id_seq', 4, true);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 8, true);


--
-- Name: chat chat_group_chat_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.chat
    ADD CONSTRAINT chat_group_chat_id_key UNIQUE (group_chat_id);


--
-- Name: chat chat_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.chat
    ADD CONSTRAINT chat_pkey PRIMARY KEY (id);


--
-- Name: chat chat_private_chat_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.chat
    ADD CONSTRAINT chat_private_chat_id_key UNIQUE (private_chat_id);


--
-- Name: file file_hash_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.file
    ADD CONSTRAINT file_hash_key UNIQUE (hash);


--
-- Name: file file_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.file
    ADD CONSTRAINT file_pkey PRIMARY KEY (id);


--
-- Name: group_chat group_chat_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.group_chat
    ADD CONSTRAINT group_chat_pkey PRIMARY KEY (id);


--
-- Name: message message_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT message_pkey PRIMARY KEY (id);


--
-- Name: private_chat private_chat_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.private_chat
    ADD CONSTRAINT private_chat_pkey PRIMARY KEY (id);


--
-- Name: reaction unique_message_user; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction
    ADD CONSTRAINT unique_message_user UNIQUE (message_id, user_id);


--
-- Name: user_group_chat unique_user_group; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_group_chat
    ADD CONSTRAINT unique_user_group UNIQUE (user_id, group_chat_id);


--
-- Name: private_chat unique_users; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.private_chat
    ADD CONSTRAINT unique_users UNIQUE (user_id01, user_id02);


--
-- Name: user user_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: user user_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_username_key UNIQUE (username);


--
-- Name: message chat_trigger_increment_messages_count; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER chat_trigger_increment_messages_count AFTER INSERT ON public.message FOR EACH ROW EXECUTE FUNCTION public.increment_messages_amount();


--
-- Name: group_chat group_chat_trigger_add_creator_to_group_chat; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER group_chat_trigger_add_creator_to_group_chat AFTER INSERT ON public.group_chat FOR EACH ROW EXECUTE FUNCTION public.add_creator_to_group_chat();


--
-- Name: private_chat private_chat_trigger_order_user_ids; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER private_chat_trigger_order_user_ids BEFORE INSERT ON public.private_chat FOR EACH ROW EXECUTE FUNCTION public.private_chat_order_user_ids();


--
-- Name: group_chat trigger_init_chat_with_group_chat; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trigger_init_chat_with_group_chat AFTER INSERT ON public.group_chat FOR EACH ROW EXECUTE FUNCTION public.init_chat_with_group_chat();


--
-- Name: private_chat trigger_init_chat_with_private_chat; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trigger_init_chat_with_private_chat AFTER INSERT ON public.private_chat FOR EACH ROW EXECUTE FUNCTION public.init_chat_with_private_chat();


--
-- Name: user fk_avatar_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT fk_avatar_id FOREIGN KEY (avatar_id) REFERENCES public.file(id);


--
-- Name: group_chat fk_avatar_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.group_chat
    ADD CONSTRAINT fk_avatar_id FOREIGN KEY (avatar_id) REFERENCES public.file(id);


--
-- Name: message fk_chat_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT fk_chat_id FOREIGN KEY (chat_id) REFERENCES public.chat(id);


--
-- Name: group_chat fk_creator_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.group_chat
    ADD CONSTRAINT fk_creator_id FOREIGN KEY (creator_id) REFERENCES public."user"(id);


--
-- Name: message fk_file_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT fk_file_id FOREIGN KEY (message_file) REFERENCES public.file(id);


--
-- Name: message fk_forward_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT fk_forward_id FOREIGN KEY (forward_message_id) REFERENCES public.message(id);


--
-- Name: chat fk_group_chat; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.chat
    ADD CONSTRAINT fk_group_chat FOREIGN KEY (group_chat_id) REFERENCES public.group_chat(id);


--
-- Name: user_group_chat fk_group_chat_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_group_chat
    ADD CONSTRAINT fk_group_chat_id FOREIGN KEY (group_chat_id) REFERENCES public.group_chat(id);


--
-- Name: reaction fk_message_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction
    ADD CONSTRAINT fk_message_id FOREIGN KEY (message_id) REFERENCES public.message(id);


--
-- Name: chat fk_private_chat; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.chat
    ADD CONSTRAINT fk_private_chat FOREIGN KEY (private_chat_id) REFERENCES public.private_chat(id);


--
-- Name: message fk_reply_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT fk_reply_id FOREIGN KEY (reply_message_id) REFERENCES public.message(id);


--
-- Name: message fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public."user"(id);


--
-- Name: user_group_chat fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_group_chat
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public."user"(id);


--
-- Name: reaction fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public."user"(id);


--
-- Name: private_chat fk_user_id01; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.private_chat
    ADD CONSTRAINT fk_user_id01 FOREIGN KEY (user_id01) REFERENCES public."user"(id);


--
-- Name: private_chat fk_user_id02; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.private_chat
    ADD CONSTRAINT fk_user_id02 FOREIGN KEY (user_id02) REFERENCES public."user"(id);


--
-- PostgreSQL database dump complete
--

