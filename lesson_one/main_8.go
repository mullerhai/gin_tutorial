package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	router := gin.Default()
	router.Use(TlsHandler())
	pem_path := "/Users/muller/Documents/gowork/src/gin_tutorial/lesson_one/server-cert.pem"
	key_path := "/Users/muller/Documents/gowork/src/gin_tutorial/lesson_one/server-key.pem"

	router.RunTLS(":8099", pem_path, key_path)
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
