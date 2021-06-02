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

//go:generate sh -c "go run ./cmd/gen-goat-facts > goat_facts.go"

package goatopsfarm

import (
	"context"
	"errors"
	"log"
	"math/rand"

	"github.com/martinohmann/goatops.farm/gen/facts"
)

type factsSvc struct {
	logger *log.Logger
	facts  []string
}

func NewFactsService(logger *log.Logger) facts.Service {
	svc := &factsSvc{
		logger: logger,
		facts:  goatFacts, // @TODO(martinohmann): support other creatures as well
	}
	return svc
}

func (s *factsSvc) List(ctx context.Context) ([]string, error) {
	s.logger.Print("facts.List")
	return s.facts, nil
}

func (s *factsSvc) ListRandom(ctx context.Context, payload *facts.ListRandomPayload) ([]string, error) {
	s.logger.Print("facts.ListRandom")

	var n int = 5
	if payload.N != nil {
		n = *payload.N
	}

	res, err := s.randomFacts(n)
	if err != nil {
		return nil, facts.MakeBadRequest(err)
	}

	return res, nil
}

func (s *factsSvc) randomFacts(n int) ([]string, error) {
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
