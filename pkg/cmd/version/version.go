package version

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)

func NewCmdVersion() *cobra.Command {
  cmd := &cobra.Command{
    Use: "version",
    Hidden: true,
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Fprintf(os.Stdout, cmd.Root().Annotations["versionInfo"])
    },
  }
  
  return cmd
}

func Format(version string, buildDate string) string {
  return fmt.Sprintf("bb version %s (%s)\n", version, buildDate)
}
