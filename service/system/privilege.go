package system

import (
	"errors"
)

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
	isGlobal, _ := s.IsGlobalUser(userId)
	if isGlobal {
		can = true
		return
	}
	// 用户拥有的应用
	applications, _ := applicationService.ApplicationsWithUserId(userId)
	isHaveApplication := false
	for _, m := range applications {
		if m.ID == application.ID {
			isHaveApplication = true
		}
	}
	if !isHaveApplication {
		err = errors.New("unregister")
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
