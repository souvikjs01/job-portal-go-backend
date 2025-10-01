package routes

import (
	"job_portal/packages/auth"
	"job_portal/packages/config"
	"job_portal/packages/handlers"
	"job_portal/packages/repository"
	"job_portal/packages/services"
	"job_portal/packages/store"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *store.DB, cfg *config.Config) {
	// CORS Setup
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	conf.AllowCredentials = true
	conf.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(conf))

	// Initialize repository layer
	userRepo := repository.NewUserRepository(db)

	// Initialize auth service
	jwtService := auth.NewJWTService(&cfg.Auth)

	// Initialize service layer
	userService := services.NewUserService(userRepo, jwtService)

	// Initialize handler layer
	userHandler := handlers.NewUserHandler(userService)

	// Public Routes
	publicAuthRoute := router.Group("/api/v1/auth")
	{
		publicAuthRoute.POST("/signup", userHandler.Register)
		publicAuthRoute.POST("/login", userHandler.Login)
	}
	privateRoute := router.Group("/api/v1")
	{
		privateRoute.GET("/user/:id", jwtService.AuthMiddleware(), userHandler.UserProfile)
		privateRoute.GET("/users", jwtService.AuthMiddleware(), userHandler.GetAllUsers)
		privateRoute.PUT("/update/:id", jwtService.AuthMiddleware(), userHandler.UpdateUser)
		privateRoute.PUT("/update_role/:user_id", jwtService.AuthMiddleware(), userHandler.UpdateRole)
	}
}
