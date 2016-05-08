package root

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type RootController struct{}

func (ctrl *RootController) Index(c *gin.Context) {
    c.HTML(http.StatusOK, "home", gin.H{
        "title":        "Home",
        "isHomeActive": true,
    })
}

func (ctrl *RootController) AboutThisSite(c *gin.Context) {
    c.HTML(http.StatusOK, "about-this-site", gin.H{
        "title":                 "About This Site",
        "isAboutThisSiteActive": true,
    })
}

func (ctrl *RootController) Resume(c *gin.Context) {
    c.HTML(http.StatusOK, "resume", gin.H{
        "title":          "Resume",
        "isResumeActive": true,
    })
}

func (ctrl *RootController) PageNotFound(c *gin.Context) {
    c.HTML(http.StatusNotFound, "page-not-found", gin.H{
        "title": "Page Not Found",
    })
}
