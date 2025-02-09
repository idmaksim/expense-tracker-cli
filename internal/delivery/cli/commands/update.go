package commands

import (
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Commands) newUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update [id] [description] [amount]",
		Short: "Update an expense",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			amount, err := strconv.ParseFloat(args[2], 32)
			if err != nil {
				return err
			}
			return c.handler.UpdateExpense(id, args[1], float32(amount))
		},
	}
}
