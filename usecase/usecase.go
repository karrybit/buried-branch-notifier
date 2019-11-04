package usecase

import (
	"buried-branch-notifier/model/git"
	"sort"
	"time"
)

func SortByLastCommitDate(branchInformations []*git.BranchInformation) {
	sort.SliceStable(branchInformations, func(i, j int) bool {
		return branchInformations[i].LastCommitDate.Before(branchInformations[j].LastCommitDate.Time)
	})
}

func TieOldBranchToAuthor(branchInformations []*git.BranchInformation) map[string][]*git.BranchInformation {
	now := time.Now()
	branchCommiterMap := make(map[string][]*git.BranchInformation)
	for _, branchInformation := range branchInformations {
		days := int(now.Sub(branchInformation.LastCommitDate.Time).Hours()) / 24
		if days >= 14 {
			branchCommiterMap[branchInformation.CommiterName] = append(branchCommiterMap[branchInformation.CommiterName], branchInformation)
		}
	}
	return branchCommiterMap
}
