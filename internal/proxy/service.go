package proxy

import (
	"LoadBalancer/internal/selector"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type ProxyHandler struct {
	selector selector.Selector
}

func NewProxyHandler(s selector.Selector) *ProxyHandler {
	return &ProxyHandler{selector: s}
}

func (p *ProxyHandler) Connect(c *gin.Context) error {
	log.Println("[ProxyHandler] Connect")
	server, err := p.selector.Select()
	if err != nil {
		log.Println(err)
	}
	target, err := url.Parse(server)
	proxy := httputil.NewSingleHostReverseProxy(target)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = target.Host
		req.Header.Set("X-Forwarded-Host", req.Host)
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Proxy Error: %v", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}
	proxy.ServeHTTP(c.Writer, c.Request)
	return nil
}
