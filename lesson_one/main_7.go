package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 去进行服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})
	//  openssl genrsa -out ca-key.pem 1024
	//  465  openssl req -new -key ca-key.pem -out ca-req.csr -subj "/C=CN/ST=BJ/L=BJ/O=fish/OU=fish/CN=CA"
	//  466  openssl x509 -req -in ca-req.csr -out ca-cert.pem -signkey ca-key.pem -days 3650
	//  467  openssl genrsa -out server-key.pem 1024
	//  468  openssl req -new -out server-req.csr -key server-key.pem -subj "/C=CN/ST=BJ/L=BJ/O=fish/OU=fish/CN=*.fish-test.com"
	//  469  openssl x509 -req -in server-req.csr -out server-cert.pem -signkey server-key.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -days 3650
	//  470  openssl verify -CAfile ca-cert.pem  server-cert.pem
	//  471  openssl genrsa -out client-key.pem 1024
	//  472  openssl req -new -out client-req.csr -key client-key.pem -subj "/C=CN/ST=BJ/L=BJ/O=fish/OU=fish/CN=dong"
	//  473  openssl x509 -req -in client-req.csr -out client-cert.pem -signkey client-key.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -days 3650
	//  474  openssl verify -CAfile ca-cert.pem client-cert.pem
	pem_path := "/Users/muller/Documents/gowork/src/gin_tutorial/lesson_one/server-cert.pem"
	key_path := "/Users/muller/Documents/gowork/src/gin_tutorial/lesson_one/server-key.pem"

	// 监听并服务于 https://127.0.0.1:8080
	//r.Run(":8099")
	r.RunTLS(":8088", pem_path, key_path)
}
