package cmd

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

// Args .
var Args = os.Args

// ICommand .
type ICommand interface {
	Execute() error
}

// Command .
type Command struct {
	// Use is the one-line usage message.
	Use string
	// Short is the short description shown in the 'help' output.
	Short string
	// Long is the long message shown in the 'help <this-command>' output.
	Long string
	// SetOptions:
	SetOptions func(c *Command) error
	// Parse:
	Parse func(c *Command) error
	// Run: Typically the actual work function. Most commands will only implement this.
	Run func(cmd *Command, args []string)

	// flags is full set of flags.
	flags *flag.FlagSet
}

// Execute .
func (c *Command) Execute() error {
	if ok := c.SetOptions(c); ok != nil {
		fmt.Println("Error in SetOptions!")
		return ok
	}
	if ok := c.Parse(c); ok != nil {
		fmt.Println("Error in Parsing!")
		return ok
	}
	c.Run(c, Args)
	return nil
}

// Flags returns the complete FlagSet that applies
// to this command (local and persistent declared here and by all parents).
func (c *Command) Flags() *flag.FlagSet {
	if c.flags == nil {
		c.flags = flag.NewFlagSet(c.Use, flag.ContinueOnError)
	}
	return c.flags
}

func init() {
	fmt.Println("init logic here...", Args)
}
