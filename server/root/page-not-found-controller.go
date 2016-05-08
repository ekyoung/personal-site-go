package root

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func PageNotFoundController(c *gin.Context) {
    c.HTML(http.StatusNotFound, "page-not-found", gin.H{
        "title": "Page Not Found",
    })
}
