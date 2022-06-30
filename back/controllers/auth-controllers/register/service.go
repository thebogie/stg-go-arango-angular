package registerAuth

import (
	model "back/models"
)

type Service interface {
	RegisterService(input *InputRegister) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input *InputRegister) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}

	resultRegister, errRegister := s.repository.RegisterRepository(&users)

	return resultRegister, errRegister
}
