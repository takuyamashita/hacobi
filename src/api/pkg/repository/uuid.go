package repository

import (
	google_uuid "github.com/google/uuid"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_staff_usecase"
)

var _ live_house_staff_usecase.UuidRepositoryIntf = (*uuidRepository)(nil)

type uuidRepository struct{}

func NewUuidRepository() live_house_staff_usecase.UuidRepositoryIntf {
	return &uuidRepository{}
}

func (uuidRepository) Generate() (string, error) {
	uuid, err := google_uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
