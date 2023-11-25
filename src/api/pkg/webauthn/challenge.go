package webauthn

import (
	"crypto/rand"
	"encoding/base64"
)

const ChallengeLength = 64

type Challenge []byte

func NewChallenge() (Challenge, error) {

	challenge := make([]byte, ChallengeLength)

	if _, err := rand.Read(challenge); err != nil {
		return nil, err
	}

	return challenge, nil
}

func (c Challenge) URLSafeBase64() string {

	return base64.RawURLEncoding.EncodeToString(c)
}
