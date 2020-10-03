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
		Env: map[string]string{
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
	Env map[string]string
	Out bytes.Buffer //  Stdout
	Err bytes.Buffer // Stderr

	args  []string
	dir   string
	stdin io.Reader
	t     T
}

func (me *TCmd) Getenv(key string) (v string) {
	v, _ = me.Env[key]
	return
}

func (me *TCmd) Args() []string         { return me.args }
func (me *TCmd) Getwd() (string, error) { return me.dir, nil }
func (me *TCmd) Stdin() io.Reader       { return me.stdin }
func (me *TCmd) Stdout() io.Writer {
	if me.t != nil {
		return writeFunc(func(data []byte) (int, error) {
			lines := bytes.Split(data, []byte("\n"))
			for _, line := range lines {
				me.t.Log("stdout:", string(line))
			}
			me.Out.Write(data)
			return len(data), nil
		})
	}
	return &me.Out
}
func (me *TCmd) Stderr() io.Writer {
	if me.t != nil {
		return writeFunc(func(data []byte) (int, error) {
			lines := bytes.Split(data, []byte("\n"))
			for _, line := range lines {
				me.t.Log("stderr:", string(line))
			}
			me.Err.Write(data)
			return len(data), nil
		})
	}
	return &me.Err
}

// Cleanup
func (me *TCmd) Cleanup() { os.RemoveAll(me.dir) }

// Use redirects stdout and stderr to t.Log.
func (me *TCmd) Use(t T) *TCmd {
	me.t = t
	return me
}

type writeFunc func([]byte) (int, error)

// Method
func (me writeFunc) Write(data []byte) (int, error) {
	return me(data)
}

type T interface {
	Log(...interface{})
}
