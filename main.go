package main

import (
	"fmt"

	"github.com/bayupermadi/mon-gearman/monitor"
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

	//for {

	monitor.LogSize(logPath, int64(logMaxSize))

	//	<-time.After(time.Second * 30)
	//}

}