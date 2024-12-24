package webhook_demo

import (
	"context"
	"fmt"
	"github.com/go-kipi/let-me-know/lib/mongo"
	"github.com/go-kipi/let-me-know/lib/utiles"
)

type Req struct {
	Content string `json:"content"`
}

func Discord(c context.Context, req Req, mongo mongo.MongoI) (interface{}, error) {
	username := "kipi"
	host := "https://discord.com/api/webhooks"
	token := "1295365649719365682/MjDpqWEwiZzSHZSJHf8RLCrHZw9j19xymsKcEl2ftAiU57fJDC-VkTkm4I4TKZOtwiBe"

	discordUrl := fmt.Sprintf("%s/%s", host, token)
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	data := map[string]interface{}{
		"username": username,
		"content":  req.Content,
	}
	return utiles.PostReq(discordUrl, headers, data)
}
