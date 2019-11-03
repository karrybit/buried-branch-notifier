package slack

import (
	"buried-branch-notifier/model/git"
	"fmt"
	"time"
)

type Attachment struct {
	Color  string  `json:"color"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func NewAttachments(branchCommiterMap map[string][]*git.BranchInformation) []Attachment {
	var attachments []Attachment
	now := time.Now()
	for commiter, branchInformations := range branchCommiterMap {
		attachment := newAttachment(commiter, branchInformations, now)
		attachments = append(attachments, *attachment)
	}
	return attachments
}

func newAttachment(author string, branchInformations []*git.BranchInformation, now time.Time) *Attachment {
	attachment := Attachment{Color: "danger"}
	message := buildMessage(branchInformations, now)
	attachment.Fields = append(attachment.Fields, Field{Title: author, Value: message, Short: false})
	return &attachment
}

func buildMessage(branchInformations []*git.BranchInformation, now time.Time) string {
	var buffer []byte
	for _, branchInformation := range branchInformations {
		message := fmt.Sprintf("%s has stopped from %s (%vdays ago)\n", branchInformation.BranchName, branchInformation.LastCommitDate.Format("2006-01-02"), int(now.Sub(branchInformation.LastCommitDate.Time).Hours())/24)
		buffer = append(buffer, message...)
	}
	return string(buffer)
}
