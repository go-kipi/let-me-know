package handlers

import (
	"context"
	"fmt"
	"github.com/go-kipi/let-me-know/lib/mongo"
)

type Req2 struct {
	Id string `json:"id"`
}

func IsAlive2(c context.Context, req Req2, mongo mongo.MongoI) (interface{}, error) {
	fmt.Println(c, req)
	return req.Id, nil
}
