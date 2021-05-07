package rest

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gin-gonic/gin"
)

// V1Delegate handles "/rest/v1" routes.
func V1Delegate(group *gin.RouterGroup, app *inits.App) {
	group.POST("/oauth/token", oauthTokenHandler(app))
}
