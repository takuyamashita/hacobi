package webauthn

import (
	"crypto/rand"
)

const ChallengeLength = 64

func NewChallenge() (BufferSource, error) {

	challenge := make([]byte, ChallengeLength)

	if _, err := rand.Read(challenge); err != nil {
		return nil, err
	}

	return challenge, nil
}
