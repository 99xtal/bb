package root

import (
  "fmt"
  "os"
  "strings"

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

  namePadding := 12

  if isRootCmd(command) {
    if versionVal, err := flags.GetBool("version"); err == nil && versionVal {
      fmt.Fprintf(os.Stdout, command.Annotations["versionInfo"])
      return
    } 
  }
  helpEntries := []helpEntry{}

  helpEntries = append(helpEntries, helpEntry{"", command.Long})
  helpEntries = append(helpEntries, helpEntry{"USAGE", "  " + command.UseLine()})

  for _, g := range GroupedCommands(command) {
		var names []string
		for _, c := range g.Commands {
			names = append(names, rpad("  "+c.Name()+":", namePadding)+c.Short)
		}
		helpEntries = append(helpEntries, helpEntry{
			Title: strings.ToUpper(g.Title),
			Body:  strings.Join(names, "\n"),
		})
	}

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

type CommandGroup struct {
	Title    string
	Commands []*cobra.Command
}

func GroupedCommands(cmd *cobra.Command) []CommandGroup {
	var res []CommandGroup

	for _, g := range cmd.Groups() {
		var cmds []*cobra.Command
		for _, c := range cmd.Commands() {
			if c.GroupID == g.ID && c.IsAvailableCommand() {
				cmds = append(cmds, c)
			}
		}
		if len(cmds) > 0 {
			res = append(res, CommandGroup{
				Title:    g.Title,
				Commands: cmds,
			})
		}
	}

	var cmds []*cobra.Command
	for _, c := range cmd.Commands() {
		if c.GroupID == "" && c.IsAvailableCommand() {
			cmds = append(cmds, c)
		}
	}
	if len(cmds) > 0 {
		defaultGroupTitle := "Additional commands"
		if len(cmd.Groups()) == 0 {
			defaultGroupTitle = "Available commands"
		}
		res = append(res, CommandGroup{
			Title:    defaultGroupTitle,
			Commands: cmds,
		})
	}

	return res
}

// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds ", padding)
	return fmt.Sprintf(template, s)
}
