package routes

import (
	"job_portal/packages/controllers"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	router := r.Group("/api/v1")

	router.POST("/user/register", controllers.RegisterUser)
	router.POST("/user/login", controllers.LogInUser)
}
