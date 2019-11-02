package main

import (
	"branch-purge-list-creator/command"
	"branch-purge-list-creator/model/git"
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

	branchOwnerMap := git.NewBranchInformations(gitLogs)
	for key, element := range branchOwnerMap {
		for _, branchInformation := range element {
			fmt.Printf("%s: %+v\n", key, branchInformation)
		}
	}
}
