package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var protoRoot string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate_proto",
	Short: "Generates Go code from proto files",
	Long: `This command scans the proto directory and generates Go and gRPC code 
in the same locations as the .proto files, without creating extra folders.`,
	Run: func(cmd *cobra.Command, args []string) {
		if protoRoot == "" {
			fmt.Println("❌ Error: Please specify the proto root folder using --path")
			os.Exit(1)
		}

		// Convert to absolute path
		absPath, err := filepath.Abs(protoRoot)
		if err != nil {
			fmt.Println("❌ Error resolving absolute path:", err)
			os.Exit(1)
		}

		fmt.Println("🔍 Scanning for .proto files in:", absPath)

		// Walk through the directory to find all .proto files
		err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filepath.Ext(path) == ".proto" {
				fmt.Println("🛠️ Processing:", path)

				// Get the directory where the .proto file is located
				protoDir := filepath.Dir(path)

				// Run protoc with --go_out and --go-grpc_out set to the same directory
				cmd := exec.Command("protoc",
					"--go_out=.",      // Keep generated files in the same folder
					"--go-grpc_out=.", // Keep gRPC file in the same folder
					"--proto_path="+absPath,   // Ensure correct import paths
					path,                      // The .proto file
				)

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Run(); err != nil {
					fmt.Println("❌ Error generating:", path, err)
					os.Exit(1)
				} else {
					fmt.Println("✅ Generated:", filepath.Join(protoDir, "*.pb.go"))
				}
			}
			return nil
		})

		if err != nil {
			fmt.Println("❌ Error while searching proto files:", err)
			os.Exit(1)
		}

		fmt.Println("🚀 All proto files processed successfully!")
	},
}

func init() {
	RootCommand.AddCommand(generateCmd)

	// Define a flag for specifying the root proto directory
	generateCmd.Flags().StringVarP(&protoRoot, "path", "p", "", "Root directory containing proto files (required)")
	generateCmd.MarkFlagRequired("path")
}
