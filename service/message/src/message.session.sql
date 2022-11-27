
-- CREATE DATABASE message;

CREATE TABLE public.user_chat(
	user_id STRING NOT NULL,
	chat_id STRING NOT NULL,
	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	UNIQUE INDEX contacts_update_time_index (contacts_id ASC, update_time DESC)
);

CREATE TABLE public.chat(
	id UUID DEFAULT gen_random_uuid()
	name STRING NOT NULL,
	last_message_id STRING NOT NULL,
	members STRING[] NOT NULL,
	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	UNIQUE INDEX contacts_update_time_index (contacts_id ASC, update_time DESC)
);

CREATE TABLE public.chat_message(
	id UUID DEFAULT gen_random_uuid()
	name STRING NOT NULL,
	last_message_id STRING NOT NULL,
	members
	message STRING NOT NULL,
	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	UNIQUE INDEX contacts_update_time_index (contacts_id ASC, update_time DESC)
);

-- CREATE INDEX update_time_index ON add_contacts_apply (update_time)
-- //may cause hotspot 
-- SET experimental_enable_hash_sharded_indexes=on;
-- //BUCKET_COUNT(2 * count(node_num)) //covering index
-- CREATE INDEX add_contacts_apply_hash_index
-- ON add_contacts_apply(update_time DESC)
-- USING HASH WITH BUCKET_COUNT=6;
-- //may not hint the index
create unique index default_unique_index on add_contacts_apply(contacts_id asc, update_time desc, user_id asc) 


CREATE TABLE contacts(
	user_id STRING NOT NULL,
	contacts_id STRING NOT NULL,
	remark_name STRING NOT NULL, 
	update_time TIMESTAMPTZ NOT NULL DEFAULT now(),
	CONSTRAINT "primary" PRIMARY KEY (user_id ASC, contacts_id ASC),
    UNIQUE INDEX default_unique_index(user_id asc, remark_name asc, contacts_id asc)
);

create unique index default_unique_index on contacts(user_id asc, remark_name asc, contacts_id asc) 