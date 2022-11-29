
-- CREATE DATABASE message;

-- CREATE TABLE public.user_chat(
-- 	user_id STRING NOT NULL,
-- 	chat_id STRING NOT NULL,
-- 	pinned BOOL DEFAULT false,
-- 	unread_message_count INT DEFAULT 0,
-- 	last_deleted_time TIMESTAMPTZ, -- whether show chat(last_deleted_time < last_message_update_time)
-- 	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
-- 	CONSTRAINT "primary" PRIMARY KEY (user_id ASC, chat_id ASC)
-- );

-- CREATE TYPE chat_type AS ENUM ('pvp', 'group');
-- CREATE TABLE public.chat(
-- 	id UUID DEFAULT gen_random_uuid(),
-- 	type CHAT_TYPE,
-- 	members STRING[] NOT NULL,
-- 	last_message JSONB,
-- 	last_message_update_time STRING AS (last_message->>'update_time') STORED,
-- 	name STRING, 
-- 	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
-- 	CONSTRAINT "primary" PRIMARY KEY (id ASC)
-- );

-- CREATE TABLE public.message(
-- 	id UUID DEFAULT gen_random_uuid()
-- 	chat_id STRING NOT NULL,
-- 	message JSONB,
-- 	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
-- 	CONSTRAINT "primary" PRIMARY KEY (id ASC),
-- 	UNIQUE INDEX (chat_id ASC, update_time DESC, id ASC)
-- );
