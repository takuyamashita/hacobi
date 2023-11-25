package webauthn

type PublicKeyCreadentialRpEntity struct {
	Id   BufferSource `json:"id,omitempty"`
	Name string       `json:"name"`
}
