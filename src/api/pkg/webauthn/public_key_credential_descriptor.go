package webauthn

import "encoding/base64"

const (
	AuthnticatorTransportUSB AuthnticatorTransport = "usb"
	AuthnticatorTransportNFC AuthnticatorTransport = "nfc"
	AuthnticatorTransportBLE AuthnticatorTransport = "ble"
	AuthnticatorTransportInt AuthnticatorTransport = "internal"
)

type AuthnticatorTransport string

type CredentialKey []byte

func (k CredentialKey) MarshalJSON() ([]byte, error) {

	return []byte(`"` + base64.RawURLEncoding.EncodeToString(k) + `"`), nil
}

type PublicKeyCredentialDescriptor struct {
	Type       PublicKeyCredentialType `json:"type"`
	Id         CredentialKey           `json:"id"`
	Transports []AuthnticatorTransport `json:"transports"`
}
