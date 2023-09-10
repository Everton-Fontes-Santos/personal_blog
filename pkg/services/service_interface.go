package services

type ServiceInterface interface {
	Execute(opts ...any) (any, error)
}
