// Copyright 2021 martinohmann
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	goatopsfarm "github.com/martinohmann/goatops.farm"
	"github.com/martinohmann/goatops.farm/gen/creatures"
	creaturessvr "github.com/martinohmann/goatops.farm/gen/http/creatures/server"
	staticsvr "github.com/martinohmann/goatops.farm/gen/http/static/server"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

func startServer(ctx context.Context, listenAddr string, creaturesEndpoints *creatures.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, redirectHTTPS, debug bool) {
	// Setup goa log adapter.
	adapter := middleware.NewLogger(logger)

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	mux := goahttp.NewMuxer()

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	fs := http.FS(goatopsfarm.StaticFS)
	eh := errorHandler(logger)

	creaturesServer := creaturessvr.New(creaturesEndpoints, mux, dec, enc, eh, nil)
	staticServer := staticsvr.New(nil, mux, dec, enc, eh, nil, fs, fs, fs)

	if debug {
		servers := goahttp.Servers{
			creaturesServer,
			staticServer,
		}
		servers.Use(httpmdlwr.Debug(mux, os.Stdout))
	}

	// Configure the mux.
	staticsvr.Mount(mux, staticServer)
	creaturessvr.Mount(mux, creaturesServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux

	if redirectHTTPS {
		handler = redirect(handler)
	}
	handler = httpmdlwr.Log(adapter)(handler)
	handler = httpmdlwr.RequestID()(handler)

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: listenAddr, Handler: handler}

	for _, m := range staticServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	for _, m := range creaturesServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", listenAddr)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", listenAddr)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}

func redirect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proto := r.Header.Get("X-Forwarded-Proto")
		if proto == "http" || proto == "HTTP" {
			http.Redirect(w, r, fmt.Sprintf("https://%s%s", r.Host, r.URL), http.StatusPermanentRedirect)
			return
		}

		h.ServeHTTP(w, r)
	})
}
