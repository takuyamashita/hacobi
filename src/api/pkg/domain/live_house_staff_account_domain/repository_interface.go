package live_house_staff_account_domain

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
)

type LiveHouseStaffAccountRepositoryIntf interface {
	FindByEmail(emailAddress domain.LiveHouseStaffEmailAddress, ctx context.Context) (LiveHouseStaffAccountIntf, error)
}
