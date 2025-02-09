package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

)

func (c *Commands) newDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete an expense by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("id is required")
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return c.handler.DeleteExpense(id)
		},
	}
}
