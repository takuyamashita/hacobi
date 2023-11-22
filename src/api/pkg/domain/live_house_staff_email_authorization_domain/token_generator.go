package live_house_staff_email_authorization_domain

type TokenGeneratorIntf interface {
	Generate() (Token, error)
}

type tokenGeneratorImpl struct {
	randomStringRepository RandomStringRepositoryIntf
}

func NewTokenGenerator(randomStringRepository RandomStringRepositoryIntf) TokenGeneratorIntf {
	return &tokenGeneratorImpl{
		randomStringRepository: randomStringRepository,
	}
}

func (t tokenGeneratorImpl) Generate() (Token, error) {
	tokenStr, err := t.randomStringRepository.Generate(tokenLength)
	if err != nil {
		return Token{}, err
	}

	tkn, err := newToken(tokenStr)
	if err != nil {
		return Token{}, err
	}

	return tkn, nil
}
