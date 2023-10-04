package main

import (
	"context"
	"flag"
	"os"

	"github.com/Dev-Time/gopbo/cmd"
	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&cmd.UnpackCmd{}, "")
	subcommands.Register(&cmd.ValidateCmd{}, "")
	subcommands.Register(&cmd.PackCmd{}, "")
	subcommands.Register(&cmd.VersionCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
