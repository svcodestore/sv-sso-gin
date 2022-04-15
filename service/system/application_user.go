package system

import (
	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model"
)

type ApplicationUserService struct {
}

func (s *ApplicationUserService) CreateApplicationUser(o *model.ApplicationUser) (applicationUser model.ApplicationUser, err error) {
	result := global.ApplicationUserMgr.Create(o)
	if result.Error != nil {
		err = result.Error
		return
	}

	applicationUser, err = s.ApplicationUserWithId(o)
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
	applicationUsers, err = global.ApplicationUserMgr.Gets()
	return
}

func (s *ApplicationUserService) ApplicationUserWithId(o *model.ApplicationUser) (applicationUser model.ApplicationUser, err error) {
	applicationUser, err = global.ApplicationUserMgr.FetchByPrimaryKey(o.ApplicationID, o.UserID)
	return
}
