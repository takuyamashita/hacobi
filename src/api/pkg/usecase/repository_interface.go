package usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type UuidRepositoryIntf interface {
	Generate() (string, error)
}

type LiveHouseStaffRepositoryIntf interface {
	Save(owner live_house_staff_domain.LiveHouseStaffIntf, ctx context.Context) error
	FindById(id live_house_staff_domain.LiveHouseStaffId) (live_house_staff_domain.LiveHouseStaffIntf, error)
}
