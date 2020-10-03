package wolf_test

import (
	"fmt"
	"os"

	"github.com/gregoryv/fox"
	"github.com/gregoryv/wolf"
)

func ExampleTCmd_Use() {
	t := fox.NewSyncLog(os.Stdout) // ie. *testing.T
	cmd := wolf.NewTCmd().Use(t)
	defer cmd.Cleanup()

	myRunFunc(cmd)
	// output:
	// stdout: wolf is howling
	// stderr: wolf plays pacman
}

func myRunFunc(cmd wolf.Command) {
	fmt.Fprint(cmd.Stdout(), "wolf is howling")
	fmt.Fprint(cmd.Stderr(), "wolf plays pacman")
}
