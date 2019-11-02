package command

import (
	"os/exec"
)

func ExecGitLog(branches []string) ([]string, error) {
	var branchLogs []string
	dateFormat := "--date=short"
	for _, branch := range branches {
		logFormat := "--pretty=format:{%n  \"branch_name\": \"" + branch + "\",%n  \"commiter_name\": \"%aN\",%n  \"last_commit_date\": \"%ad\"%n}"
		branchLog, err := exec.Command("git", "log", "-n", "1", branch, logFormat, dateFormat).Output()
		if err != nil {
			return []string{}, err
		}
		branchLogs = append(branchLogs, string(branchLog))
	}
	return branchLogs, nil
}
