package controllers

import (
	"fmt"
	"job_portal/packages/models"
	"job_portal/packages/store"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateJob(c *gin.Context) {
	userIdVal, exists := c.Get("user_id")
	isAdminVal, adminExists := c.Get("is_admin")
	fmt.Println(userIdVal)
	if !exists || !adminExists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	isAdmin, ok := isAdminVal.(bool)
	if !ok || !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Only admin can create jobs",
		})
		return
	}

	userId, ok := userIdVal.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID in token"})
		return
	}

	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job data"})
		return
	}
	job.UserID = int(userId)
	job.CreatedAt = time.Now()

	if err := store.DB.Create(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create job"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Job created successfully",
		"job":     job,
	})

}

/*
	"title": "frontend",
  	"description": "handle ui/ux part of a team",
  	"location": "bangalore",
  	"company": "hp",
  	"min_salary": 16,
  	"experience_level": "0-2 years",
	"skills": "html, css, js, react",
	"max_salary": 20,
	"type": "engineering"
*/
