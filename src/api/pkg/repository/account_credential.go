package repository

import (
	"context"
	"fmt"

	"github.com/takuyamashita/hacobi/src/api/pkg/db"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/account_credential_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

type AccountCredentialRepositoryIntf interface {
	usecase.AccountCredentialRepositoryIntf
}

type AccountCredentialRepositoryImpl struct {
	db *db.MySQL
}

func NewAccountCredentialRepository(db *db.MySQL) AccountCredentialRepositoryIntf {
	return &AccountCredentialRepositoryImpl{
		db: db,
	}
}

func getTransport() [5]account_credential_domain.Transport {
	return [5]account_credential_domain.Transport{
		account_credential_domain.USB,
		account_credential_domain.NFC,
		account_credential_domain.BLE,
		account_credential_domain.Hybrid,
		account_credential_domain.Internal,
	}
}

type transportKey uint

const (
	USB transportKey = iota
	NFC
	BLE
	Hybrid
	Internal
)

func (repo AccountCredentialRepositoryImpl) Save(
	credential account_credential_domain.AccountCredentialIntf,
	ctx context.Context,
) error {

	/*
		+-----------------------------+--------------+------+-----+----------------------+-------------------+
		| public_key_id               | varchar(128) | NO   | PRI | NULL                 |                   |
		| public_key                  | varchar(128) | NO   |     | NULL                 |                   |
		| attestation_type            | varchar(128) | NO   |     | NULL                 |                   |
		| transport                   | varchar(128) | NO   |     | NULL                 |                   |
		| user_present                | tinyint(1)   | NO   |     | NULL                 |                   |
		| user_verified               | tinyint(1)   | NO   |     | NULL                 |                   |
		| backup_eligible             | tinyint(1)   | NO   |     | NULL                 |                   |
		| backup_state                | tinyint(1)   | NO   |     | NULL                 |                   |
		| aaguid                      | varchar(128) | NO   |     | NULL                 |                   |
		| sign_count                  | int unsigned | NO   |     | NULL                 |                   |
		| attachment                  | varchar(128) | NO   |     | NULL                 |                   |
		| created_at                  | datetime(6)  | NO   |     | CURRENT_TIMESTAMP(6) | DEFAULT_GENERATED |
		+-----------------------------+--------------+------+-----+----------------------+-------------------+
	*/

	bool2int := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}

	var transportFlag uint = 0
	for _, transport := range credential.Transport() {
		switch transport {
		case account_credential_domain.USB:
			transportFlag |= 1 << USB
		case account_credential_domain.NFC:
			transportFlag |= 1 << NFC
		case account_credential_domain.BLE:
			transportFlag |= 1 << BLE
		case account_credential_domain.Hybrid:
			transportFlag |= 1 << Hybrid
		case account_credential_domain.Internal:
			transportFlag |= 1 << Internal
		default:
			return fmt.Errorf("invalid transport: %s", transport)
		}
	}

	_, err := repo.db.ExecContext(
		ctx,
		`
			INSERT INTO account_credentials
				(
					public_key_id,
					public_key,
					attestation_type,
					transport,
					user_present,
					user_verified,
					backup_eligible,
					backup_state,
					aaguid,
					sign_count,
					attachment,
					created_at
				)
			VALUES
				(
					?,
					?,
					?,
					?,
					?,
					?,
					?,
					?,
					?,
					?,
					?,
					?
				)
		`,
		credential.PublicKeyId().String(),
		credential.PublicKey().String(),
		credential.AttestationType().String(),
		transportFlag,
		bool2int(credential.Flags().UserPresent),
		bool2int(credential.Flags().UserVerified),
		bool2int(credential.Flags().BackupEligible),
		bool2int(credential.Flags().BackupState),
		credential.Authenticator().AAGUID().String(),
		credential.Authenticator().SignCount(),
		string(credential.Authenticator().Attachment()),
		credential.CreatedAt(),
	)
	if err != nil {
		return err
	}

	return nil
}
