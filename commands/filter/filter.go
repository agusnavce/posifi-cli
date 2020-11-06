package filter

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCmdFilter creates the version command
func NewCmdFilter(f *pkg.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "filter",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test")
		},
	}

	return cmd
}
