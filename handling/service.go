package handling

import (
	"belajar-api/database"
	"belajar-api/students"
)

type Service interface {
	CreateStudent(student students.StudentsRequest) (students.Students, error)
	GetAllStudent() ([]students.Students, error)
	GetStudentsByNIM(NIM string) (students.Students, error)
	DelStudentsByNIM(NIM string) (students.Students, error)
	PutStudentsByNIM(NIM string, student students.StudentsRequest) (students.Students, error)
}

type service struct {
	repository database.Repository
}

func (s *service) DelStudentsByNIM(NIM string) (students.Students, error) {
	student, err:=s.repository.DelStudentsByNIM(NIM)
	return student, err
}

func NewService(repository database.Repository) *service {
	return &service{repository}
}

func (s *service) GetAllStudent() ([]students.Students, error) {
	students, err := s.repository.GetAllStudent()
	return students, err
}

func (s *service) GetStudentsByNIM(NIM string) (students.Students, error) {
	student, err := s.repository.GetStudentsByNIM(NIM)
	return student, err
}

func (s *service) CreateStudent(studentRequest students.StudentsRequest) (students.Students, error) {
	student := students.Students{
		Name:      studentRequest.Name,
		NIM:       studentRequest.NIM,
		Major:     studentRequest.Major,
		GPA:       studentRequest.GPA,
		Graduated: studentRequest.Graduated,
	}
	newStudent, err := s.repository.CreateStudent(student)
	return newStudent, err
}

func (s *service)  PutStudentsByNIM(NIM string, studentRequest students.StudentsRequest) (students.Students, error){
	student:=students.Students{
		Name: studentRequest.Name,
		NIM: studentRequest.NIM,
		Major: studentRequest.Major,
		GPA: studentRequest.GPA,
		Graduated: studentRequest.Graduated,
	}
	updateStudent, err:=s.repository.PutStudentsByNIM(NIM, student)
	return updateStudent, err
}