package model

import (
	"time"
)

// ApplicationUser [...]
type ApplicationUser struct {
	ApplicationID int64               `gorm:"primaryKey;index:application_user_application_id_index;column:application_id;type:bigint;not null" json:"-"`
	Applications  Applications        `gorm:"joinForeignKey:application_id;foreignKey:ApplicationID;reference:ApplicationID" json:"applicationsList"`
	UserID        int64               `gorm:"primaryKey;index:application_user_user_id_index;column:user_id;type:bigint;not null" json:"-"`
	Users         UsersWithoutModInfo `gorm:"joinForeignKey:user_id;foreignKey:UserID;reference:UserID" json:"usersList"`
	Status        bool                `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt     time.Time           `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy     int64               `gorm:"column:created_by;type:bigint;not null" json:"-"`
	CreatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:CreatedBy;reference:CreatedBy" json:"createdByUser"`
	UpdatedAt     time.Time           `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy     int64               `gorm:"column:updated_by;type:bigint;not null" json:"-"`
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
	ID            string              `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	Code          string              `gorm:"unique;column:code;type:varchar(64);not null" json:"code"`
	Name          string              `gorm:"column:name;type:varchar(255)" json:"name"`
	InternalURL   string              `gorm:"column:internal_url;type:varchar(255)" json:"internalUrl"`
	HomepageURL   string              `gorm:"column:homepage_url;type:varchar(255)" json:"homepageUrl"`
	Status        bool                `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	ClientID      string              `gorm:"column:client_id;type:varchar(255)" json:"clientId"`
	ClientSecret  string              `gorm:"column:client_secret;type:varchar(255)" json:"clientSecret"`
	RedirectURIs  string              `gorm:"column:redirect_uris;type:varchar(255)" json:"redirectUris"`
	TokenFormat   string              `gorm:"column:token_format;type:varchar(100);default:JWT" json:"tokenFormat"`
	CreatedAt     time.Time           `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy     int64               `gorm:"column:created_by;type:bigint;not null" json:"-"`
	CreatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:CreatedBy;reference:CreatedBy" json:"createdByUser"`
	UpdatedAt     time.Time           `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy     int64               `gorm:"column:updated_by;type:bigint;not null" json:"-"`
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
	TokenFormat:  "token_format",
	CreatedAt:    "created_at",
	CreatedBy:    "created_by",
	UpdatedAt:    "updated_at",
	UpdatedBy:    "updated_by",
}

// OrganizationApplication [...]
type OrganizationApplication struct {
	OrganizationID string              `gorm:"primaryKey;index:organization_application_organization_id_index;column:organization_id;type:bigint;not null" json:"-"`
	Organizations  Organizations       `gorm:"joinForeignKey:organization_id;foreignKey:OrganizationID;reference:OrganizationID" json:"organizationsList"`
	ApplicationID  string              `gorm:"primaryKey;index:organization_application_application_id_index;column:application_id;type:bigint;not null" json:"-"`
	Applications   Applications        `gorm:"joinForeignKey:application_id;foreignKey:ApplicationID;reference:ApplicationID" json:"applicationsList"`
	Status         bool                `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt      time.Time           `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy      int64               `gorm:"column:created_by;type:bigint;not null" json:"-"`
	CreatedByUser  UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:CreatedBy;reference:CreatedBy" json:"createdByUser"`
	UpdatedAt      time.Time           `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy      int64               `gorm:"column:updated_by;type:bigint;not null" json:"-"`
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
	ID            string              `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	Code          string              `gorm:"unique;column:code;type:varchar(64);not null" json:"code"`
	Name          string              `gorm:"column:name;type:varchar(255)" json:"name"`
	Status        bool                `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt     time.Time           `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy     int64               `gorm:"column:created_by;type:bigint;not null" json:"-"`
	CreatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:CreatedBy;reference:CreatedBy" json:"createdByUser"`
	UpdatedAt     time.Time           `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy     int64               `gorm:"column:updated_by;type:bigint;not null" json:"-"`
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

// Users [...]
//type Users struct {
//	ID        string    `gorm:"primaryKey;column:id;type:bigint;not null" json:"-"`
//	UUID      uuid.UUID `gorm:"index:user_uuid_index;column:uuid;type:binary(16);not null" json:"uuid"`
//	LoginID   string    `gorm:"unique;column:login_id;type:varchar(16);not null" json:"loginId"`
//	Password  string    `gorm:"column:password;type:varchar(1024)" json:"password"`
//	Name      string    `gorm:"column:name;type:varchar(32)" json:"name"`
//	Alias     string    `gorm:"column:alias;type:varchar(32)" json:"alias"`
//	Phone     string    `gorm:"column:phone;type:varchar(16)" json:"phone"`
//	Email     string    `gorm:"column:email;type:varchar(1024)" json:"email"`
//	Lang      string    `gorm:"column:lang;type:char(5);default:zh_CN" json:"lang"`
//	Status    bool      `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
//	CreatedAt time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
//	CreatedBy int64     `gorm:"column:created_by;type:bigint;not null" json:"createdBy"`
//	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
//	UpdatedBy int64     `gorm:"column:updated_by;type:bigint;not null" json:"updatedBy"`
//}

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
	Alias     string
	Phone     string
	Email     string
	Lang      string
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
	Alias:     "alias",
	Phone:     "phone",
	Email:     "email",
	Lang:      "lang",
	Status:    "status",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
}
