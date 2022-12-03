
-- CREATE DATABASE message;
-- CREATE TYPE chat_type AS ENUM ('p2p', 'group');

CREATE TABLE public.chat(
	id STRING NOT NULL,
	type CHAT_TYPE NOT NULL,
	members STRING[] NOT NULL,
	name STRING, 
	last_message JSONB,
	--last_message_update_time STRING AS (last_message->>'update_time') STORED,
	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (id ASC)
	--INVERTED INDEX (members)
);
--这种设计，每次chat发送一条消息，会更新每个members,排序相关,需要更新时间
--创建chat事务创建
--查询需要join
--个人的配置不影响主体chat
--主体chat信息更改也不影响个人配置
--没办法在一个chat表里面遍历更改所有的user setting(JSONB方式)
--影响用户对chat特定状态+影响排序的一个表， chat通用信息一个表
--还要一个方案，每次排序user_chat,都全部查询user所有chat,再排序
--但每次查询都全量查找
--还要就是未读信息数量，也是每个用户特定，又和每条消息关联的
--每次有新消息，就得更新用户特定状态。。顺带也更新时间
--最终选择last_message_time在user_chat维护的方案
CREATE TABLE public.user_chat(
	user_id STRING NOT NULL,
	chat_id STRING NOT NULL,
	priority INT DEFAULT 0, --default: 0, pinned: 10
	last_message_time TIMESTAMPTZ,
	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (user_id ASC, chat_id ASC),
	UNIQUE INDEX (user_id ASC, priority DESC, last_message_time DESC, chat_id ASC)
	INDEX (chat_id ASC)
);
CREATE INDEX ON user_chat (chat_id) STORING (priority);

CREATE TABLE public.message(
	id UUID DEFAULT gen_random_uuid(),
	chat_id STRING NOT NULL,
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