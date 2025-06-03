package main

import (
	"jadwalin/config"
	"jadwalin/controller"
	"jadwalin/repository"
	"jadwalin/routes"
	"jadwalin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	userRepo := repository.NewUserRepository(config.DB)
	scheduleRepo := repository.NewScheduleRepository(config.DB)
	blacklistTokenRepo := repository.NewBlacklistTokenRepository(config.DB)
	taskRepo := repository.NewTaskRepository(config.DB)
	studentTaskRepo := repository.NewStudentTaskRepository(config.DB)
	authRepo := repository.NewAuthRepository(config.DB)
	prodiRepo := repository.NewProdiRepository(config.DB)
	gradeRepo := repository.NewGradeRepository(config.DB)
	subjectRepo := repository.NewSubjectRepository(config.DB)

	userService := services.NewUserService(userRepo)
	scheduleService := services.NewScheduleService(scheduleRepo)
	blacklistTokenService := services.NewBlacklistTokenService(blacklistTokenRepo)
	studentTaskService := services.NewStudentTaskService(studentTaskRepo)
	taskService := services.NewTaskService(taskRepo, userService, studentTaskService)
	authService := services.NewAuthService(authRepo, userRepo)
	prodiService := services.NewProdiService(prodiRepo)
	gradeService := services.NewGradeService(gradeRepo)
	subjectService := services.NewSubjectService(subjectRepo)

	authController := controller.NewAuthController(authService, blacklistTokenService)
	scheduleController := controller.NewScheduleController(scheduleService)
	userController := controller.NewUserController(userService)
	taskController := controller.NewTaskController(taskService)
	studentTaskController := controller.NewStudentTaskController(studentTaskService)
	prodiController := controller.NewProdiController(prodiService)
	gradeController := controller.NewGradeController(gradeService)
	subjectController := controller.NewSubjectController(subjectService)


	router := gin.Default()

	routes.AuthRoutes(router, authController, blacklistTokenService)
	routes.UserRoutes(router, userController, blacklistTokenService)
	routes.ScheduleRoutes(router, scheduleController, blacklistTokenService)	
	routes.TaskRoutes(router, taskController, blacklistTokenService)
	routes.StudentTask(router, studentTaskController, blacklistTokenService)
	routes.ProdiRoute(router, prodiController, blacklistTokenService)
	routes.GradeRoute(router, gradeController, blacklistTokenService)
	routes.SubjectRoute(router, subjectController, blacklistTokenService)
	

	router.Run()
}