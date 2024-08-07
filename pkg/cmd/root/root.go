package root

import (
  browseCmd "github.com/99xtal/bb/pkg/cmd/browse"
  versionCmd "github.com/99xtal/bb/pkg/cmd/version"
  "github.com/spf13/cobra"
)

func NewCmdRoot(version string, buildDate string) (*cobra.Command, error) {
  cmd := &cobra.Command{
    Use: "bb <command> <subcommand> [flags]",
    Short: "Bitbucket CLI",
    Long: "Work seamlessly with Bitbucket from the command line.",
    Annotations: map[string]string{
      "versionInfo": versionCmd.Format(version, buildDate),
    },
  }

  cmd.Flags().BoolP("version", "v", false, "Show bb version")
  cmd.SetHelpFunc(rootHelpFunc)

  cmd.AddCommand(browseCmd.NewCmdBrowse())
  cmd.AddCommand(versionCmd.NewCmdVersion())
  return cmd, nil
}
