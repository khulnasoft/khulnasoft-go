// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package request_tracers

import (
	"github.com/khulnasoft/khulnasoft-go/option"
)

// RequestTracerService contains methods and other services that help with
// interacting with the khulnasoft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestTracerService] method instead.
type RequestTracerService struct {
	Options []option.RequestOption
	Traces  *TraceService
}

// NewRequestTracerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequestTracerService(opts ...option.RequestOption) (r *RequestTracerService) {
	r = &RequestTracerService{}
	r.Options = opts
	r.Traces = NewTraceService(opts...)
	return
}