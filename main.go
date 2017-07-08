package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/google/subcommands"
)

type keepCmd struct {
}

func (*keepCmd) Name() string     { return "keep" }
func (*keepCmd) Synopsis() string { return "Keep a diary." }
func (*keepCmd) Usage() string {
	return `keep:
  Keep a diary
`
}
func (p *keepCmd) SetFlags(f *flag.FlagSet) {
}
func (p *keepCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}

type entryCmd struct {
	date string
}

func (*entryCmd) Name() string     { return "entry" }
func (*entryCmd) Synopsis() string { return "Display an entry." }
func (*entryCmd) Usage() string {
	return `entry:
  Display an entry.
`
}
func (p *entryCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.date, "date", time.Now().Format("2006-01-02"), "date of entry")
}
func (p *entryCmd) Execute(_ context.Context, f *flag.FlagSet, argv ...interface{}) subcommands.ExitStatus {
	var lgbk = (argv[0]).(*Logbook)

	t, err := time.Parse("2006-01-02", p.date)
	if err != nil {
		fmt.Println(p.date + " is invalud date.")
	} else {
		fmt.Println(t.Format(lgbk.name + ": " + "2006-01-02"))
	}

	return subcommands.ExitSuccess
}

func main() {
	lgbk, _ := LoadDiary("")

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&keepCmd{}, "")
	subcommands.Register(&entryCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx, lgbk)))
}
