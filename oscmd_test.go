package wolf

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestNewCmd(t *testing.T) {
	cmd := NewOSCmd()
	assert := asserter.Wrap(t).Assert

	assert(cmd.Getenv("PWD") != "").Error(`cmd.Getenv("PWD") failed `)
	assert(len(cmd.Args()) != 0).Error("empty cmd.Args")
	wd, _ := cmd.Getwd()
	assert(wd != "").Error("empty cmd.Getwd()")
	assert(cmd.Stdin() != nil).Error("nil cmd.Stdin")
	assert(cmd.Stdout() != nil).Error("nil cmd.Stdout")
	assert(cmd.Stderr() != nil).Error("nil cmd.Stderr")
}
