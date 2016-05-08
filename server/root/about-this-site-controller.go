package root

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func AboutThisSiteController(c *gin.Context) {
    c.HTML(http.StatusOK, "about-this-site", gin.H{
        "title":                 "About This Site",
        "isAboutThisSiteActive": true,
    })
}
