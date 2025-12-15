package main

import (
	"LoadBalancer/internal/proxy"
	"LoadBalancer/internal/selector"
	"log"

	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	r := gin.Default()

	r.Any("/*path", func(c *gin.Context) {
		s := selector.NewStaticSelector()
		proxy := proxy.NewProxyHandler(s)

		err := proxy.Connect(c)
		if err != nil {
			log.Println(err)
		}
	})
	r.Run(":8085")
}
