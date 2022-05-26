package main

import (
	"demo_project/config"
	"demo_project/controller"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	//. "dao"
	//. "logs"
	//. "models"
	"demo_project/util"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Credentials", "true")
		// c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		// if c.Request.Method == "OPTIONS" {
		// 	c.AbortWithStatus(204)
		// 	return
		// }
		c.Next()
	}
}

func main() {

	config.InitDB()
	util.LogSet()
	r := gin.Default()
	r.Use(Cors())

	v1 := r.Group("demoapi/v1")
	{
		// Event Builder //
		v1.POST("/contact/us", controller.ContactUs)
	}

	r.Run(":7000")
	//r.RunTLS(":7002", "/var/www/html/go_projects/src/event_builder/4b3c8a054fb0c0ec.crt", "/var/www/html/go_projects/src/event_builder/multitv.key")

	//http.HandleFunc("/channels", controller.ChannelIndex)

}

/*
	func main() {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
*/
