package live_house_staff_domain

import "context"

type LiveHouseStaffRepository interface {
	FindByEmail(emailAddress LiveHouseStaffEmailAddress, ctx context.Context) (LiveHouseStaffIntf, error)
}
