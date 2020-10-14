package main

import (
	"flag"
	"net/http"

	"cfg/handler"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	http.HandleFunc("/group", handler.GetDefaultGroup().Handler)
	http.HandleFunc("/tcp", handler.GetDefaultTCP().Handler)
	http.HandleFunc("/dns", handler.GetDefaultDNS().Handler)
	addr := ":8443"
	err := http.ListenAndServeTLS(addr, "cert/server.crt", "cert/server.key", nil)
	if err != nil {
		glog.Info(err)
	}
}
