package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cobra_example",
	Short: "A simple Cobra CLI example!",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Cobra CLI example!")
	},
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print a message",
	Run: func(cmd *cobra.Command, args []string) {
		message := "Hello, Cobra CLI!"
		fmt.Println(message)
	},
}

func init() {
	// Add the subcommand to the root command
	RootCmd.AddCommand(printCmd)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
