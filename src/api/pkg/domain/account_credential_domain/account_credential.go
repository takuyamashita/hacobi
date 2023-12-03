package account_credential_domain

import (
	"time"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
)

type AccountCredentialIntf interface {
	PublicKeyId() domain.PublicKeyId
	PublicKey() PublicKey
	AttestationType() AttestationType
	Transport() []Transport
	Flags() Flags
	Authenticator() Authenticator
	CreatedAt() time.Time
}

type accountCredentialImpl struct {
	publicKeyID     domain.PublicKeyId
	publicKey       PublicKey
	attestationType AttestationType
	transport       []Transport
	flags           Flags
	authenticator   Authenticator
	createdAt       time.Time
}

type NewAccountCredentialParams struct {
	PublicKeyID     string
	PublicKey       string
	AttestationType string
	Transport       []string
	Flags           struct {
		UserPresent    bool
		UserVerified   bool
		BackupEligible bool
		BackupState    bool
	}
	Authenticator struct {
		AAGUID       string
		SignCount    uint32
		Attachment   string
		CloneWarning bool
	}
	CreatedAt time.Time
}

func NewAccountCredential(
	params NewAccountCredentialParams,
) (AccountCredentialIntf, error) {

	publicKeyId, err := domain.NewPublicKeyId(params.PublicKeyID)

	publicKey, err := NewPublicKey(params.PublicKey)
	if err != nil {
		return nil, err
	}

	attestationType, err := NewAttestationType(params.AttestationType)
	if err != nil {
		return nil, err
	}

	transport, err := NewTransports(params.Transport)
	if err != nil {
		return nil, err
	}

	flags := Flags{
		UserPresent:    params.Flags.UserPresent,
		UserVerified:   params.Flags.UserVerified,
		BackupEligible: params.Flags.BackupEligible,
		BackupState:    params.Flags.BackupState,
	}

	aaguid, err := NewAAGUID(params.Authenticator.AAGUID)
	if err != nil {
		return nil, err
	}

	authenticator := NewAuthenticator(
		aaguid,
		params.Authenticator.SignCount,
		params.Authenticator.CloneWarning,
		AuthenticatorAttachment(params.Authenticator.Attachment),
	)

	return &accountCredentialImpl{
		publicKeyID:     publicKeyId,
		publicKey:       publicKey,
		attestationType: attestationType,
		transport:       transport,
		flags:           flags,
		authenticator:   authenticator,
		createdAt:       params.CreatedAt,
	}, nil
}

func (l *accountCredentialImpl) PublicKeyId() domain.PublicKeyId {
	return l.publicKeyID
}

func (l *accountCredentialImpl) PublicKey() PublicKey {

	return l.publicKey
}

func (l *accountCredentialImpl) AttestationType() AttestationType {
	return l.attestationType
}

func (l *accountCredentialImpl) Transport() []Transport {
	return l.transport
}

func (l *accountCredentialImpl) Flags() Flags {
	return l.flags
}

func (l *accountCredentialImpl) Authenticator() Authenticator {
	return l.authenticator
}

func (l *accountCredentialImpl) CreatedAt() time.Time {
	return l.createdAt
}
