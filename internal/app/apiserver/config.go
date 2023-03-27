package apiserver

type Config struct {
	BindAddress string `json:"bind_address"`
}

func NewConfig() *Config {
	return &Config{
		BindAddress: "0.0.0.0:8080",
	}
}
