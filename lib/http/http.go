package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kipi/let-me-know/lib/http/reply"
	"github.com/go-kipi/let-me-know/lib/mongo"
	"reflect"
	"strings"
)

type Server struct {
	mongo mongo.MongoI
	//appwrite
}

func (s *Server) Handle(ht interface{}) func(context *gin.Context) {
	return func(context *gin.Context) {
		request := createInnerHandlers(reflect.ValueOf(getHandlerRequestStruct(ht)))
		if context.Request.Method != "GET" && context.Request.Method != "DELETE" && !strings.Contains(context.GetHeader("Content-Type"), "multipart/form-data") {
			if err := context.ShouldBindJSON(&request); err != nil {
				fmt.Println("Cannot parse request to struct ")
				reply.ErrorReplay(context, 500, fmt.Errorf("Cannot parse request to struct ", err))
				return
			}
		}
		exec := func() (interface{}, error) {
			c := reflect.ValueOf(context.Request.Context())
			req := reflect.Indirect(reflect.ValueOf(request))
			m := (reflect.ValueOf(s.mongo))

			deps := []reflect.Value{c, req, m}

			responseArr := reflect.ValueOf(ht).Call(deps)
			if len(responseArr) == 1 {
				if !responseArr[0].IsNil() {
					return nil, fmt.Errorf("responseArr == 1")
				} else {
					return nil, nil
				}
			} else if len(responseArr) == 2 {
				if !responseArr[1].IsNil() {
					return responseArr[1].Interface(), fmt.Errorf("responseArr == 2")
				}
				return responseArr[0].Interface(), nil
			} else {
				err := fmt.Errorf("invalid response")
				return nil, err
			}
		}

		if response, err := exec(); err != nil {
			reply.ErrorReplay(context, 500, response)
		} else {
			reply.GinSuccessReply(context, response)
		}

	}

}

func getHandlerRequestStruct(f interface{}) interface{} {
	fType := reflect.TypeOf(f)
	argType := fType.In(1)
	return reflect.New(argType).Interface()
}

func createInnerHandlers(v reflect.Value) interface{} {
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	n := reflect.New(v.Type())
	return n.Interface()
}

func NewServer(mongo mongo.MongoI) *Server {
	return &Server{mongo: mongo}
}
