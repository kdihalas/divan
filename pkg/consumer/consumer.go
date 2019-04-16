package consumer

type ConsumerProvider interface {
	Get(string) interface{}
	GetKeys(string, int) []string
	DeleteKey(string)
}