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
	result := global.OrganizationMgr.Create(o)
	if result.Error != nil {
		err = result.Error
		return
	}

	return s.OrganizationWithId(o)
}

func (s *OrganizationService) DeleteOrganizationWithId(o *model.Organizations) (isDeleted bool) {
	db := global.DB.Where("id = ?", o.ID).Delete(o)
	isDeleted = db.RowsAffected == 1
	return
}

func (s *OrganizationService) UpdateOrganizationWithId(o *model.Organizations) (organization model.Organizations, err error) {
	db := global.OrganizationMgr.Where("id = ?", o.ID).Updates(o)
	if db.RowsAffected == 1 {
		return s.OrganizationWithId(o)
	}
	err = db.Error
	return
}

func (s *OrganizationService) AllOrganization() (organizations []*model.Organizations, err error) {
	organizations, err = global.OrganizationMgr.Gets()
	return
}

func (s *OrganizationService) OrganizationWithId(o *model.Organizations) (organization model.Organizations, err error) {
	organization, err = global.OrganizationMgr.GetFromID(o.ID)
	return
}
