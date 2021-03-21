package wolf

import (
	"io"
	"log"
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

// Stop returns the given exit code
func (me *OSCmd) Stop(exitCode int) int { return exitCode }
func (me *TCmd) Fatal(v ...interface{}) { log.Fatal(v...) }
