package controllers

// import (
// 	"fmt"
// 	"job_portal/packages/models"
// 	"job_portal/packages/store"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func CreateJob(c *gin.Context) {
// 	userIdVal, exists := c.Get("user_id")
// 	isAdminVal, adminExists := c.Get("is_admin")
// 	fmt.Println(userIdVal)
// 	if !exists || !adminExists {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Unauthorized",
// 		})
// 		return
// 	}

// 	isAdmin, ok := isAdminVal.(bool)
// 	if !ok || !isAdmin {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"error": "Only admin can create jobs",
// 		})
// 		return
// 	}

// 	userId, ok := userIdVal.(float64)
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID in token"})
// 		return
// 	}

// 	var job models.Job
// 	if err := c.ShouldBindJSON(&job); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job data"})
// 		return
// 	}
// 	job.UserID = int(userId)
// 	job.CreatedAt = time.Now()

// 	if err := store.DB.Create(&job).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create job"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "Job created successfully",
// 		"job":     job,
// 	})

// }

// /*
// 	"title": "frontend",
//   	"description": "handle ui/ux part of a team",
//   	"location": "bangalore",
//   	"company": "hp",
//   	"min_salary": 16,
//   	"experience_level": "0-2 years",
// 	"skills": "html, css, js, react",
// 	"max_salary": 20,
// 	"type": "engineering"
// */

// // update job:
// func UpdateJob(c *gin.Context) {
// 	isAdminVal, adminExists := c.Get("is_admin")

// 	if !adminExists || !isAdminVal.(bool) {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"error": "Only admin can update jobs",
// 		})
// 		return
// 	}

// 	jobIdStr := c.Param("id")
// 	jobId, err := strconv.Atoi(jobIdStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
// 		return
// 	}

// 	// find the job
// 	var job models.Job
// 	if err := store.DB.First(&job, jobId).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
// 		return
// 	}

// 	// bind the request body
// 	var updateData map[string]interface{}
// 	if err := c.ShouldBindJSON(&updateData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid request body",
// 		})
// 	}

// 	if err := store.DB.Model(&job).Updates(updateData).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job " + err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Job updated successfully",
// 		"job":     job,
// 	})
// }

// // get job by its id:
// func GetJobById(c *gin.Context) {
// 	jobIdStr := c.Param("id")
// 	jobId, err := strconv.Atoi(jobIdStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid job ID",
// 		})
// 		return
// 	}

// 	var job models.Job
// 	if err := store.DB.First(&job, jobId).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "job not found",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, job)
// }

// // get job by its id and delete
// func DeleteJobById(c *gin.Context) {
// 	jobIdStr := c.Param("id")
// 	jobId, err := strconv.Atoi(jobIdStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid job ID",
// 		})
// 		return
// 	}

// 	// check if the user is admin or not
// 	isAdminVal, exist := c.Get("is_admin")
// 	if !exist {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Unauthorized",
// 		})
// 		return
// 	}

// 	isAdmin, ok := isAdminVal.(bool)
// 	if !ok || !isAdmin {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"error": "Only admin can delete jobs",
// 		})
// 		return
// 	}
// 	var job models.Job
// 	result := store.DB.Delete(&job, jobId)
// 	if result.RowsAffected == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
// 		return
// 	}

// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete job"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
// }

// // fetch all the jobs
// func GetAllJobs(c *gin.Context) {
// 	var jobs []models.Job

// 	result := store.DB.Find(&jobs)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "failed to fetch jobs",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"jobs": jobs,
// 	})
// }
