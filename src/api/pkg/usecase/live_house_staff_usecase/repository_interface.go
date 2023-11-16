package live_house_staff_usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type UuidRepository interface {
	Generate() (string, error)
}

type LiveHouseStaffRepository interface {
	Save(owner live_house_staff_domain.LiveHouseStaffIntf, ctx context.Context) error
}
