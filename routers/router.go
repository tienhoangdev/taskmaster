package routers

import (
	"golang_project_base/controllers"
	"golang_project_base/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	habitService := &services.HabitService{}
	habitController := &controllers.HabitController{Service: habitService}

	r.POST("/habits", habitController.CreateHabit)

	return r
}
