/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/achintya-7/genie-gen/processes"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Builds a new project with the given name",
	Long: `
	Builds a new project with the given name.
	eg. genie-gen init github.com/achintya-7/genie-gen will create a new project with the name genie-gen in the github.com/achintya-7 directory.
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		fmt.Printf("Initialize project as github.com/%s : [y/n] ", projectName)

		var response string
		fmt.Scanln(&response)

		if response == "n" || response == "N" {
			fmt.Println("Aborting...")
			os.Exit(0)
		}

		fmt.Println("Initializing project...")
		err := processes.Start(projectName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Creating util package...")
		err = processes.Util()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Creating app.env file...")
		err = processes.Env()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Creating db package...")
		err = processes.Db()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Creating sqlc yaml file...")
		err = processes.Sqlc()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Creating api package...")
		err = processes.Api()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Creating Makefile...")
		err = processes.Makefile()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Running go mod tidy...")
		err = processes.End()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
