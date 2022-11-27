
-- CREATE DATABASE userdb

CREATE TABLE public.user(
	id STRING NOT NULL,
	data JSONB,
	name STRING AS (data->>'name') STORED,  
	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	INDEX name_index (name ASC) STORING (data, update_time)
);


-- CREATE INDEX ON "user" (name) STORING (data, update_time); 
-- DROP INDEX "user"@name_index;