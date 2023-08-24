package root

import (
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

  cmd.SetHelpFunc(rootHelpFunc)

  cmd.AddCommand(versionCmd.NewCmdVersion())
  return cmd, nil
}
