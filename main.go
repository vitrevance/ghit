package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	path := flag.String("p", "", "path to web application root directory")
	addr := flag.String("a", "0.0.0.0:80", "address to listen on")

	flag.Parse()

	if *path == "" {
		log.Fatal("path to web application is not provided")
	}

	server := echo.New()
	server.GET("/*", echo.WrapHandler(http.FileServer(http.FS(os.DirFS(*path)))))

	server.Start(*addr)
}
