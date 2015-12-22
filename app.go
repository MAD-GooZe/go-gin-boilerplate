package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"os"
	"net/http"
	"github.com/MAD-GooZe/jaderender"
)

func main() {

	APP_PORT := os.Getenv("PORT")
	if (APP_PORT == "") {
		APP_PORT = "8080"
	}

	r := gin.Default()

	r.NoRoute(static.Serve("/", static.LocalFile("./public", true)))
	//r := gin.New()
	r.HTMLRender = jaderender.Default()

	r.GET("/", index)
	r.Run(":" + APP_PORT) // listen and serve on 0.0.0.0:APP_PORT
}

// Home page handler.
func index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", gin.H{
		"test": "test",
	})
}