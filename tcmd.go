package wolf

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/gregoryv/nexus"
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
	Out bytes.Buffer // Stdout
	Err bytes.Buffer // Stderr

	args  []string
	dir   string
	stdin io.Reader
}

func (me *TCmd) Getenv(key string) (v string) {
	v, _ = me.Env[key]
	return
}

func (me *TCmd) Args() []string         { return me.args }
func (me *TCmd) Getwd() (string, error) { return me.dir, nil }
func (me *TCmd) Stdin() io.Reader       { return me.stdin }
func (me *TCmd) Stdout() io.Writer      { return &me.Out }
func (me *TCmd) Stderr() io.Writer      { return &me.Err }

// Cleanup
func (me *TCmd) Cleanup() { os.RemoveAll(me.dir) }

// Dump returns a dump of the command, see DumpTo
func (me *TCmd) Dump() string {
	var b strings.Builder
	me.DumpTo(&b)
	return b.String()
}

// DumpTo writes argument, stdout and stderr if any to the given writer
func (me *TCmd) DumpTo(w io.Writer) error {
	p, err := nexus.NewPrinter(w)

	p.Print("> ")
	p.Print(strings.Join(me.Args(), " "))
	p.Println()
	io.Copy(p, &me.Out)

	if me.Err.Len() > 0 {
		p.Println()
		p.Println("STDERR:")
		io.Copy(p, &me.Err)
	}
	return *err
}
