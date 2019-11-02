package model

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

type BranchInfo struct {
	BranchName     string         `json:"branch_name"`
	CommiterName   string         `json:"commiter_name"`
	LastCommitDate CustomDateTime `json:"last_commit_date"`
}

func New(gitLog string) BranchInfo {
	bi := BranchInfo{}
	json.Unmarshal([]byte(gitLog), &bi)
	return bi
}
