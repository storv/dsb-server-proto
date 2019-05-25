package config

type AppConfig struct {
	ZkConnInfo       string `toml:"zk_conn_info"`
	ZkPushServerPath string `toml:"zk_push_server_path"`
	ZkServerDataPath string `toml:"zk_server_data_path"`
	IdcName          string `toml:"idc_name"`
	GrpcPort         string `toml:"grpc_port"`
	HttpPort         string `toml:"http_port"`
	RedisHost        string `toml:"redis_host"`
	RedisPass        string `toml:"redis_pass"`
}
