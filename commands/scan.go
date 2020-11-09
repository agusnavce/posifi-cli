package commands

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCmdScan creates the version command
func NewCmdScan(f *pkg.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "scan",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test")
		},
	}

	return cmd
}
