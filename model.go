package main

import "time"

type GitLog struct {
	branchName     string
	commitHash     string
	lastCommitDate time.Time
	commiter       string
}
