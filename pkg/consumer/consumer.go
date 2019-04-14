package consumer

type ConsumerProvider interface {
	Get(string) interface{}
	GetKeys(string, int64) []string
	DeleteKey(string)
}