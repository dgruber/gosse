package jobfunc

import ()

type JobFuncError struct {
	message string
}

func (j JobFuncError) Error() string {
	return j.message
}

var (
	AlreadedStartedError = JobFuncError{"Registration not possible since app already started"}
	UnknownFunctionError = JobFuncError{"The given function was not found"}
)
