package wolf

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// NewTCmd returns a command with temporary working directory and
// buffered outputs, useful during testing.
// os.Chdir is called to change working directory to the temporary directory.
// The first argument should be name of command, just as in os.Args. If ommited
// /noname-tcmd is used. Temporary directory is based on that name.
func NewTCmd(args ...string) *TCmd {
	if len(args) == 0 {
		args = []string{"/noname-tcmd"}
	}
	wd, err := ioutil.TempDir("", path.Base(args[0]))
	handleErr(err)
	cmd := TCmd{
		env: map[string]string{
			"PWD": wd,
		},
		args:  args,
		dir:   wd,
		stdin: strings.NewReader(""),
	}
	os.Chdir(cmd.dir)
	return &cmd
}

var handleErr = func(err error) {
	if err != nil {
		panic(err)
	}
}

type TCmd struct {
	env    map[string]string
	args   []string
	dir    string
	stdin  io.Reader
	stdout bytes.Buffer
	stderr bytes.Buffer
}

func (me *TCmd) Getenv(key string) (v string) {
	v, _ = me.env[key]
	return
}

func (me *TCmd) Args() []string         { return me.args }
func (me *TCmd) Getwd() (string, error) { return me.dir, nil }
func (me *TCmd) Stdin() io.Reader       { return me.stdin }
func (me *TCmd) Stdout() io.Writer      { return &me.stdout }
func (me *TCmd) Stderr() io.Writer      { return &me.stderr }

// Cleanup
func (me *TCmd) Cleanup() { os.RemoveAll(me.dir) }
