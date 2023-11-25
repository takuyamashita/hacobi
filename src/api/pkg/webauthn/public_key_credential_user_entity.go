package webauthn

type PublicKeyCredentialUserEntity struct {
	Id          BufferSource `json:"id"`
	Name        string       `json:"name"`
	DisplayName string       `json:"displayName"`
}
