package exception

import (
	"fmt"
	"os"
)

// FatalIfError checks if the err parameter is not nil. If err is not nil,
// it prints the error message in red color and exits the program.
func FatalIfError(err error) {
	if err != nil {
		fmt.Printf("\033[31m%v\033[0m\n", err)
		os.Exit(1)
	}
}
