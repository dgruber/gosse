// Package workflow defines helper functions for creating job
// workflows

package workflow

import (
	"fmt"
	"github.com/dgruber/drmaa2interface"
	"os"
)

// CreateJobTemplate creates a job template and fills out the remote command
// and rguments so that the function can be launched from the jobfunc package.
func CreateJobTemplate(function string, args []string) (jt drmaa2interface.JobTemplate) {
	workingdir, _ := os.Getwd()
	appname := os.Args[0]
	jt.RemoteCommand = fmt.Sprintf("%s/%s", workingdir, appname)
	jt.Args = make([]string, 0, len(args)+1)
	jt.Args = append(jt.Args, function, args...)
	return jt
}
