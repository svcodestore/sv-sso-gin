package system

import (
	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model"
)

type OrganizationApplicationService struct {
}

func (s *OrganizationApplicationService) CreateOrganizationApplication(o *model.OrganizationApplication) (organizationApplication model.OrganizationApplication, err error) {
	result := model.OrganizationApplicationMgr(global.DB).Create(o)
	if result.Error != nil {
		err = result.Error
		return
	}

	organizationApplication, err = s.OrganizationApplicationWithId(o)
	return
}

func (s *OrganizationApplicationService) DeleteOrganizationApplicationWithId(o *model.OrganizationApplication) (isDeleted bool) {
	db := global.DB.Where("organization_id = ? and application_id = ?", o.OrganizationID, o.ApplicationID).Delete(o)
	isDeleted = db.RowsAffected == 1
	return
}

func (s *OrganizationApplicationService) UpdateOrganizationApplicationWithId(o *model.OrganizationApplication) (organizationApplication model.OrganizationApplication, err error) {
	db := model.OrganizationApplicationMgr(global.DB).Where("organization_id = ? and application_id = ?", o.OrganizationID, o.ApplicationID).Updates(o)
	if db.RowsAffected == 1 {
		organizationApplication, err = s.OrganizationApplicationWithId(o)
		return
	}
	err = db.Error
	return
}

func (s *OrganizationApplicationService) AllOrganizationApplication() (organizationApplications []*model.OrganizationApplication, err error) {
	organizationApplications, err = model.OrganizationApplicationMgr(global.DB).Gets()
	return
}

func (s *OrganizationApplicationService) OrganizationApplicationWithId(o *model.OrganizationApplication) (organizationApplication model.OrganizationApplication, err error) {
	organizationApplication, err = model.OrganizationApplicationMgr(global.DB).FetchByPrimaryKey(o.OrganizationID, o.ApplicationID)
	return
}

func (s *OrganizationApplicationService) AvailableOrganizationApplications() (results []*model.OrganizationApplication, err error) {
	organizationApplications, err := s.AllOrganizationApplication()
	if err != nil {
		return
	}

	for i := 0; i < len(organizationApplications); i++ {
		if organizationApplications[i].Status == 1 {
			results = append(results, organizationApplications[i])
		}
	}

	return
}
