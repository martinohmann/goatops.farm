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
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

var (
	factsFile  = flag.String("facts-file", "goatfacts.json", "Path to goatfacts")
	listenAddr = flag.String("listen-address", ":8080", "Listen address")
)

//go:generate sh -c "go run ./tools/gen_goatfacts.go > goatfacts.json"

func main() {
	rand.Seed(time.Now().UnixNano())

	flag.Parse()

	var facts []string

	buf, err := os.ReadFile(*factsFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(buf, &facts); err != nil {
		log.Fatal(err)
	}

	tpl, err := template.New("index").Parse(indexTpl)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		tpl.Execute(rw, facts)
	})

	http.HandleFunc("/fact", func(rw http.ResponseWriter, r *http.Request) {
		writeJSON(rw, facts[rand.Intn(len(facts))])
	})

	http.HandleFunc("/facts", func(rw http.ResponseWriter, r *http.Request) {
		writeJSON(rw, facts)
	})

	log.Printf("goatops.farm listening on %s", *listenAddr)
	http.ListenAndServe(*listenAddr, nil)
}

func writeJSON(rw http.ResponseWriter, v interface{}) {
	rw.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(rw)
	enc.SetIndent("", "  ")

	if err := enc.Encode(v); err != nil {
		log.Printf("error: %v", err)
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

var indexTpl = `<!DOCTYPE html>
<html lang="en">
<head>
  <title>goatops.farm</title>
  <style>body { font-family: sans-serif; } li { margin-bottom: 10px; }</style>
</head>
<body>
  <h1>Important facts about goats</h1>
  <ol>
{{- range $fact := . }}
    <li>{{ $fact }}</li>
{{- end }}
  </ol>
  <h2>API</h2>
  <ul>
    <li><a href="/fact">/fact</a> - random fact</li>
    <li><a href="/facts">/facts</a> - all facts</li>
  </ul>
  <p>
    Source code on <a href="https://github.com/martinohmann/goatops.farm">GitHub</a>.
    Goat facts generated from <a href="https://github.com/binford2k/goatops">https://github.com/binford2k/goatops</a>.
  </p>
</body>
</html>`
