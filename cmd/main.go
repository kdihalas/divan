package main

import (
	"fmt"
	"github.com/kdihalas/divan/pkg/consumer"
	"github.com/kdihalas/divan/pkg/provider"
	"github.com/spf13/viper"
	"time"

	log "github.com/sirupsen/logrus"
)

var consumers = make(chan *consumer.ConsumerConfig)

var Consumers = make(map[string]consumer.ConsumerProvider)

var Backends = make(map[string]provider.Provider)

func main() {
	time.Sleep(1 * time.Second)
	c := viper.Get("consumers").([]interface{})
	b := viper.Get("backends").([]interface{})
	f := viper.Get("funnels").([]interface{})

	for _, conf := range c {
		var consumerConf = make(map[string]interface{})
		for key, val := range conf.(map[interface{}]interface{}) {
			consumerConf[key.(string)] = val
		}
		Consumers[consumerConf["name"].(string)] = GetConsumerProvider(consumerConf)
	}

	for _, conf := range b {
		var backendConf = make(map[string]interface{})
		for key, val := range conf.(map[interface{}]interface{}) {
			backendConf[key.(string)] = val
		}
		Backends[backendConf["name"].(string)] = GetProvider(backendConf)
	}

	controller := consumer.NewConsumerController(consumers, Consumers, Backends)
	go controller.Start()

	for _, funnel := range f {
		var funConfig = make(map[string]interface{})
		for key, val := range funnel.(map[interface{}]interface{}) {
			funConfig[key.(string)] = val
		}

		c := &consumer.ConsumerConfig{
			ID:     fmt.Sprintf("%d", time.Now().Unix()),
			Name:   funConfig["name"].(string),
			Config: funConfig,
		}

		consumers <- c
	}

	log.Info("Running divan..")

	select {}
}