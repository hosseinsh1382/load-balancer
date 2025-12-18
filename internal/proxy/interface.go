package proxy

import "github.com/gin-gonic/gin"

type Handler interface {
	Connect(c *gin.Context, chn chan error)
}
