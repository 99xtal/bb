package browse

import (
  "github.com/pkg/browser"
  "github.com/spf13/cobra"
)

type BrowseOptions struct {
  Branch        string
	Commit        string
	NoBrowserFlag bool
}

func NewCmdBrowse() *cobra.Command {
  opts := &BrowseOptions{}
  cmd := &cobra.Command{
    Long:  "Open the Bitbucket repository in the web browser.",
		Short: "Open the repository in the browser",
		Use:   "browse [<number> | <path> | <commit-SHA>]",
		Args:  cobra.MaximumNArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
      return runBrowse(opts)
    },
  }

	cmd.Flags().BoolVarP(&opts.NoBrowserFlag, "no-browser", "n", false, "Print destination URL instead of opening the browser")
	cmd.Flags().StringVarP(&opts.Commit, "commit", "c", "", "Select another commit by passing in the commit SHA, default is the last commit")
	cmd.Flags().StringVarP(&opts.Branch, "branch", "b", "", "Select another branch by passing in the branch name")
  return cmd
}

func runBrowse(opts *BrowseOptions) error {
  return browser.OpenURL("http://bitbucket.com/exclusiveresortsit/thesourcev2/")
}
