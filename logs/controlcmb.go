// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

import (
	"github.com/khulnasoft/khulnasoft-go/option"
)

// ControlCmbService contains methods and other services that help with interacting
// with the khulnasoft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewControlCmbService] method instead.
type ControlCmbService struct {
	Options []option.RequestOption
	Config  *ControlCmbConfigService
}

// NewControlCmbService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewControlCmbService(opts ...option.RequestOption) (r *ControlCmbService) {
	r = &ControlCmbService{}
	r.Options = opts
	r.Config = NewControlCmbConfigService(opts...)
	return
}