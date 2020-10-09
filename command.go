/* Package wolf provides a generic command implementation.

Type Command wraps the way a command was called. NewOSCmd returns a
wrapper for various os.X methods, whereas NewTCmd uses temporary
directories and mocked environment for easy testing.

   func Test_myRunFunc(t *testing.T) {
       cmd := wolf.NewTCmd()
       defer cmd.Cleanup()

       myRunFunc(cmd)
       t.Log(cmd.Dump())
   }

*/
package wolf

import "io"

// Command defines a command execution context.
type Command interface {
	Getenv(string) string
	Args() []string
	Getwd() (string, error)
	Stdin() io.Reader
	Stdout() io.Writer
	Stderr() io.Writer
	Stop(exitCode int) int
}
