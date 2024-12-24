package reply

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	UnauthorizedStatus = "unauthorized"
	RejectedStatus     = "rejected"
	NoMatchStatus      = "noMatch"
	ErrorStatus        = "error"
	InvalidStatus      = "invalid"
	SuccessStatus      = "success"
)

type BaseResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message,omitempty"`
}

type Response struct {
	BaseResponse
	Data interface{} `json:"data,omitempty"`
}

func ErrorReplay(c *gin.Context, statusCode int, err interface{}) {
	r := Response{
		BaseResponse: BaseResponse{
			Status:  ErrorStatus,
			Message: fmt.Sprint(err),
		},
	}
	c.JSON(statusCode, r)
}

func GinSuccessReply(c *gin.Context, r interface{}) {
	serviceReply := Response{}
	serviceReply.Status = SuccessStatus
	serviceReply.Data = r
	c.JSON(http.StatusOK, serviceReply)
}
