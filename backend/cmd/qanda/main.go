package main

import (
	"flag"
	qanda "gitlab.secoder.net/bauhinia/qanda/backend/pkg"
)

var listen = flag.String("listen", "127.0.0.1:8080", "listen address")

func main() {
	flag.Parse()
	e := qanda.New()
	e.Logger.Fatal(e.Start(*listen))
}
