package webauthn

// registration
// https://www.w3.org/TR/webauthn/#sctn-registering-a-new-credential

// https://www.w3.org/TR/webauthn/#dictdef-publickeycredentialcreationoptions
type PublicKeyCredentialCreationOptions struct {
	rp   PublicKeyCreadentialRpEntity
	user PublicKeyCredentialUserEntity

	challenge        Challenge
	pubKeyCredParams []PublicKeyCredentialParameters

	timeout uint

	// このメンバは、単一の認証子で同じアカウントに対する複数のクレデンシャルの作成を制限したい依拠当事者による使用を意図している。
	// 新しいクレデンシャルが、このパラメータに列挙されたクレデンシャルのいずれかを含む認証機で作成される場合、クライアントはエラーを返すように要求される。
	excludeCredentials []PublicKeyCredentialDescriptor
}
