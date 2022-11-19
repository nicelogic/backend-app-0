
-- for test sql & DDL/schema/table design (sqltool needed)
-- CREATE DATABASE contacts ;

-- CREATE TABLE add_contacts_apply (
--         user_id STRING NOT NULL,
--         contacts_id STRING NOT NULL,
--         message STRING NOT NULL,
-- 		update_time TIMESTAMPTZ NOT NULL DEFAULT now(),   //Sort-Optimizing Indexes
--         CONSTRAINT "primary" PRIMARY KEY (contacts_id, user_id)
-- );
-- CREATE INDEX add_contacts_apply_hash_index
-- ON add_contacts_apply(update_time DESC)
-- USING HASH WITH BUCKET_COUNT=6;

-- //CREATE INDEX update_time_index ON add_contacts_apply (update_time) //may cause hotspot
-- SET experimental_enable_hash_sharded_indexes=on;
-- //BUCKET_COUNT(2 * count(node_num)) //covering index
-- from dbeaver INDEX recommend
-- CREATE INDEX ON add_contacts_apply (message) STORING (update_time);


CREATE TABLE contacts(
	user_id STRING NOT NULL,
	contacts_id STRING NOT NULL,
	remark_name STRING NOT NULL, 
	update_time TIMESTAMPTZ NOT NULL DEFAULT now(),
	CONSTRAINT "primary" PRIMARY KEY (user_id ASC, contacts_id ASC),
    INDEX remark_name_index(remark_name ASC)
);