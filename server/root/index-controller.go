package root

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func IndexController(c *gin.Context) {
    c.HTML(http.StatusOK, "home", gin.H{
        "title":        "Home",
        "isHomeActive": true,
    })
}
