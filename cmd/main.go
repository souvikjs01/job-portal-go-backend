package main

import (
	"job_portal/packages/config"
	"job_portal/packages/routes"
	"job_portal/packages/store"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	// connect database
	db, err := store.ConnectDB(cfg.Database.DB_URL)
	if err != nil {
		log.Fatalf("Failed to database connection %s", err)
	}

	defer db.Close()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Response": "Everything is ok",
		})
	})
	// allow all the origin :
	r.Use(cors.Default())

	routes.SetupRoutes(r, db, cfg)

	r.Run(":" + cfg.Server.Port)
}
