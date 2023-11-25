package webauthn

// registration
// https://www.w3.org/TR/webauthn/#sctn-registering-a-new-credential

// https://www.w3.org/TR/webauthn/#dictdef-publickeycredentialcreationoptions
type PublicKeyCredentialCreationOptions struct {
	rp   PublicKeyCreadentialRpEntity
	user PublicKeyCredentialUserEntity
}
