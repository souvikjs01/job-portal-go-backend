package main

import (
	"job_portal/packages/config"
	"job_portal/packages/models"
	"job_portal/packages/routes"
	"job_portal/packages/store"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	dbUrl := config.GetEnv("DATABASE_URL")
	port := config.GetEnv("PORT")

	store.ConnectDB(dbUrl)
	store.DB.AutoMigrate(
		&models.User{},
		&models.Job{},
		&models.ChangePasswordRequest{},
		&models.PasswordValidation{},
		&models.ForgotPasswordRequest{},
	)
	r := gin.Default()

	// allow all the origin :
	r.Use(cors.Default())
	routes.Routers(r)

	r.Run(":" + port)
}
