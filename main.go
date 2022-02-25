package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

func index(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, strings.Join(v, " "))
	}

	w.Header().Set("env", strings.Join(os.Environ(), "\n"))

	remoteAddr := r.RemoteAddr
	if ip := r.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = r.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	req := make(map[string]interface{}, 0)
	req["status"] = 200
	req["msg"] = "success"
	req["ip"] = remoteAddr
	reqByte, _ := json.Marshal(req)
	w.Write(reqByte)
}

func main() {

	http.HandleFunc("/healthz", index)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
