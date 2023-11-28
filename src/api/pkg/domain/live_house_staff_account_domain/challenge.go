package live_house_staff_account_domain

import "encoding/base64"

type Challenge struct {
	value []byte
}

func NewChallenge(base64URLSafeString string) (Challenge, error) {

	bytes, err := base64.RawURLEncoding.DecodeString(base64URLSafeString)
	if err != nil {
		return Challenge{}, err
	}

	return Challenge{bytes}, nil
}

func NewChallengeFromBytes(bytes []byte) Challenge {
	return Challenge{bytes}
}

func (c Challenge) String() string {
	return base64.RawURLEncoding.EncodeToString(c.value)
}
