package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)


var env string
var envs=[]string{"local","dev","stage"}

var RootCommand = &cobra.Command{
	Use: "auth",
	Run: func(cmd *cobra.Command, args []string) {
		// dbURL := os.Getenv("DB_URL")
		if env == "" {
			log.Fatal("env is not provided")
		}
		// _ = os.Getenv("DB_TYPE")
	},
}

func Execute() error {
	if err := RootCommand.Execute(); err != nil {
		os.Exit(1)
	}
	return nil
}

func init() {
	RootCommand.PersistentFlags().StringVarP(&env, "env", "e", "", "Specify the environment: local, dev, stage, prod")
}

func AddCommand(cmd *cobra.Command){
	RootCommand.AddCommand(cmd)
}