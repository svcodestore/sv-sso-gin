package model

import (
	"time"
)

// Organizations [...]
type Organizations struct {
	ID        string    `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	Code      string    `gorm:"unique;column:code;type:varchar(64);not null" json:"code"`
	Name      string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Status    bool      `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy int64     `gorm:"column:created_by;type:bigint;not null" json:"createdBy"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy int64     `gorm:"column:updated_by;type:bigint;not null" json:"updatedBy"`
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
