package seeder

import (
	"fmt"
	"github.com/xian1367/layout-go-zero/orm/factory"
	"github.com/xian1367/layout-go-zero/pkg/console"
	"github.com/xian1367/layout-go-zero/pkg/seed"
	"gorm.io/gorm"
)

func init() {
	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		users := factory.MakeUsers(10)

		result := db.Table("users").Create(&users)

		console.ExitIf(result.Error)

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
