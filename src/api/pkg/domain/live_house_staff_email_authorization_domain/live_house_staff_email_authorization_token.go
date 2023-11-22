package live_house_staff_email_authorization_domain

import "fmt"

const tokenLength = 64

type Token struct {
	value string
}

func newToken(token string) (Token, error) {

	if len(token) != tokenLength {
		return Token{}, fmt.Errorf("token length must be %d", tokenLength)
	}

	return Token{token}, nil
}

func (t Token) String() string {
	return t.value
}
