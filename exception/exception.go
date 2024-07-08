package exception

import (
	"fmt"
	"os"
)

func FatalIfError(err error) {
	if err != nil {
		fmt.Printf("\033[31m%v\033[0m\n", err)
		os.Exit(1)
	}
}
