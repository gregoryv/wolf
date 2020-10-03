package wolf

import (
	"io"
	"os"
)

// NewOSCmd returns a command from the OS values.
func NewOSCmd() *OSCmd {
	return &OSCmd{}
}

type OSCmd struct{}

func (me *OSCmd) Getenv(key string) string { return os.Getenv(key) }
func (me *OSCmd) Args() []string           { return os.Args }
func (me *OSCmd) Getwd() (string, error)   { return os.Getwd() }
func (me *OSCmd) Stdin() io.Reader         { return os.Stdin }
func (me *OSCmd) Stdout() io.Writer        { return os.Stdout }
func (me *OSCmd) Stderr() io.Writer        { return os.Stderr }
