package command

import (
	"fmt"
	"os/exec"
)

func ExecGitLog(branches []string) ([]string, error) {
	var branchLogs []string
	dateFormat := "--date=iso-strict"
	for _, branch := range branches {
		logFormat := fmt.Sprintf(`--pretty=format:{"branch_name":"%s","commiter_name":"%s","last_commit_date":"%s"}`, branch, "%aN", "%ad")
		branchLog, err := exec.Command("git", "log", "-n", "1", branch, logFormat, dateFormat).Output()
		if err != nil {
			return []string{}, err
		}
		branchLogs = append(branchLogs, string(branchLog))
	}
	return branchLogs, nil
}
