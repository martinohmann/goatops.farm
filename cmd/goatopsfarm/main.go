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
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	goatopsfarm "github.com/martinohmann/goatops.farm"
	"github.com/martinohmann/goatops.farm/gen/facts"
)

func main() {
	listenAddr := flag.String("listen-addr", "0.0.0.0:8080", "Address to listen on")
	debug := flag.Bool("debug", false, "Log request and response bodies")
	redirectHTTPS := flag.Bool("redirect-https", true, "Redirect HTTP to HTTPS")

	flag.Parse()

	logger := log.New(os.Stderr, "[goatopsfarm] ", log.Ltime)

	factsSvc := goatopsfarm.NewFactsService(logger)
	factsEndpoints := facts.NewEndpoints(factsSvc)

	errc := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	startServer(ctx, *listenAddr, factsEndpoints, &wg, errc, logger, *redirectHTTPS, *debug)

	logger.Printf("exiting (%v)", <-errc)

	cancel()

	wg.Wait()
	logger.Println("exited")
}
