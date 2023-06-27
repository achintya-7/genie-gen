/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "genie-gen",
	Short: "A CLI tool to generate a backend for you",
	Long: `
	This is a tool which will generate a backend implementation for you using gin-gonic and sqlc with postgresql.
	It also has an app.env file which you can use to store your environment variables.
	You can use this tool to generate a backend for your project quickly and focus on writing the apis.
	It also has a makefile which you can use to run the project and generate various stuff like migrations and sqlc files.
	You have to do a little bit of configuration to get it working for your project.	
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
