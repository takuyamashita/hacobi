package webauthn

// https://www.w3.org/TR/webauthn/#sctn-attestation
type AttestationObject struct {
	AuthData     AuthenticatorData
	RawAuthData  []byte                 `json:"authData"`
	Format       string                 `json:"fmt"`
	AttStatement map[string]interface{} `json:"attStmt,omitempty"`
}
