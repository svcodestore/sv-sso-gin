package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UsersMgr struct {
	*_BaseMgr
}

// UsersMgr open func
func UsersMgr(db *gorm.DB) *_UsersMgr {
	if db == nil {
		panic(fmt.Errorf("UsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Model(Users{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UsersMgr) GetTableName() string {
	return "users"
}

// Get 获取
func (obj *_UsersMgr) Get() (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UsersMgr) Gets() (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UsersMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUUID uuid获取
func (obj *_UsersMgr) WithUUID(uuid []byte) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithLoginID login_id获取
func (obj *_UsersMgr) WithLoginID(loginID string) Option {
	return optionFunc(func(o *options) { o.query["login_id"] = loginID })
}

// WithPassword password获取
func (obj *_UsersMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithName name获取
func (obj *_UsersMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAlias alias获取
func (obj *_UsersMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

// WithPhone phone获取
func (obj *_UsersMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithEmail email获取
func (obj *_UsersMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithLang lang获取
func (obj *_UsersMgr) WithLang(lang string) Option {
	return optionFunc(func(o *options) { o.query["lang"] = lang })
}

// WithStatus status获取
func (obj *_UsersMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_UsersMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_UsersMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_UsersMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_UsersMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_UsersMgr) GetByOption(opts ...Option) (result Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UsersMgr) GetByOptions(opts ...Option) (results []*Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_UsersMgr) GetFromID(id string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UsersMgr) GetBatchFromID(ids []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUUID 通过uuid获取内容
func (obj *_UsersMgr) GetFromUUID(uuid []byte) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`uuid` = ?", uuid).Find(&results).Error

	return
}

// GetBatchFromUUID 批量查找
func (obj *_UsersMgr) GetBatchFromUUID(uuids [][]byte) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`uuid` IN (?)", uuids).Find(&results).Error

	return
}

// GetFromLoginID 通过login_id获取内容
func (obj *_UsersMgr) GetFromLoginID(loginID string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`login_id` = ?", loginID).Find(&result).Error

	return
}

// GetBatchFromLoginID 批量查找
func (obj *_UsersMgr) GetBatchFromLoginID(loginIDs []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`login_id` IN (?)", loginIDs).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_UsersMgr) GetFromPassword(password string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_UsersMgr) GetBatchFromPassword(passwords []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_UsersMgr) GetFromName(name string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_UsersMgr) GetBatchFromName(names []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAlias 通过alias获取内容
func (obj *_UsersMgr) GetFromAlias(alias string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`alias` = ?", alias).Find(&results).Error

	return
}

// GetBatchFromAlias 批量查找
func (obj *_UsersMgr) GetBatchFromAlias(aliass []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`alias` IN (?)", aliass).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容
func (obj *_UsersMgr) GetFromPhone(phone string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找
func (obj *_UsersMgr) GetBatchFromPhone(phones []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_UsersMgr) GetFromEmail(email string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找
func (obj *_UsersMgr) GetBatchFromEmail(emails []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromLang 通过lang获取内容
func (obj *_UsersMgr) GetFromLang(lang string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`lang` = ?", lang).Find(&results).Error

	return
}

// GetBatchFromLang 批量查找
func (obj *_UsersMgr) GetBatchFromLang(langs []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`lang` IN (?)", langs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_UsersMgr) GetFromStatus(status bool) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_UsersMgr) GetBatchFromStatus(statuss []bool) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_UsersMgr) GetFromCreatedAt(createdAt time.Time) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_UsersMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_UsersMgr) GetFromCreatedBy(createdBy string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_UsersMgr) GetBatchFromCreatedBy(createdBys []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_UsersMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_UsersMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容
func (obj *_UsersMgr) GetFromUpdatedBy(updatedBy string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量查找
func (obj *_UsersMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UsersMgr) FetchByPrimaryKey(id string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByUserLoginIDUIndex primary or index 获取唯一内容
func (obj *_UsersMgr) FetchUniqueByUserLoginIDUIndex(loginID string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`login_id` = ?", loginID).Find(&result).Error

	return
}

// FetchIndexByUserUUIDIndex  获取多个内容
func (obj *_UsersMgr) FetchIndexByUserUUIDIndex(uuid []byte) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`uuid` = ?", uuid).Find(&results).Error

	return
}
