package webauthn

// https://www.w3.org/TR/webauthn/#dictdef-publickeycredentialcreationoptions
type PublicKeyCredentialCreationOptions struct {
	rp   PublicKeyCreadentialRpEntity
	user PublicKeyCredentialUserEntity
}
