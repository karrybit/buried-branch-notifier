package git

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func BenchmarkUnmarshalDefault(b *testing.B) {
	unmarshalLogs(unmarshalTarget)
}

func BenchmarkUnmarshalLogsWithGoroutine(b *testing.B) {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	var branchInformations []*BranchInformation
	for _, gitLog := range unmarshalTarget {
		wg.Add(1)
		go func(log string) {
			defer wg.Done()
			branchInformation := BranchInformation{}
			json.Unmarshal([]byte(log), &branchInformation)
			mutex.Lock()
			branchInformations = append(branchInformations, &branchInformation)
			mutex.Unlock()
		}(gitLog)
	}
	wg.Wait()
}

func BenchmarkTieOldBranchToAuthorDefault(b *testing.B) {
	branchInformations := unmarshalLogs(unmarshalTarget)
	tieOldBranchToAuthor(branchInformations)
}
func BenchmarkTieOldBranchToAuthorWithGoroutine(b *testing.B) {
	branchInformations := unmarshalLogs(unmarshalTarget)
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	now := time.Now()
	branchCommiterMap := make(map[string][]*BranchInformation)
	for _, branchInformation := range branchInformations {
		wg.Add(1)
		go func(branchInformation *BranchInformation, now time.Time) {
			defer wg.Done()
			days := int(now.Sub(branchInformation.LastCommitDate.Time).Hours()) / 24
			if days >= 14 {
				mutex.Lock()
				branchCommiterMap[branchInformation.CommiterName] = append(branchCommiterMap[branchInformation.CommiterName], branchInformation)
				mutex.Unlock()
			}
		}(branchInformation, now)
	}
	wg.Wait()
}

const logCount = 10000

var unmarshalTarget = makeDummyLogs(logCount)

func makeDummyLogs(n int) []string {
	dummyLogs := make([]string, n)
	for {
		if n < 0 {
			break
		}
		dummyLogs = append(dummyLogs, fmt.Sprintf(`{
			"branch_name": "%s",
			"commiter_name": "%s",
			"last_commit_date": "2018-11-01T22:49:32+09:00"
		}`, randString(16), randString(16)))
		n--
	}
	return dummyLogs
}

var randSrc = rand.NewSource(time.Now().UnixNano())

const (
	rs6Letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rs6LetterIdxBits = 6
	rs6LetterIdxMask = 1<<rs6LetterIdxBits - 1
	rs6LetterIdxMax  = 63 / rs6LetterIdxBits
)

func randString(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), rs6LetterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rs6LetterIdxMax
		}
		idx := int(cache & rs6LetterIdxMask)
		if idx < len(rs6Letters) {
			b[i] = rs6Letters[idx]
			i--
		}
		cache >>= rs6LetterIdxBits
		remain--
	}
	return string(b)
}
