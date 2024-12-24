package server

import (
	"github.com/gin-gonic/gin"
	h "github.com/go-kipi/let-me-know/lib/http"
	"github.com/go-kipi/let-me-know/lib/http/handlers"
	webhook_demo "github.com/go-kipi/let-me-know/lib/http/handlers/webhook-demo"
	"github.com/mandrigin/gin-spa/spa"
)

func routes(r *gin.Engine, s *h.Server) {
	webhook := r.Group("/webhook")
	{
		webhook.POST("/")
	}
	webhookTest := r.Group("/webhook-demo")
	{
		webhookTest.POST("/teams", s.Handle(webhook_demo.Teams))
		webhookTest.POST("/discord", s.Handle(webhook_demo.Discord))
	}
	r.GET("/isAlive", s.Handle(handlers.IsAlive))
	r.POST("/isAlive2", s.Handle(handlers.IsAlive2))
	//r.POST("/getText", s.Handle(handler.GetText))
	//r.POST("/setText", s.Handle(handler.SetText))

	r.Use(spa.Middleware("/", "./frontend/dist/"))
}
