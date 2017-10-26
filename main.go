// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// OO design for Command.go just like cobra
package main

import (
	"fmt"

	cmd "github.com/pmlpml/agenda/cmd"
)

func main() {
	//cmd.Execute()
	var RootCmd = &cmd.Command{
		Use:   "test",
		Short: "A brief description of your application",
		Long:  "A longer description",
	}
	RootCmd.SetOptions = func(c *cmd.Command) error {
		fmt.Println("Set Options here ...")
		c.Flags().StringP("user", "u", "Anonymous", "Help message for username")
		return nil
	}
	RootCmd.Parse = func(c *cmd.Command) error {
		fmt.Println("Parse here ...")
		c.Flags().Parse(cmd.Args)
		return nil
	}
	RootCmd.Run = func(c *cmd.Command, a []string) {
		fmt.Println("Do comamnd here ...")
		username, _ := c.Flags().GetString("user")
		fmt.Println("myCommand called by " + username)
	}
	RootCmd.Execute()
}
