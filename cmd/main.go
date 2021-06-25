package main

import (
	"flag"
	"os"

	"Auth/initiator"
	//_ "github.com/lib/pq"
)

var testInit bool

func init() {
	flag.BoolVar(&testInit, "test", false, "initialize test mode without serving")
	flag.Parse()

	os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	initiator.User(testInit)
}
