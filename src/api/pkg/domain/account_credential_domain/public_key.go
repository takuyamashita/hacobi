package account_credential_domain

import "encoding/base64"

type PublicKey []byte

func NewPublicKey(key string) (PublicKey, error) {

	s, err := base64.RawURLEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	return PublicKey(s), nil
}

func (key PublicKey) String() string {
	return base64.RawURLEncoding.EncodeToString(key)
}
