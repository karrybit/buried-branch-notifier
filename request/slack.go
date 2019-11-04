package request

import (
	"buried-branch-notifier/model/git"
	"buried-branch-notifier/model/slack"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Requester struct {
	httpClient        *http.Client
	branchCommiterMap map[string][]*git.BranchInformation
	branchCount       int
}

const urlString = ""

// New is function to initialize Client
func NewRequester(branchCommiterMap map[string][]*git.BranchInformation, branchCount int) (*Requester, error) {
	requester := Requester{httpClient: &http.Client{Timeout: time.Duration(10) * time.Second}, branchCommiterMap: branchCommiterMap, branchCount: branchCount}
	return &requester, nil
}

func (r *Requester) Notify() error {
	bodyByte, _ := json.Marshal(struct {
		UserName    string             `json:"username"`
		IconEmoji   string             `json:"icon_emoji"`
		Text        string             `json:"text"`
		Attachments []slack.Attachment `json:"attachments"`
	}{
		"Buried Branch Notifier",
		":zawazawa:",
		fmt.Sprintf("*The following branches have not developed for more than 2 weeks. There are a total of %d brunches. Let's purge!!*", r.branchCount),
		slack.NewAttachments(r.branchCommiterMap),
	})
	bodyReader := bytes.NewReader(bodyByte)

	request, err := http.NewRequest(http.MethodPost, urlString, bodyReader)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err := r.httpClient.Do(request)
	if err != nil {
		return err
	}

	response.Body.Close()
	return nil
}
