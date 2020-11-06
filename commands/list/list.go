package list

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCmdList creates the version command
func NewCmdList(f *pkg.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test")
		},
	}

	return cmd
}
