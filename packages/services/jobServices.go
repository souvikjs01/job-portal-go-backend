package services

import (
	"job_portal/packages/models"
	"job_portal/packages/repository"
)

type JobService interface {
	CreateJob(req *models.CreateJob, userId string) (*models.Job, error)
	GetJobByID(id string) (*models.Job, error)
	GetAllJob() ([]models.Job, error)
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
	job, err := s.jobRepo.Create(req, userId)

	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s *jobService) GetJobByID(id string) (*models.Job, error) {
	job, err := s.jobRepo.GetJobByID(id)

	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s *jobService) GetAllJob() ([]models.Job, error) {
	job, err := s.jobRepo.GetAllJob()

	if err != nil {
		return nil, err
	}

	return job, nil
}
