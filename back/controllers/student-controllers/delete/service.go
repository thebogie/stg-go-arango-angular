package deleteStudent

import (
	model "back/models"
)

type Service interface {
	DeleteStudentService(input *InputDeleteStudent) (*model.EntityStudent, string)
}

type service struct {
	repository Repository
}

func NewServiceDelete(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) DeleteStudentService(input *InputDeleteStudent) (*model.EntityStudent, string) {

	students := model.EntityStudent{
		ID: input.ID,
	}

	resultCreateStudent, errCreateStudent := s.repository.DeleteStudentRepository(&students)

	return resultCreateStudent, errCreateStudent
}
