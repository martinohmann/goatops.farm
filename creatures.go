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

//go:generate sh -c "go run ./cmd/gen-goat-facts > ./data/goat.go"

package goatopsfarm

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/martinohmann/goatops.farm/gen/creatures"
)

type creatureSvc struct {
	logger    *log.Logger
	rand      *rand.Rand
	creatures map[string]*creatures.Creature
}

func NewCreatureService(logger *log.Logger, data []*creatures.Creature) creatures.Service {
	creaturesMap := make(map[string]*creatures.Creature)
	for _, creature := range data {
		creaturesMap[creature.Name] = creature
	}

	return &creatureSvc{
		logger:    logger,
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		creatures: creaturesMap,
	}
}

// List implements list.
func (s *creatureSvc) List(ctx context.Context) (*creatures.ListResult, error) {
	s.logger.Print("creatures.List")

	var res creatures.ListResult
	for _, creature := range s.creatures {
		res.Creatures = append(res.Creatures, creature)
	}

	sort.Slice(res.Creatures, func(i, j int) bool {
		return res.Creatures[i].Name < res.Creatures[j].Name
	})

	return &res, nil
}

// Get implements get.
func (s *creatureSvc) Get(ctx context.Context, payload *creatures.GetPayload) (*creatures.GetResult, error) {
	s.logger.Print("creatures.Get")

	if payload.Name == "" {
		return nil, creatures.MakeBadRequest(errors.New("creatures are not nameless"))
	}

	creature, ok := s.creatures[payload.Name]
	if !ok {
		return nil, creatures.MakeNotFound(errors.New("creature was not seen here"))
	}

	return &creatures.GetResult{Creature: creature}, nil
}

// RandomFacts implements random-facts.
func (s *creatureSvc) RandomFacts(ctx context.Context, payload *creatures.RandomFactsPayload) (*creatures.RandomFactsResult, error) {
	s.logger.Print("creatures.RandomFacts")

	res, err := s.Get(ctx, &creatures.GetPayload{Name: payload.Name})
	if err != nil {
		return nil, err
	}

	n := 3
	if payload.N != nil {
		n = *payload.N
	}

	facts, err := randomSample(s.rand, res.Creature.Facts, n)
	if err != nil {
		return nil, creatures.MakeBadRequest(err)
	}

	return &creatures.RandomFactsResult{Facts: facts}, nil
}

func randomSample(r *rand.Rand, data []string, n int) ([]string, error) {
	if n < 0 {
		return nil, errors.New("n must be >= 0")
	}

	if n > len(data) {
		n = len(data)
	}

	switch {
	case n == 0:
		return []string{}, nil
	case n == 1:
		return []string{data[r.Intn(len(data))]}, nil
	default:
		sample := make([]string, len(data))

		copy(sample, data)

		r.Shuffle(len(data), func(i, j int) {
			sample[i], sample[j] = sample[j], sample[i]
		})

		return sample[:n], nil
	}
}
