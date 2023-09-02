package route

import (
	createStudent "gazo/controllers/student-controllers/create"
	deleteStudent "gazo/controllers/student-controllers/delete"
	resultStudent "gazo/controllers/student-controllers/result"
	resultsStudent "gazo/controllers/student-controllers/results"
	updateStudent "gazo/controllers/student-controllers/update"
	handlerCreateStudent "gazo/handlers/student-handlers/create"
	handlerDeleteStudent "gazo/handlers/student-handlers/delete"
	handlerResultStudent "gazo/handlers/student-handlers/result"
	handlerResultsStudent "gazo/handlers/student-handlers/results"
	handlerUpdateStudent "gazo/handlers/student-handlers/update"
	middleware "gazo/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitStudentRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Student
	*/
	createStudentRepository := createStudent.NewRepositoryCreate(db)
	createStudentService := createStudent.NewServiceCreate(createStudentRepository)
	createStudentHandler := handlerCreateStudent.NewHandlerCreateStudent(createStudentService)

	resultsStudentRepository := resultsStudent.NewRepositoryResults(db)
	resultsStudentService := resultsStudent.NewServiceResults(resultsStudentRepository)
	resultsStudentHandler := handlerResultsStudent.NewHandlerResultsStudent(resultsStudentService)

	resultStudentRepository := resultStudent.NewRepositoryResult(db)
	resultStudentService := resultStudent.NewServiceResult(resultStudentRepository)
	resultStudentHandler := handlerResultStudent.NewHandlerResultStudent(resultStudentService)

	deleteStudentRepository := deleteStudent.NewRepositoryDelete(db)
	deleteStudentService := deleteStudent.NewServiceDelete(deleteStudentRepository)
	deleteStudentHandler := handlerDeleteStudent.NewHandlerDeleteStudent(deleteStudentService)

	updateStudentRepository := updateStudent.NewRepositoryUpdate(db)
	updateStudentService := updateStudent.NewServiceUpdate(updateStudentRepository)
	updateStudentHandler := handlerUpdateStudent.NewHandlerUpdateStudent(updateStudentService)

	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/student", createStudentHandler.CreateStudentHandler)
	groupRoute.GET("/student", resultsStudentHandler.ResultsStudentHandler)
	groupRoute.GET("/student/:id", resultStudentHandler.ResultStudentHandler)
	groupRoute.DELETE("/student/:id", deleteStudentHandler.DeleteStudentHandler)
	groupRoute.PUT("/student/:id", updateStudentHandler.UpdateStudentHandler)
}
