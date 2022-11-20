
-- for test sql & DDL/schema/table design (sqltool needed)
-- CREATE DATABASE contacts ;

CREATE TABLE public.add_contacts_apply (
	user_id STRING NOT NULL,
	contacts_id STRING NOT NULL,
	message STRING NOT NULL,
	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (contacts_id ASC, user_id ASC),
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
-- create unique index contacts_update_time_index on add_contacts_apply(contacts_id asc, update_time desc) 


CREATE TABLE contacts(
	user_id STRING NOT NULL,
	contacts_id STRING NOT NULL,
	remark_name STRING NOT NULL, 
	update_time TIMESTAMPTZ NOT NULL DEFAULT now(),
	CONSTRAINT "primary" PRIMARY KEY (user_id ASC, contacts_id ASC),
    INDEX remark_name_index(remark_name ASC)
);