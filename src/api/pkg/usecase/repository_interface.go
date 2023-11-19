package usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type UuidRepositoryIntf interface {
	Generate() (string, error)
}

type TransationRepositoryIntf interface {
	Begin(ctx context.Context) (commit func() error, rollback func() error, err error)
}

type LiveHouseStaffRepositoryIntf interface {
	Save(owner live_house_staff_domain.LiveHouseStaffIntf, ctx context.Context) error
	FindById(id live_house_staff_domain.LiveHouseStaffId, ctx context.Context) (live_house_staff_domain.LiveHouseStaffIntf, error)
}

type LiveHouseAccountRepositoryIntf interface {
	Save(account live_house_account_domain.LiveHouseAccountIntf, ctx context.Context) error
}
