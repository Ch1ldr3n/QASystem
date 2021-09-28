package main

import (
	"flag"
	qanda "gitlab.secoder.net/bauhinia/qanda/backend/pkg"
)

var listen = flag.String("listen", "127.0.0.1:8080", "listen address")

func main() {
	flag.Parse()
	qanda.Echo.Logger.Fatal(qanda.Echo.Start(*listen))
}
