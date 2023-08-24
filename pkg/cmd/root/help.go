package root

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)

type helpEntry struct {
  Title string
  Body string
}

func isRootCmd(command *cobra.Command) bool {
  return command != nil && !command.HasParent()
}

func rootHelpFunc(command *cobra.Command, args []string) {
  flags := command.Flags()

  if isRootCmd(command) {
    if versionVal, err := flags.GetBool("version"); err == nil && versionVal {
      fmt.Fprintf(os.Stdout, command.Annotations["versionInfo"])
      return
    } 
  }
  helpEntries := []helpEntry{}

  helpEntries = append(helpEntries, helpEntry{"", command.Long})
  helpEntries = append(helpEntries, helpEntry{"USAGE", "  " + command.UseLine()})

  flagUsages := command.LocalFlags().FlagUsages()
  if flagUsages != "" {
    helpEntries = append(helpEntries, helpEntry{"FLAGS", flagUsages})
  }

  for _, e := range(helpEntries) {
    if e.Title != "" {
      fmt.Fprintln(os.Stdout, e.Title)
      fmt.Fprintln(os.Stdout, e.Body)
    } else {
      fmt.Fprintln(os.Stdout, e.Body)
    }
    fmt.Fprintln(os.Stdout)
  }
}

