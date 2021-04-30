package wolf

import (
	"io"
	"testing"
)

func TestNewTCmd(t *testing.T) {
	cases := map[string]*TCmd{
		"no arguments":       NewTCmd(),
		"single argument":    NewTCmd("-h"),
		"multiple arguments": NewTCmd("-p", "-v"),
	}

	for name, cmd := range cases {
		t.Run(name, func(t *testing.T) {
			if cmd == nil {
				t.FailNow()
			}
			cmd.Cleanup()
		})
	}
}

func Test_handleErr(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Error("expected panic")
		}
	}()
	handleErr(io.EOF)
}
