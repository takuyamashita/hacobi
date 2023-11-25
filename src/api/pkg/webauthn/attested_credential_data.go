package webauthn

// https://www.w3.org/TR/webauthn/#sctn-attested-credential-data
type AttestedCredentialData struct {
	AAGUID       []byte `json:"aaguid"`
	CredentialID []byte `json:"credential_id"`

	// The raw credential public key bytes received from the attestation data.
	CredentialPublicKey []byte `json:"public_key"`
}
