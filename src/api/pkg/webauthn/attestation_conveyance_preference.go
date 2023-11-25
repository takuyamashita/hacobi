package webauthn

const (
	AttestationConveyancePreferenceNone       AttestationConveyancePreference = "none"
	AttestationConveyancePreferenceIndirect   AttestationConveyancePreference = "indirect"
	AttestationConveyancePreferenceDirect     AttestationConveyancePreference = "direct"
	AttestationConveyancePreferenceEnterprise AttestationConveyancePreference = "enterprise"
)

// https://www.w3.org/TR/webauthn/#enumdef-attestationconveyancepreference
type AttestationConveyancePreference string
