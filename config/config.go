package config

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	// config
	config *Config

	// config file path
	config_path = "config.yaml"
)

var (
	ErrFileNotFound = errors.New("config file not found")
	ErrConfigParse  = errors.New("config parse error")
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

type Logger struct {
	Enable bool   `yaml:"enable" json:"enable"`
	Format string `yaml:"format" json:"format"`
	Level  string `yaml:"level" json:"level"`
	Path   string `yaml:"path" json:"path"`
}

type Mysql struct {
	Dsn             string        `yaml:"dsn" json:"dsn"`
	MaxIdleConns    int           `yaml:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int           `yaml:"maxOpenConns" json:"maxOpenConns"`
	SlowThreshold   time.Duration `yaml:"slowThreshold" json:"slowThreshold"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime" json:"connMaxLifetime"`
	EnableLog       bool          `yaml:"enableLog" json:"enableLog"`
}

type Middleware struct {
	EnableCircuitBreaker bool `yaml:"enableCircuitBreaker" json:"enableCircuitBreaker"` // circuit breaker
	EnableRateLimit      bool `yaml:"enableRateLimit" json:"enableRateLimit"`           // rate limit
	EnableMetrics        bool `yaml:"enableMetrics" json:"enableMetrics"`               // metrics
	EnableTrace          bool `yaml:"enableTrace" json:"enableTrace"`                   // trace
}

func (c *Config) parse() error {
	// check file exist
	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		return ErrFileNotFound
	}

	// use viper
	filePath, fileName := filepath.Split(config_path)
	ext := strings.TrimLeft(path.Ext(fileName), ".")
	fileName = strings.Replace(fileName, "."+ext, "", 1)

	viper.AddConfigPath(filePath)
	viper.AddConfigPath(".")
	viper.SetConfigName(fileName)
	viper.SetConfigType(ext)
	err := viper.ReadInConfig()
	if err != nil {
		return ErrConfigParse
	}

	err = viper.Unmarshal(c)
	if err != nil {
		return ErrConfigParse
	}

	return nil
}
