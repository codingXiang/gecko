package cmd

import (
	"github.com/codingXiang/gecko/module/repository/builder"
	"github.com/spf13/cobra"
)

//var Root = &cobra.Command{
//	Use:     "repository",
//	Aliases: []string{"repo"},
//	Short:   "Repository 指令",
//	Long:    `Repository 指令`,
//}

var (
	source      string
	filename    string
	destination string
	pkg         string
)

var (
	Root = &cobra.Command{
		Use:     "repository",
		Aliases: []string{"repo"},
		Short:   "Repository 指令",
		Long:    `Repository 指令`,
		Run: func(cmd *cobra.Command, args []string) {
			if pkg == "" || filename == "" || destination == "" {
				cmd.Help()
				return
			}
			repo := builder.NewRepositoryBuilder(source, filename, destination, pkg)
			repo.Save(repo.General())
		},
	}
	//Create = &cobra.Command{
	//	Use:   "create",
	//	Short: "建立 Repository",
	//	Long:  "",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		if pkg == "" || filename == "" || destination == "" {
	//			cmd.Help()
	//			return
	//		}
	//		repo := builder.NewRepositoryBuilder(source, filename, destination, pkg)
	//		repo.Create()
	//	},
	//}
)

func init() {
	Root.Flags().StringVarP(&pkg, "package", "p", "", "Repository 的 Package 名稱")
	Root.Flags().StringVarP(&filename, "filename", "f", "", "Model 的檔案名稱")
	Root.Flags().StringVarP(&source, "source", "s", "./model", "讀取 Model 的檔案路徑")
	Root.Flags().StringVarP(&destination, "destination", "d", "./module", "建立 Repository 的路徑")

}

//func autoAddFlags(cmds ...*cobra.Command) {
//	for _, cmd := range cmds {
//		cmd.Flags().StringP("path", "p", "./module", "建立 Repository 的路徑")
//		cmd.Flags().StringP("name", "n", "", "Package 的名稱")
//	}
//}
