// Package jobfunc contains helper functions for creating single binary
// applications which can be executed as job workflow orchestration or
// as jobs which then executes only a specific function.
package jobfunc

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

type JobFunction func([]string) int

type Function struct {
	F JobFunction
}

type Launcher struct {
	started  bool // function registration is only possible before calling Main()
	fs       map[string]Function
	workflow Function
}

func NewLauncher() *Launcher {
	return &Launcher{fs: make(map[string]Function)}
}

func (l *Launcher) RegisterFunction(f JobFunction) error {
	if l.started {
		return AlreadedStartedError
	}

	// get function name
	fullname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	parts := strings.Split(fullname, ".")
	n := parts[len(parts)-1]

	// get args

	// save name
	fn := Function{
		F: f,
	}
	l.fs[n] = fn
	return nil
}

func (l *Launcher) ListFunctions() []string {
	keys := make([]string, 0, len(l.fs))
	for key := range l.fs {
		keys = append(keys, key)
	}
	return keys
}

// Main should be called in the applications main function in order
// to start the application. It chooses given the argements to main
// either to run only a specific function or to start the Launcher(),
// which executes the workflow. This is done when the argument is
// "launch". It blocks until the workflow is finished.
func (l *Launcher) Main(args []string) (int, error) {
	var ret int
	if args == nil || len(args) < 2 {
		// todo print help about all registered jobs
		fmt.Println("Requires \"launch\" or job name as parameter")
		keys := l.ListFunctions()
		if len(keys) > 0 {
			fmt.Println("Following job names available:")
			for _, name := range keys {
				fmt.Println(name)
			}
		}
		return 0, UnknownFunctionError
	}
	fname := args[1]
	if fname == "launch" {
		if fn, found := l.fs["workflow"]; found == false {
			fmt.Println("No workflow function registered")
			os.Exit(1)
		} else {
			fn.F(args[2:])
		}
	}
	fn, found := l.fs[fname]
	if !found {
		return 0, UnknownFunctionError
	}
	// execute function
	ret = fn.F(args[2:])
	return ret, nil
}
