package system

import (
	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/utils"
)

type OrganizationService struct {
}

func (s *OrganizationService) CreateOrganization(o *model.Organizations) (organization model.Organizations, err error) {
	o.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
	result := model.OrganizationsMgr(global.DB).Create(o)
	if result.Error != nil {
		err = result.Error
		return
	}

	return s.OrganizationWithId(o.ID)
}

func (s *OrganizationService) DeleteOrganizationWithId(o *model.Organizations) (isDeleted bool) {
	db := global.DB.Where("id = ?", o.ID).Delete(o)
	isDeleted = db.RowsAffected == 1
	return
}

func (s *OrganizationService) UpdateOrganizationWithId(o *model.Organizations) (organization model.Organizations, err error) {
	id := o.ID
	o.ID = ""
	db := model.OrganizationsMgr(global.DB).Where("id = ?", id).Updates(o)
	if db.RowsAffected == 1 {
		return s.OrganizationWithId(id)
	}
	err = db.Error
	return
}

func (s *OrganizationService) UpdateOrganizationStatusWithId(status bool, id, updatedBy string) (organization model.Organizations, err error) {
	err = model.OrganizationsMgr(global.DB).Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error
	if err != nil {
		return
	}
	organization, err = model.OrganizationsMgr(global.DB).GetFromID(id)
	return
}

func (s *OrganizationService) AllOrganization() (organizations []*model.Organizations, err error) {
	organizations, err = model.OrganizationsMgr(global.DB).Gets()
	return
}

func (s *OrganizationService) OrganizationWithId(id string) (organization model.Organizations, err error) {
	organization, err = model.OrganizationsMgr(global.DB).GetFromID(id)
	return
}

func (s *OrganizationService) OrganizationsWithApplicationIds(applicationIds ...string) (organizations []*model.Organizations, err error) {
	results, err := model.OrganizationApplicationMgr(global.DB).GetBatchFromApplicationID(applicationIds)
	if err != nil {
		return
	}

	l := len(results)
	var ids = make([]string, l)
	for i := 0; i < l; i++ {
		ids[i] = results[i].OrganizationID
	}
	organizations, err = model.OrganizationsMgr(global.DB).GetBatchFromID(ids)

	return
}

func (s *OrganizationService) AvailableOrganizations() (organizations []*model.Organizations, err error) {
	results, err := s.AllOrganization()
	if err != nil {
		return
	}

	for i := 0; i < len(results); i++ {
		if results[i].Status == 1 {
			organizations = append(organizations, results[i])
		}
	}

	return
}

func (s *OrganizationService) IsAvailableOrganizations(organizationIds ...string) (organizations []*model.Organizations, err error) {
	o, err := model.OrganizationsMgr(global.DB).GetBatchFromID(organizationIds)
	if err != nil {
		return
	}

	for i := 0; i < len(o); i++ {
		if o[i].Status == 1 {
			organizations = append(organizations, o[i])
		}
	}
	return
}
