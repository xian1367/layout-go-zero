package seed

import (
	"github.com/xian1367/layout-go-zero/orm/seeder"
	"github.com/xian1367/layout-go-zero/pkg/console"
	"github.com/xian1367/layout-go-zero/pkg/seed"

	"github.com/spf13/cobra"
)

var CmdSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert fake data to the orm",
	Run:   runSeeders,
	Args:  cobra.MaximumNArgs(1), // 只允许 1 个参数
}

func runSeeders(cmd *cobra.Command, args []string) {
	seeder.Initialize()
	if len(args) > 0 {
		// 有传参数的情况
		name := args[0]
		if len(seed.GetSeeder(name).Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + name)
		}
	} else {
		// 默认运行全部迁移
		seed.RunAll()
		console.Success("Done seeding.")
	}
}
