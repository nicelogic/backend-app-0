
-- for test sql & DDL/schema/table design (sqltool needed)
-- CREATE DATABASE contacts ;

-- CREATE TABLE add_contacts_apply (
--         user_id STRING NOT NULL,
--         contacts_id STRING NOT NULL,
--         message STRING NOT NULL,
-- 		update_time TIMESTAMPTZ NOT NULL DEFAULT now(),   //Sort-Optimizing Indexes
--         CONSTRAINT "primary" PRIMARY KEY (contacts_id, user_id)
-- );
-- //CREATE INDEX update_time_index ON add_contacts_apply (update_time) //may cause hotspot
-- SET experimental_enable_hash_sharded_indexes=on;
-- CREATE INDEX add_contacts_apply_hash_index ON add_contacts_apply(update_time)
-- USING HASH WITH BUCKET_COUNT=6; //BUCKET_COUNT(2 * count(node_num)) //covering index

-- DROP TABLE add_contacts_apply;

-- CREATE TABLE contacts(
--         id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
--         user_id STRING NOT NULL,
--         contacts_id STRING NOT NULL,
--         remark_name STRING NOT NULL, //Sort-Optimizing Indexes
-- 		update_time TIMESTAMPTZ NOT NULL,
--         CONSTRAINT "primary" PRIMARY KEY (user_id ASC, contacts_id ASC),
--         FAMILY "primary" (id, city, name, address, credit_card)
-- );