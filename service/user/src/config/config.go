package config

type Config struct{
	Path string
	Listen_address string
	Db_name string
	Db_config_file_path string
	Db_pool_connections_num int32
	S3_config_file_path string
	S3_bucket_name string
	S3_signed_object_expired_seconds int32
	Public_key_file_path string
	Private_key_file_path string
}

