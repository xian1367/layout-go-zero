package make

import (
	"github.com/spf13/cobra"
	"github.com/xian1367/layout-go-zero/pkg/orm"
	"gorm.io/gen"
)

var CmdMakeGen = &cobra.Command{
	Use:   "gen",
	Short: "Generate file and code, example: make gen",
	Run: func(cmd *cobra.Command, args []string) {
		g := gen.NewGenerator(gen.Config{
			OutPath:      "orm/gen/query",
			ModelPkgPath: "gen",
		})

		g.UseDB(orm.DB)

		fieldOpts := []gen.ModelOpt{
			gen.FieldIgnore("id", "created_at", "updated_at", "deleted_at"),
		}

		allModel := g.GenerateAllTable(fieldOpts...)

		g.ApplyBasic(allModel...)

		g.Execute()
	},
}
