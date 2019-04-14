package main

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")       // name of config file (without extension)
	viper.AddConfigPath("/etc/divan/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.divan") // call multiple times to add many search paths
	viper.AddConfigPath(".")            // optionally look for config in the working directory
	err := viper.ReadInConfig()         // Find and read the config file
	if err != nil {                     // Handle errors reading the config file
		log.Errorf("Fatal error config file: %s \n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Warnf("Config file changed:", e.Name)
	})
}
