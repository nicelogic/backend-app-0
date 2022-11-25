
-- CREATE DATABASE auth;

CREATE TABLE public.auth(
	auth_id STRING NOT NULL,
	auth_id_type STRING NOT NULL,
	user_id STRING NOT NULL,
	auth_id_type_username_pwd STRING,
	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (auth_id ASC, auth_id_type ASC)
);
