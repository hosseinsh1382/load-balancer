package main

import (
	"LoadBalancer/internal/handlers"
	"LoadBalancer/internal/proxy"
	"LoadBalancer/internal/selector"
	"LoadBalancer/internal/server"
	"LoadBalancer/pkg"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	fmt.Println("Starting Server")
	errorLogger := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger := log.New(os.Stdout, "INFO  ", log.Ldate|log.Ltime|log.Lshortfile)
	r := pkg.NewDefaultJsonReader[[]server.Registry]("./configs/servers.json")
	hl := server.NewJsonHolder(r)
	s := selector.NewRoundRobin(hl)
	p := proxy.NewProxyHandler(s)
	h := handlers.NewHandler(p, errorLogger, infoLogger)
	e := gin.Default()
	e.Any("/*path", h.HandleAll)
	e.Run(":8085")
}
