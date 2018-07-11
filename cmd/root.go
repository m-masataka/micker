package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use: "micker",
	Short: "docker minimum impl",
	Long: `runc`,
	Run: func(cmd *cobra.Command, args []string){
	},
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print version",
	Long: `Print version`,
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("micker v0.1")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
