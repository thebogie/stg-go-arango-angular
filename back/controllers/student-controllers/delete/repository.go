package deleteStudent

import (
	config "back/configs"
	model "back/models"
)

type Repository interface {
	DeleteStudentRepository(input *model.EntityStudent) (*model.EntityStudent, string)
}

type repository struct {
	db config.DBconnection
}

func NewRepositoryDelete(db config.DBconnection) *repository {
	return &repository{db: db}
}

func (r *repository) DeleteStudentRepository(input *model.EntityStudent) (*model.EntityStudent, string) {

	var students model.EntityStudent
	//db := r.db.Model(&students)
	errorCode := make(chan string, 1)

	/*
		checkStudentId := db.Debug().Select("*").Where("id = ?", input.ID).Find(&students)

		if checkStudentId.RowsAffected < 1 {
			errorCode <- "DELETE_STUDENT_NOT_FOUND_404"
			return &students, <-errorCode
		}

		deleteStudentId := db.Debug().Select("*").Where("id = ?", input.ID).Find(&students).Delete(&students)

		if deleteStudentId.Error != nil {
			errorCode <- "DELETE_STUDENT_FAILED_403"
			return &students, <-errorCode
		} else {
			errorCode <- "nil"
		}
	*/
	return &students, <-errorCode
}
