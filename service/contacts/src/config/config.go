package config

type Config struct{
	Db_name string
	Db_config_file_path string
	Db_pool_connections_num int32
	Pulsar_url string
	Path string
	Listen_address string
}

