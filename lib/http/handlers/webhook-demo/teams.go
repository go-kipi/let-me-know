package webhook_demo

import (
	"context"
	"fmt"
	"github.com/go-kipi/let-me-know/lib/mongo"
	"github.com/go-kipi/let-me-know/lib/utiles"
)

type ReqTeams struct {
	Content string `json:"content"`
}

type TeamsBody struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type TeamsContent struct {
	Schema  string      `json:"$schema"`
	Type    string      `json:"type"`
	Version string      `json:"version"`
	Body    []TeamsBody `json:"body"`
}
type TeamsAttachments struct {
	ContentType string       `json:"contentType"`
	ContentUrl  interface{}  `json:"contentUrl"`
	Content     TeamsContent `json:"content"`
}

type TeamsReq struct {
	Type        string             `json:"type"`
	Attachments []TeamsAttachments `json:"attachments"`
}

func Teams(c context.Context, req ReqTeams, mongo mongo.MongoI) (interface{}, error) {
	//get User Details
	//
	host := "https://heilasystems.webhook.office.com/webhookb2"
	token := "cac31881-d766-4f89-a9d3-0dab3a3f138d@ddd0ba21-5c5e-4b3f-8ccf-2ca607869bec/IncomingWebhook/3cbb8fcf8f574ffaa5205ff68a7fe588/1fe824bb-c981-44a2-9ff2-673e3da89d58/V2mSFYnhfVVxxsIDy9x3PbpR2cER3Ho-22IxrGAlSiWKM1"

	url := fmt.Sprintf("%s/%s", host, token)
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	body := TeamsBody{
		Type: "TextBlock",
		Text: req.Content,
	}
	teamsContent := TeamsContent{
		Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
		Type:    "AdaptiveCard",
		Version: "1.2",
		Body:    append([]TeamsBody{}, body),
	}

	teamsAttachments := TeamsAttachments{
		ContentType: "application/vnd.microsoft.card.adaptive",
		ContentUrl:  nil,
		Content:     teamsContent,
	}

	data := TeamsReq{
		Type:        "message",
		Attachments: append([]TeamsAttachments{}, teamsAttachments),
	}

	res, err := utiles.PostReq(url, headers, data)
	if err != nil {
		return nil, err
	}

	fmt.Println(res)

	return nil, nil
}
