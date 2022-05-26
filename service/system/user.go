package system

import (
	"errors"
	uuid "github.com/satori/go.uuid"

	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/utils"
)

type UserService struct {
}

func (s *UserService) AccountDecrypt(usr, pwd string) (u, p string) {
	var c CryptoService
	u, _ = c.AesDecrypt(usr)
	p, _ = c.AesDecrypt(pwd)
	return
}

func (s *UserService) Login(usr, pwd string) (model.Users, error) {
	username, password := s.AccountDecrypt(usr, pwd)
	u := model.Users{LoginID: username, Password: password}

	return s.DoLogin(u)
}

func (s *UserService) DoLogin(u model.Users) (user model.Users, err error) {
	var c CryptoService

	user, err = model.UsersMgr(global.DB).GetFromLoginID(u.LoginID)
	if c.PasswordVerify(u.Password, user.Password) {
		return
	}

	return user, errors.New("invalid password")
}

func (s *UserService) RegisterUser(u model.UsersToSave) (user model.Users, err error) {
	var c CryptoService
	username, password := s.AccountDecrypt(u.LoginID, u.Password)
	u.LoginID = username
	u.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
	u.UUID = uuid.NewV4().Bytes()
	p, _ := c.PasswordHash(password)
	u.Password = p
	result := global.DB.Create(u)
	if result.Error != nil {
		err = result.Error
		return
	}

	user, err = model.UsersMgr(global.DB).GetFromID(u.ID)
	return
}

func (s *UserService) CreateUser(u *model.UsersToSave) (user model.Users, err error) {
	var c CryptoService

	u.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
	u.UUID = uuid.NewV4().Bytes()
	p, _ := c.PasswordHash(u.Password)
	u.Password = p
	result := global.DB.Create(u)
	if result.Error != nil {
		err = result.Error
		return
	}

	user, err = model.UsersMgr(global.DB).GetFromID(u.ID)
	return
}

func (s *UserService) DeleteUserWithId(u model.Users) bool {
	db := global.DB.Where("id = ?", u.ID).Delete(u)
	return db.RowsAffected == 1
}

func (s *UserService) UpdateUser(u model.UsersToSave) (user model.Users, err error) {
	id := u.ID
	u.ID = ""
	if u.Password != "" {
		var c CryptoService
		p, _ := c.PasswordHash(u.Password)
		u.Password = p
	}
	db := model.UsersMgr(global.DB).Where("id = ?", id).Updates(u)

	if db.RowsAffected == 1 {
		user, err = model.UsersMgr(global.DB).GetFromID(id)
		return
	}

	err = db.Error
	return
}

func (s *UserService) UpdateUserStatus(status bool, id, updatedBy string) (user model.Users, err error) {
	err = model.UsersMgr(global.DB).Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error
	if err != nil {
		return
	}
	user, err = model.UsersMgr(global.DB).GetFromID(id)
	return
}

func (s *UserService) AllUser() (users []*model.Users, err error) {
	users, err = model.UsersMgr(global.DB).Gets()
	return
}

func (s *UserService) UserWithId(id string) (user model.Users, err error) {
	user, err = model.UsersMgr(global.DB).GetFromID(id)
	return
}

func (s *UserService) UserWithIdAndApplicationId(id, applicationId string) (user model.Users, err error) {
	users, err := model.ApplicationUserMgr(global.DB).GetFromApplicationID(applicationId)
	if err != nil {
		return
	}
	userCount := len(users)
	if userCount == 0 {
		err = errors.New("not registered")
		return
	}
	for i := 0; i < userCount; i++ {
		if users[i].UserID == id {
			user, err = userService.UserWithId(id)
		} else {
			err = errors.New("not registered")
			return
		}
	}
	return
}

func (s *UserService) UsersWithApplicationIds(applicationIds ...string) (users []*model.Users, err error) {
	maps, err := model.ApplicationUserMgr(global.DB).GetBatchFromApplicationID(applicationIds)
	if err != nil {
		return
	}
	mapCount := len(maps)
	if mapCount == 0 {
		return
	}
	var ids = make([]string, mapCount)
	for i := 0; i < mapCount; i++ {
		ids[i] = maps[i].UserID
	}
	users, err = model.UsersMgr(global.DB).GetBatchFromID(ids)
	return
}

func (s *UserService) AvailableUsers() (users []*model.Users, err error) {
	results, err := s.AllUser()
	if err != nil {
		return
	}

	for i := 0; i < len(results); i++ {
		if results[i].Status == 1 {
			users = append(users, results[i])
		}
	}

	return
}

func (s *UserService) IsAvailableUsers(userIds ...string) (users []*model.Users, err error) {
	u, err := model.UsersMgr(global.DB).GetBatchFromID(userIds)
	if err != nil {
		return
	}

	for i := 0; i < len(u); i++ {
		if u[i].Status == 1 {
			users = append(users, u[i])
		}
	}

	return
}
