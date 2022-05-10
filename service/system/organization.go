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
	db := global.OrganizationMgr.Where("id = ?", id).Updates(o)
	if db.RowsAffected == 1 {
		global.OrganizationMgr = model.OrganizationsMgr(utils.Gorm())
		return s.OrganizationWithId(id)
	}
	err = db.Error
	return
}

func (s *OrganizationService) UpdateOrganizationStatusWithId(status bool, id, updatedBy string) (organization model.Organizations, err error) {
	err = global.OrganizationMgr.Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error
	if err != nil {
		return
	}
	global.OrganizationMgr = model.OrganizationsMgr(utils.Gorm())
	organization, err = global.OrganizationMgr.GetFromID(id)
	return
}

func (s *OrganizationService) AllOrganization() (organizations []*model.Organizations, err error) {
	organizations, err = global.OrganizationMgr.Gets()
	return
}

func (s *OrganizationService) OrganizationWithId(id string) (organization model.Organizations, err error) {
	organization, err = global.OrganizationMgr.GetFromID(id)
	return
}

func (s *OrganizationService) OrganizationsWithApplicationId(applicationId string) (organizations []*model.Organizations, err error) {
	results, err := global.OrganizationApplicationMgr.GetFromApplicationID(applicationId)
	if err != nil {
		return
	}

	l := len(results)
	var ids = make([]string, l)
	for i := 0; i < l; i++ {
		ids[i] = results[i].OrganizationID
	}
	organizations, err = global.OrganizationMgr.GetBatchFromID(ids)

	return
}