package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name    string `yaml:"name"`
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Prefork bool   `yaml:"prefork"`
		Debug   bool   `yaml:"debug"`
		Secret  string `yaml:"secret"`
	} `yaml:"app"`
	Auth struct {
		AccessTokenExpMins  int `yaml:"access_token_exp_mins"`
		RefreshTokenExpDays int `yaml:"refresh_token_exp_days"`
	} `yaml:"auth"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
	ImageKit struct {
		PrivateKey  string `yaml:"private_key"`
		PublicKey   string `yaml:"public_key"`
		URLEndpoint string `yaml:"url_endpoint"`
	} `yaml:"imagekit"`
}

func NewConfig() *viper.Viper {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")

	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	return config
}
