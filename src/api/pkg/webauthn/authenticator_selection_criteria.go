package webauthn

const (
	AuthenticatorAttachmentPlatform      AuthenticatorAttachment = "platform"
	AuthenticatorAttachmentCrossPlatform AuthenticatorAttachment = "cross-platform"

	ResidentKeyRequirementRequired    ResidentKeyRequirement = "required"
	ResidentKeyRequirementPreferred   ResidentKeyRequirement = "preferred"
	ResidentKeyRequirementDiscouraged ResidentKeyRequirement = "discouraged"

	UserVerificationRequirementRequired    UserVerificationRequirement = "required"
	UserVerificationRequirementPreferred   UserVerificationRequirement = "preferred"
	UserVerificationRequirementDiscouraged UserVerificationRequirement = "discouraged"
)

type AuthenticatorAttachment string

type ResidentKeyRequirement string

type UserVerificationRequirement string

type AuthenticatorSelectionCriteria struct {
	AuthenticatorAttachment AuthenticatorAttachment
	ResidentKey             ResidentKeyRequirement

	// このメンバは、WebAuthn Level 1 との後方互換性のために保持され、歴史的な理由から、その命名は、発見可能なクレデンシャルに関する非推奨の「常駐」用語を保持する。
	// 依拠当事者は、residentKey が required に設定されている場合にのみ、このメンバを true に設定す べきである。
	RequiredResidentKey bool
	UserVerification    UserVerificationRequirement
}
