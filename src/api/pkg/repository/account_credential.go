package repository

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/takuyamashita/hacobi/src/api/pkg/db"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
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
			AS new
			ON DUPLICATE KEY UPDATE
				public_key = new.public_key,
				attestation_type = new.attestation_type,
				transport = new.transport,
				user_present = new.user_present,
				user_verified = new.user_verified,
				backup_eligible = new.backup_eligible,
				backup_state = new.backup_state,
				aaguid = new.aaguid,
				sign_count = new.sign_count,
				attachment = new.attachment,
				created_at = new.created_at
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

func (repo AccountCredentialRepositoryImpl) FindByIds(
	ids []domain.PublicKeyId,
	ctx context.Context,
) ([]account_credential_domain.AccountCredentialIntf, error) {

	var account = make([]account_credential_domain.AccountCredentialIntf, len(ids))

	t := bytes.Repeat([]byte("?,"), len(ids))

	var args = make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id.String()
	}

	rows, err := repo.db.QueryContext(
		ctx,
		fmt.Sprintf(
			`
				SELECT
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
				FROM
					account_credentials
				WHERE
					public_key_id IN (%s)
			`,
			t[:len(t)-1],
		),
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {

		var (
			publicKeyId     string
			publicKey       string
			attestationType string
			transport       uint
			userPresent     int
			userVerified    int
			backupEligible  int
			backupState     int
			aaguid          string
			signCount       uint32
			attachment      string
			createdAt       time.Time
		)

		if err := rows.Scan(
			&publicKeyId,
			&publicKey,
			&attestationType,
			&transport,
			&userPresent,
			&userVerified,
			&backupEligible,
			&backupState,
			&aaguid,
			&signCount,
			&attachment,
			&createdAt,
		); err != nil {
			return nil, err
		}

		var transportArray = make([]string, 0, 5)
		for i, t := range getTransport() {
			switch transport & (1 << transportKey(i)) {
			case 1 << USB:
				transportArray = append(transportArray, string(t))
			case 1 << NFC:
				transportArray = append(transportArray, string(t))
			case 1 << BLE:
				transportArray = append(transportArray, string(t))
			case 1 << Hybrid:
				transportArray = append(transportArray, string(t))
			case 1 << Internal:
				transportArray = append(transportArray, string(t))
			}
		}

		accountCredential, err := account_credential_domain.NewAccountCredential(
			account_credential_domain.NewAccountCredentialParams{
				PublicKeyID:     publicKeyId,
				PublicKey:       publicKey,
				AttestationType: attestationType,
				Transport:       transportArray,
				Flags: account_credential_domain.Flags{
					UserPresent:    userPresent == 1,
					UserVerified:   userVerified == 1,
					BackupEligible: backupEligible == 1,
					BackupState:    backupState == 1,
				},
				Authenticator: struct {
					AAGUID       string
					SignCount    uint32
					Attachment   string
					CloneWarning bool
				}{
					AAGUID:       aaguid,
					SignCount:    signCount,
					Attachment:   attachment,
					CloneWarning: false,
				},
				CreatedAt: createdAt,
			},
		)
		if err != nil {
			return nil, err
		}

		account[i] = accountCredential
	}

	return account, nil
}
