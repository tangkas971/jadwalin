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
	subjectRepo := repository.NewSubjectRepository(config.DB)
	taskRepo := repository.NewTaskRepository(config.DB)
	studentTaskRepo := repository.NewStudentTaskRepository(config.DB)
	authRepo := repository.NewAuthRepository(config.DB)

	userService := services.NewUserService(userRepo)
	scheduleService := services.NewScheduleService(scheduleRepo)
	blacklistTokenService := services.NewBlacklistTokenService(blacklistTokenRepo)
	subjectService := services.NewSubjectServices(subjectRepo)
	studentTaskService := services.NewStudentTaskService(studentTaskRepo)
	taskService := services.NewTaskService(taskRepo, userService, studentTaskService)
	authService := services.NewAuthService(authRepo, userRepo)
	
	authController := controller.NewAuthController(authService, blacklistTokenService)
	scheduleController := controller.NewScheduleController(scheduleService)
	userController := controller.NewUserController(userService)
	subjectController := controller.NewSubjectController(subjectService)
	taskController := controller.NewTaskController(taskService)
	studentTaskController := controller.NewStudentTaskController(studentTaskService)


	router := gin.Default()

	routes.AuthRoutes(router, authController, blacklistTokenService)
	routes.UserRoutes(router, userController, blacklistTokenService)
	routes.ScheduleRoutes(router, scheduleController, blacklistTokenService)
	routes.SubjectRoutes(router, subjectController, blacklistTokenService)
	routes.TaskRoutes(router, taskController, blacklistTokenService)
	routes.StudentTask(router, studentTaskController, blacklistTokenService)
	

	router.Run()
}