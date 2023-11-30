package live_house_staff_account_domain

import "time"

type CredentialChallengeIntf interface {
	Challenge() Challenge
}

type credentialChallengeImpl struct {
	challenge Challenge
	createdAt time.Time
}

func NewCredentialChallenge(challenge string, createdAt time.Time) (CredentialChallengeIntf, error) {

	c, err := NewChallenge(challenge)
	if err != nil {
		return nil, err
	}

	return &credentialChallengeImpl{
		challenge: c,
		createdAt: createdAt,
	}, nil
}

func (c credentialChallengeImpl) Challenge() Challenge {
	return c.challenge
}
