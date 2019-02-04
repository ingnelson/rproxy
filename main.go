package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("No!, need {{TARGET_URL}} {{LISTEN_PORT}}")
		return
	}
	remote, err := url.Parse(os.Args[1])
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.Handle("/", &ProxyHandler{proxy})
	err = http.ListenAndServe(":"+os.Args[2], nil)
	if err != nil {
		panic(err)
	}
}

type ProxyHandler struct {
	p *httputil.ReverseProxy
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Header().Set("X-Ben", "Rad")
	ph.p.ServeHTTP(w, r)
}
