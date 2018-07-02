// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// cars endpoints
//
// Command:
// $ goa gen goa.design/goa/examples/streaming/design -o
// $(GOPATH)/src/goa.design/goa/examples/streaming

package carssvc

import (
	"context"

	goa "goa.design/goa"
	"goa.design/goa/security"
)

// Endpoints wraps the "cars" service endpoints.
type Endpoints struct {
	Login goa.Endpoint
	List  goa.Endpoint
}

// ListEndpointInput is the input type of "list" endpoint that holds the method
// payload and the server stream.
type ListEndpointInput struct {
	// Payload is the method payload.
	Payload *ListPayload
	// Stream is the server stream used by the "list" method to send data.
	Stream ListServerStream
}

// NewEndpoints wraps the methods of the "cars" service with endpoints.
func NewEndpoints(s Service, authBasicFn security.AuthBasicFunc, authJWTFn security.AuthJWTFunc) *Endpoints {
	return &Endpoints{
		Login: NewLoginEndpoint(s, authBasicFn),
		List:  NewListEndpoint(s, authJWTFn),
	}
}

// Use applies the given middleware to all the "cars" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Login = m(e.Login)
	e.List = m(e.List)
}

// NewLoginEndpoint returns an endpoint function that calls the method "login"
// of service "cars".
func NewLoginEndpoint(s Service, authBasicFn security.AuthBasicFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LoginPayload)
		var err error
		sc := security.BasicScheme{
			Name: "basic",
		}
		ctx, err = authBasicFn(ctx, p.User, p.Password, &sc)
		if err != nil {
			return nil, err
		}
		return s.Login(ctx, p)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "cars".
func NewListEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ep := req.(*ListEndpointInput)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"stream:read", "stream:write"},
			RequiredScopes: []string{"stream:read"},
		}
		ctx, err = authJWTFn(ctx, ep.Payload.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.List(ctx, ep.Payload, ep.Stream)
	}
}
