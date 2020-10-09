package wolf

import (
	"io"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestNewTestCmd(t *testing.T) {
	cmd := NewTCmd()
	defer cmd.Cleanup()

	assert := asserter.Wrap(t).Assert
	assert(cmd.Getenv("PWD") != "").Error(`cmd.Getenv("PWD") failed `)
	assert(len(cmd.Args()) != 0).Error("empty cmd.Args")
	wd, _ := cmd.Getwd()
	assert(wd != "").Error("empty cmd.Getwd()")
	assert(cmd.Stdin() != nil).Error("nil cmd.Stdin")
	assert(cmd.Stdout() != nil).Error("nil cmd.Stdout")
	assert(cmd.Stderr() != nil).Error("nil cmd.Stderr")

	cmd.Stdout().Write([]byte("line1\nline2\nline 3"))
	cmd.Stderr().Write([]byte("line1\nline2\nline 3"))
}

func Test_handleErr(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Error("expected panic")
		}
	}()
	handleErr(io.EOF)
}
