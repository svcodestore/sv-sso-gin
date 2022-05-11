package system

import "errors"

type PrivilegeService struct {
}

func (s *PrivilegeService) IsGlobalUser(userId string) (isGlobalUser bool, err error) {
	_, isGlobalUser, err = privilegeApplicationService.AccessibleApplications(userId)
	return
}

func (s *PrivilegeService) CanAccessSystem(userId, clientId string) (can bool, err error) {
	application, e := applicationService.ApplicationWithClientId(clientId)
	if e != nil {
		err = e
		return
	}
	if application.ClientID != clientId {
		err = errors.New("application nonexistent")
		return
	}

	// 此用户是否是全局用户
	isGlobal, er := s.IsGlobalUser(userId)
	if isGlobal {
		can = true
		return
	}
	_, e = userService.UserWithIdAndApplicationId(userId, application.ID)
	// 当前应用没有此用户并且不是全局用户
	if e != nil && er != nil {
		err = e
		return
	}
	// 当前应用是否已注册
	organizations, e := organizationService.OrganizationsWithApplicationIds(application.ID)
	if e != nil {
		err = e
		return
	}
	organizationCount := len(organizations)
	if organizationCount == 0 {
		err = errors.New("this system is not registered")
		return
	}
	can = true

	return
}
