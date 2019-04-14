package main

import (
	"fmt"
	"github.com/kdihalas/divan/pkg/consumer"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var consumers = make(chan *consumer.ConsumerConfig)

func main() {
	time.Sleep(1 * time.Second)
	c := viper.Get("consumer").(map[string]interface{})
	s := viper.Get("sync").([]interface{})[0]

	controller := consumer.NewConsumerController(consumers, c)
	go controller.Start()

	for name, config := range s.(map[string]interface{}) {
		c := &consumer.ConsumerConfig{
			ID:     fmt.Sprintf("%d", time.Now().Unix()),
			Name:   name,
			Config: config,
		}
		consumers <- c
	}

	log.Info("Running divan..")

	select {}
}
