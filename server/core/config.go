package core

import (
	"fmt"
	"gin-vue-admin/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const defaultConfigFile = "config.dev.yaml"

func init() {
	v := viper.New()
	var configFile = defaultConfigFile
	mode := os.Args[1]
	if len(mode) > 0 {
		configFile = strings.Replace(defaultConfigFile, "dev", mode, 1)
	}
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s , %s \n", configFile, err))
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
