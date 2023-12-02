package domain

import "encoding/base64"

type PublicKeyId []byte

func NewPublicKeyId(id string) (PublicKeyId, error) {

	s, err := base64.RawURLEncoding.DecodeString(id)
	if err != nil {
		return nil, err
	}

	return PublicKeyId(s), nil
}

func (id PublicKeyId) String() string {
	return base64.RawURLEncoding.EncodeToString(id)
}
