// Package utils defines some useful functions for all the project
package utils

import (
	"fmt"
	"os"
)

func CheckError(err error, errMsg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, errMsg, ":", err)
	}
}
