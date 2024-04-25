package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeApiRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file, example make request user_http user",
	Args:  cobra.ExactArgs(2), // 只允许且必须传 2 个参数
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		model := makeModelFromString(args[1])

		// 组建目标目录
		filePath := fmt.Sprintf("app/%s/request/%s_request.go", serviceName, model.PackageName)

		// 基于模板创建文件（做好变量替换）
		createFileFromStub(filePath, "request", model)
	},
}
