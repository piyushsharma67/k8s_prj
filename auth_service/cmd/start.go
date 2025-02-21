package cmd

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var port string
var dbType string

func loadEnv() error{
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

var startServer = &cobra.Command{
	Use: "start",
	Run:func(cmd *cobra.Command, args []string) {
		// if env == ""{
		// 	log.Fatal("Env not provided")
		// }
		// configPrj, err := configPkg.Loadconfig(env)
		// if err!=nil{
		// 	log.Fatal(err)
		// }

		// var repoInstance *repository.Repositories

		// if dbType ==utils.Postgres{
		// 	pgxpool, err := pgxpool.New(context.Background(), configPrj.GetDSN())
		// 	repoInstance=repository.Ini()
		// }else{
			
		// }
		fmt.Println("i am called")
	},
}

func init(){
	startServer.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
	startServer.Flags().StringVar(&dbType, "db-type", "postgres", "Database type (postgres or mongo)")
	RootCommand.AddCommand(startServer) 
}