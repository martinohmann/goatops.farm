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
	"text/template"
	"time"
)

//go:generate sh -c "go run ./tools/gen_facts.go > facts.go"

var listenAddr = flag.String("listen-address", ":8080", "Listen address")

func main() {
	rand.Seed(time.Now().UnixNano())

	flag.Parse()

	factTpl, err := template.New("fact").Parse(factTemplate)
	if err != nil {
		log.Fatal(err)
	}

	factsTpl, err := template.New("facts").Parse(factsTemplate)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		factTpl.Execute(rw, facts[rand.Intn(len(facts))])
	})

	http.HandleFunc("/fact", func(rw http.ResponseWriter, r *http.Request) {
		factTpl.Execute(rw, facts[rand.Intn(len(facts))])
	})

	http.HandleFunc("/facts", func(rw http.ResponseWriter, r *http.Request) {
		factsTpl.Execute(rw, facts)
	})

	http.HandleFunc("/fact.json", func(rw http.ResponseWriter, r *http.Request) {
		writeJSON(rw, facts[rand.Intn(len(facts))])
	})

	http.HandleFunc("/facts.json", func(rw http.ResponseWriter, r *http.Request) {
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

var header = `<!DOCTYPE html>
<html lang="en">
<head>
  <title>goatops.farm</title>
  <style>body { font-family: sans-serif; } li { margin-bottom: 10px; }</style>
</head>
<body>
`

var footer = `<h2>API</h2>
  <ul>
    <li><a href="/fact">/fact</a> - random fact</li>
    <li><a href="/facts">/facts</a> - all facts</li>
    <li><a href="/fact.json">/fact.json</a> - random fact (JSON)</li>
    <li><a href="/facts.json">/facts.json</a> - all facts (JSON)</li>
  </ul>
  <p>
    Source code on <a href="https://github.com/martinohmann/goatops.farm">GitHub</a>.
  </p>
  <p>
    Goat facts generated from <a href="https://github.com/binford2k/goatops">https://github.com/binford2k/goatops</a>.
  </p>
</body>
</html>`

var factsTemplate = header + `<h1>Important facts about goats</h1>
  <ol>
{{- range $fact := . }}
    <li>{{ $fact }}</li>
{{- end }}
  </ol>
` + footer

var factTemplate = header + `<h1>Something you should know about goats</h1>
  <p>{{ . }}</p>
` + footer
