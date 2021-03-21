package wolf

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestNewOSCmd(t *testing.T) {
	cmd := NewOSCmd()
	cmd.exit = func(v int) {
		if v != 1 {
			t.Error("got exit code", v)
		}
	}
	testCommand(t, cmd)
	testCommand(t, NewTCmd())
}

func testCommand(t *testing.T, cmd Command) {
	t.Helper()
	assert := asserter.Wrap(t).Assert

	assert(cmd.Getenv("PWD") != "").Error(`cmd.Getenv("PWD") failed `)
	assert(len(cmd.Args()) != 0).Error("empty cmd.Args")
	wd, _ := cmd.Getwd()
	assert(wd != "").Error("empty cmd.Getwd()")
	assert(cmd.Stdin() != nil).Error("nil cmd.Stdin")
	assert(cmd.Stdout() != nil).Error("nil cmd.Stdout")
	assert(cmd.Stderr() != nil).Error("nil cmd.Stderr")
	cmd.Fatal()
	cmd.Exit(1)
}
