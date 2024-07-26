package orm

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(100);not null json:username"`
	Password string `gorm:"type:varchar(100);not null json:password"`
	Fullname string `gorm:"type:varchar(100);not null json:fullname"`
	gorm.Model
}
