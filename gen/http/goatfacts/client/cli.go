// Code generated by goa v3.4.2, DO NOT EDIT.
//
// goatfacts HTTP client CLI support package
//
// Command:
// $ goa gen github.com/martinohmann/goatops.farm/design

package client

import (
	"fmt"
	"strconv"

	goatfacts "github.com/martinohmann/goatops.farm/gen/goatfacts"
)

// BuildRandomFactsPayload builds the payload for the goatfacts RandomFacts
// endpoint from CLI flags.
func BuildRandomFactsPayload(goatfactsRandomFactsN string) (*goatfacts.RandomFactsPayload, error) {
	var err error
	var n *int
	{
		if goatfactsRandomFactsN != "" {
			var v int64
			v, err = strconv.ParseInt(goatfactsRandomFactsN, 10, 64)
			val := int(v)
			n = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for n, must be INT")
			}
		}
	}
	v := &goatfacts.RandomFactsPayload{}
	v.N = n

	return v, nil
}
