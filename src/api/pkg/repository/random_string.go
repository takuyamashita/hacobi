package repository

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_email_authorization_domain"
)

var _ live_house_staff_email_authorization_domain.RandomStringRepositoryIntf = (*randomStringRepositoryImpl)(nil)

type RandomStringRepositoryIntf interface {
	live_house_staff_email_authorization_domain.RandomStringRepositoryIntf
}

type randomStringRepositoryImpl struct{}

func NewRandomStringRepository() RandomStringRepositoryIntf {
	return &randomStringRepositoryImpl{}
}

func (r randomStringRepositoryImpl) Generate(length int) (URLSafeString string, err error) {

	s := make([]byte, length)

	if _, err = rand.Read(s); err != nil {
		return "", err
	}

	URLSafeString = hex.EncodeToString(s)

	return
}
