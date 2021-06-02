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
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	cli "github.com/martinohmann/goatops.farm/gen/http/cli/goatopsfarm"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	serverURL := flag.String("url", "https://goatops.farm", "URL to service host")
	verbose := flag.Bool("verbose", false, "Print request and response details")
	v := flag.Bool("v", false, "Print request and response details")
	timeout := flag.Int("timeout", 30, "Maximum number of seconds to wait for response")

	flag.Usage = usage
	flag.Parse()

	debug := *verbose || *v

	u, err := url.Parse(*serverURL)
	if err != nil {
		return err
	}

	endpoint, payload, err := parseEndpoint(u.Scheme, u.Host, *timeout, debug)
	if err != nil {
		if err == flag.ErrHelp {
			return nil
		}

		return fmt.Errorf("%w: run '%s --help' for detailed usage", err, os.Args[0])
	}

	data, err := endpoint(context.Background(), payload)
	if err != nil {
		return err
	}

	if data != nil {
		m, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(m))
	}

	return nil
}

func parseEndpoint(scheme, host string, timeout int, debug bool) (goa.Endpoint, interface{}, error) {
	var doer goahttp.Doer
	{
		doer = &http.Client{Timeout: time.Duration(timeout) * time.Second}
		if debug {
			doer = goahttp.NewDebugDoer(doer)
		}
	}

	return cli.ParseEndpoint(
		scheme,
		host,
		doer,
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		debug,
	)
}

func usage() {
	fmt.Fprintf(os.Stderr, `%s is a command line client for the goatopsfarm API.

Usage:
    %s [-url URL][-timeout SECONDS][-verbose|-v] SERVICE ENDPOINT [flags]

    -url URL     specify service URL overriding host URL (https://goatops.farm)
    -timeout     maximum number of seconds to wait for response (30)
    -verbose|-v  print request and response details (false)

Commands:
%s
Additional help:
    %s SERVICE [ENDPOINT] --help

Example:
%s
`, os.Args[0], os.Args[0], indent(cli.UsageCommands()), os.Args[0], indent(cli.UsageExamples()))
}

func indent(s string) string {
	if s == "" {
		return ""
	}

	return "    " + strings.Replace(s, "\n", "\n    ", -1)
}
