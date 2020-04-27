package cmd

import (
	"github.com/codingXiang/gecko/module/model/cmd"
	repo "github.com/codingXiang/gecko/module/repository/cmd"
	svc "github.com/codingXiang/gecko/module/service/cmd"
	"github.com/spf13/cobra"
)

var General = &cobra.Command{
	Use:     "general",
	Aliases: []string{"g"},
	Short:   "自動產生程式碼",
	Long:    "自動產生程式碼",
}

func init() {
	Root.AddCommand(General)
	General.AddCommand(model.Root, repo.Root, svc.Root)
}
