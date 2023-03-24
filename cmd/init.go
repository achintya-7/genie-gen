/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/achintya-7/genie-gen/constant"
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
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string

		fmt.Println(constant.GetArt())

		var respone1 string
		fmt.Print("Do you want to add a GitHub tag? (y/n) : ")
		fmt.Scanln(&respone1)

		if respone1 == "n" || respone1 == "N" {
			projectName = ""
		} else {
			projectName = "github.com/"

			var username string
			fmt.Print("Enter the GitHub username: ")
			fmt.Scanln(&username)

			if username == "" {
				fmt.Println("Username cannot be empty")
				os.Exit(1)
			} else {
				projectName += username + "/"
			}

		}

		var pkgName string
		fmt.Print("Enter the project name: ")
		fmt.Scanln(&pkgName)

		if pkgName == "" {
			fmt.Println("Project name cannot be empty")
			os.Exit(1)
		} else {
			projectName += pkgName
		}

		fmt.Println("Project name: ", projectName)
		fmt.Println("Package name: ", pkgName)
		fmt.Print("Continue? (y/n) : ")
		fmt.Scanln(&respone1)

		if respone1 == "n" || respone1 == "N" {
			os.Exit(1)
		}

		fmt.Println("Initializing project...")
		err := processes.Start(pkgName, projectName)
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
