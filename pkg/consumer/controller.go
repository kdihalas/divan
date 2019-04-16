package consumer

import (
	"strings"
	"time"

	"github.com/kdihalas/divan/pkg/provider"
	log "github.com/sirupsen/logrus"
)

type Consumer struct {
	consumer chan *ConsumerConfig
	consumers map[string]ConsumerProvider
	backends map[string]provider.Provider
}

type ConsumerConfig struct {
	ID string
	Name string
	Config interface{}
}

func (c *Consumer) Start() {
	log.Info("Starting controller")
	for consumerConf := range c.consumer {
		go c.spawn(consumerConf)
	}

}

func (c *Consumer) spawn(consumerConf *ConsumerConfig) {
	log.WithField("ID", consumerConf.ID).Info("New consumer started for datbase/table/prefix:", consumerConf.Name)
	conf := consumerConf.Config.(map[string]interface{})

	interval := conf["interval"].(int)
	buffer := conf["buffer"].(int)

	consumerClient := *c.getConsumerProvider(conf["from"].(string))
	backendClient := *c.getBackendProvider(conf["to"].(string))

	for {
		keys := consumerClient.GetKeys(conf["name"].(string), buffer)
		for _, key := range keys {
			id := strings.Split(key, "/")[1]
			doc := consumerClient.Get(key)
			err := backendClient.Update(id, doc)
			if err == nil {
				consumerClient.DeleteKey(key)
			}
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func (c *Consumer) getConsumerProvider(name string) *ConsumerProvider{
	for consumerName, consumerProvider := range c.consumers {
		if name == consumerName {
			return &consumerProvider
		}
	}
	return nil
}

func (c *Consumer) getBackendProvider(name string) *provider.Provider{
	for backendName, backendProvider := range c.backends {
		if name == backendName {
			return &backendProvider
		}
	}
	return nil
}

func NewConsumerController(consumer chan *ConsumerConfig, consumers map[string]ConsumerProvider, backends map[string]provider.Provider) *Consumer {
	con := &Consumer{
		consumer: consumer,
		consumers: consumers,
		backends: backends,
	}
	return con
}