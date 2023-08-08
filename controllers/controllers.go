package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seew0/jobscheduler/models"
	"github.com/seew0/jobscheduler/services"
)

func CreateJobHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var req struct {
		Duration int `json:"duration"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	if req.Duration <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Duration: duration should always be positive"})
		return
	}

	jobID := services.CreateJob(req.Duration)
	c.JSON(http.StatusCreated, gin.H{"jobId": jobID})
}

func StartJobHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	jobIDStr := c.Param("id")

	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	job, exists := services.GetJobByID(jobID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found!"})
		return
	}

	if job.Status == models.Idle {
		services.StartJob(job)
		c.JSON(http.StatusOK, gin.H{"message": "Job has started"})
		return
	}
	if job.Status == models.Done {
		c.JSON(http.StatusOK, gin.H{"message": "Job is already done"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Job already running"})
	return
}

func GetJobStatusHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	jobIDStr := c.Param("id")
	var status string
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	job, exists := services.GetJobByID(jobID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	if job.Status == models.Idle {
		status = "idle"
	}
	if job.Status == models.Running {
		currentTime := time.Now().Unix()
		status = services.TrackJob(job, currentTime)
	}
	if job.Status == models.Done {
		status = "done"
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}
