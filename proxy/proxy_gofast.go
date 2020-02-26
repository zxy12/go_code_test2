package main

import (
	_ "fmt"
	"log"
	"net/http"
	_ "os"

	"github.com/yookoala/gofast"
)

func main() {
	// Get fastcgi application server tcp address
	// from env FASTCGI_ADDR. Then configure
	// connection factory for the address.
	address := "127.0.0.1:9000"
	connFactory := gofast.SimpleConnFactory("tcp", address)

	// route all requests to a single php file
	http.Handle("/", gofast.NewHandler(
		//gofast.NewFileEndpoint("/Users/zhouxinyu/www/localhost/src/github.com/zxy12/go_code_test/proxy/html/index.php")(gofast.BasicSession),
		gofast.NewPHPFS("/Users/zhouxinyu/www/localhost/src/github.com/zxy12/go_code_test/proxy/html/")(gofast.BasicSession),
		gofast.SimpleClientFactory(connFactory, 0),
	))

	// serve at 8080 port
	log.Fatal(http.ListenAndServe(":8089", nil))
}
