package repository

import "github.com/abdullahjankhan/go_sample/models"

type Store interface {
	BeginTx() (Store, error)
	Rollback() error
	CommitTx() error
	// mention your database layer function
	SampleCreate(*models.SampleRequest) error
}
