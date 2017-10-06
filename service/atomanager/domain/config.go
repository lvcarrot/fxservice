package domain

type DatabaseConf struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type RedisConf struct {
	Host string
	Port int
	DB   int
}

type ServerConf struct {
	ExternalListenAddress string
	InternalListenAddress string
	AppKey                string
}
