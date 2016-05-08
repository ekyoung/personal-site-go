package root

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func ResumeController(c *gin.Context) {
    c.HTML(http.StatusOK, "resume", gin.H{
        "title":          "Resume",
        "isResumeActive": true,
    })
}
