package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	branches, err := exec.Command("git", "branch", "-rv").Output()

	if err != nil {
		fmt.Println("%v", err)
		os.Exit(1)
	}

	fmt.Println(string(branches))
}
