package main

import (
	"flag"
	qanda "gitlab.secoder.net/bauhinia/qanda/backend/pkg"
)

var listen = flag.String("listen", "127.0.0.1:8080", "listen address")
var serve = flag.String("serve", "/usr/share/qanda", "static artifacts to serve")
var storage = flag.String("storage", "sqlite3", "database type")
var database = flag.String("database", "file:ent?mode=memory&cache=shared&_fk=1", "database connection string")
var key = flag.String("key", "super-secret-key", "jwt secret key for tokens")

func main() {
	flag.Parse()
	e := qanda.New(*serve, *storage, *database, *key)
	e.Logger.Fatal(e.Start(*listen))
}
