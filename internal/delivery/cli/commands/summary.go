package commands

import "github.com/spf13/cobra"

func (c *Commands) newSummaryCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "summary",
		Short: "Show the summary of all expenses",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.handler.Summary()
		},
	}
}
