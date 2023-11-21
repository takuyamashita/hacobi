package live_house_staff_domain

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
)

type LiveHouseStaffRepositoryIntf interface {
	FindByEmail(emailAddress domain.LiveHouseStaffEmailAddress, ctx context.Context) (LiveHouseStaffIntf, error)
}
