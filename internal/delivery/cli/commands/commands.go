package commands

import (
	"github.com/idmaksim/expense-tracker-cli/internal/delivery/cli/handlers"
	"github.com/spf13/cobra"
)

type Commands struct {
	handler *handlers.ExpenseHandler
	rootCmd *cobra.Command
}

func NewCommands(handler *handlers.ExpenseHandler) *Commands {
	cmd := &Commands{
		handler: handler,
		rootCmd: &cobra.Command{
			Use:   "expense-cli",
			Short: "Cli for tracking expense",
		},
	}
	cmd.setup()
	return cmd
}

func (c *Commands) Execute() error {
	return c.rootCmd.Execute()
}

func (c *Commands) setup() {
	c.rootCmd.AddCommand(
		c.newAddCmd(),
		c.newListCmd(),
		c.newDeleteCmd(),
		c.newSummaryCmd(),
		c.newUpdateCmd(),
	)
}
