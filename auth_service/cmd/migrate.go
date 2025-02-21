package cmd

import (
	"database/sql"
	"fmt"
	configPkg "k8s_project/auth_service/config"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

var RunMigration=&cobra.Command{
	Use : "migrate [direction]",
	Short: "Run database migrations",
	Long:  "Run database migrations either up or down",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		direction:=args[0]
		runMigration(direction)

	},
}

func init(){
	AddCommand(RunMigration)
}


func runMigration(direction string){
	config,err:=configPkg.Loadconfig(env)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("connection string",config.GetDSN())
	dbConn, err := sql.Open("postgres", config.GetDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	err = goose.SetDialect("postgres")

	if err != nil {
		log.Fatalf("Failed to set Goose dialect: %v", err)
	}

	// Apply migrations from the "db/migrations" directory
	migrationsDir := "sql_db/migrations"
	switch direction {
	case "up":
		err = goose.Up(dbConn, migrationsDir)
	case "down":
		err = goose.Down(dbConn, migrationsDir)
	case "redo":
		err = goose.Redo(dbConn, migrationsDir)
	case "status":
		err = goose.Status(dbConn, migrationsDir)
	default:
		log.Fatalf("Invalid migration direction: %s. Use 'up', 'down', or 'redo'.", direction)
	}
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("✅ Database migration applied successfully!")
}