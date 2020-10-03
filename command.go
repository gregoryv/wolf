// Package wolf provides a generic command implementation
package wolf

import "io"

// Command defines a command execution context.
type Command interface {
	Getenv(string) string
	Args() []string
	Getwd() (string, error)
	Stdin() io.Reader
	Stdout() io.Writer
	Stderr() io.Writer
}
