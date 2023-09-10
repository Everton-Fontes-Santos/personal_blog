package services

type ServiceInterface[T any] interface {
	Execute(opts ...any) (T, error)
}
