package proxy

import "github.com/gin-gonic/gin"

type ProxyHandler interface {
	Connect(c *gin.Context) error
}
