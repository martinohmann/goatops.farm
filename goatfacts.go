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

//go:generate sh -c "go run ./cmd/gen-facts > facts.go"

package goatopsfarm

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"log"
	"math/rand"

	goatfacts "github.com/martinohmann/goatops.farm/gen/goatfacts"
	goa "goa.design/goa/v3/pkg"
)

// goatfacts service example implementation.
// The example methods log the requests and return zero values.
type goatFactsSvc struct {
	logger   *log.Logger
	indexTpl *template.Template
	facts    []string
}

// NewGoatFactsService returns the goatfacts service implementation.
func NewGoatFactsService(logger *log.Logger) (goatfacts.Service, error) {
	tpl, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		return nil, err
	}
	svc := &goatFactsSvc{
		logger:   logger,
		indexTpl: tpl,
		facts:    facts,
	}
	return svc, nil
}

func (s *goatFactsSvc) ListFacts(ctx context.Context) ([]string, error) {
	s.logger.Print("goatfacts.ListFacts")
	return s.facts, nil
}

func (s *goatFactsSvc) RandomFacts(ctx context.Context, payload *goatfacts.RandomFactsPayload) ([]string, error) {
	s.logger.Print("goatfacts.RandomFacts")

	var n int = 5
	if payload.N != nil {
		n = *payload.N
	}

	facts, err := s.randomFacts(n)
	if err != nil {
		return nil, goatfacts.MakeBadRequest(err)
	}

	return facts, nil
}

func (s *goatFactsSvc) Index(context.Context) ([]byte, error) {
	var buf bytes.Buffer

	facts, _ := s.randomFacts(5)

	data := struct {
		Facts []string
	}{
		Facts: facts,
	}

	if err := s.indexTpl.Execute(&buf, data); err != nil {
		return nil, goa.Fault(err.Error())
	}

	return buf.Bytes(), nil
}

func (s *goatFactsSvc) randomFacts(n int) ([]string, error) {
	if n <= 0 {
		return nil, errors.New("n must be > 0")
	}

	if n > len(s.facts) {
		n = len(s.facts)
	}

	selection := make([]string, 0, n)
	for i := 0; i < n; i++ {
		fact := s.facts[rand.Intn(len(s.facts))]
		selection = append(selection, fact)
	}

	return selection, nil
}

var indexTemplate = `<!DOCTYPE html>
<html>
<head>
  <title>goatops.farm</title>
</head>
<body>
  <h1>Something you should know about goats</h1>
{{ with .Facts }}
  <ol>
{{ range $fact := . }}
    <li>{{ $fact }}</li>
{{ end }}
  </ol>
{{ end }}

  <h2>JSON API</h2>
  <p>
    See <a href="/static/swagger">swagger-ui</a> for API endpoint docs.
  </p>
  <p>
    Source code on <a href="https://github.com/martinohmann/goatops.farm">GitHub</a>.
  </p>
  <p>
    Goat facts generated from <a href="https://github.com/binford2k/goatops">https://github.com/binford2k/goatops</a>.
  </p>
</body>
</html>`
