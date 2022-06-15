package system

import (
	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model"
	"gorm.io/gorm/clause"
)

type UserLoginRecordService struct {
}

func (s *UserLoginRecordService) UserLoginWithUserId(userId string) (loginInfo model.UserLogin, err error) {
	loginInfo, err = model.UserLoginMgr(global.DB).GetFromUserID(userId)

	return
}

func (s *UserLoginRecordService) UpsertUserLogin(userId, applicationId, ip4, ip6, device string) (loginInfo model.UserLogin, err error) {
	if device == "" {
		device = "WEB"
	}
	result := global.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"application_id": applicationId,
			"ipv4":           ip4,
			"ipv6":           ip6,
			"device":         device,
		}),
	}).Create(&model.UserLogin{
		UserID:        userId,
		ApplicationID: applicationId,
		IPv4:          ip4,
		IPv6:          ip6,
		Device:        device,
	})

	if result.Error != nil {
		err = result.Error
		return
	}

	return s.UserLoginWithUserId(userId)
}
