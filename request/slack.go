package request

import (
	"buried-branch-notifyer/model/git"
	"buried-branch-notifyer/model/slack"
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type Requester struct {
	url               *url.URL
	httpClient        *http.Client
	branchCommiterMap map[string][]*git.BranchInformation
}

const urlString = ""

// New is function to initialize Client
func NewRequester(branchCommiterMap map[string][]*git.BranchInformation) (*Requester, error) {
	url, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	requester := Requester{httpClient: &http.Client{Timeout: time.Duration(10) * time.Second}, branchCommiterMap: branchCommiterMap}
	requester.url = url
	return &requester, nil
}

func (r *Requester) Notify() error {
	bodyByte, _ := json.Marshal(struct {
		UserName    string             `json:"username"`
		IconEmoji   string             `json:"icon_emoji"`
		Text        string             `json:"text"`
		Attachments []slack.Attachment `json:"attachments"`
	}{
		"Stalin",
		":zawazawa:",
		"*The following branches have not moved for more than 2 weeks. Let's purge!!*",
		slack.NewAttachments(r.branchCommiterMap),
	})
	bodyReader := bytes.NewReader(bodyByte)

	request, err := http.NewRequest(http.MethodPost, r.url.String(), bodyReader)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err := r.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	return nil
}
