package commands

import (
	"github.com/MakeNowJust/heredoc"
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

	cmd.AddCommand(NewCmdAccuracy(f))
	cmd.AddCommand(NewCmdFilter(f))
	cmd.AddCommand(NewCmdList(f))
	cmd.AddCommand(NewCmdScan(f))
	cmd.AddCommand(NewCmdTrain(f))
	cmd.AddCommand(NewCmdInsert(f))
	cmd.AddCommand(NewCmdDelete(f))

	// Help topics
	// cmd.AddCommand(NewHelpTopic("environment"))

	return cmd
}
