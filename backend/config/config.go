package config

import (
	"os"

	"github.com/CloudDetail/metadata/configs"
	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Logger struct {
		Level         string `mapstructure:"level"`
		EnableConsole bool   `mapstructure:"console_enable"`
		EnableFile    bool   `mapstructure:"file_enable"`
		FilePath      string `mapstructure:"file_path"`
		FileNum       int    `mapstructure:"file_num"`
		FileSize      int    `mapstructure:"file_size_mb"`
	} `mapstructure:"logger"`
	Database struct {
		Connection string `mapstructure:"connection"`
		MaxOpen    int    `mapstructure:"max_open"`
		MaxIdle    int    `mapstructure:"max_idle"`
		MaxLife    int    `mapstructure:"max_life_second"`
		MySql      struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			Database string `mapstructure:"database"`
			UserName string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			Charset  string `mapstructure:"charset"`
		} `mapstructure:"mysql"`
		Sqllite struct {
			Database string `mapstructure:"database"`
		} `mapstructure:"sqllite"`
	} `mapstructure:"database"`
	ClickHouse struct {
		Address  string `mapstructure:"address"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
	} `mapstructure:"clickhouse"`
	Promethues struct {
		Address string `mapstructure:"address"`
		Storage string `mapstructure:"storage"`
	} `mapstructure:"promethues"`
	Language struct {
		Local string `mapstructure:"local"`
	} `mapstructure:"language"`
	MetaServer struct {
		Enable           bool                     `mapstructure:"enable"`
		MetaSourceConfig configs.MetaSourceConfig `mapstructure:"meta_source_config"`
	} `mapstructure:"meta_server"`
}

func Get() *Config {
	if config == nil {
		viper.SetConfigType("yaml")
		configFile, found := os.LookupEnv("APO_CONFIG")
		if !found {
			configFile = "./config/apo.yml"
		}
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}
	}
	return config
}
