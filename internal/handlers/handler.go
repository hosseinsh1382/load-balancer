package handlers

import (
	"LoadBalancer/internal/proxy"
	"log"

	"github.com/gin-gonic/gin"
)

type handler struct {
	proxyHandler proxy.Handler
	errorLogger  *log.Logger
	infoLogger   *log.Logger
}

func NewHandler(proxyHandler proxy.Handler, errorLogger *log.Logger, infoLogger *log.Logger) *handler {
	return &handler{
		proxyHandler: proxyHandler,
		errorLogger:  errorLogger,
		infoLogger:   infoLogger,
	}
}

func (h *handler) HandleAll(c *gin.Context) {
	chn := make(chan error)
	go h.proxyHandler.Connect(c, chn)
	err := <-chn
	if err != nil {
		log.Println(err)
	}
}
