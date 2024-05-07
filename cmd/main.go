package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xian1367/layout-go-zero/cmd/make"
	"github.com/xian1367/layout-go-zero/cmd/migrate"
	"github.com/xian1367/layout-go-zero/cmd/seed"
	"github.com/xian1367/layout-go-zero/cmd/service"
	"github.com/xian1367/layout-go-zero/config"
	"github.com/xian1367/layout-go-zero/pkg/console"
	"github.com/xian1367/layout-go-zero/pkg/orm"
	"github.com/zeromicro/go-zero/core/logc"
	"os"
)

func main() {
	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   config.Get().Name,
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			config.Init("cmd")
			logc.MustSetup(config.Get().Cmd.Log)
			orm.Init()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		make.CmdMake,
		migrate.CmdMigrate,
		seed.CmdSeed,
		service.CmdService,
	)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
