package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

type Redis struct {
	config map[string]interface{}
	client *redis.Client
}


func (r *Redis) connect() {
	r.client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", r.config["server"].(string), r.config["port"]),
		Password: viper.GetString(r.config["password"].(string)),
		DB: r.config["db"].(int),
	})
}

func (r *Redis) Get(key string) interface{} {
	var doc interface{}
	val, _ := r.client.Get(key).Result()
	err := json.Unmarshal([]byte(val), &doc)
	if err != nil {
		log.Error(err)
	}
	return doc
}

func (r *Redis) GetKeys(prefix string, buffer int) []string{
	var cursor uint64
	var n int
	var keys []string
	for {
		var err error
		keys, cursor, err = r.client.Scan(cursor, fmt.Sprintf("%s*", prefix), int64(buffer)).Result()
		if err != nil {
			log.Error(err)
		}
		n += len(keys)
		if cursor == 0 {
			break
		}
	}

	log.Infof("found %d keys", n)
	return keys
}

func (r *Redis) DeleteKey(key string) {
	r.client.Del(key)
}

func NewRedisConsumerProvider(config map[string]interface{}) *Redis{
	re := &Redis{
		config: config,
	}
	re.connect()
	return re
}