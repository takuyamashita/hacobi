package webauthn

// registration
// https://www.w3.org/TR/webauthn/#sctn-registering-a-new-credential

// https://www.w3.org/TR/webauthn/#dictdef-publickeycredentialcreationoptions
type PublicKeyCredentialCreationOptions struct {
	Rp   PublicKeyCreadentialRpEntity
	User PublicKeyCredentialUserEntity

	Challenge        Challenge
	PubKeyCredParams []PublicKeyCredentialParameters

	Timeout uint

	// このメンバは、単一の認証子で同じアカウントに対する複数のクレデンシャルの作成を制限したい依拠当事者による使用を意図している。
	// 新しいクレデンシャルが、このパラメータに列挙されたクレデンシャルのいずれかを含む認証機で作成される場合、クライアントはエラーを返すように要求される。
	ExcludeCredentials    []PublicKeyCredentialDescriptor
	AuthnticatorSelection AuthenticatorSelectionCriteria
	Attestation           AttestationConveyancePreference

	// 5 Web Authentication API」で定義されている、公開鍵クレデンシャルを生成するメカニズム、 およびAuthenticationアサーションをリクエストおよび生成するメカニズムは、 特定の使用ケースに合わせて拡張することができる。各ケースには、登録拡張および/または認証拡張を定義することで対応する。
	//すべての拡張はクライアント拡張である。これは、拡張がクライアントとの通信およびク ライアントによる処理を伴うことを意味する。クライアント拡張は、以下のステップとデータを定義する：
	// navigator.credentials.create()エクステンションリクエストパラメータと、登録エクステンションのレスポンス値。
	// navigator.credentials.get() 認証エクステンションのエクステンションリクエストパラメータとレスポンス値。
	// 登録エクステンションと認証エクステンションのクライアントエクステンション処理。
	// 公開鍵クレデンシャルを作成するとき、または認証アサーションを要求するとき、WebAuthn 依拠当事者は一連の拡張の使用を要求できる。これらの拡張は、クライアントや WebAuthn 認証機能でサポートされていれば、要求された操作中に呼び出される。依拠当事者は、get()呼び出し(認証拡張の場合)または create()呼び出し(登録拡張の場合)で、各拡張のクライアント拡張入力をクライアントに送信する。クライアントは、クライアントプラットフォームがサポートする各拡 張のクライアント拡張処理を実行し、拡張識別子およびクライアント拡張出力値を含む、各拡 張で指定されるクライアントデータを拡張する。
	// Extentions AuthenticationExtensionsClientInputs
}
