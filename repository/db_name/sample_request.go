package db_name

import (
	"github.com/abdullahjankhan/go_sample/models"
)

func (s *Store) SampleCreate(req *models.SampleRequest) error {
	err := s.db.Create(req).Error
	if err != nil {
		return &models.StandardError{
			Code:        models.INTERNAL_SERVER_ERROR,
			ActualError: err,
			Line:        "Store/sample_request: SampleCreate():10",
			Message:     models.DB_OPERATION_ERROR_MESSAGE,
		}
	}
	return nil
}
