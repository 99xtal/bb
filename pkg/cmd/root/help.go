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

func rootHelpFunc(command *cobra.Command, args []string) {
  helpEntries := []helpEntry{}

  helpEntries = append(helpEntries, helpEntry{"", command.Long})
  helpEntries = append(helpEntries, helpEntry{"USAGE", command.UseLine()})

  for _, e := range(helpEntries) {
    if e.Title != "" {
      fmt.Fprintln(os.Stdout, e.Title)
      fmt.Fprintln(os.Stdout, "  " + e.Body)
    } else {
      fmt.Fprintln(os.Stdout, e.Body)
    }
    fmt.Fprintln(os.Stdout)
  }
}

