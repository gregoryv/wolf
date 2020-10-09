package wolf_test

import (
	"fmt"

	"github.com/gregoryv/wolf"
)

func ExampleTCmd_Use() {
	cmd := wolf.NewTCmd()
	defer cmd.Cleanup()

	myRunFunc(cmd)
	fmt.Println("stdout:", cmd.Out.String())
	fmt.Println("stderr:", cmd.Err.String())
	// output:
	// stdout: wolf is howling
	// stderr: wolf plays pacman
}

func myRunFunc(cmd wolf.Command) {
	fmt.Fprint(cmd.Stdout(), "wolf is howling")
	fmt.Fprint(cmd.Stderr(), "wolf plays pacman")
}
