package webauthn

const (
	AttestationConveyancePreferenceNone       AttestationConveyancePreference = "none"
	AttestationConveyancePreferenceIndirect   AttestationConveyancePreference = "indirect"
	AttestationConveyancePreferenceDirect     AttestationConveyancePreference = "direct"
	AttestationConveyancePreferenceEnterprise AttestationConveyancePreference = "enterprise"
)

type AttestationConveyancePreference string
