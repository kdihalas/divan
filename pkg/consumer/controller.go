package consumer

import (
	"strings"
	"time"

	"github.com/kdihalas/divan/pkg/provider"
	log "github.com/sirupsen/logrus"
)

type Consumer struct {
	consumers chan *ConsumerConfig
	config map[string]interface{}
	client ConsumerProvider
}

type ConsumerConfig struct {
	ID string
	Name string
	Config interface{}
}

func (c *Consumer) Start() {
	log.Info("Starting controller")
	select {
	case consumer := <- c.consumers:
		c.spawn(consumer)
	}

}

func (c *Consumer) spawn(consumer *ConsumerConfig) {
	log.WithField("ID", consumer.ID).Info("New consumer started for datbase/table/prefix:", consumer.Name)
	conf := consumer.Config.(map[string]interface{})

	interval := conf["interval"].(int64)
	buffer := conf["buffer"].(int64)
	backend := c.getProvider(conf)
	for {
		keys := c.client.GetKeys(consumer.Name, buffer)
		for _, key := range keys {
			id := strings.Split(key, "/")[1]
			doc := c.client.Get(key)
			err := backend.Update(id, doc)
			if err == nil {
				c.client.DeleteKey(key)
			}
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func (c *Consumer) getConsumerProvider(conf map[string]interface{}) ConsumerProvider {
	if conf["provider"] == "redis" {
		pro := NewRedisConsumerProvider(conf)
		return pro
	}
	return nil
}

func (c *Consumer) SetConsumerProvider() {
	c.client = c.getConsumerProvider(c.config)
}

func (c *Consumer) getProvider(conf map[string]interface{}) provider.Provider {
	if conf["provider"] == "couchdb" {
		pro := provider.NewCouchDbProvider(conf)
		return pro
	}

	return nil
}

func NewConsumerController(consumers chan *ConsumerConfig, config map[string]interface{}) *Consumer {
	con := &Consumer{
		consumers: consumers,
		config: config,
	}
	con.SetConsumerProvider()
	return con
}