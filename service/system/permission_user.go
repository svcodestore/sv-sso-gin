package system

import "github.com/svcodestore/sv-sso-gin/model"

type PermissionUserService struct {
}

func (s *PermissionUserService) AvailableUsersWithApplicationIds(appIds ...string) (users []*model.Users, err error) {
	availableApplications, err := privilegeApplicationService.AvailableApplications()
	if err != nil {
		return
	}

	cnt := len(availableApplications)
	if cnt == 0 {
		return
	}

	var applicationIds = make([]string, cnt)
	for i := 0; i < cnt; i++ {
		applicationIds = append(applicationIds, availableApplications[i].ID)
	}

	if len(appIds) > 0 {
		applicationIds = appIds
	}

	u, err := userService.UsersWithApplicationIds(applicationIds...)
	if err != nil {
		return
	}
	for i := 0; i < len(u); i++ {
		if u[i].Status {
			users = append(users, u[i])
		}
	}

	return
}
