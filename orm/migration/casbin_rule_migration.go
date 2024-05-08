package migration

import (
	"github.com/xian1367/layout-go-zero/orm/model"
	"github.com/xian1367/layout-go-zero/pkg/migrate"
)

func init() {
	type CasbinRule struct {
		model.IDField

		Ptype string `gorm:"size:100;uniqueIndex:unique_index;comment:'规则类型'"`
		V0    string `gorm:"size:100;uniqueIndex:unique_index;comment:'主角色ID'"`
		V1    string `gorm:"size:100;uniqueIndex:unique_index;comment:'子角色ID'"`
		V2    string `gorm:"size:100;uniqueIndex:unique_index;comment:'api路径'"`
		V3    string `gorm:"size:100;uniqueIndex:unique_index;comment:'rest模式'"`
		V4    string `gorm:"size:100;uniqueIndex:unique_index;"`
		V5    string `gorm:"size:100;uniqueIndex:unique_index;"`

		model.CommonTimestampsField
	}

	migrate.Add(&CasbinRule{})
}
