package main

import (
	"errors"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	urlVal := os.Getenv("PROX_URL")
  if urlVal == "" {
    panic(errors.New("Expecting an PROXY_URL value"))
  }
	remote, err := url.Parse(urlVal)
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)
			r.Host = remote.Host
			w.Header().Set("X-Ben", "Rad")
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":8086", nil)
	if err != nil {
		panic(err)
	}
}
