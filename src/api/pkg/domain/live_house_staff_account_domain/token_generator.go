package live_house_staff_account_domain

import "github.com/takuyamashita/hacobi/src/api/pkg/domain"

type TokenGeneratorIntf interface {
	Generate() (*Token, error)
}

type tokenGeneratorImpl struct {
	randomStringRepository domain.RandomStringRepositoryIntf
}

func NewTokenGenerator(randomStringRepository domain.RandomStringRepositoryIntf) TokenGeneratorIntf {
	return &tokenGeneratorImpl{
		randomStringRepository: randomStringRepository,
	}
}

func (t tokenGeneratorImpl) Generate() (*Token, error) {

	tokenStr, err := t.randomStringRepository.Generate(TokenLength)
	if err != nil {
		return nil, err
	}

	tkn, err := NewTokenFromHexString(tokenStr)
	if err != nil {
		return nil, err
	}

	return &tkn, nil
}
