package commands

import (
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Commands) newAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [description] [amount]",
		Short: "Add a new expense",
		RunE: func(cmd *cobra.Command, args []string) error {
			description := args[0]
			amount, err := strconv.ParseFloat(args[1], 32)
			if err != nil {
				return err
			}
			return c.handler.CreateExpense(description, float32(amount))
		},
	}
}
