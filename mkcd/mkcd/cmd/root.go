/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mkcd",
	Short: "mkdir + cd",
	Long: `mkdir and cd into a directory name
	Example usage :
	mkcd <dir-name>
	`,
	Run:  MkcdRunner,
	Args: cobra.ExactArgs(1),
}

func MkcdRunner(cmd *cobra.Command, args []string) {
	// now the args[0] is the directory name
	dirName := args[0]

	err := os.Mkdir(dirName, 0o755)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(dirName)
	if err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
