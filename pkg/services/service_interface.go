package services

type ServiceInterface interface {
	Execute(params ...any) (any, error)
}
