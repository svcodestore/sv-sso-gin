package system

import (
	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/utils"
)

type ApplicationService struct {
}

func (s *ApplicationService) CreateApplication(a *model.Applications) (application model.Applications, err error) {
	a.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
	a.ClientID = utils.GenerateClientId()
	a.ClientSecret = utils.GenerateClientSecret()

	result := global.ApplicationMgr.Create(a)
	if result.Error != nil {
		err = result.Error
		return
	}

	return s.ApplicationWithId(a.ID)
}

func (s *ApplicationService) DeleteApplication(a *model.Applications) (isDeleted bool) {
	db := global.DB.Where("id = ?", a.ID).Delete(a)
	isDeleted = db.RowsAffected == 1
	return
}

func (s *ApplicationService) UpdateApplicationWithId(a *model.Applications) (application model.Applications, err error) {
	id := a.ID
	a.ID = ""
	db := global.ApplicationMgr.Where("id = ?", id).Updates(a)
	if db.RowsAffected == 1 {
		global.ApplicationMgr = model.ApplicationsMgr(utils.Gorm())
		return s.ApplicationWithId(id)
	}
	err = db.Error
	return
}

func (s ApplicationService) UpdateApplicationStatusWithId(status bool, id, updatedBy string) (application model.Applications, err error) {
	err = global.ApplicationMgr.Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error
	if err != nil {
		return
	}
	global.ApplicationMgr = model.ApplicationsMgr(utils.Gorm())
	application, err = global.ApplicationMgr.GetFromID(id)
	return
}

func (s *ApplicationService) AllApplication() (applications []*model.Applications, err error) {
	applications, err = global.ApplicationMgr.Gets()
	return
}

func (s *ApplicationService) ApplicationWithId(id string) (application model.Applications, err error) {
	application, err = global.ApplicationMgr.GetFromID(id)
	return
}

func (s *ApplicationService) ApplicationWithClientId(clientId string) (application model.Applications, err error) {
	application, err = global.ApplicationMgr.GetFromClientID(clientId)
	return
}

func (s *ApplicationService) ApplicationClientSecretWithClientId(clientId string) (clientSecret string, err error) {
	var application *model.Applications
	err = global.DB.Table(global.ApplicationMgr.GetTableName()).Select("client_secret").Where("client_id = ?", clientId).Find(&application).Error
	clientSecret = application.ClientSecret

	return
}

func (s *ApplicationService) ApplicationsWithOrganizationIds(organizationIds ...string) (applications []*model.Applications, err error) {
	results, err := global.OrganizationApplicationMgr.GetBatchFromOrganizationID(organizationIds)
	if err != nil {
		return
	}

	l := len(results)
	var ids = make([]string, l)
	for i := 0; i < l; i++ {
		ids[i] = results[i].ApplicationID
	}
	applications, err = global.ApplicationMgr.GetBatchFromID(ids)

	return
}

func (s *ApplicationService) ApplicationsWithUserId(userId string) (applications []*model.Applications, err error) {
	results, err := global.ApplicationUserMgr.GetFromUserID(userId)
	if err != nil {
		return
	}

	l := len(results)
	var ids = make([]string, l)
	for i := 0; i < l; i++ {
		ids[i] = results[i].ApplicationID
	}
	applications, err = global.ApplicationMgr.GetBatchFromID(ids)

	return
}

func (s *ApplicationService) AvailableApplications() (applications []*model.Applications, err error) {
	results, err := s.AllApplication()
	if err != nil {
		return
	}

	for i := 0; i < len(results); i++ {
		if results[i].Status {
			applications = append(applications, results[i])
		}
	}

	return
}

func (s *ApplicationService) IsAvailableApplications(applicationIds ...string) (applications []*model.Applications, err error) {
	apps, err := global.ApplicationMgr.GetBatchFromID(applicationIds)
	if err != nil {
		return
	}

	for i := 0; i < len(apps); i++ {
		if apps[i].Status {
			applications = append(applications, apps[i])
		}
	}

	return
}
