package opt

import "errors"

var (
	// ErrUsage is really no error, but signals that help text should be printed.
	ErrUsage = errors.New("unknown options")
	// ErrNoCommand is a warning that no command was supplied.
	ErrNoCommand = errors.New("no command specified")
)
