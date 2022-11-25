package config

type Config struct{
	Path string
	Listen_address string
	Db_name string
	Db_config_file_path string
	Db_pool_connections_num int32
	// Pulsar_config_file_path string
	// Pulsar_topic string
}

