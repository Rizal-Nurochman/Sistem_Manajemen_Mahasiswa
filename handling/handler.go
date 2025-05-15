package handling

import (
	"belajar-api/students"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	studentService Service
}

func NewStudentHandler(studentService Service) *StudentHandler {
	return &StudentHandler{studentService}
}

func (handler *StudentHandler) GetStudentsHandler(c *gin.Context) {
	students,err:=handler.studentService.GetAllStudent()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":students,
	})
}

func (handler *StudentHandler) GetStudentsByNIMHandler(c *gin.Context)  {
	NIM:=c.Param("NIM")
	student, err:=handler.studentService.GetStudentsByNIM(NIM)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":student,
	})
}

func (handler *StudentHandler) PostStudentsHandler(c *gin.Context)  {
	var StudentsRequest students.StudentsRequest

	err:=c.ShouldBindJSON(&StudentsRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	newStudent, err:=handler.studentService.CreateStudent(StudentsRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":newStudent,
	})
}

func (handler *StudentHandler) PutStudentsByNIMHandler(c *gin.Context)  {
	NIM:=c.Param("NIM")
	var studentReq students.StudentsRequest
	err:=c.ShouldBindJSON(&studentReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	student, err:=handler.studentService.PutStudentsByNIM(NIM, studentReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":student,
		"message":"Update student completed",
	})
}

func (handler *StudentHandler) DelStudentsByNIMHandler(c *gin.Context)  {
	NIM:=c.Param("NIM")
	_, err:=handler.studentService.DelStudentsByNIM(NIM)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"Delete student completed!",
	})
}