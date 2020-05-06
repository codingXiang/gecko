package cmd

import (
	builder2 "github.com/codingXiang/gecko/module/delivery/builder"
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
		Use:     "delivery",
		Aliases: []string{"dv"},
		Short:   "Delivery 指令",
		Long:    `Delivery 指令`,
		Run: func(cmd *cobra.Command, args []string) {
			if pkg == "" || filename == "" || destination == "" {
				cmd.Help()
				return
			}
		},
	}

	HttpRoot = &cobra.Command{
		Use:     "http",
		Aliases: []string{"api"},
		Short:   "Http/API 指令",
		Long:    "Http/API 指令",
		Run: func(cmd *cobra.Command, args []string) {
			if pkg == "" || filename == "" || destination == "" {
				cmd.Help()
				return
			}
			repo := builder2.NewHttpBuilder(source+"/"+pkg, filename, destination, pkg)
			repo.Save(repo.General())
			//repo.Save(repo.Create())
		},
	}
	gRPCRoot = &cobra.Command{
		Use:   "grpc",
		Short: "gRPC 指令",
		Long:  "gRPC 指令",
		Run: func(cmd *cobra.Command, args []string) {
			if pkg == "" || filename == "" || destination == "" {
				cmd.Help()
				return
			}
		},
	}
	CliRoot = &cobra.Command{
		Use:   "cli",
		Short: "Cli 指令",
		Long:  "Cli 指令",
		Run: func(cmd *cobra.Command, args []string) {
			if pkg == "" || filename == "" || destination == "" {
				cmd.Help()
				return
			}
		},
	}
)

func init() {
	Root.PersistentFlags().StringVarP(&pkg, "package", "p", "", "Delivery 的 Package 名稱")
	Root.PersistentFlags().StringVarP(&filename, "filename", "f", "service.go", "Service 的檔案名稱")
	Root.PersistentFlags().StringVarP(&source, "source", "s", "./module", "讀取 Service 的檔案路徑")
	Root.PersistentFlags().StringVarP(&destination, "destination", "d", "./module", "建立 Delivery 的路徑")
	Root.AddCommand(HttpRoot, gRPCRoot, CliRoot)
}
