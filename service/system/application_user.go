package system

import (
	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model"
)

type ApplicationUserService struct {
}

func (s *ApplicationUserService) CreateApplicationUser(o *model.ApplicationUser) (applicationUser model.ApplicationUser, err error) {
	result := model.ApplicationUserMgr(global.DB).Create(o)
	if result.Error != nil {
		err = result.Error
		return
	}

	applicationUser, err = s.ApplicationUserWithId(o.ApplicationID, o.UserID)
	return
}

func (s *ApplicationUserService) DeleteApplicationUserWithId(o *model.ApplicationUser) (isDeleted bool) {
	db := global.DB.Where("application_id = ? and user_id = ?", o.ApplicationID, o.UserID).Delete(o)
	isDeleted = db.RowsAffected == 1
	return
}

func (s *ApplicationUserService) UpdateApplicationUserWithId(o *model.ApplicationUser) (applicationUser model.ApplicationUser, err error) {
	db := global.DB.Where("application_id = ? and user_id = ?", o.ApplicationID, o.UserID).Updates(o)
	err = db.Error
	return
}

func (s *ApplicationUserService) AllApplicationUser() (applicationUsers []*model.ApplicationUser, err error) {
	applicationUsers, err = model.ApplicationUserMgr(global.DB).Gets()
	return
}

func (s *ApplicationUserService) ApplicationUserWithId(applicationId, userId string) (applicationUser model.ApplicationUser, err error) {
	applicationUser, err = model.ApplicationUserMgr(global.DB).FetchByPrimaryKey(applicationId, userId)
	return
}

func (s *ApplicationUserService) AvailableApplicationUsers() (results []*model.ApplicationUser, err error) {
	applicationUsers, err := s.AllApplicationUser()
	if err != nil {
		return
	}

	for i := 0; i < len(applicationUsers); i++ {
		if applicationUsers[i].Status == 1 {
			results = append(results, applicationUsers[i])
		}
	}

	return
}
