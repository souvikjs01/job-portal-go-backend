package handlers

import (
	"job_portal/packages/models"
	"job_portal/packages/services"
	"job_portal/packages/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	jobService services.JobService
}

func NewJobHandler(jobService services.JobService) *JobHandler {
	return &JobHandler{
		jobService: jobService,
	}
}

func (h *JobHandler) CreateJob(c *gin.Context) {
	user_role, existUserRole := c.Get("role")
	user_id, existsUserId := c.Get("userId")

	if !existUserRole || !existsUserId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "role not found",
		})
		return
	}

	role, ok1 := user_role.(models.Role)
	userId, ok2 := user_id.(string)

	if !ok1 || !ok2 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "invalid role type",
		})
		return
	}

	if role == models.RoleUser {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "access denied",
		})
		return
	}

	var req models.CreateJob
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	if err := validation.ValidateStruct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	newJob, err := h.jobService.CreateJob(&req, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newJob,
	})

}

func (h *JobHandler) FindJobByID(c *gin.Context) {
	id := c.Param("id")

	// Call service layer
	job, err := h.jobService.GetJobByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    job,
	})
}

func (h *JobHandler) FindAllJob(c *gin.Context) {
	// Call service layer
	jobs, err := h.jobService.GetAllJob()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    jobs,
	})
}
