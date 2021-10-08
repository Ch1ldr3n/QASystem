package main

import (
	"flag"
	qanda "gitlab.secoder.net/bauhinia/qanda/backend/pkg"
)

var listen = flag.String("listen", "127.0.0.1:8080", "listen address")
var serve = flag.String("serve", "/usr/share/qanda", "static artifacts to serve")

func main() {
	flag.Parse()
	e := qanda.New(*serve)
	e.Logger.Fatal(e.Start(*listen))
}
