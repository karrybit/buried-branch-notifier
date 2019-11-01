package main

import (
	"os/exec"
)

func execGitBranch() (string, error) {
	branches, err := exec.Command("git", "branch", "-rv").Output()
	if err != nil {
		return "", err
	}

	return string(branches), nil
}
