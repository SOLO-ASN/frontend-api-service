package config

import (
	"errors"
)

var (
	// config
	config *Config

	// config file path
	config_path = "config.yaml"
)

var (
	ErrFileNotFound = errors.New("config file not found")
)

func Init(path string) {
	config = &Config{}
	if path != "" {
		config_path = path
	}
	err := config.parse()
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	if config == nil {
		panic("config is nil")
	}
	return config
}

type Config struct {
	App        App        `yaml:"app" json:"app"`
	Logger     Logger     `yaml:"logger" json:"logger"`
	Redis      Redis      `yaml:"cache" json:"cache"`
	Mysql      Mysql      `yaml:"mysql" json:"mysql"`
	MiddleWare Middleware `yaml:"middleware" json:"middleware"`
}

type App struct {
	Name string `yaml:"name" json:"name"`
	Host string `yaml:"host" json:"host"`
}

type Redis struct {
	EnableCluster bool     `yaml:"enableCluster" json:"enableCluster"`
	AddressList   []string `yaml:"addressList" json:"addressList"`
	Password      string   `yaml:"password" json:"password"`
	DB            int      `yaml:"db" json:"db"`
	Prefix        string   `yaml:"prefix" json:"prefix"`
}

type Logger struct{}

type Mysql struct {
}

type Middleware struct {
	EnableCircuitBreaker bool `yaml:"enableCircuitBreaker" json:"enableCircuitBreaker"` // circuit breaker
	EnableRateLimit      bool `yaml:"enableRateLimit" json:"enableRateLimit"`           // rate limit
	EnableMetrics        bool `yaml:"enableMetrics" json:"enableMetrics"`               // metrics
	EnableTrace          bool `yaml:"enableTrace" json:"enableTrace"`                   // trace
}

func (c *Config) parse() error {
	//TODO implement me
	panic("implement me")
	//f, err := os.Open(config_path)
	//defer f.Close()
	//if err != nil {
	//	return ErrFileNotFound
	//}
	//
	return nil
}
