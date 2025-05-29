package routes

import (
	"job_portal/packages/controllers"
	"job_portal/packages/middleware"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	router := r.Group("/api/v1")

	// auth routes:
	router.POST("/user/register", controllers.RegisterUser)
	router.POST("/user/login", controllers.LogInUser)

	// job routes:
	router.POST("/job/new", middleware.Authenticated, controllers.CreateJob)

}
