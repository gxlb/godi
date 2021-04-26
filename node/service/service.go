package service

type IService interface {
	Name() string
}

type Service struct {
	entity IService // entity of this service
	name   string   // name of this serice
}
