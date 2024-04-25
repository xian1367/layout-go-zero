package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Crate model file, example: make model user",
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	Run: func(cmd *cobra.Command, args []string) {
		// 格式化模型名称，返回一个 Model 对象
		model := makeModelFromString(args[0])

		dir := fmt.Sprintf("orm/model/%s/", model.PackageName)

		// 替换变量
		createFileFromStub(dir+model.PackageName+"_model.go", "model/model", model)
		createFileFromStub(dir+model.PackageName+"_util.go", "model/model_util", model)
		createFileFromStub(dir+model.PackageName+"_hooks.go", "model/model_hooks", model)
	},
}
