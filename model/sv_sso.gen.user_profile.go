package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _UserProfileMgr struct {
	*_BaseMgr
}

// UserProfileMgr open func
func UserProfileMgr(db *gorm.DB) *_UserProfileMgr {
	if db == nil {
		panic(fmt.Errorf("UserProfileMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserProfileMgr{_BaseMgr: &_BaseMgr{DB: db.Model(UserProfile{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserProfileMgr) GetTableName() string {
	return "user_profile"
}

// Get 获取
func (obj *_UserProfileMgr) Get() (result UserProfile, err error) {
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
func (obj *_UserProfileMgr) Gets() (results []*UserProfile, err error) {
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
func (obj *_UserProfileMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithAvatar avatar获取
func (obj *_UserProfileMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithWechatUId wechat_uid获取
func (obj *_UserProfileMgr) WithWechatUId(wechatUId string) Option {
	return optionFunc(func(o *options) { o.query["wechat_uid"] = wechatUId })
}

// WithWechatName wechat_name获取
func (obj *_UserProfileMgr) WithWechatName(wechatName string) Option {
	return optionFunc(func(o *options) { o.query["wechat_name"] = wechatName })
}

// WithQqUId qq_uid获取
func (obj *_UserProfileMgr) WithQqUId(qqUId string) Option {
	return optionFunc(func(o *options) { o.query["qq_uid"] = qqUId })
}

// WithQqName qq_name获取
func (obj *_UserProfileMgr) WithQqName(qqName string) Option {
	return optionFunc(func(o *options) { o.query["qq_name"] = qqName })
}

// WithSkypeUId skype_uid获取
func (obj *_UserProfileMgr) WithSkypeUId(skypeUId string) Option {
	return optionFunc(func(o *options) { o.query["skype_uid"] = skypeUId })
}

// WithSkypeName skype_name获取
func (obj *_UserProfileMgr) WithSkypeName(skypeName string) Option {
	return optionFunc(func(o *options) { o.query["skype_name"] = skypeName })
}

// WithGoogleUId google_uid获取
func (obj *_UserProfileMgr) WithGoogleUId(googleUId string) Option {
	return optionFunc(func(o *options) { o.query["google_uid"] = googleUId })
}

// WithGoogleName google_name获取
func (obj *_UserProfileMgr) WithGoogleName(googleName string) Option {
	return optionFunc(func(o *options) { o.query["google_name"] = googleName })
}

// WithGender gender获取
func (obj *_UserProfileMgr) WithGender(gender string) Option {
	return optionFunc(func(o *options) { o.query["gender"] = gender })
}

// GetByOption 功能选项模式获取
func (obj *_UserProfileMgr) GetByOption(opts ...Option) (result UserProfile, err error) {
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
func (obj *_UserProfileMgr) GetByOptions(opts ...Option) (results []*UserProfile, err error) {
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
func (obj *_UserProfileMgr) GetFromUserID(userID string) (result UserProfile, err error) {
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
func (obj *_UserProfileMgr) GetBatchFromUserID(userIDs []string) (results []*UserProfile, err error) {
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

// GetFromAvatar 通过avatar获取内容
func (obj *_UserProfileMgr) GetFromAvatar(avatar string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`avatar` = ?", avatar).Find(&results).Error
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

// GetBatchFromAvatar 批量查找
func (obj *_UserProfileMgr) GetBatchFromAvatar(avatars []string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`avatar` IN (?)", avatars).Find(&results).Error
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

// GetFromWechatUId 通过wechat_uid获取内容
func (obj *_UserProfileMgr) GetFromWechatUId(wechatUId string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`wechat_uid` = ?", wechatUId).Find(&results).Error
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

// GetBatchFromWechatUId 批量查找
func (obj *_UserProfileMgr) GetBatchFromWechatUId(wechatUIds []string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`wechat_uid` IN (?)", wechatUIds).Find(&results).Error
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

// GetFromWechatName 通过wechat_name获取内容
func (obj *_UserProfileMgr) GetFromWechatName(wechatName string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`wechat_name` = ?", wechatName).Find(&results).Error
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

// GetBatchFromWechatName 批量查找
func (obj *_UserProfileMgr) GetBatchFromWechatName(wechatNames []string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`wechat_name` IN (?)", wechatNames).Find(&results).Error
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

// GetFromQqUId 通过qq_uid获取内容
func (obj *_UserProfileMgr) GetFromQqUId(qqUId string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`qq_uid` = ?", qqUId).Find(&results).Error
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

// GetBatchFromQqUId 批量查找
func (obj *_UserProfileMgr) GetBatchFromQqUId(qqUIds []string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`qq_uid` IN (?)", qqUIds).Find(&results).Error
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

// GetFromQqName 通过qq_name获取内容
func (obj *_UserProfileMgr) GetFromQqName(qqName string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`qq_name` = ?", qqName).Find(&results).Error
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

// GetBatchFromQqName 批量查找
func (obj *_UserProfileMgr) GetBatchFromQqName(qqNames []string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`qq_name` IN (?)", qqNames).Find(&results).Error
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

// GetFromSkypeUId 通过skype_uid获取内容
func (obj *_UserProfileMgr) GetFromSkypeUId(skypeUId string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`skype_uid` = ?", skypeUId).Find(&results).Error
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

// GetBatchFromSkypeUId 批量查找
func (obj *_UserProfileMgr) GetBatchFromSkypeUId(skypeUIds []string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`skype_uid` IN (?)", skypeUIds).Find(&results).Error
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

// GetFromSkypeName 通过skype_name获取内容
func (obj *_UserProfileMgr) GetFromSkypeName(skypeName string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`skype_name` = ?", skypeName).Find(&results).Error
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

// GetBatchFromSkypeName 批量查找
func (obj *_UserProfileMgr) GetBatchFromSkypeName(skypeNames []string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`skype_name` IN (?)", skypeNames).Find(&results).Error
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

// GetFromGoogleUId 通过google_uid获取内容
func (obj *_UserProfileMgr) GetFromGoogleUId(googleUId string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`google_uid` = ?", googleUId).Find(&results).Error
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

// GetBatchFromGoogleUId 批量查找
func (obj *_UserProfileMgr) GetBatchFromGoogleUId(googleUIds []string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`google_uid` IN (?)", googleUIds).Find(&results).Error
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

// GetFromGoogleName 通过google_name获取内容
func (obj *_UserProfileMgr) GetFromGoogleName(googleName string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`google_name` = ?", googleName).Find(&results).Error
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

// GetBatchFromGoogleName 批量查找
func (obj *_UserProfileMgr) GetBatchFromGoogleName(googleNames []string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`google_name` IN (?)", googleNames).Find(&results).Error
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

// GetFromGender 通过gender获取内容
func (obj *_UserProfileMgr) GetFromGender(gender string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`gender` = ?", gender).Find(&results).Error
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

// GetBatchFromGender 批量查找
func (obj *_UserProfileMgr) GetBatchFromGender(genders []string) (results []*UserProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`gender` IN (?)", genders).Find(&results).Error
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
func (obj *_UserProfileMgr) FetchByPrimaryKey(userID string) (result UserProfile, err error) {
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
