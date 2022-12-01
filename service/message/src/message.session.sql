
-- CREATE DATABASE message;
-- CREATE TYPE chat_type AS ENUM ('p2p', 'group');

CREATE TABLE public.chat(
	id UUID DEFAULT gen_random_uuid(),
	type CHAT_TYPE NOT NULL,
	members STRING[] NOT NULL,
	members_setting JSONB,
	last_message JSONB,
	last_message_update_time STRING AS (last_message->>'update_time') STORED,
	name STRING, 
	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	INVERTED INDEX (members)
);

CREATE TABLE public.message(
	id UUID DEFAULT gen_random_uuid(),
	chat_id UUID NOT NULL,
	content JSONB,
	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	UNIQUE INDEX (chat_id ASC, update_time DESC, id ASC)
);




-- CREATE INVERTED INDEX ON chat (members);
-- deprecate
-- CREATE TABLE public.user_chat(
-- 	user_id STRING NOT NULL,
-- 	chat_id UUID NOT NULL,
-- 	pinned BOOL DEFAULT false,
-- 	unread_message_count INT DEFAULT 0,
-- 	last_deleted_time TIMESTAMPTZ, -- whether show chat(last_deleted_time < last_message_update_time)
-- 	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
-- 	CONSTRAINT "primary" PRIMARY KEY (user_id ASC, chat_id ASC)
-- );






{"1": {"last_deleted_time": "", "pinned": true, "unread_message_count": 0}, "2": {"last_deleted_time": "", "pinned": true, "unread_message_count": 0}}
{"1": {"last_deleted_time": "", "pinned": true, "unread_message_count": 0}, "2": {"last_deleted_time": "", "pinned": true, "unread_message_count": 0}, "3": {"last_deleted_time": "", "pinned": true, "unread_message_count": 0}}