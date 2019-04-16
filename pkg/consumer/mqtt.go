package consumer

import (
	mqttClient "github.com/yosssi/gmq/mqtt/client"
)

type MQTT struct {
	config map[string]interface{}
	client *mqttClient.Client
}