package services

import (
	"sync"
	"time"

	"github.com/seew0/jobscheduler/models"
)

var (
	jobs           = make(map[int]*models.Job)
	jobIDLock      sync.Mutex
	jobStarterLock sync.Mutex
)

func CreateJob(duration int) int {
	jobIDLock.Lock()
	defer jobIDLock.Unlock()

	jobID := len(jobs) + 1 
	job := &models.Job{
		Duration:  duration,
		Status:    models.Idle,
		StartTime: -1,
	}
	jobs[jobID] = job
	return jobID
}

func GetJobByID(jobID int) (*models.Job, bool) {
	job, exists := jobs[jobID]
	return job, exists
}

func StartJob(job *models.Job) {
	jobStarterLock.Lock()
	defer jobStarterLock.Unlock()

	job.StartTime = time.Now().Unix()
	job.Status = models.Running
}

func TrackJob(job *models.Job, currentTime int64) string {
	var status string
	if job.StartTime+int64(job.Duration) <= currentTime {
		job.Status = models.Done
		status = "done"
	} else {
		status = "running"
	}
	return status
}
