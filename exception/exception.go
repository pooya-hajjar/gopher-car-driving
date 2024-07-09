package exception

import (
	"fmt"
	"os"
)

// FatalIfError checks err parameter then print red-color error if it is not nil
func FatalIfError(err error) {
	if err != nil {
		fmt.Printf("\033[31m%v\033[0m\n", err)
		os.Exit(1)
	}
}
