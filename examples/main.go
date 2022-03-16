// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"log"
	"net/http"
	"time"

	prom_http_exporter "github.com/pan-net-security/go-prometheus-http-exporter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func foo(w http.ResponseWriter, r *http.Request) {
	time.Sleep(200 * time.Millisecond)
	w.WriteHeader(301)
	w.Write([]byte("foo"))
}

func bar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	w.Write([]byte("bar"))
}

func main() {
	e := prom_http_exporter.New()

	r := http.NewServeMux()
	r.Handle(e.Metric("/foo", foo))
	r.Handle(e.Metric("/bar", bar))
	r.Handle("/metrics", promhttp.Handler())

	s := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      r,
	}
	log.Fatalln(http.ListenAndServe(":8080", s.Handler))
}
