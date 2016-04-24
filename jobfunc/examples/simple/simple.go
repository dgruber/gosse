package main

import (
	"fmt"
	"github.com/dgruber/gosse/jobfunc"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func getFunctionName(f func([]string) int) string {
	fullname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	parts := strings.Split(fullname, ".")
	return parts[len(parts)-1]
}

func createCmd(f func([]string) int) *exec.Cmd {
	cmd := exec.Command(os.Args[0], getFunctionName(f))
	cmd.Stdout = os.Stdout
	return cmd
}

func workflow(args []string) int {
	// do something
	createCmd(job1).Run()
	createCmd(job2).Run()
	createCmd(job3).Run()
	return 0
}
func job1(args []string) int {
	// do something
	fmt.Println("Job1 sleeps 1 second...")
	time.Sleep(time.Second * 1)
	return 0
}
func job2(args []string) int {
	// do something
	fmt.Println("Job2 sleeps 2 second...")
	time.Sleep(time.Second * 2)
	return 0
}
func job3(args []string) int {
	// do something
	fmt.Println("Job3 sleeps 3 second...")
	time.Sleep(time.Second * 3)
	return 0
}

func main() {
	launcher := jobfunc.NewLauncher()
	launcher.RegisterFunction(workflow)
	launcher.RegisterFunction(job1)
	launcher.RegisterFunction(job2)
	launcher.RegisterFunction(job3)
	launcher.Main(os.Args)
}
