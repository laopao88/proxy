package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"net/url"
)

var remoteUrl string

func StartProxy(e *echo.Echo) {
	sensitiveUrl, err := url.Parse(remoteUrl)
	if err != nil {
	} else {
		e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
			Skipper: func(c echo.Context) bool {
				return false
			},
			Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{{URL: sensitiveUrl}}),
			Rewrite: map[string]string{
				"/*": "/$1",
			},
			ModifyResponse: func(response *http.Response) error {
				response.Header.Set(echo.HeaderAccessControlAllowOrigin, "*")
				response.Header.Set(echo.HeaderAccessControlAllowHeaders, "*")
				response.Header.Set(echo.HeaderAccessControlAllowMethods, "*")
				response.Header.Set(echo.HeaderAccessControlMaxAge, "86400")
				return nil
			},
		}))
	}
}

func main() {
	var localPort string
	flag.StringVar(&localPort, "lp", "", "Local Port")
	flag.StringVar(&remoteUrl, "rp", "", "Remote Url")
	flag.Parse()
	if localPort == "" || remoteUrl == "" {
		flag.Usage()
		return
	}
	e := echo.New()
	StartProxy(e)
	if err := e.Start(":" + localPort); err != nil {
	}
}
