// Package utils defines some useful functions for all the project
package utils

import (
	"fmt"
	"os"
)

func CheckError(err error, comment string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "In: %s, Fatal error: %s", comment, err.Error())
		os.Exit(1)
	}
}
