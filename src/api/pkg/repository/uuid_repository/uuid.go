package uuid_repository

import (
	google_uuid "github.com/google/uuid"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_staff_usecase"
)

var _ live_house_staff_usecase.UuidRepository = (*uuid)(nil)

type uuid struct{}

func NewUuid() uuid {
	return uuid{}
}

func (uuid) Generate() (string, error) {
	uuid, err := google_uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
