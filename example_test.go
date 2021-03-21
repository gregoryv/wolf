package wolf_test

import (
	"fmt"

	"github.com/gregoryv/wolf"
)

func ExampleTCmd_Use() {
	cmd := wolf.NewTCmd("mycmd", "-h")
	defer cmd.Cleanup()

	myRunFunc(cmd)
	fmt.Println(cmd.Dump())
	// output:
	// $ mycmd -h
	// wolf is howling
	// exit 0
	// STDERR:
	// wolf plays pacman
}

func myRunFunc(cmd wolf.Command) {
	fmt.Fprint(cmd.Stdout(), "wolf is howling")
	fmt.Fprint(cmd.Stderr(), "wolf plays pacman")
	cmd.Exit(0)
}
