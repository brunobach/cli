package main

import (
	"github.com/brunobach/cli/internal/config"

	"github.com/brunobach/cli/internal/command/create"
	"github.com/spf13/cobra"
)

var cmdRoot = &cobra.Command{
	Use:     "",
	Example: "",
	Short:   "",
}

var (
	tplPath string
)

func init() {
	cmdRoot.AddCommand(cmdCreate)

	cmdCreate.AddCommand(cmdCreateController)
	cmdCreate.AddCommand(cmdCreateUsecase)
	cmdCreate.AddCommand(cmdCreateRepository)
	cmdCreate.AddCommand(cmdCreateModel)
	cmdCreate.AddCommand(cmdCreateAll)
}

func Execute() error {
	return cmdRoot.Execute()
}

var cmdCreate = &cobra.Command{
	Use:     "create [type] [controller_name]",
	Short:   "Create a new controller",
	Example: "cli create controller user",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cmdCreateController.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	cmdCreateUsecase.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	cmdCreateRepository.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	cmdCreateModel.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	cmdCreateAll.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
}

var cmdCreateController = &cobra.Command{
	Use:     "controller",
	Short:   "Create a new Controller",
	Example: "ca-starters-go create controller user",
	Args:    cobra.ExactArgs(1),
	Run:     parseCreate,
}
var cmdCreateUsecase = &cobra.Command{
	Use:     "usecase",
	Short:   "Create a new Usecase",
	Example: "ca-starters-go create usecase user",
	Args:    cobra.ExactArgs(1),
	Run:     parseCreate,
}
var cmdCreateRepository = &cobra.Command{
	Use:     "repository",
	Short:   "Create a new repository",
	Example: "ca-starters-go create repository user",
	Args:    cobra.ExactArgs(1),
	Run:     parseCreate,
}
var cmdCreateModel = &cobra.Command{
	Use:     "model",
	Short:   "Create a new model",
	Example: "ca-starters-go create model user",
	Args:    cobra.ExactArgs(1),
	Run:     parseCreate,
}
var cmdCreateAll = &cobra.Command{
	Use:     "all",
	Short:   "Create a new controller & usecase & repository & model",
	Example: "ca-starters-go create all user",
	Args:    cobra.ExactArgs(1),
	Run:     parseCreate,
}

func parseCreate(cmd *cobra.Command, args []string) {
	cfg := &config.Cfg{
		TplPath:    tplPath,
		CreateType: cmd.Use,
		Args:       args,
	}
	create.Run(cfg)
}

func main() {
	if err := Execute(); err != nil {
		panic(err)
	}
}
