package remove

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCmdremove creates the version command
func NewCmdremove(f *pkg.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "remove",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test")
		},
	}

	return cmd
}
