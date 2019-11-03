package main

import (
	"branch-purge-list-creator/command"
	"branch-purge-list-creator/model/git"
	"branch-purge-list-creator/request"
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
