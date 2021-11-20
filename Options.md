# Command line options and flags

For more detailed option usage, see the [tag document](Tags.md).

## How it works
The `opt` package is a fairly GNU-like command line option parser with additional tool command functionality. It interprets options as encountered, then switches context to a new set of options for each tool command encountered.

Single-character options can be merged into a string, and parsing moves on to the next long/short/tool option string when a non-flag (non-boolean) option is encountered.

### Special notes
- The Usage() function and tool commands need to be called manually. The package only parses and won't execute anything automatucally.

## Using it

A minimal example of use:
```go
var Options struct {
	opt.DefaultHelp
	Version bool `opt:"" short:"V" long:"version" help:"Display the version string and exit."`
}

func main() {
	a := opt.Parse(&Options)
	if Options.Help {
		a.Usage()
		return
	}

	if Options.Version {
		fmt.Println("Version whatever")
		return
	}

	println("Twiddling thumbs.")
}
```

This exits if the `-h` or `--help` flag is specified, showing pretty-printed usage options.

A tool command variant:
```go
var Options struct {
	Version bool `opt:"" short:"V" long:"version" help:"Display the version string and exit."`
	Help HelpCmd `short:"c" help:"Show help on a subject."`
}

func main() {
	a := opt.Parse(&Options)
	if Options.Help {
		a.Usage()
	}

	err := a.RunCommand(false)
	if err != nil {
		if err == opt.ErrNoCommand {
			a.Usage()
			return
		}

		pr("Error running command: %s", err.Error())
		os.Exit(2)
	}
}

type HelpCmd struct {
	Subject string `help:"The subject." placeholder:"SUBJECT"`
}

func (cmd *HelpCmd) Run(args []string) error {
	if cmd.Subject == "" {
		return opt.ErrUsage
	}

	// Do help stuff here
}
```

### Positional options
Options can be tagged with only placeholder and help tags, which will make them positional arguments.
