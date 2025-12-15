package proxy

import "github.com/gin-gonic/gin"

type proxy_handler interface {
	Connect(c *gin.Context)
}
