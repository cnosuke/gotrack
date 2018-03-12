package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cnosuke/gotrack/recorder"
	"github.com/cnosuke/gotrack/recorder/stdout"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

var (
	Name     string
	Version  string
	Revision string

	binding      string
	cookieKey    string
	cookieDomain string
	cookiePath   string
	cookieSecure bool
	cookieMaxAge int
)

func main() {
	app := cli.NewApp()
	app.Version = fmt.Sprintf("%s (%s)", Version, Revision)
	app.Name = Name

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "binding, b",
			Usage:       "Server binding address",
			Value:       "127.0.0.1:8080",
			Destination: &binding,
		},
		cli.StringFlag{
			Name:        "key",
			Usage:       "Cookie Key",
			Value:       "gotrack",
			Destination: &cookieKey,
		},
		cli.StringFlag{
			Name:        "domain",
			Usage:       "Cookie domain",
			Value:       "",
			Destination: &cookieDomain,
		},
		cli.StringFlag{
			Name:        "path",
			Usage:       "Cookie pass",
			Value:       "/",
			Destination: &cookiePath,
		},
		cli.BoolFlag{
			Name:        "secure",
			Usage:       "Cookie secure",
			Destination: &cookieSecure,
		},
		cli.IntFlag{
			Name:        "maxAge",
			Usage:       "Cookie Max age",
			Value:       10 * 365 * 24 * 60 * 60, // 10 years
			Destination: &cookieMaxAge,
		},
	}

	app.Action = func(c *cli.Context) error {
		setupRecorder()

		r := gin.Default()
		r.GET("/v1/i/:group/:key", blankImageHandler)
		r.GET("/v1/p/:group/:key", blankPageHandler)

		// Health check
		r.GET("/site/sha", func(c *gin.Context) {
			c.String(http.StatusOK, Revision)
		})

		r.Run(binding)

		return nil
	}

	app.Run(os.Args)
}

func setupRecorder() {
	// Sample: Set STDOUT as recorder.
	rec := stdout.NewStdoutRecorder()
	recorder.SetGlobal(rec)
}
