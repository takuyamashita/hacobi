package account_credential_domain

import (
	"encoding/base64"
	"fmt"
)

type AuthenticatorAttachment string

type AAGUID []byte

func (a AAGUID) String() string {
	return base64.RawURLEncoding.EncodeToString(a)
}

func NewAAGUID(aaguid string) (AAGUID, error) {

	s, err := base64.RawURLEncoding.DecodeString(aaguid)
	if err != nil {
		return nil, err
	}

	return AAGUID(s), nil
}

const (
	Platform      AuthenticatorAttachment = "platform"
	CrossPlatform AuthenticatorAttachment = "cross-platform"
)

type Authenticator struct {
	aAGUID       AAGUID
	signCount    uint32
	cloneWarning bool
	attachment   AuthenticatorAttachment
}

func NewAuthenticator(
	aaguid AAGUID,
	signCount uint32,
	cloneWarning bool,
	attachment AuthenticatorAttachment,
) Authenticator {
	return Authenticator{
		aAGUID:       aaguid,
		signCount:    signCount,
		cloneWarning: cloneWarning,
		attachment:   attachment,
	}
}

func (a Authenticator) AAGUID() AAGUID {
	return a.aAGUID
}

func (a Authenticator) SignCount() uint32 {
	return a.signCount
}

func (a Authenticator) CloneWarning() bool {
	return a.cloneWarning
}

func (a Authenticator) Attachment() AuthenticatorAttachment {
	return a.attachment
}

func NewAuthenticatorAttachment(attachment string) (AuthenticatorAttachment, error) {

	switch attachment {
	case "platform":
		return Platform, nil
	case "cross-platform":
		return CrossPlatform, nil
	default:
		return "", fmt.Errorf("unknown attachment: %s", attachment)
	}
}
