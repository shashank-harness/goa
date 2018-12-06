// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// chatter HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/goa/examples/chatter/design -o
// $(GOPATH)/src/goa.design/goa/examples/chatter

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	chattersvc "goa.design/goa/examples/chatter/gen/chatter"
	chattersvcviews "goa.design/goa/examples/chatter/gen/chatter/views"
	goahttp "goa.design/goa/http"
)

// BuildLoginRequest instantiates a HTTP request object with method and path
// set to call the "chatter" service "login" endpoint
func (c *Client) BuildLoginRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LoginChatterPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chatter", "login", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeLoginRequest returns an encoder for requests sent to the chatter login
// server.
func EncodeLoginRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chattersvc.LoginPayload)
		if !ok {
			return goahttp.ErrInvalidType("chatter", "login", "*chattersvc.LoginPayload", v)
		}
		req.SetBasicAuth(p.User, p.Password)
		return nil
	}
}

// DecodeLoginResponse returns a decoder for responses returned by the chatter
// login endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeLoginResponse may return the following errors:
//	- "unauthorized" (type chattersvc.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeLoginResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "login", err)
			}
			return body, nil
		case http.StatusUnauthorized:
			var (
				body LoginUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "login", err)
			}
			return nil, NewLoginUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chatter", "login", resp.StatusCode, string(body))
		}
	}
}

// BuildEchoerRequest instantiates a HTTP request object with method and path
// set to call the "chatter" service "echoer" endpoint
func (c *Client) BuildEchoerRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	scheme := c.scheme
	switch c.scheme {
	case "http":
		scheme = "ws"
	case "https":
		scheme = "wss"
	}
	u := &url.URL{Scheme: scheme, Host: c.host, Path: EchoerChatterPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chatter", "echoer", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeEchoerRequest returns an encoder for requests sent to the chatter
// echoer server.
func EncodeEchoerRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chattersvc.EchoerPayload)
		if !ok {
			return goahttp.ErrInvalidType("chatter", "echoer", "*chattersvc.EchoerPayload", v)
		}
		if !strings.Contains(p.Token, " ") {
			req.Header.Set("Authorization", "Bearer "+p.Token)
		} else {
			req.Header.Set("Authorization", p.Token)
		}
		return nil
	}
}

// DecodeEchoerResponse returns a decoder for responses returned by the chatter
// echoer endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeEchoerResponse may return the following errors:
//	- "invalid-scopes" (type chattersvc.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type chattersvc.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeEchoerResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "echoer", err)
			}
			return body, nil
		case http.StatusForbidden:
			var (
				body EchoerInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "echoer", err)
			}
			return nil, NewEchoerInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body EchoerUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "echoer", err)
			}
			return nil, NewEchoerUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chatter", "echoer", resp.StatusCode, string(body))
		}
	}
}

// BuildListenerRequest instantiates a HTTP request object with method and path
// set to call the "chatter" service "listener" endpoint
func (c *Client) BuildListenerRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	scheme := c.scheme
	switch c.scheme {
	case "http":
		scheme = "ws"
	case "https":
		scheme = "wss"
	}
	u := &url.URL{Scheme: scheme, Host: c.host, Path: ListenerChatterPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chatter", "listener", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListenerRequest returns an encoder for requests sent to the chatter
// listener server.
func EncodeListenerRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chattersvc.ListenerPayload)
		if !ok {
			return goahttp.ErrInvalidType("chatter", "listener", "*chattersvc.ListenerPayload", v)
		}
		if !strings.Contains(p.Token, " ") {
			req.Header.Set("Authorization", "Bearer "+p.Token)
		} else {
			req.Header.Set("Authorization", p.Token)
		}
		return nil
	}
}

// DecodeListenerResponse returns a decoder for responses returned by the
// chatter listener endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeListenerResponse may return the following errors:
//	- "invalid-scopes" (type chattersvc.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type chattersvc.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeListenerResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		case http.StatusForbidden:
			var (
				body ListenerInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "listener", err)
			}
			return nil, NewListenerInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body ListenerUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "listener", err)
			}
			return nil, NewListenerUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chatter", "listener", resp.StatusCode, string(body))
		}
	}
}

// BuildSummaryRequest instantiates a HTTP request object with method and path
// set to call the "chatter" service "summary" endpoint
func (c *Client) BuildSummaryRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	scheme := c.scheme
	switch c.scheme {
	case "http":
		scheme = "ws"
	case "https":
		scheme = "wss"
	}
	u := &url.URL{Scheme: scheme, Host: c.host, Path: SummaryChatterPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chatter", "summary", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSummaryRequest returns an encoder for requests sent to the chatter
// summary server.
func EncodeSummaryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chattersvc.SummaryPayload)
		if !ok {
			return goahttp.ErrInvalidType("chatter", "summary", "*chattersvc.SummaryPayload", v)
		}
		if !strings.Contains(p.Token, " ") {
			req.Header.Set("Authorization", "Bearer "+p.Token)
		} else {
			req.Header.Set("Authorization", p.Token)
		}
		return nil
	}
}

// DecodeSummaryResponse returns a decoder for responses returned by the
// chatter summary endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeSummaryResponse may return the following errors:
//	- "invalid-scopes" (type chattersvc.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type chattersvc.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeSummaryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SummaryResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "summary", err)
			}
			p := NewSummaryChatSummaryCollectionOK(body)
			view := "default"
			vres := chattersvcviews.ChatSummaryCollection{p, view}
			if err = vres.Validate(); err != nil {
				return nil, goahttp.ErrValidationError("chatter", "summary", err)
			}
			res := chattersvc.NewChatSummaryCollection(vres)
			return res, nil
		case http.StatusForbidden:
			var (
				body SummaryInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "summary", err)
			}
			return nil, NewSummaryInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body SummaryUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "summary", err)
			}
			return nil, NewSummaryUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chatter", "summary", resp.StatusCode, string(body))
		}
	}
}

// BuildHistoryRequest instantiates a HTTP request object with method and path
// set to call the "chatter" service "history" endpoint
func (c *Client) BuildHistoryRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	scheme := c.scheme
	switch c.scheme {
	case "http":
		scheme = "ws"
	case "https":
		scheme = "wss"
	}
	u := &url.URL{Scheme: scheme, Host: c.host, Path: HistoryChatterPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chatter", "history", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeHistoryRequest returns an encoder for requests sent to the chatter
// history server.
func EncodeHistoryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chattersvc.HistoryPayload)
		if !ok {
			return goahttp.ErrInvalidType("chatter", "history", "*chattersvc.HistoryPayload", v)
		}
		if !strings.Contains(p.Token, " ") {
			req.Header.Set("Authorization", "Bearer "+p.Token)
		} else {
			req.Header.Set("Authorization", p.Token)
		}
		values := req.URL.Query()
		if p.View != nil {
			values.Add("view", *p.View)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeHistoryResponse returns a decoder for responses returned by the
// chatter history endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeHistoryResponse may return the following errors:
//	- "invalid-scopes" (type chattersvc.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type chattersvc.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeHistoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body HistoryResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "history", err)
			}
			p := NewHistoryChatSummaryOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &chattersvcviews.ChatSummary{p, view}
			if err = vres.Validate(); err != nil {
				return nil, goahttp.ErrValidationError("chatter", "history", err)
			}
			res := chattersvc.NewChatSummary(vres)
			return res, nil
		case http.StatusForbidden:
			var (
				body HistoryInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "history", err)
			}
			return nil, NewHistoryInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body HistoryUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "history", err)
			}
			return nil, NewHistoryUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chatter", "history", resp.StatusCode, string(body))
		}
	}
}
