package rooms

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/commands/filters/remove"
	"github.com/agusnavce/posifi-cli/commands/rooms/create"
	"github.com/agusnavce/posifi-cli/commands/rooms/list"
	"github.com/agusnavce/posifi-cli/commands/rooms/update"
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

	cmd.AddCommand(create.NewCmdCreate(f))
	cmd.AddCommand(list.NewCmdList(f))
	cmd.AddCommand(remove.NewCmdRemove(f))
	cmd.AddCommand(update.NewCmdUpdate(f))

	return cmd
}
