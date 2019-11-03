package command

import (
	"os/exec"
	"strings"
)

func ExecGitBranch() ([]string, error) {
	branchesString, err := exec.Command("git", "branch", "-r").Output()
	if err != nil {
		return []string{}, err
	}

	branches := splitBranches(strings.TrimSpace(string(branchesString)))
	return removeHeadBranch(branches), nil
}

func splitBranches(branchesString string) []string {
	splitedBranches := strings.Split(branchesString, "\n")
	var trimedBranches []string
	for _, branch := range splitedBranches {
		trimedBranches = append(trimedBranches, strings.TrimSpace(branch))
	}
	return trimedBranches
}

func removeHeadBranch(branches []string) []string {
	for i, branch := range branches {
		if strings.Contains(branch, "origin/HEAD") {
			return append(branches[:i], branches[i+1:]...)
		}
	}
	return branches
}
