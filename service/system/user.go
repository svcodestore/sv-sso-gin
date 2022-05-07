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

	user, err = global.UserMgr.GetFromLoginID(u.LoginID)
	if c.PasswordVerify(u.Password, user.Password) {
		return user, err
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

	user, err = global.UserMgr.GetFromID(u.ID)
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

	user, err = global.UserMgr.GetFromID(u.ID)
	return
}

func (s *UserService) DeleteUserWithId(u model.Users) bool {
	db := global.DB.Where("id = ?", u.ID).Delete(u)
	return db.RowsAffected == 1
}

func (s *UserService) UpdateUser(u model.UsersToSave) (user model.Users, err error) {
	id := u.ID
	u.ID = ""
	db := global.UserMgr.Where("id = ?", id).Updates(u)

	if db.RowsAffected == 1 {
		user, err = global.UserMgr.GetFromID(id)
		return
	}

	err = db.Error
	return
}

func (s *UserService) UpdateUserStatus(status bool, id, updatedBy string) (user model.Users, err error) {
	err = global.UserMgr.Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error
	if err != nil {
		return
	}
	user, err = global.UserMgr.GetFromID(id)
	return
}

func (s *UserService) AllUser() (users []*model.Users, err error) {
	// problem
	global.UserMgr = model.UsersMgr(utils.Gorm())
	users, err = global.UserMgr.Gets()
	return
}

func (s *UserService) UserWithId(u model.Users) (user model.Users, err error) {
	user, err = global.UserMgr.GetFromID(u.ID)
	return
}
