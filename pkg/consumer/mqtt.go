package consumer

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/yosssi/gmq/mqtt"
	mqttClient "github.com/yosssi/gmq/mqtt/client"

)

type MQTT struct {
	config map[string]interface{}
	client *mqttClient.Client
	messages map[string]interface{}
}

func (m *MQTT) connect() {
	m.messages = make(map[string]interface{}, m.config["buffer"].(int))
	m.client = mqttClient.New(&mqttClient.Options{
		ErrorHandler: func(e error) {
			log.Error(e)
		},
	})

	err := m.client.Connect(&mqttClient.ConnectOptions{
		Network:  m.config["network"].(string),
		Address:  fmt.Sprintf("%s:%d", m.config["server"].(string), m.config["port"].(int)),
		ClientID: []byte(m.config["clientid"].(string)),
		UserName: []byte(m.config["username"].(string)),
		Password: []byte(m.config["password"].(string)),
	})

	if err != nil {
		log.Error(err)
	}

	err = m.client.Subscribe(&mqttClient.SubscribeOptions{
		SubReqs: []*mqttClient.SubReq{
			&mqttClient.SubReq{
				TopicFilter: []byte(m.config["topic"].(string)),
				QoS:         mqtt.QoS0,
				Handler: func(topicName, message []byte) {
					var value interface{}

					newId, _ := uuid.NewUUID()
					id := md5.Sum([]byte(newId.String()))
					prefix := strings.Trim(m.config["topic"].(string), "/")
					err = json.Unmarshal(message, &value)

					if err != nil {
						log.Error(err)
					} else {
						m.messages[fmt.Sprintf("%s/%x", prefix,id)] = value
					}
				},
			},
		},
	})
	if err != nil {
		log.Error(err)
	}
}

func (m *MQTT) Get(key string) interface{} {
	return m.messages[key]
}

func (m *MQTT) GetKeys(prefix string, buffer int) []string{
	var keys []string
	if len(m.messages) <= m.config["buffer"].(int) {
		log.Debug("messages less than buffer waiting for next run")
		return keys
	}
	for key, _ := range m.messages {
		keys = append(keys, key)
	}
	log.Infof("found %d keys", len(m.messages))
	return keys
}

func (m *MQTT) DeleteKey(key string) {
	delete(m.messages, key)
}


func NewMqttConsumerProvider(config map[string]interface{}) *MQTT{
	mq := &MQTT{
		config: config,
	}

	go mq.connect()
	return mq
}