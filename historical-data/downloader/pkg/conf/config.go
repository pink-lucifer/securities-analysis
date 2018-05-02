package conf

import (
	"../persist"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func Init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("application")       // name of config file (without extension)
	viper.AddConfigPath("/etc/downloader/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.downloader") // call multiple times to add many search paths
	viper.AddConfigPath(".")                 // optionally look for config in the working directory
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {                          // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	cfg := new(Config)
	cfg, err = cfg.loadDataSourceConfig()
	if err != nil {
		panic(err)
	}
}

func (cfg *Config) loadDataSourceConfig() (*Config, error) {
	dstype := viper.GetString("datasource.type")
	dsUrl := viper.GetString("datasource.url")
	dsUserName := viper.GetString("datasource.username")
	dsPwd := viper.GetString("datasource.password")
	maxIdle := viper.GetInt("datasource.max-idle")
	maxOpen := viper.GetInt("datasource.max-open")
	maxLifeTime := viper.GetInt64("datasource.max-life-time")

	dsConfig := &persist.DataSourceConfig{
		DsType:      dstype,
		Url:         dsUrl,
		Username:    dsUserName,
		Password:    dsPwd,
		MaxIdle:     maxIdle,
		MaxOpen:     maxOpen,
		MaxLifetime: time.Duration(maxLifeTime),
	}
	cfg.DataSourceConfig = dsConfig
	return cfg, nil
}
