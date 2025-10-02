package services

import (
	"job_portal/packages/models"
	"job_portal/packages/repository"
)

type JobService interface {
	CreateJob(req *models.CreateJob, userId string) (*models.Job, error)
}

type jobService struct {
	jobRepo repository.JobRepository
}

func NewJobService(jobRepo repository.JobRepository) JobService {
	return &jobService{
		jobRepo: jobRepo,
	}
}

func (s *jobService) CreateJob(req *models.CreateJob, userId string) (*models.Job, error) {

}
