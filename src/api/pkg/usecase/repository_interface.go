package usecase

import (
	"context"
	"io"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type CredentialKeyIntf interface {
	//xxx: TODO return []byte, error
	CreateChallenge() (protocol.URLEncodedBase64, error)

	//xxx: TODO return []byte or struct, error
	ParseCredentialKey(body io.Reader) (*protocol.ParsedCredentialCreationData, error)
}

type UuidRepositoryIntf interface {
	Generate() (string, error)
}

type TransationRepositoryIntf interface {
	Begin(ctx context.Context) (commit func() error, rollback func() error, err error)
}

type LiveHouseStaffAccountRepositoryIntf interface {
	Save(account live_house_staff_account_domain.LiveHouseStaffAccountIntf, ctx context.Context) error
	FindByProvisionalRegistrationToken(token live_house_staff_account_domain.Token, ctx context.Context) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error)
}

type LiveHouseStaffRepositoryIntf interface {
	Save(owner live_house_staff_domain.LiveHouseStaffIntf, ctx context.Context) error
	FindById(id live_house_staff_domain.LiveHouseStaffId, ctx context.Context) (live_house_staff_domain.LiveHouseStaffIntf, error)
}

type LiveHouseAccountRepositoryIntf interface {
	Save(account live_house_account_domain.LiveHouseAccountIntf, ctx context.Context) error
}
