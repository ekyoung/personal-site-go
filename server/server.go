package server

import (
    "github.com/gin-gonic/gin"

    "github.com/ekyoung/personal-site-go/server/root"
)

func RenderPageNotFound(c *gin.Context) {
    root.PageNotFoundController(c)
}
