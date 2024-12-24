package config

type httpConfig struct {
	ListenAddress string
}

type Config struct {
	HTTP  httpConfig
	Mongo mongo
}

type mongo struct {
	Url string `json:"url"`
	Db  string `json:"db"`
}
