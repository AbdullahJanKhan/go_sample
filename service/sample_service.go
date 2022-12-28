package service

import (
	"github.com/abdullahjankhan/go_sample/models"
	"github.com/abdullahjankhan/go_sample/repository"
)

type SampleService interface {
	SampleFunc(*models.SampleRequest) error
}

type sampleService struct {
	store repository.Store
}

func NewSampleService(store repository.Store) SampleService {
	return &sampleService{
		store: store,
	}
}

func (ss *sampleService) SampleFunc(req *models.SampleRequest) error {
	return ss.store.SampleCreate(req)
}
