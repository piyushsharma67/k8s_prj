package cmd

import (
	"database/sql"
	"fmt"
	configPkg "notification_service/config"
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
	config,err:=configPkg.LoadConfig(env)
	if err!=nil{
		log.Fatal(err)
	}
	
	dbConn, err := sql.Open("postgres", config.GetDSN())
	fmt.Println("error is",err)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	if err = dbConn.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Database ping successful")

	err = goose.SetDialect("postgres")
	if err != nil {
		log.Fatalf("Failed to set Goose dialect: %v", err)
	}
	fmt.Println("Goose dialect set to postgres")
	// Apply migrations from the "db/migrations" directory
	migrationsDir := "sql_db/migrations"
	switch direction {
	case "up":
		fmt.Println("Running goose.Up")
		err = goose.Up(dbConn, migrationsDir)
	case "down":
		fmt.Println("Running goose.Down")
		err = goose.Down(dbConn, migrationsDir)
	case "redo":
		fmt.Println("Running goose.Redo")
		err = goose.Redo(dbConn, migrationsDir)
	case "status":
		fmt.Println("Running goose.Status")
		err = goose.Status(dbConn, migrationsDir)
	default:
		log.Fatalf("Invalid migration direction: %s. Use 'up', 'down', 'redo', or 'status'.", direction)
	}
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("✅ Database migration applied successfully!")
}