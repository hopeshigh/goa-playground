// Code generated by goa v3.12.4, DO NOT EDIT.
//
// todo HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/hopeshigh/goa-playground/design

package server

import (
	"context"
	"io"
	"net/http"

	todoviews "github.com/hopeshigh/goa-playground/gen/todo/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateResponse returns an encoder for responses returned by the todo
// create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the todo create
// endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreatePayload(&body)

		return payload, nil
	}
}

// EncodeCompleteResponse returns an encoder for responses returned by the todo
// complete endpoint.
func EncodeCompleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodeCompleteRequest returns a decoder for requests sent to the todo
// complete endpoint.
func DecodeCompleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewCompletePayload(id)

		return payload, nil
	}
}

// EncodeViewResponse returns an encoder for responses returned by the todo
// view endpoint.
func EncodeViewResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*todoviews.Todoresult)
		enc := encoder(ctx, w)
		body := NewViewResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeViewRequest returns a decoder for requests sent to the todo view
// endpoint.
func DecodeViewRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewViewPayload(id)

		return payload, nil
	}
}

// EncodeListResponse returns an encoder for responses returned by the todo
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(todoviews.TodoresultCollection)
		enc := encoder(ctx, w)
		body := NewTodoresultResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// marshalTodoviewsTodoresultViewToTodoresultResponse builds a value of type
// *TodoresultResponse from a value of type *todoviews.TodoresultView.
func marshalTodoviewsTodoresultViewToTodoresultResponse(v *todoviews.TodoresultView) *TodoresultResponse {
	res := &TodoresultResponse{
		ID:          *v.ID,
		Title:       *v.Title,
		Description: *v.Description,
		Status:      *v.Status,
	}

	return res
}