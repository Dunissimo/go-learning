package main

import (
	"errors"
	"fmt"
)

type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
	return "validation error"
}

var (
	errBadConnection = errors.New("bad connection")
	errBadRequest    = errors.New("bad request")
	unknownErrorMsg  = "unknown error"
)

// could be done with an array of critical errors
func GetErrorMsg(err error) string {
	if errors.As(err, &nonCriticalError{}) {
		return ""
	}

	if errors.Is(err, errBadConnection) {
		return errBadConnection.Error()
	}

	if errors.Is(err, errBadRequest) {
		return errBadRequest.Error()
	}

	return unknownErrorMsg
}

func main() {
	fmt.Println(GetErrorMsg(errBadRequest))
	fmt.Println(GetErrorMsg(errBadConnection))
	fmt.Println(GetErrorMsg(nonCriticalError{}))
	fmt.Println(GetErrorMsg(errors.New("unknown")))
	fmt.Println(GetErrorMsg(fmt.Errorf("wrap: %w", errBadRequest)))
	fmt.Println(GetErrorMsg(fmt.Errorf("wrap: %w", errBadConnection)))
	fmt.Println(GetErrorMsg(fmt.Errorf("wrap: %w", nonCriticalError{})))
}
