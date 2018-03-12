package main

import (
	"mime"
	"net/http"
	"strings"

	"github.com/cnosuke/go_blank_image/blank_image"
	"github.com/cnosuke/gotrack/recorder"
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
)

func isValidUUID(t string) bool {
	return len(t) != 0
}

func generateUUID() string {
	return uuid.New()
}

func setUserUUID(c *gin.Context) (id string) {
	k, _ := c.Cookie(cookieKey)

	if isValidUUID(k) {
		id = k
	} else {
		id = generateUUID()
	}

	c.SetCookie(
		cookieKey,
		id,
		cookieMaxAge,
		cookiePath,
		cookieDomain,
		cookieSecure,
		false,
	)

	return
}

func blankPageHandler(c *gin.Context) {
	trackGroup := c.Param("group")
	trackKey := c.Param("key")

	id := setUserUUID(c)
	recorder.Get().Post(trackGroup, trackKey, id)

	c.String(http.StatusNoContent, "")
}

func isValidExt(ext string) bool {
	return ext == "png" || ext == "gif"
}

func splitKeyAndExt(s string) (key string, ext string) {
	pair := strings.SplitN(s, ".", 2)

	if len(pair) == 2 {
		key = pair[0]
		ext = pair[1]
	} else {
		key = pair[0]
		ext = "png"
	}

	return
}

func blankImageHandler(c *gin.Context) {
	trackGroup := c.Param("group")
	trackKey, ext := splitKeyAndExt(c.Param("key"))

	if !isValidExt(ext) {
		c.String(http.StatusNotFound, "NotFound")
		return
	}

	id := setUserUUID(c)
	recorder.Get().Post(trackGroup, trackKey, id)

	switch ext {
	case "png":
		c.Data(
			http.StatusOK,
			mime.TypeByExtension(".png"),
			blank_image.PngBytes(),
		)
	case "gif":
		c.Data(
			http.StatusOK,
			mime.TypeByExtension(".gif"),
			blank_image.GifBytes(),
		)
	}

	return
}
