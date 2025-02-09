package commands

import "github.com/spf13/cobra"

func (c *Commands) newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all expenses",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.handler.ListExpenses()
		},
	}
}
