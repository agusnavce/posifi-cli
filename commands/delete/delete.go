package delete

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCmdDelete creates the version command
func NewCmdDelete(f *pkg.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test")
		},
	}

	return cmd
}
