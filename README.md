[wolf](https://godoc.org/github.com/gregoryv/wolf) - Package wolf provides a generic command implementation

## Quick start

    go get github.com/gregoryv/wolf
	
in your func main

    func main() {
       cmd := wolf.NewOSCmd()
       app := NewApp(cmd)
	   code := app.Run()
	   os.Exit(code)
    }

and in your tests

	func Test_myRunFunc(t *testing.T) {
       cmd := wolf.NewTCmd()
       defer cmd.Cleanup()

       app := NewApp(cmd)
	   code := app.Run()
	   if code != 0 {
           t.Error(cmd.Out.String(), "\n", cmd.Err.String())
       }
    }
