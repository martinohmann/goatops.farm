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
	"time"

	"github.com/martinohmann/goatops.farm/gen/facts"
)

type factsSvc struct {
	logger *log.Logger
	facts  []string
	rand   *rand.Rand
}

func NewFactsService(logger *log.Logger) facts.Service {
	return &factsSvc{
		logger: logger,
		facts:  goatFacts, // @TODO(martinohmann): support other creatures as well
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *factsSvc) List(ctx context.Context) (*facts.ListResult, error) {
	s.logger.Print("facts.List")
	return &facts.ListResult{Facts: s.facts}, nil
}

func (s *factsSvc) ListRandom(ctx context.Context, payload *facts.ListRandomPayload) (*facts.ListRandomResult, error) {
	s.logger.Print("facts.ListRandom")

	var n int = 5
	if payload.N != nil {
		n = *payload.N
	}

	fs, err := s.randomFacts(n)
	if err != nil {
		return nil, facts.MakeBadRequest(err)
	}

	res := &facts.ListRandomResult{
		Facts: fs,
	}

	return res, nil
}

func (s *factsSvc) randomFacts(n int) ([]string, error) {
	if n < 0 {
		return nil, errors.New("n must be >= 0")
	}

	if n > len(s.facts) {
		n = len(s.facts)
	}

	switch {
	case n == 0:
		return []string{}, nil
	case n == 1:
		fact := s.facts[s.rand.Intn(len(s.facts))]
		return []string{fact}, nil
	default:
		res := make([]string, len(s.facts))

		copy(res, s.facts)

		rand.Shuffle(len(s.facts), func(i, j int) {
			res[i], res[j] = res[j], res[i]
		})

		return res[:n], nil
	}
}
