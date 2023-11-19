package live_house_staff_domain

import "context"

type LiveHouseStaffRepositoryIntf interface {
	FindByEmail(emailAddress LiveHouseStaffEmailAddress, ctx context.Context) (LiveHouseStaffIntf, error)
}
