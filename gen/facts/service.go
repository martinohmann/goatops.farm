// Code generated by goa v3.4.2, DO NOT EDIT.
//
// facts service
//
// Command:
// $ goa gen github.com/martinohmann/goatops.farm/design

package facts

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The facts service provides you with important facts about goats and other
// creatures.
type Service interface {
	// List implements list.
	List(context.Context) (res []string, err error)
	// ListRandom implements list-random.
	ListRandom(context.Context, *ListRandomPayload) (res []string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "facts"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"list", "list-random"}

// ListRandomPayload is the payload type of the facts service list-random
// method.
type ListRandomPayload struct {
	// Number of random facts
	N *int
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "bad_request",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
