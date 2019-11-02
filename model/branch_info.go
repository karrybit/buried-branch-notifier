package model

import (
	"encoding/json"
)

type BranchInfo struct {
	BranchName           string `json:"branch_name"`
	CommiterName         string `json:"commiter_name"`
	LastCommitDateString string `json:"last_commit_date"`
}

func New(gitLog string) BranchInfo {
	bi := BranchInfo{}
	json.Unmarshal([]byte(gitLog), &bi)
	return bi
}
