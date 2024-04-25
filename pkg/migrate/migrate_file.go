package migrate

import (
	"database/sql"
	"fmt"
	"github.com/xian1367/layout-go-zero/pkg/console"
	"github.com/xian1367/layout-go-zero/pkg/helper"
	"gorm.io/gorm"
	"strings"
)

// migrationFunc 定义 up 和 down 回调方法的类型
type migrationFunc func(gorm.Migrator, *sql.DB)

// migrationFiles 所有的迁移文件数组
var migrationFiles []MigrationFile

// MigrationFile 代表着单个迁移文件
type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

// Add 新增一个迁移文件，所有的迁移文件都需要调用此方法来注册
func Add(model interface{}) {
	name := fmt.Sprintf("%T", model)
	lastInd := strings.LastIndex(name, ".")
	name = name[lastInd+1:]
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: helper.StrSnake(name),
		Up: func(migrator gorm.Migrator, DB *sql.DB) {
			console.ExitIf(migrator.AutoMigrate(model))
		},
		Down: func(migrator gorm.Migrator, DB *sql.DB) {
			console.ExitIf(migrator.DropTable(model))
		},
	})
}

// getMigrationFile 通过迁移文件的名称来获取到 MigrationFile 对象
func getMigrationFile(name string) MigrationFile {
	for _, mFile := range migrationFiles {
		if strings.Replace(name, "_migration", "", -1) == mFile.FileName {
			return mFile
		}
	}
	return MigrationFile{}
}

// isNotMigrated 判断迁移是否已执行
func (mFile MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == mFile.FileName {
			return false
		}
	}
	return true
}
