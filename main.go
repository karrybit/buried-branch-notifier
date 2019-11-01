package main

import (
	"fmt"
	"os"
)

func main() {
	branches, err := execGitBranch()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(branches)
}
