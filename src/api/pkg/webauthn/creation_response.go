package webauthn

type CreationResponse struct {
	Response                AuthenticatorAttestationResponse      `json:"response"`
	ID                      string                                `json:"id"`
	RawID                   BufferSource                          `json:"rawId"`
	AuthenticatorAttachment AuthenticatorAttachment               `json:"authenticatorAttachment"`
	ClientExtensionResults  AuthenticationExtensionsClientOutputs `json:"clientExtensionResults,omitempty"`
}
