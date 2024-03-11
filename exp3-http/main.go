package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "Essa é uma requisição GET")
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "Essa é uma requisição POST")
	})

	r.GET("/unauthorized", func(c *gin.Context) {
		c.String(http.StatusUnauthorized, "Unauthorized Access (401)")
	})

	r.GET("/forbidden", func(c *gin.Context) {
		c.String(http.StatusForbidden, "Forbidden Access (403)")
	})

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Resource Not Found (404)")
	})

	r.GET("/error", func(c *gin.Context) {
		c.String(http.StatusInternalServerError, "Internal Server Error (500)")
	})

	r.Run("127.0.0.1:7000")
}
