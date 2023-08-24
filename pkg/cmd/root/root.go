package root

import (
  "github.com/spf13/cobra"
)

func NewCmdRoot() (*cobra.Command, error) {
  cmd := &cobra.Command{
    Use: "bb <command> <subcommand> [flags]",
    Short: "Bitbucket CLI",
    Long: "Work seamlessly with Bitbucket from the command line.",
  }

  cmd.SetHelpFunc(rootHelpFunc)
  return cmd, nil
}
