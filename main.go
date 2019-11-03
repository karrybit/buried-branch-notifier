package main

import (
	"buried-branch-notifier/command"
	"buried-branch-notifier/model/git"
	"buried-branch-notifier/request"
	"fmt"
	"os"
)

func main() {
	branches, err := command.ExecGitBranch()
	exitIfError(err)

	gitLogs, err := command.ExecGitLog(branches)
	exitIfError(err)

	branchCommiterMap := git.NewBranchCommiterMap(gitLogs)

	requester, err := request.NewRequester(branchCommiterMap)
	exitIfError(err)

	err = requester.Notify()
	exitIfError(err)
}

func exitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
