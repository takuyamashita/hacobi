package repository

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
)

var _ domain.RandomStringRepositoryIntf = (*randomStringRepositoryImpl)(nil)

type RandomStringRepositoryIntf interface {
	domain.RandomStringRepositoryIntf
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
