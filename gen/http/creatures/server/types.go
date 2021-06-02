// Code generated by goa v3.4.2, DO NOT EDIT.
//
// creatures HTTP server types
//
// Command:
// $ goa gen github.com/martinohmann/goatops.farm/design

package server

import (
	creatures "github.com/martinohmann/goatops.farm/gen/creatures"
	goa "goa.design/goa/v3/pkg"
)

// ListResponseBody is the type of the "creatures" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// List of creatures
	Creatures []*CreatureResponseBody `form:"creatures" json:"creatures" xml:"creatures"`
}

// GetResponseBody is the type of the "creatures" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// The creature
	Creature *CreatureResponseBody `form:"creature" json:"creature" xml:"creature"`
}

// RandomFactsResponseBody is the type of the "creatures" service
// "random-facts" endpoint HTTP response body.
type RandomFactsResponseBody struct {
	// Random facts about the creature
	Facts []string `form:"facts" json:"facts" xml:"facts"`
}

// GetNotFoundResponseBody is the type of the "creatures" service "get"
// endpoint HTTP response body for the "not_found" error.
type GetNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RandomFactsBadRequestResponseBody is the type of the "creatures" service
// "random-facts" endpoint HTTP response body for the "bad_request" error.
type RandomFactsBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RandomFactsNotFoundResponseBody is the type of the "creatures" service
// "random-facts" endpoint HTTP response body for the "not_found" error.
type RandomFactsNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreatureResponseBody is used to define fields on response body types.
type CreatureResponseBody struct {
	// Name of the creature
	Name string `form:"name" json:"name" xml:"name"`
	// Facts about the creature
	Facts []string `form:"facts" json:"facts" xml:"facts"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "creatures" service.
func NewListResponseBody(res *creatures.ListResult) *ListResponseBody {
	body := &ListResponseBody{}
	if res.Creatures != nil {
		body.Creatures = make([]*CreatureResponseBody, len(res.Creatures))
		for i, val := range res.Creatures {
			body.Creatures[i] = marshalCreaturesCreatureToCreatureResponseBody(val)
		}
	}
	return body
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "creatures" service.
func NewGetResponseBody(res *creatures.GetResult) *GetResponseBody {
	body := &GetResponseBody{}
	if res.Creature != nil {
		body.Creature = marshalCreaturesCreatureToCreatureResponseBody(res.Creature)
	}
	return body
}

// NewRandomFactsResponseBody builds the HTTP response body from the result of
// the "random-facts" endpoint of the "creatures" service.
func NewRandomFactsResponseBody(res *creatures.RandomFactsResult) *RandomFactsResponseBody {
	body := &RandomFactsResponseBody{}
	if res.Facts != nil {
		body.Facts = make([]string, len(res.Facts))
		for i, val := range res.Facts {
			body.Facts[i] = val
		}
	}
	return body
}

// NewGetNotFoundResponseBody builds the HTTP response body from the result of
// the "get" endpoint of the "creatures" service.
func NewGetNotFoundResponseBody(res *goa.ServiceError) *GetNotFoundResponseBody {
	body := &GetNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRandomFactsBadRequestResponseBody builds the HTTP response body from the
// result of the "random-facts" endpoint of the "creatures" service.
func NewRandomFactsBadRequestResponseBody(res *goa.ServiceError) *RandomFactsBadRequestResponseBody {
	body := &RandomFactsBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRandomFactsNotFoundResponseBody builds the HTTP response body from the
// result of the "random-facts" endpoint of the "creatures" service.
func NewRandomFactsNotFoundResponseBody(res *goa.ServiceError) *RandomFactsNotFoundResponseBody {
	body := &RandomFactsNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetPayload builds a creatures service get endpoint payload.
func NewGetPayload(name string) *creatures.GetPayload {
	v := &creatures.GetPayload{}
	v.Name = name

	return v
}

// NewRandomFactsPayload builds a creatures service random-facts endpoint
// payload.
func NewRandomFactsPayload(name string, n *int) *creatures.RandomFactsPayload {
	v := &creatures.RandomFactsPayload{}
	v.Name = name
	v.N = n

	return v
}
