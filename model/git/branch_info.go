package git

import (
	"encoding/json"
	"strings"
	"time"
)

type CustomDateTime struct {
	time.Time
}

func (cdt *CustomDateTime) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse("2006-01-02T15:04:05-07:00", strInput)
	if err != nil {
		return err
	}

	cdt.Time = newTime
	return nil
}

type BranchInformation struct {
	BranchName     string         `json:"branch_name"`
	CommiterName   string         `json:"commiter_name"`
	LastCommitDate CustomDateTime `json:"last_commit_date"`
}

func NewBranchInformations(gitLogs []string) []*BranchInformation {
	return unmarshalLogs(gitLogs)
}

func unmarshalLogs(gitLogs []string) []*BranchInformation {
	var branchInformations []*BranchInformation
	for _, gitLog := range gitLogs {
		branchInformation := BranchInformation{}
		json.Unmarshal([]byte(gitLog), &branchInformation)
		branchInformations = append(branchInformations, &branchInformation)
	}
	return branchInformations
}
