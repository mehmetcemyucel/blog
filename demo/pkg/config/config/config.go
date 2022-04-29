package config

import (
	"fmt"
	"github.com/spf13/viper"
	"helloWorld/pkg/util"
)

const (
	defaultConfigPath     = "config"
	defaultConfigFileName = "application"
	defaultConfigFileType = "yaml"
	ymlTagName            = "mapstructure"
)

type (
	LogConfig struct {
		KafkaLogConfig KafkaLogConfig `mapstructure:"kafka"`
		FileLogConfig  FileLogConfig  `mapstructure:"file"`
	}

	KafkaLogConfig struct {
		KafkaCluster string `mapstructure:"kafka-cluster"`
		Topic        string `mapstructure:"topic"`
	}

	FileLogConfig struct {
		LogPath string `mapstructure:"log-path"`
	}

	PostgreConfig struct {
		Host        string `mapstructure:"host"`
		Port        string `mapstructure:"port"`
		Username    string `mapstructure:"username"`
		Password    string `mapstructure:"password"`
		Name        string `mapstructure:"name"`
		TablePrefix string `mapstructure:"table-prefix"`
		Ssl         bool   `mapstructure:"ssl"`
	}

	ServerConfig struct {
		Port string   `mapstructure:"port"`
		Cors []string `mapstructure:"cors"`
	}

	ExampleConfig struct {
		ExampleParam1 string `mapstructure:"example-param1"`
		ExampleParam2 string `mapstructure:"example-param2"`
	}

	Config struct {
		Postgre PostgreConfig `mapstructure:"postgre"`
		Server  ServerConfig  `mapstructure:"app-server"`
		Example ExampleConfig `mapstructure:"example"`
		Log     LogConfig     `mapstructure:"log"`
	}
)

func NewConfig() *Config {
	return ReadConfig(&Config{}, util.ReadEnv())
}

var ReadConfig = func(c *Config, env string) *Config {
	fmt.Println("Configuration read initiated...")
	v := viper.New()
	switch {
	case env == "ENV":
		v = readFromEnv(v)
	case env == "REMOTE":
		v = readFromConfigServer(v)
	case env == "LOCAL":
		v = readFromLocalAppYml(v)
	default:
		v = readFromLocalAppYml(v)
	}
	if err := v.Unmarshal(c); err != nil {
		panic("Configuration unmarshalling occured an error, terminating" + err.Error())
	}
	return c
}
