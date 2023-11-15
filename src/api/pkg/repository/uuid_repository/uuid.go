package uuid_repository

import google_uuid "github.com/google/uuid"

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
