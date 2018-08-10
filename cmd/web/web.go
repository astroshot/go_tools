package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	web "modules/pkg/web"
)

func main() {
	addr := "localhost:%s"
	port := flag.String("port", "8000", "http listen port")
	flag.Parse()
	addr = fmt.Sprintf(addr, *port)
	fmt.Println("Listening port %s...", *port)
	web.RegistHandlers()
	log.Fatal(http.ListenAndServe(addr, nil))
}
