package main

import (
	"buried-branch-notifyer/command"
	"buried-branch-notifyer/model/git"
	"buried-branch-notifyer/request"
	"fmt"
	"os"
)

func main() {
	branches, err := command.ExecGitBranch()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gitLogs, err := command.ExecGitLog(branches)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	branchCommiterMap := git.NewBranchCommiterMap(gitLogs)

	requester, err := request.NewRequester(branchCommiterMap)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	requester.Notify()
}
