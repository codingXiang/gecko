package cmd

import (
	builder3 "github.com/codingXiang/gecko/module/delivery/builder"
	"github.com/codingXiang/gecko/module/repository/builder"
	builder2 "github.com/codingXiang/gecko/module/service/builder"
	"github.com/spf13/cobra"
)

var (
	source      string
	filename    string
	destination string
	pkg         string
)

var (
	Root = &cobra.Command{
		Use:     "module",
		Aliases: []string{"m"},
		Short:   "module 指令",
		Long:    `module 指令`,
		Run: func(cmd *cobra.Command, args []string) {
			if pkg == "" || filename == "" || destination == "" || source == "" {
				cmd.Help()
				return
			}
			repo := builder.NewRepositoryBuilder(source, filename, destination, pkg)
			repo.Save(repo.General())
			svc := builder2.NewServiceBuilder(source, filename, destination, pkg)
			svc.Save(svc.General())
			delivery := builder3.NewHttpBuilder(destination+"/"+pkg, "service.go", destination, pkg)
			delivery.Save(delivery.General())
		},
	}
)

func init() {
	Root.PersistentFlags().StringVarP(&pkg, "package", "p", "", "Module Package 名稱")
	Root.PersistentFlags().StringVarP(&filename, "filename", "f", "", "Model 的檔案名稱")
	Root.PersistentFlags().StringVarP(&source, "source", "s", "./module", "讀取 Model 的檔案路徑")
	Root.PersistentFlags().StringVarP(&destination, "destination", "d", "./module", "建立 Module 的路徑")
}
