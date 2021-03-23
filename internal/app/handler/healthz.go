package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHealthz(c *gin.Context) {
	c.String(http.StatusOK, "i am ok\n")
}
