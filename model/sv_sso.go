package model

import (
	"time"
)

// ApplicationUser [...]
type ApplicationUser struct {
	ApplicationID string              `gorm:"primaryKey;index:application_user_application_id_index;column:application_id;type:bigint unsigned;not null" json:"applicationId"`
	Applications  Applications        `gorm:"joinForeignKey:application_id;foreignKey:ApplicationID;reference:ApplicationID" json:"applicationsList"`
	UserID        string              `gorm:"primaryKey;index:application_user_user_id_index;column:user_id;type:bigint unsigned;not null" json:"userId"`
	Users         UsersWithoutModInfo `gorm:"joinForeignKey:user_id;foreignKey:UserID;reference:UserID" json:"usersList"`
	Status        uint8               `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"`
	CreatedAt     time.Time           `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy     string              `gorm:"column:created_by;type:bigint;not null" json:"createdBy"`
	CreatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:CreatedBy;reference:CreatedBy" json:"createdByUser"`
	UpdatedAt     time.Time           `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy     string              `gorm:"column:updated_by;type:bigint;not null" json:"updatedBy"`
	UpdatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:UpdatedBy;reference:UpdatedBy" json:"updatedByUser"`
}

// TableName get sql table name.获取数据库表名
func (m *ApplicationUser) TableName() string {
	return "application_user"
}

// ApplicationUserColumns get sql column name.获取数据库列名
var ApplicationUserColumns = struct {
	ApplicationID string
	UserID        string
	Status        string
	CreatedAt     string
	CreatedBy     string
	UpdatedAt     string
	UpdatedBy     string
}{
	ApplicationID: "application_id",
	UserID:        "user_id",
	Status:        "status",
	CreatedAt:     "created_at",
	CreatedBy:     "created_by",
	UpdatedAt:     "updated_at",
	UpdatedBy:     "updated_by",
}

// Applications [...]
type Applications struct {
	ID            string              `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	Code          string              `gorm:"unique;column:code;type:varchar(64);not null" json:"code"`
	Name          string              `gorm:"unique;column:name;type:varchar(255);not null" json:"name"`
	InternalURL   string              `gorm:"column:internal_url;type:varchar(255)" json:"internalUrl"`
	HomepageURL   string              `gorm:"column:homepage_url;type:varchar(255)" json:"homepageUrl"`
	Status        uint8               `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"`
	ClientID      string              `gorm:"unique;column:client_id;type:varchar(255);not null" json:"clientId"`
	ClientSecret  string              `gorm:"column:client_secret;type:varchar(255)" json:"clientSecret"`
	RedirectURIs  string              `gorm:"column:redirect_uris;type:varchar(255);not null" json:"redirectUris"`
	LoginURIs     string              `gorm:"column:login_uris;type:varchar(255);not null" json:"loginUris"`
	TokenFormat   string              `gorm:"column:token_format;type:varchar(100);default:JWT" json:"tokenFormat"`
	CreatedAt     time.Time           `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy     string              `gorm:"index:applications_fk_created_by;column:created_by;type:bigint unsigned;not null" json:"createdBy"`
	CreatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:CreatedBy;reference:CreatedBy" json:"createdByUser"`
	UpdatedAt     time.Time           `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy     string              `gorm:"column:updated_by;type:bigint;not null" json:"updatedBy"`
	UpdatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:UpdatedBy;reference:UpdatedBy" json:"updatedByUser"`
}

// TableName get sql table name.获取数据库表名
func (m *Applications) TableName() string {
	return "applications"
}

// ApplicationsColumns get sql column name.获取数据库列名
var ApplicationsColumns = struct {
	ID           string
	Code         string
	Name         string
	InternalURL  string
	HomepageURL  string
	Status       string
	ClientID     string
	ClientSecret string
	RedirectURIs string
	LoginURIs    string
	TokenFormat  string
	CreatedAt    string
	CreatedBy    string
	UpdatedAt    string
	UpdatedBy    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	InternalURL:  "internal_url",
	HomepageURL:  "homepage_url",
	Status:       "status",
	ClientID:     "client_id",
	ClientSecret: "client_secret",
	RedirectURIs: "redirect_uris",
	LoginURIs:    "login_uris",
	TokenFormat:  "token_format",
	CreatedAt:    "created_at",
	CreatedBy:    "created_by",
	UpdatedAt:    "updated_at",
	UpdatedBy:    "updated_by",
}

// OrganizationApplication [...]
type OrganizationApplication struct {
	OrganizationID string              `gorm:"primaryKey;index:organization_application_organization_id_index;column:organization_id;type:bigint unsigned;not null" json:"organizationId"`
	Organizations  Organizations       `gorm:"joinForeignKey:organization_id;foreignKey:OrganizationID;reference:OrganizationID" json:"organizationsList"`
	ApplicationID  string              `gorm:"primaryKey;index:organization_application_application_id_index;column:application_id;type:bigint unsigned;not null" json:"applicationId"`
	Applications   Applications        `gorm:"joinForeignKey:application_id;foreignKey:ApplicationID;reference:ApplicationID" json:"applicationsList"`
	Status         uint8               `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"`
	CreatedAt      time.Time           `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy      string              `gorm:"index:organization_application_fk_created_by;column:created_by;type:bigint unsigned;not null" json:"createdBy"`
	CreatedByUser  UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:CreatedBy;reference:CreatedBy" json:"createdByUser"`
	UpdatedAt      time.Time           `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy      string              `gorm:"column:updated_by;type:bigint;not null" json:"updatedBy"`
	UpdatedByUser  UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:UpdatedBy;reference:UpdatedBy" json:"updatedByUser"`
}

// TableName get sql table name.获取数据库表名
func (m *OrganizationApplication) TableName() string {
	return "organization_application"
}

// OrganizationApplicationColumns get sql column name.获取数据库列名
var OrganizationApplicationColumns = struct {
	OrganizationID string
	ApplicationID  string
	Status         string
	CreatedAt      string
	CreatedBy      string
	UpdatedAt      string
	UpdatedBy      string
}{
	OrganizationID: "organization_id",
	ApplicationID:  "application_id",
	Status:         "status",
	CreatedAt:      "created_at",
	CreatedBy:      "created_by",
	UpdatedAt:      "updated_at",
	UpdatedBy:      "updated_by",
}

// Organizations [...]
type Organizations struct {
	ID            string              `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	Code          string              `gorm:"unique;column:code;type:varchar(64);not null" json:"code"`
	Name          string              `gorm:"unique;column:name;type:varchar(255);not null" json:"name"`
	Status        uint8               `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"`
	CreatedAt     time.Time           `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy     string              `gorm:"index:organizations_fk_created_by;column:created_by;type:bigint unsigned;not null" json:"createdBy"`
	CreatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:CreatedBy;reference:CreatedBy" json:"createdByUser"`
	UpdatedAt     time.Time           `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy     string              `gorm:"column:updated_by;type:bigint;not null" json:"updatedBy"`
	UpdatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:UpdatedBy;reference:UpdatedBy" json:"updatedByUser"`
}

// TableName get sql table name.获取数据库表名
func (m *Organizations) TableName() string {
	return "organizations"
}

// OrganizationsColumns get sql column name.获取数据库列名
var OrganizationsColumns = struct {
	ID        string
	Code      string
	Name      string
	Status    string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
}{
	ID:        "id",
	Code:      "code",
	Name:      "name",
	Status:    "status",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
}

// UserLogin [...]
type UserLogin struct {
	UserID        string              `gorm:"primaryKey;column:user_id;type:bigint unsigned;not null" json:"userId"`
	Users         UsersWithoutModInfo `gorm:"joinForeignKey:user_id;foreignKey:UserID" json:"usersList"`
	ApplicationID string              `gorm:"index:user_login_application_id_index;column:application_id;type:bigint unsigned;not null" json:"applicationId"`
	Applications  Applications        `gorm:"joinForeignKey:application_id;foreignKey:ApplicationID" json:"applicationsList"`
	IPv4          string              `gorm:"column:ipv4;type:varchar(15)" json:"ipv4"`
	IPv6          string              `gorm:"column:ipv6;type:varchar(46)" json:"ipv6"`
	Time          time.Time           `gorm:"column:time;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"time"`
	Device        string              `gorm:"column:device;type:varchar(1024);default:WEB" json:"device"`
}

// TableName get sql table name.获取数据库表名
func (m *UserLogin) TableName() string {
	return "user_login"
}

// UserLoginColumns get sql column name.获取数据库列名
var UserLoginColumns = struct {
	UserID        string
	ApplicationID string
	IPv4          string
	IPv6          string
	Time          string
	Device        string
}{
	UserID:        "user_id",
	ApplicationID: "application_id",
	IPv4:          "ipv4",
	IPv6:          "ipv6",
	Time:          "time",
	Device:        "device",
}

// UserProfile [...]
type UserProfile struct {
	UserID     string              `gorm:"primaryKey;column:user_id;type:bigint unsigned;not null" json:"userId"`
	Users      UsersWithoutModInfo `gorm:"joinForeignKey:user_id;foreignKey:UserID" json:"usersList"`
	WechatUId  string              `gorm:"column:wechat_uid;type:varchar(255)" json:"wechatUid"`
	WechatName string              `gorm:"column:wechat_name;type:varchar(64)" json:"wechatName"`
	QqUId      string              `gorm:"column:qq_uid;type:varchar(255)" json:"qqUid"`
	QqName     string              `gorm:"column:qq_name;type:varchar(64)" json:"qqName"`
	SkypeUId   string              `gorm:"column:skype_uid;type:varchar(255)" json:"skypeUid"`
	SkypeName  string              `gorm:"column:skype_name;type:varchar(64)" json:"skypeName"`
	GoogleUId  string              `gorm:"column:google_uid;type:varchar(255)" json:"googleUid"`
	GoogleName string              `gorm:"column:google_name;type:varchar(64)" json:"googleName"`
}

// TableName get sql table name.获取数据库表名
func (m *UserProfile) TableName() string {
	return "user_profile"
}

// UserProfileColumns get sql column name.获取数据库列名
var UserProfileColumns = struct {
	UserID     string
	WechatUId  string
	WechatName string
	QqUId      string
	QqName     string
	SkypeUId   string
	SkypeName  string
	GoogleUId  string
	GoogleName string
}{
	UserID:     "user_id",
	WechatUId:  "wechat_uid",
	WechatName: "wechat_name",
	QqUId:      "qq_uid",
	QqName:     "qq_name",
	SkypeUId:   "skype_uid",
	SkypeName:  "skype_name",
	GoogleUId:  "google_uid",
	GoogleName: "google_name",
}

// Users [...]
// type Users struct {
// 	ID        string    `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
// 	UUID      []byte    `gorm:"unique;index:user_uuid_index;column:uuid;type:binary(16);not null" json:"uuid"`
// 	LoginID   string    `gorm:"unique;column:login_id;type:varchar(16);not null" json:"loginId"`
// 	Password  string    `gorm:"column:password;type:varchar(1024)" json:"password"`
// 	Name      string    `gorm:"column:name;type:varchar(32)" json:"name"`
// 	Alias     string    `gorm:"column:alias;type:varchar(32)" json:"alias"`
// 	Phone     string    `gorm:"column:phone;type:varchar(16)" json:"phone"`
// 	Email     string    `gorm:"column:email;type:varchar(1024)" json:"email"`
// 	Lang      string    `gorm:"column:lang;type:char(5);default:zh_CN" json:"lang"`
// 	Status    uint8     `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"`
// 	CreatedAt time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
// 	CreatedBy string    `gorm:"index:users_fk_created_by;column:created_by;type:bigint unsigned;not null" json:"createdBy"`
// 	Users     Users     `gorm:"joinForeignKey:created_by;foreignKey:id" json:"usersList"`
// 	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
// 	UpdatedBy string    `gorm:"index:users_fk_updated_by;column:updated_by;type:bigint unsigned;not null" json:"updatedBy"`
// }

// TableName get sql table name.获取数据库表名
func (m *Users) TableName() string {
	return "users"
}

// UsersColumns get sql column name.获取数据库列名
var UsersColumns = struct {
	ID        string
	UUID      string
	LoginID   string
	Password  string
	Name      string
	Avatar    string
	Alias     string
	Phone     string
	Email     string
	Gender    string
	Lang      string
	HomePath  string
	Status    string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
}{
	ID:        "id",
	UUID:      "uuid",
	LoginID:   "login_id",
	Password:  "password",
	Name:      "name",
	Avatar:    "avatar",
	Alias:     "alias",
	Phone:     "phone",
	Email:     "email",
	Gender:    "gender",
	Lang:      "lang",
	HomePath:  "home_path",
	Status:    "status",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
}
