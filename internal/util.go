package internal

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func PrintOut(str string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, str+"\n", a...)
}
func PrintError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
}

func PrintErrorMessage(msg string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "ERROR: "+msg+"\n", a...)
}

func PrintErrorMessageCascade(err error, msg string, a ...interface{}) {
	PrintError(CascadeError(err, msg, a))
}

func PrintInfo(str string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "INFO: "+str+"\n", a...)
}

func PrintVerbose(str string, a ...interface{}) {
	if Verbose {
		_, _ = fmt.Fprintf(os.Stderr, "DEBUG: "+str+"\n", a...)
	}
}

func CascadeError(err error, str string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(str, a...) + "(" + err.Error() + ")")
}

func closeReader(z VirtualFile, reader io.ReadCloser) {
	err := reader.Close()
	if err != nil {
		PrintErrorMessageCascade(err, "[1017] Impossible to close reader %v", z.FullPath())
	}
}
