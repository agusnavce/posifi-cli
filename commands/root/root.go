package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/agusnavce/posifi-cli/commands/filters"
	"github.com/agusnavce/posifi-cli/commands/rooms"
	"github.com/agusnavce/posifi-cli/commands/train"
	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCmdRoot creates the command root
func NewCmdRoot(f *pkg.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pcli <command> <subcommand> [flags]",
		Short: "Posifi CLI",
		Long:  `Work seamlessly with Posifi from the command line.`,

		SilenceErrors: true,
		SilenceUsage:  true,
		Example: heredoc.Doc(`
			$ pcli issue create
		`),
		Annotations: map[string]string{
			"help:environment": heredoc.Doc(`
				See 'pcli help environment' for the list of supported environment variables.
			`),
		},
	}

	cmd.SetOut(f.IOStreams.Out)
	cmd.SetErr(f.IOStreams.ErrOut)

	cmd.PersistentFlags().Bool("help", false, "Show help for command")
	cmd.SetHelpFunc(rootHelpFunc)
	// cmd.SetUsageFunc(rootUsageFunc)
	// cmd.SetFlagErrorFunc(rootFlagErrrorFunc)

	// formattedVersion := versionCmd.Format(version, buildDate)
	// cmd.SetVersionTemplate(formattedVersion)
	// cmd.Version = formattedVersion
	// cmd.Flags().Bool("version", false, "Show gh version")

	cmd.AddCommand(train.NewCmdTrain(f))
	cmd.AddCommand(rooms.NewCmdRooms(f))
	cmd.AddCommand(filters.NewCmdFilters(f))

	// Help topics
	// cmd.AddCommand(NewHelpTopic("environment"))

	return cmd
}
