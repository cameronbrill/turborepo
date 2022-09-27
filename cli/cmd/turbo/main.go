package main

import (
	"os"

	"github.com/cameronbrill/turborepo/cli/internal/cmd"
)

func main() {
	os.Exit(cmd.RunWithArgs(os.Args[1:], turboVersion))
}
