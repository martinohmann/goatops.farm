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
	"context"
	"errors"
	"log"
	"math/rand"
	"strconv"

	goatfacts "github.com/martinohmann/goatops.farm/gen/goatfacts"
)

var (
	goatFacts    []*goatfacts.Fact
	goatFactsMap map[string]*goatfacts.Fact
)

func init() {
	goatFactsMap = make(map[string]*goatfacts.Fact)

	for i, fact := range facts {
		goatFact := &goatfacts.Fact{
			ID:   strconv.Itoa(i + 1),
			Text: fact,
		}
		goatFacts = append(goatFacts, goatFact)
		goatFactsMap[goatFact.ID] = goatFact
	}
}

// goatfacts service example implementation.
// The example methods log the requests and return zero values.
type goatFactsSvc struct {
	logger *log.Logger
}

// NewGoatFactsService returns the goatfacts service implementation.
func NewGoatFactsService(logger *log.Logger) goatfacts.Service {
	return &goatFactsSvc{logger}
}

// GetFact implements get-fact.
func (s *goatFactsSvc) GetFact(ctx context.Context, p *goatfacts.GetFactPayload) (*goatfacts.Fact, error) {
	s.logger.Print("goatfacts.get-fact")

	fact, ok := goatFactsMap[p.ID]
	if !ok {
		return nil, goatfacts.MakeNotFound(errors.New("not found"))
	}

	return fact, nil
}

// ListFacts implements list-facts.
func (s *goatFactsSvc) ListFacts(ctx context.Context) ([]*goatfacts.Fact, error) {
	s.logger.Print("goatfacts.list-facts")
	return goatFacts, nil
}

// GetRandomFact implements get-random-fact.
func (s *goatFactsSvc) GetRandomFact(ctx context.Context) (*goatfacts.Fact, error) {
	s.logger.Print("goatfacts.get-random-fact")

	if len(goatFacts) == 0 {
		return nil, goatfacts.MakeNotFound(errors.New("not found"))
	}

	return goatFacts[rand.Intn(len(goatFacts))], nil
}
