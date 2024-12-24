package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kipi/let-me-know/lib/config"
	"github.com/go-kipi/let-me-know/lib/http"
	"github.com/go-kipi/let-me-know/lib/http/server"
	"github.com/go-kipi/let-me-know/lib/mongo"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		mongo.Module,
		http.Module,
		fx.Provide(
			server.ServerHTTP,
		),
		fx.Invoke(func(*gin.Engine) {}),
	)
	app.Run()
}
