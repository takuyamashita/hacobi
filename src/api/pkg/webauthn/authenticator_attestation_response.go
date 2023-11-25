package webauthn

type AuthenticatorAttestationResponse struct {
	AttestationObject BufferSource `json:"attestationObject"`
	ClientDataJSON    BufferSource `json:"clientDataJSON"`
	Transports        []string     `json:"transports"`
}
