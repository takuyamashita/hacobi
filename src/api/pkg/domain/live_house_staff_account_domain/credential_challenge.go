package live_house_staff_account_domain

import "time"

type CredentialChallengeIntf interface {
	Challenge() Challenge
}

type CredentialChallengeImpl struct {
	challenge Challenge
	createdAt time.Time
}

func NewCredentialChallenge(challenge Challenge, createdAt time.Time) (CredentialChallengeIntf, error) {

	return &CredentialChallengeImpl{
		challenge: challenge,
		createdAt: createdAt,
	}, nil
}

func (c CredentialChallengeImpl) Challenge() Challenge {
	return c.challenge
}
