package root

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ekyoung/personal-site-go/gin-wrapper"
)

type RootController struct{}

func (ctrl *RootController) Index(c wrapper.Context) {
	c.HTML(http.StatusOK, "home", gin.H{
		"title":        "Home",
		"isHomeActive": true,
	})
}
func (ctrl *RootController) AboutThisSite(c wrapper.Context) {
	c.HTML(http.StatusOK, "about-this-site", gin.H{
		"title":                 "About This Site",
		"isAboutThisSiteActive": true,
	})
}

func (ctrl *RootController) Resume(c wrapper.Context) {
	c.HTML(http.StatusOK, "resume", gin.H{
		"title":          "Resume",
		"isResumeActive": true,
	})
}

func (ctrl *RootController) PageNotFound(c wrapper.Context) {
	c.HTML(http.StatusNotFound, "page-not-found", gin.H{
		"title": "Page Not Found",
	})
}

func (ctrl *RootController) RecoveryHandler(c *gin.Context, err interface{}) {
	c.HTML(http.StatusInternalServerError, "error", gin.H{
		"title": "Error",
	})
}
