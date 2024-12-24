package server

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-kipi/let-me-know/lib/config"
	h "github.com/go-kipi/let-me-know/lib/http"
	"go.uber.org/fx"
	"net"
	"net/http"
	"os"
	"time"
)

func ServerHTTP(lc fx.Lifecycle, config *config.Config, s *h.Server) *gin.Engine {
	serviceName := os.Args[0]
	router := gin.Default()
	router.TrustedPlatform = gin.PlatformGoogleAppEngine
	router.TrustedPlatform = "X-CDN-IP"
	router.SetTrustedProxies([]string{})
	router.RemoteIPHeaders = []string{"X-Forwarded-For", "X-Real-IP"}
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(cors.Default())

	routes(router, s)
	srv := &http.Server{Addr: ":" + config.HTTP.ListenAddress, Handler: router} // define a web server

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr) // the web server starts listening on 8080
			if err != nil {

				fmt.Println(fmt.Sprintf("[%s]Failed to start HTTP Server at %s", serviceName, srv.Addr))
				return err
			}
			go srv.Serve(ln) // process an incoming request in a go routine pid :=
			fmt.Println(fmt.Sprintf("pid:%d: [%s]Succeeded to start HTTP Server at %s", os.Getpid(), serviceName, srv.Addr))

			return nil

		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(ctx) // stop the web server
			fmt.Println(fmt.Sprintf("[%s] HTTP Server is stopped", serviceName))
			return nil
		},
	})

	return router
}
