package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _UserLoginMgr struct {
	*_BaseMgr
}

// UserLoginMgr open func
func UserLoginMgr(db *gorm.DB) *_UserLoginMgr {
	if db == nil {
		panic(fmt.Errorf("UserLoginMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserLoginMgr{_BaseMgr: &_BaseMgr{DB: db.Model(UserLogin{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserLoginMgr) GetTableName() string {
	return "user_login"
}

// Get 获取
func (obj *_UserLoginMgr) Get() (result UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_UserLoginMgr) Gets() (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取
func (obj *_UserLoginMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithApplicationID application_id获取
func (obj *_UserLoginMgr) WithApplicationID(applicationID string) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithIPv4 ipv4获取
func (obj *_UserLoginMgr) WithIPv4(ipv4 string) Option {
	return optionFunc(func(o *options) { o.query["ipv4"] = ipv4 })
}

// WithIPv6 ipv6获取
func (obj *_UserLoginMgr) WithIPv6(ipv6 string) Option {
	return optionFunc(func(o *options) { o.query["ipv6"] = ipv6 })
}

// WithTime time获取
func (obj *_UserLoginMgr) WithTime(time time.Time) Option {
	return optionFunc(func(o *options) { o.query["time"] = time })
}

// WithDevice device获取
func (obj *_UserLoginMgr) WithDevice(device string) Option {
	return optionFunc(func(o *options) { o.query["device"] = device })
}

// GetByOption 功能选项模式获取
func (obj *_UserLoginMgr) GetByOption(opts ...Option) (result UserLogin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserLoginMgr) GetByOptions(opts ...Option) (results []*UserLogin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromUserID 通过user_id获取内容
func (obj *_UserLoginMgr) GetFromUserID(userID string) (result UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` = ?", userID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromUserID 批量查找
func (obj *_UserLoginMgr) GetBatchFromUserID(userIDs []string) (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicationID 通过application_id获取内容
func (obj *_UserLoginMgr) GetFromApplicationID(applicationID string) (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ?", applicationID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicationID 批量查找
func (obj *_UserLoginMgr) GetBatchFromApplicationID(applicationIDs []string) (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` IN (?)", applicationIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIPv4 通过ipv4获取内容
func (obj *_UserLoginMgr) GetFromIPv4(ipv4 string) (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`ipv4` = ?", ipv4).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIPv4 批量查找
func (obj *_UserLoginMgr) GetBatchFromIPv4(ipv4s []string) (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`ipv4` IN (?)", ipv4s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIPv6 通过ipv6获取内容
func (obj *_UserLoginMgr) GetFromIPv6(ipv6 string) (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`ipv6` = ?", ipv6).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIPv6 批量查找
func (obj *_UserLoginMgr) GetBatchFromIPv6(ipv6s []string) (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`ipv6` IN (?)", ipv6s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTime 通过time获取内容
func (obj *_UserLoginMgr) GetFromTime(time time.Time) (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`time` = ?", time).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTime 批量查找
func (obj *_UserLoginMgr) GetBatchFromTime(times []time.Time) (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`time` IN (?)", times).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDevice 通过device获取内容
func (obj *_UserLoginMgr) GetFromDevice(device string) (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`device` = ?", device).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDevice 批量查找
func (obj *_UserLoginMgr) GetBatchFromDevice(devices []string) (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`device` IN (?)", devices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserLoginMgr) FetchByPrimaryKey(userID string) (result UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` = ?", userID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByUserLoginApplicationIDIndex  获取多个内容
func (obj *_UserLoginMgr) FetchIndexByUserLoginApplicationIDIndex(applicationID string) (results []*UserLogin, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ?", applicationID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
