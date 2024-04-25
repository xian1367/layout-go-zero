package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xian1367/layout-go-zero/pkg/console"
)

var CmdMakeCmdService = &cobra.Command{
	Use:   "cmd_service",
	Short: "Create a command, should be snake_case, example: make cmd_service user user",
	Args:  cobra.ExactArgs(2), // 至少传 2 个参数
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		model := makeModelFromString(args[1])

		// 组建目标目录
		filePath := fmt.Sprintf("cmd/service/%s/%s.go", serviceName, model.PackageName)

		// 基于模板创建文件（做好变量替换）
		createFileFromStub(filePath, "cmd_service", model)

		// 友好提示
		console.Success("command name:" + model.PackageName)
		console.Success("command variable name: cmd.Cmd" + model.StructName)
		console.Warning("please edit main.go's helper.Commands slice to register command")
	},
}
