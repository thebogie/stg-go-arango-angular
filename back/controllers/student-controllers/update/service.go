package updateStudent

import (
	model "back/models"
)

type Service interface {
	UpdateStudentService(input *InputUpdateStudent) (*model.EntityStudent, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateStudentService(input *InputUpdateStudent) (*model.EntityStudent, string) {

	students := model.EntityStudent{
		ID:   input.ID,
		Name: input.Name,
		Npm:  input.Npm,
		Fak:  input.Fak,
		Bid:  input.Bid,
	}

	resultUpdateStudent, errUpdateStudent := s.repository.UpdateStudentRepository(&students)

	return resultUpdateStudent, errUpdateStudent
}
