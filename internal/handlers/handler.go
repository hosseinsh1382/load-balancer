package handlers

import (
	"LoadBalancer/internal/proxy"
	"log"

	"github.com/gin-gonic/gin"
)

type handler struct {
	proxyHandler proxy.ProxyHandler
	errorLogger  *log.Logger
	infoLogger   *log.Logger
}

func NewHandler(proxyHandler proxy.ProxyHandler, errorLogger *log.Logger, infoLogger *log.Logger) *handler {
	return &handler{
		proxyHandler: proxyHandler,
		errorLogger:  errorLogger,
		infoLogger:   infoLogger,
	}
}

func (h *handler) HandleAll(c *gin.Context) {
	err := h.proxyHandler.Connect(c)
	if err != nil {
		log.Println(err)
	}
}
