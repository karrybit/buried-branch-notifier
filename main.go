package main

import (
	"buried-branch-notifier/command"
	"buried-branch-notifier/model/git"
	"buried-branch-notifier/request"
	"buried-branch-notifier/usecase"
	"fmt"
	"os"
)

func main() {
	branches, err := command.ExecGitBranch()
	exitIfError(err)

	gitLogs, err := command.ExecGitLog(branches)
	exitIfError(err)

	branchInformations := git.NewBranchInformations(gitLogs)

	usecase.SortByLastCommitDate(branchInformations)
	branchCommiterMap := usecase.TieOldBranchToAuthor(branchInformations)

	requester, err := request.NewRequester(branchCommiterMap, len(branchInformations))
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
