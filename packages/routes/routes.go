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

	// user management:
	router.GET("/user/all", controllers.GetAllUsers)
	router.GET("/user/:id", controllers.GetUserById)
	router.DELETE("/user/remove/:id", middleware.Authenticated, controllers.RemoveUserById)

	// job routes:
	router.POST("/job/new", middleware.Authenticated, controllers.CreateJob)
	router.PUT("/job/update/:id", middleware.Authenticated, controllers.UpdateJob)
	router.GET("/job/:id", controllers.GetJobById)
	router.DELETE("/job/remove/:id", middleware.Authenticated, controllers.DeleteJobById)
	router.GET("/job/all", controllers.GetAllJobs)
}
