package resultStudent

import (
	model "back/models"
)

type Service interface {
	ResultStudentService(input *InputResultStudent) (*model.EntityStudent, string)
}

type service struct {
	repository Repository
}

func NewServiceResult(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResultStudentService(input *InputResultStudent) (*model.EntityStudent, string) {

	students := model.EntityStudent{
		ID: input.ID,
	}

	resultCreateStudent, errCreateStudent := s.repository.ResultStudentRepository(&students)

	return resultCreateStudent, errCreateStudent
}
