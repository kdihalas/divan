package provider

type Provider interface {
	Update(string, interface{}) error
}