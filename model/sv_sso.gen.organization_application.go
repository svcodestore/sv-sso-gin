package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrganizationApplicationMgr struct {
	*_BaseMgr
}

// OrganizationApplicationMgr open func
func OrganizationApplicationMgr(db *gorm.DB) *_OrganizationApplicationMgr {
	if db == nil {
		panic(fmt.Errorf("OrganizationApplicationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrganizationApplicationMgr{_BaseMgr: &_BaseMgr{DB: db.Model(OrganizationApplication{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrganizationApplicationMgr) GetTableName() string {
	return "organization_application"
}

// Get 获取
func (obj *_OrganizationApplicationMgr) Get() (result OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("organizations").Where("id = ?", result.OrganizationID).Find(&result.Organizations).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("applications").Where("id = ?", result.ApplicationID).Find(&result.Applications).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("users").Where("id = ?", result.CreatedBy).Find(&result.CreatedByUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("users").Where("id = ?", result.UpdatedBy).Find(&result.UpdatedByUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_OrganizationApplicationMgr) Gets() (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithOrganizationID organization_id获取
func (obj *_OrganizationApplicationMgr) WithOrganizationID(organizationID int64) Option {
	return optionFunc(func(o *options) { o.query["organization_id"] = organizationID })
}

// WithApplicationID application_id获取
func (obj *_OrganizationApplicationMgr) WithApplicationID(applicationID int64) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithStatus status获取
func (obj *_OrganizationApplicationMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_OrganizationApplicationMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_OrganizationApplicationMgr) WithCreatedBy(createdBy int64) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_OrganizationApplicationMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_OrganizationApplicationMgr) WithUpdatedBy(updatedBy int64) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_OrganizationApplicationMgr) GetByOption(opts ...Option) (result OrganizationApplication, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("organizations").Where("id = ?", result.OrganizationID).Find(&result.Organizations).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("applications").Where("id = ?", result.ApplicationID).Find(&result.Applications).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("users").Where("id = ?", result.CreatedBy).Find(&result.CreatedByUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("users").Where("id = ?", result.UpdatedBy).Find(&result.UpdatedByUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_OrganizationApplicationMgr) GetByOptions(opts ...Option) (results []*OrganizationApplication, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromOrganizationID 通过organization_id获取内容
func (obj *_OrganizationApplicationMgr) GetFromOrganizationID(organizationID int64) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`organization_id` = ?", organizationID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOrganizationID 批量查找
func (obj *_OrganizationApplicationMgr) GetBatchFromOrganizationID(organizationIDs []int64) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`organization_id` IN (?)", organizationIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicationID 通过application_id获取内容
func (obj *_OrganizationApplicationMgr) GetFromApplicationID(applicationID int64) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ?", applicationID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicationID 批量查找
func (obj *_OrganizationApplicationMgr) GetBatchFromApplicationID(applicationIDs []int64) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` IN (?)", applicationIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容
func (obj *_OrganizationApplicationMgr) GetFromStatus(status bool) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量查找
func (obj *_OrganizationApplicationMgr) GetBatchFromStatus(statuss []bool) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_OrganizationApplicationMgr) GetFromCreatedAt(createdAt time.Time) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_OrganizationApplicationMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_OrganizationApplicationMgr) GetFromCreatedBy(createdBy int64) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_OrganizationApplicationMgr) GetBatchFromCreatedBy(createdBys []int64) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_OrganizationApplicationMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_OrganizationApplicationMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedBy 通过updated_by获取内容
func (obj *_OrganizationApplicationMgr) GetFromUpdatedBy(updatedBy int64) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedBy 批量查找
func (obj *_OrganizationApplicationMgr) GetBatchFromUpdatedBy(updatedBys []int64) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
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
func (obj *_OrganizationApplicationMgr) FetchByPrimaryKey(organizationID int64, applicationID int64) (result OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`organization_id` = ? AND `application_id` = ?", organizationID, applicationID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("organizations").Where("id = ?", result.OrganizationID).Find(&result.Organizations).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("applications").Where("id = ?", result.ApplicationID).Find(&result.Applications).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("users").Where("id = ?", result.CreatedBy).Find(&result.CreatedByUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("users").Where("id = ?", result.UpdatedBy).Find(&result.UpdatedByUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByOrganizationApplicationOrganizationIDIndex  获取多个内容
func (obj *_OrganizationApplicationMgr) FetchIndexByOrganizationApplicationOrganizationIDIndex(organizationID int64) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`organization_id` = ?", organizationID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByOrganizationApplicationApplicationIDIndex  获取多个内容
func (obj *_OrganizationApplicationMgr) FetchIndexByOrganizationApplicationApplicationIDIndex(applicationID int64) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ?", applicationID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByOrganizationApplicationFkCreatedBy  获取多个内容
func (obj *_OrganizationApplicationMgr) FetchIndexByOrganizationApplicationFkCreatedBy(createdBy int64) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByOrganizationApplicationFkUpdatedBy  获取多个内容
func (obj *_OrganizationApplicationMgr) FetchIndexByOrganizationApplicationFkUpdatedBy(updatedBy int64) (results []*OrganizationApplication, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("organizations").Where("id = ?", results[i].OrganizationID).Find(&results[i].Organizations).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("applications").Where("id = ?", results[i].ApplicationID).Find(&results[i].Applications).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].CreatedBy).Find(&results[i].CreatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UpdatedBy).Find(&results[i].UpdatedByUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
