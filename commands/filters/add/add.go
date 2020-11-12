package add

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCmdAdd creates the version command
func NewCmdAdd(f *pkg.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "add",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test")
		},
	}

	return cmd
}
