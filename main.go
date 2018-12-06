package main

import (
	"fmt"

	"github.com/bayupermadi/mon-gearmand/monitor"
	"github.com/spf13/viper"
)

func main() {
	// config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	logPath := viper.Get("app.log.files").(string)
	logMaxSize := viper.Get("app.log.max-size").(int)

	monitor.LogSize(logPath, int64(logMaxSize))

}
