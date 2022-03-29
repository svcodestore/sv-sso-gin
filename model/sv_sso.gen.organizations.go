package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _OrganizationsMgr struct {
	*_BaseMgr
}

// OrganizationsMgr open func
func OrganizationsMgr(db *gorm.DB) *_OrganizationsMgr {
	if db == nil {
		panic(fmt.Errorf("OrganizationsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrganizationsMgr{_BaseMgr: &_BaseMgr{DB: db.Model(Organizations{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrganizationsMgr) GetTableName() string {
	return "organizations"
}

// Get 获取
func (obj *_OrganizationsMgr) Get() (result Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrganizationsMgr) Gets() (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_OrganizationsMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取
func (obj *_OrganizationsMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取
func (obj *_OrganizationsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithStatus status获取
func (obj *_OrganizationsMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_OrganizationsMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_OrganizationsMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_OrganizationsMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_OrganizationsMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_OrganizationsMgr) GetByOption(opts ...Option) (result Organizations, err error) {
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
func (obj *_OrganizationsMgr) GetByOptions(opts ...Option) (results []*Organizations, err error) {
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
func (obj *_OrganizationsMgr) GetFromID(id string) (result Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_OrganizationsMgr) GetBatchFromID(ids []string) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容
func (obj *_OrganizationsMgr) GetFromCode(code string) (result Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` = ?", code).Find(&result).Error

	return
}

// GetBatchFromCode 批量查找
func (obj *_OrganizationsMgr) GetBatchFromCode(codes []string) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_OrganizationsMgr) GetFromName(name string) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_OrganizationsMgr) GetBatchFromName(names []string) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_OrganizationsMgr) GetFromStatus(status bool) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_OrganizationsMgr) GetBatchFromStatus(statuss []bool) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_OrganizationsMgr) GetFromCreatedAt(createdAt time.Time) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_OrganizationsMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_OrganizationsMgr) GetFromCreatedBy(createdBy string) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_OrganizationsMgr) GetBatchFromCreatedBy(createdBys []string) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_OrganizationsMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_OrganizationsMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容
func (obj *_OrganizationsMgr) GetFromUpdatedBy(updatedBy string) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量查找
func (obj *_OrganizationsMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_OrganizationsMgr) FetchByPrimaryKey(id string) (result Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByOrganizationsCodeUIndex primary or index 获取唯一内容
func (obj *_OrganizationsMgr) FetchUniqueByOrganizationsCodeUIndex(code string) (result Organizations, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` = ?", code).Find(&result).Error

	return
}
