package handlers

import (
	"context"
	"github.com/go-kipi/let-me-know/lib/mongo"
	"time"
)

type Req struct {
	Id string `json:"id"`
}

func IsAlive(c context.Context, req Req, mongo mongo.MongoI) (interface{}, error) {

	return time.Now(), nil
}
