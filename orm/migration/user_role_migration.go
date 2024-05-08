package migration

import (
	"github.com/xian1367/layout-go-zero/orm/model"
	"github.com/xian1367/layout-go-zero/pkg/migrate"
)

func init() {
    type UserRole struct {
		model.IDField

        Name     string `gorm:"type:varchar(255);not null;index"`

		model.CommonTimestampsField
    }

    migrate.Add(&UserRole{})
}