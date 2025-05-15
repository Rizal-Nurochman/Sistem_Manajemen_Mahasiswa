package database

import (
	"belajar-api/students"

	"gorm.io/gorm"
)

type Repository interface {
	CreateStudent(student students.Students) (students.Students, error)
	GetAllStudent() ([]students.Students, error)
	GetStudentsByNIM(NIM string) (students.Students, error)
	DelStudentsByNIM(NIM string) (students.Students, error)
	PutStudentsByNIM(NIM string, student students.Students) (students.Students, error)
}

type repository struct {
	DB *gorm.DB
}

func (r *repository) DelStudentsByNIM(NIM string) (students.Students, error) {
	var student students.Students
	err:=r.DB.Where("nim = ?", NIM).Delete(&student).Error
	return student, err
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{DB}
}

func (r *repository) GetAllStudent() ([]students.Students, error) {
	var students []students.Students
	err := r.DB.Find(&students).Error
	return students, err
}

func (r *repository) GetStudentsByNIM(NIM string) (students.Students, error) {
	var student students.Students
	err := r.DB.Where("nim = ?", NIM).First(&student).Error
	return student, err
}

func (r *repository) CreateStudent(student students.Students) (students.Students, error) {
	err := r.DB.Create(&student).Error
	return student, err
}

func (r *repository) PutStudentsByNIM(NIM string, student students.Students) (students.Students, error) {
	var isStudentExist students.Students
	err:=r.DB.Where("nim = ?", NIM).First(&isStudentExist).Error
	if err != nil {
		return isStudentExist, err
	}

	isStudentExist.Name = student.Name
	isStudentExist.NIM = student.NIM
	isStudentExist.Major = student.Major
	isStudentExist.GPA = student.GPA
	isStudentExist.Graduated = student.Graduated

	err = r.DB.Save(&isStudentExist).Error
	return isStudentExist, err
}