package migration

import (
	"github.com/xian1367/layout-go-zero/orm/model"
	"github.com/xian1367/layout-go-zero/pkg/migrate"
)

func init() {
	type User struct {
		model.IDField

		Name     string `gorm:"type:varchar(32);not null;default:'';comment:姓名"`
		Phone    string `gorm:"type:varchar(16);not null;default:'';comment:手机号"`
		Gender   string `gorm:"type:varchar(16);not null;default:'';comment:性别[Male:男,Female:女]"`
		Password string `gorm:"type:varchar(128);not null;default:'';comment:密码"`

		model.CommonTimestampsField
	}

	migrate.Add(&User{})
}
