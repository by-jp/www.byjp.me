package shared

import (
	"fmt"
	"os"
)

func Check(err error, msg string) {
	if err != nil {
		ExitMessage(fmt.Sprintf("%s\n  %v", msg, err))
	}
}

func ExitMessage(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

type closer interface {
	Close() error
}

func DoClose(c closer, msg string) {
	Check(c.Close(), msg)
}
