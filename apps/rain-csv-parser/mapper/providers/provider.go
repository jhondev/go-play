package providers

type Provider[T any] interface {
	GetData() (*T, error)
}
