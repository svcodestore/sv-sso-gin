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
			apps, err = applicationService.AllApplication()
			return
		}
		organizationIds[i] = organizations[i].ID
	}
	apps, err = applicationService.ApplicationsWithOrganizationIds(organizationIds...)

	return
}

func (s *PrivilegeApplicationService) AvailableApplications() (apps []*model.Applications, err error) {
	organizationApplication, err := organizationApplicationService.AllOrganizationApplication()
	if err != nil {
		return
	}

	cnt := len(organizationApplication)
	if cnt == 0 {
		return
	}

	for i := 0; i < cnt; i++ {
		if organizationApplication[i].Status {
			o, _ := organizationService.IsAvailableOrganizations(organizationApplication[i].OrganizationID)
			a, _ := applicationService.IsAvailableApplications(organizationApplication[i].ApplicationID)
			if len(o) == 1 && len(a) == 1 {
				apps = append(apps, a...)
			}
		}
	}

	return
}
