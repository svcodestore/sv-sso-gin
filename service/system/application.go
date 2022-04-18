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

	return s.ApplicationWithId(a)
}

func (s *ApplicationService) DeleteApplication(a *model.Applications) (isDeleted bool) {
	db := global.DB.Where("id = ?", a.ID).Delete(a)
	isDeleted = db.RowsAffected == 1
	return
}

func (s *ApplicationService) UpdateApplicationWithId(a *model.Applications) (application model.Applications, err error) {
	db := global.ApplicationMgr.Where("id = ?", a.ID).Updates(a)
	if db.RowsAffected == 1 {
		return s.ApplicationWithId(a)
	}
	err = db.Error
	return
}

func (s *ApplicationService) AllApplication() (applications []*model.Applications, err error) {
	applications, err = global.ApplicationMgr.Gets()
	return
}

func (s *ApplicationService) ApplicationWithId(a *model.Applications) (application model.Applications, err error) {
	application, err = global.ApplicationMgr.GetFromID(a.ID)
	return
}

func (s *ApplicationService) ApplicationWithClientId(a *model.Applications) (application model.Applications, err error) {
	application, err = global.ApplicationMgr.GetFromClientID(a.ClientID)
	return
}
