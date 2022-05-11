package system

import (
	"errors"
	"github.com/svcodestore/sv-sso-gin/model"
)

type PrivilegeApplicationService struct {

}

func (s *PrivilegeApplicationService) AccessibleApplications(userId string) (apps []*model.Applications, isGlobalUser bool, err error) {
	applications, err := applicationService.ApplicationsWithUserId(userId)
	if err != nil {
		return
	}

	applicationCount := len(applications)
	if applicationCount == 0 {
		err = errors.New("not registered")
		return
	}
	var applicationIds = make([]string, applicationCount)
	for i := 0; i < applicationCount; i++ {
		applicationIds[i] = applications[i].ID
	}

	organizations, e := organizationService.OrganizationsWithApplicationIds(applicationIds...)
	if e != nil {
		err = e
		return
	}
	organizationCount := len(organizations)
	if organizationCount == 0 {
		err = errors.New("not registered")
		return
	}
	var organizationIds = make([]string, organizationCount)
	for i := 0; i < organizationCount; i++ {
		if organizations[i].ID == "0" {
			isGlobalUser = true
		}
		organizationIds[i] = organizations[i].ID
	}
	apps, err = applicationService.ApplicationsWithOrganizationIds(organizationIds...)

	return
}
