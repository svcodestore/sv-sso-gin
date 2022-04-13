package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ApplicationsMgr struct {
	*_BaseMgr
}

// ApplicationsMgr open func
func ApplicationsMgr(db *gorm.DB) *_ApplicationsMgr {
	if db == nil {
		panic(fmt.Errorf("ApplicationsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ApplicationsMgr{_BaseMgr: &_BaseMgr{DB: db.Model(Applications{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ApplicationsMgr) GetTableName() string {
	return "applications"
}

// Get 获取
func (obj *_ApplicationsMgr) Get() (result Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
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
func (obj *_ApplicationsMgr) Gets() (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// WithID id获取
func (obj *_ApplicationsMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取
func (obj *_ApplicationsMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取
func (obj *_ApplicationsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithInternalURL internal_url获取
func (obj *_ApplicationsMgr) WithInternalURL(internalURL string) Option {
	return optionFunc(func(o *options) { o.query["internal_url"] = internalURL })
}

// WithHomepageURL homepage_url获取
func (obj *_ApplicationsMgr) WithHomepageURL(homepageURL string) Option {
	return optionFunc(func(o *options) { o.query["homepage_url"] = homepageURL })
}

// WithStatus status获取
func (obj *_ApplicationsMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithClientID client_id获取
func (obj *_ApplicationsMgr) WithClientID(clientID string) Option {
	return optionFunc(func(o *options) { o.query["client_id"] = clientID })
}

// WithClientSecret client_secret获取
func (obj *_ApplicationsMgr) WithClientSecret(clientSecret string) Option {
	return optionFunc(func(o *options) { o.query["client_secret"] = clientSecret })
}

// WithRedirectURIs redirect_uris获取
func (obj *_ApplicationsMgr) WithRedirectURIs(redirectURIs string) Option {
	return optionFunc(func(o *options) { o.query["redirect_uris"] = redirectURIs })
}

// WithTokenFormat token_format获取
func (obj *_ApplicationsMgr) WithTokenFormat(tokenFormat string) Option {
	return optionFunc(func(o *options) { o.query["token_format"] = tokenFormat })
}

// WithCreatedAt created_at获取
func (obj *_ApplicationsMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_ApplicationsMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_ApplicationsMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_ApplicationsMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_ApplicationsMgr) GetByOption(opts ...Option) (result Applications, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
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
func (obj *_ApplicationsMgr) GetByOptions(opts ...Option) (results []*Applications, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetFromID 通过id获取内容
func (obj *_ApplicationsMgr) GetFromID(id string) (result Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
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

// GetBatchFromID 批量查找
func (obj *_ApplicationsMgr) GetBatchFromID(ids []string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetFromCode 通过code获取内容
func (obj *_ApplicationsMgr) GetFromCode(code string) (result Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` = ?", code).Find(&result).Error
	if err == nil && obj.isRelated {
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

// GetBatchFromCode 批量查找
func (obj *_ApplicationsMgr) GetBatchFromCode(codes []string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` IN (?)", codes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetFromName 通过name获取内容
func (obj *_ApplicationsMgr) GetFromName(name string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetBatchFromName 批量查找
func (obj *_ApplicationsMgr) GetBatchFromName(names []string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetFromInternalURL 通过internal_url获取内容
func (obj *_ApplicationsMgr) GetFromInternalURL(internalURL string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`internal_url` = ?", internalURL).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetBatchFromInternalURL 批量查找
func (obj *_ApplicationsMgr) GetBatchFromInternalURL(internalURLs []string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`internal_url` IN (?)", internalURLs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetFromHomepageURL 通过homepage_url获取内容
func (obj *_ApplicationsMgr) GetFromHomepageURL(homepageURL string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`homepage_url` = ?", homepageURL).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetBatchFromHomepageURL 批量查找
func (obj *_ApplicationsMgr) GetBatchFromHomepageURL(homepageURLs []string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`homepage_url` IN (?)", homepageURLs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_ApplicationsMgr) GetFromStatus(status bool) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_ApplicationsMgr) GetBatchFromStatus(statuss []bool) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetFromClientID 通过client_id获取内容
func (obj *_ApplicationsMgr) GetFromClientID(clientID string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`client_id` = ?", clientID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetBatchFromClientID 批量查找
func (obj *_ApplicationsMgr) GetBatchFromClientID(clientIDs []string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`client_id` IN (?)", clientIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetFromClientSecret 通过client_secret获取内容
func (obj *_ApplicationsMgr) GetFromClientSecret(clientSecret string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`client_secret` = ?", clientSecret).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetBatchFromClientSecret 批量查找
func (obj *_ApplicationsMgr) GetBatchFromClientSecret(clientSecrets []string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`client_secret` IN (?)", clientSecrets).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetFromRedirectURIs 通过redirect_uris获取内容
func (obj *_ApplicationsMgr) GetFromRedirectURIs(redirectURIs string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`redirect_uris` = ?", redirectURIs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetBatchFromRedirectURIs 批量查找
func (obj *_ApplicationsMgr) GetBatchFromRedirectURIs(redirectURIss []string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`redirect_uris` IN (?)", redirectURIss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetFromTokenFormat 通过token_format获取内容
func (obj *_ApplicationsMgr) GetFromTokenFormat(tokenFormat string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`token_format` = ?", tokenFormat).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetBatchFromTokenFormat 批量查找
func (obj *_ApplicationsMgr) GetBatchFromTokenFormat(tokenFormats []string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`token_format` IN (?)", tokenFormats).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_ApplicationsMgr) GetFromCreatedAt(createdAt time.Time) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_ApplicationsMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_ApplicationsMgr) GetFromCreatedBy(createdBy string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_ApplicationsMgr) GetBatchFromCreatedBy(createdBys []string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_ApplicationsMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_ApplicationsMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_ApplicationsMgr) GetFromUpdatedBy(updatedBy string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_ApplicationsMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_ApplicationsMgr) FetchByPrimaryKey(id string) (result Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
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

// FetchUniqueByApplicationsCodeUIndex primary or index 获取唯一内容
func (obj *_ApplicationsMgr) FetchUniqueByApplicationsCodeUIndex(code string) (result Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` = ?", code).Find(&result).Error
	if err == nil && obj.isRelated {
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

// FetchIndexByApplicationsFkCreatedBy  获取多个内容
func (obj *_ApplicationsMgr) FetchIndexByApplicationsFkCreatedBy(createdBy string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// FetchIndexByApplicationsFkUpdatedBy  获取多个内容
func (obj *_ApplicationsMgr) FetchIndexByApplicationsFkUpdatedBy(updatedBy string) (results []*Applications, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
