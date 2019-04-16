package main

import "github.com/kdihalas/divan/pkg/consumer"

func GetConsumerProvider(conf map[string]interface{}) consumer.ConsumerProvider {
	if conf["provider"] == "redis" {
		pro := consumer.NewRedisConsumerProvider(conf)
		return pro
	}
	if conf["provider"] == "mqtt" {
		pro := consumer.NewMqttConsumerProvider(conf)
		return pro
	}
	return nil
}