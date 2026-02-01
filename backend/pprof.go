package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
)

func init() {
	if os.Getenv("PPROF") == "1" {
		go http.ListenAndServe(":8080", nil)
	}
}
