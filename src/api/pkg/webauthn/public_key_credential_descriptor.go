package webauthn

const (
	AuthnticatorTransportUSB AuthnticatorTransport = "usb"
	AuthnticatorTransportNFC AuthnticatorTransport = "nfc"
	AuthnticatorTransportBLE AuthnticatorTransport = "ble"
	AuthnticatorTransportInt AuthnticatorTransport = "internal"
)

type AuthnticatorTransport string

type PublicKeyCredentialDescriptor struct {
	Type       PublicKeyCredentialType
	ID         []byte
	Transports []AuthnticatorTransport
}
