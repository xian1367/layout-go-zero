package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeApi = &cobra.Command{
	Use:   "api",
	Short: "Create api, example: make api user_http user",
	Args:  cobra.ExactArgs(2), // 至少传 2 个参数
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		model := makeModelFromString(args[1])

		dir := fmt.Sprintf("app/%s/", serviceName)
		createFileFromStub(dir+"main.go", "http/main", model, map[string]string{
			"{{ServiceName}}": serviceName,
		})
		createFileFromStub(dir+"route/route.go", "http/route", model)

		CmdMakeApiController.Run(cmd, args)
		CmdMakeApiRequest.Run(cmd, args)
		CmdMakeApiRoute.Run(cmd, args)
	},
}
