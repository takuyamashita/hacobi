package live_house_staff_email_authorization_domain

import "fmt"

const TokenLength = 64

type Token struct {
	value string
}

func newToken(token string) (Token, error) {

	if len(token) != TokenLength*2 {
		return Token{}, fmt.Errorf("token length must be %d", TokenLength)
	}

	return Token{token}, nil
}

func (t Token) String() string {
	return t.value
}
