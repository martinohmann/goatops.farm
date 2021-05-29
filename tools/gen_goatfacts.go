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

//go:build tools
// +build tools

package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/microcosm-cc/bluemonday"
)

var factRe = regexp.MustCompile(`^\s+<li>(.+)</?li>`)

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/binford2k/goatops/master/index.html")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	p := bluemonday.NewPolicy()

	r := bufio.NewScanner(resp.Body)

	var facts []string

	for r.Scan() {
		line := r.Text()

		if factRe.MatchString(line) {
			fact := factRe.ReplaceAllString(line, `$1`)
			facts = append(facts, p.Sanitize(fact))
		}
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	if err := enc.Encode(facts); err != nil {
		log.Fatal(err)
	}
}
