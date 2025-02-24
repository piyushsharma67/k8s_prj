package cmd

import (
	"fmt"

	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

var migrateCreateCmd = &cobra.Command{
	Use: "create [name]",
	Short: "Create a new migration file",
	Long:  "Create a new migration file with the specified name inside the postgres/migrations folder",
	Args:  cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		dir := "sql_db/migrations"
		err := goose.Create(nil, dir, name, "sql")
		if err != nil {
			fmt.Printf("Error creating migration: %v\n", err)
			return
		}

		fmt.Printf("✅ Migration file created successfully: %s\n", name)
	},
}

func init() {
	AddCommand(migrateCreateCmd)
}