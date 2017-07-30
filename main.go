package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/google/subcommands"
)

var version = "0.0.1"

type keepCmd struct {
	date string
}

func (*keepCmd) Name() string     { return "keep" }
func (*keepCmd) Synopsis() string { return "Keep a diary." }
func (*keepCmd) Usage() string {
	return `keep:
  Keep a diary
`
}
func (p *keepCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.date, "date", time.Now().Format(DATE_FORMAT), "date of entry")
}
func (p *keepCmd) Execute(_ context.Context, f *flag.FlagSet, argv ...interface{}) subcommands.ExitStatus {
	var lgbk = (argv[0]).(*Logbook)

	lgbk.Keep(p.date, argv[1].(string))

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
	f.StringVar(&p.date, "date", time.Now().Format(DATE_FORMAT), "date of entry")
}
func (p *entryCmd) Execute(_ context.Context, f *flag.FlagSet, argv ...interface{}) subcommands.ExitStatus {
	var lgbk = (argv[0]).(*Logbook)

	lgbk.Entry(p.date)

	return subcommands.ExitSuccess
}

func getWorkpath() string {
	var err error

	var u *user.User
	u, err = user.Current()
	if err != nil {
		return ""
	}

	return filepath.Join(u.HomeDir, ".logbook")
}

func main() {
	var err error

	var isShowVersion bool
	flag.BoolVar(&isShowVersion, "v", false, "show version")
	flag.BoolVar(&isShowVersion, "version", false, "show version")

	var workpath string
	workpath = getWorkpath()

	var lgbk *Logbook
	lgbk, err = LoadLogbook(workpath)
	if err != nil {
		fmt.Println("failed LoadLogbook() [err = ", err.Error(), "]")
		return
	}

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&keepCmd{}, "")
	subcommands.Register(&entryCmd{}, "")

	flag.Parse()
	if isShowVersion {
		fmt.Println("version:", version)
		return
	}
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx, lgbk, workpath)))
}
