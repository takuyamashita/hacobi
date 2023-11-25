package webauthn

type PublicKeyCredentialUserEntity struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}
