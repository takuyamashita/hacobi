package account_credential_domain

type AttestationType string

func (t AttestationType) String() string {
	return string(t)
}

func NewAttestationType(t string) (AttestationType, error) {

	return AttestationType(t), nil
}
