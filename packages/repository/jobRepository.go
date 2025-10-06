package repository

import (
	"fmt"
	"job_portal/packages/models"
	"job_portal/packages/store"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobRepository interface {
	Create(job *models.CreateJob, userId string) (*models.Job, error)
	GetJobByID(id string) (*models.Job, error)
	GetAllJob() ([]models.Job, error)
	Update(id string, job *models.UpdateJob) (*models.Job, error)
	Delete(id string) error
	SearchJobs(query string) ([]models.Job, error)
}

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *store.DB) JobRepository {
	return &jobRepository{
		db: db.DB,
	}
}

func (r *jobRepository) Create(job *models.CreateJob, userId string) (*models.Job, error) {
	newJob := models.Job{
		Id:              uuid.New().String(),
		Title:           job.Title,
		Description:     job.Description,
		Location:        job.Location,
		Company:         job.Company,
		MinSalary:       job.MinSalary,
		MaxSalary:       job.MaxSalary,
		ExperienceLevel: job.ExperienceLevel,
		Skills:          job.Skills,
		Type:            job.Type,
		ApplyLink:       job.ApplyLink,
		UserID:          userId,
	}

	if err := r.db.Create(&newJob).Error; err != nil {
		return nil, fmt.Errorf("failed to create job %s", err)
	}

	return &newJob, nil
}

func (r *jobRepository) GetJobByID(id string) (*models.Job, error) {
	var job models.Job
	if err := r.db.Where("id = ?", id).First(&job).Error; err != nil {
		return nil, fmt.Errorf("not found")
	}

	return &job, nil
}

func (r *jobRepository) GetAllJob() ([]models.Job, error) {
	var jobs []models.Job
	if err := r.db.Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("not found")
	}

	return jobs, nil
}

func (r *jobRepository) Update(id string, job *models.UpdateJob) (*models.Job, error) {
	var existJob models.Job

	if err := r.db.First(&existJob, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("job not found: %w", err)
	}

	if err := r.db.Model(&existJob).Updates(job).Error; err != nil {
		return nil, fmt.Errorf("failed updating job: %w", err)
	}

	return &existJob, nil
}

func (r *jobRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.Job{}).Error
}

func (r *jobRepository) SearchJobs(query string) ([]models.Job, error) {
	var jobs []models.Job

	err := r.db.Where(`
		title ILIKE ? OR 
		description ILIKE ? OR 
		company ILIKE ? OR 
		skills ILIKE ? OR 
		location ILIKE ?`,
		"%"+query+"%",
		"%"+query+"%",
		"%"+query+"%",
		"%"+query+"%",
		"%"+query+"%",
	).Find(&jobs).Error

	if err != nil {
		return nil, err
	}

	return jobs, nil
}
