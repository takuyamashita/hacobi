package live_house_staff_email_authorization_domain

type Token struct {
	value string
}

func newToken(token string) Token {
	return Token{token}
}
