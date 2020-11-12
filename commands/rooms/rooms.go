package rooms

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCmdRooms creates the version command
func NewCmdRooms(f *pkg.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "rooms",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test")
		},
	}

	return cmd
}
