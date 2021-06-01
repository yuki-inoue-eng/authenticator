package models

import "time"

type UserIF interface {
	TableName() string // gormで使用するテーブルを強制指定
	NeedsValidation() bool
}

type User struct {
	ID          int        `gorm:"primaryKey"`
	LoginID     string     `gorm:"column:login_id"`
	Password    string     `gorm:"column:password"`
	LastLoginAt *time.Time `gorm:"column:last_login_at"`
	Qos         Qos        `gorm:"column:qos"`
	CreatedAt   time.Time  // column name is `created_at`
	UpdatedAt   time.Time  // column name is `updated_at`
	DeletedAt   *time.Time // column name is `deleted_at`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) NeedsValidation() bool {
	return u.Qos == QosDefault
}

const (
	QosDefault = Qos("default")
	QosRaw     = Qos("raw")
)

// Qos defines type of quality of service
type Qos string

func ToQos(s string) Qos {
	if Qos(s) == QosRaw {
		return QosRaw
	}
	return QosDefault
}
