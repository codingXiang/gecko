package model

import (
	"github.com/codingXiang/gecko/module/model/builder"
	"github.com/spf13/cobra"
)

var (
	source      string
	destination string
	filename    string
)

var (
	Root = &cobra.Command{
		Use:   "model",
		Short: "建立 model interface 與方法",
		Long:  `透過解析 Struct 進而自動產出 interface 與 method`,
		Run: func(cmd *cobra.Command, args []string) {
			if source == "" || destination == "" || filename == "" {
				cmd.Help()
				return
			}

			m := builder.NewModelBuilder(source, destination, filename)
			m.General()
			m.Save()
		},
	}
)

func init() {
	Root.Flags().StringVarP(&source, "source", "s", "", "讀取 Model 的檔案路徑")
	Root.Flags().StringVarP(&destination, "destination", "d", "", "建立 Model Interface 的檔案路徑")
	Root.Flags().StringVarP(&filename, "filename", "f", "", "Model 的檔名")
}
