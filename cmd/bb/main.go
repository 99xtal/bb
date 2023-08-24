package main

import (
  "context"
  "fmt"
  "os"
  
  "github.com/99xtal/bb/pkg/cmd/root"
)

type exitCode int

const (
	exitOK     exitCode = 0
	exitError  exitCode = 1
	exitCancel exitCode = 2
	exitAuth   exitCode = 4
)

func run() exitCode {
   ctx := context.Background()

  rootCmd, err := root.NewCmdRoot()
  if err != nil {
    fmt.Fprintf(os.Stderr, "failed to create root command: %s\n", err)
    return exitError
  }

  _, err = rootCmd.ExecuteContextC(ctx)
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
    return exitError
  }
  
  return exitOK
}

func main() {
  code := run()
  os.Exit(int(code))
}
