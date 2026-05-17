package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rtnl/blaze/pkg/provider/craft"
	"github.com/samber/lo"
)

var (
	LOGIC = new(Logic)
)

func main() {
	r := gin.Default()

	r.GET("/craft/version/:version_query/:target", func(c *gin.Context) {
		versionQuery := strings.ToLower(c.Param("version_query"))
		if versionQuery == "" {
			versionQuery = VersionQueryLatest
		}

		downloadTarget := strings.ToLower(c.Param("target"))
		if !lo.Contains([]string{VersionDownloadTargetServer, VersionDownloadTargetClient}, downloadTarget) {
			c.Status(400)
			return
		}

		versionKind, ok := craft.VersionKindFromString(c.DefaultQuery("kind", string(craft.VersionKindRelease))).Get()
		if !ok {
			c.Status(400)
			return
		}

		downloadUrl, err := LOGIC.GetVersionDownloadUrl(versionQuery, versionKind, downloadTarget).Get()
		if err != nil {
			c.Status(500)
			return
		}

		c.Redirect(302, downloadUrl)
		return
	})

	r.Run()
}
