package core

import (
	"fmt"
	"gin-vue-admin/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var configFile = "config_local.yaml"

func init() {
	env := os.Getenv("GIN_ENV")
	fmt.Println("env:"+env)
	if  env != "" {
		configFile = "config_" + env + ".yaml"
	}
	v := viper.New()
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	global.GVA_VP = v
}
