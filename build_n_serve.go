package main

// this is a localhost server to install PWA on local computer. also it builds the package in the folder 'dist'.
// command line flag '--buildonly' = only build package. no flag = build and serve at localhost.
// this will not build the server. to build the server, just goto 'backend' folder and `go build`.

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const LOCAL_SERVER_PORT = ":8080"

func buildFrontend() {
	// tanim uses 'esbuild' to bundle everything and it should have been installed by npm.

}

func watch() {

}

func serveLocal() {
	router := httprouter.New()
	router.ServeFiles("/*filepath", http.Dir("./dist/"))
	http.ListenAndServe(LOCAL_SERVER_PORT, router)
}

func main() {
	serveLocal()
}
