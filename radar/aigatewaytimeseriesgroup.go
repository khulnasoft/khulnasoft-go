// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/khulnasoft/khulnasoft-go/option"
)

// AIGatewayTimeseriesGroupService contains methods and other services that help
// with interacting with the khulnasoft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIGatewayTimeseriesGroupService] method instead.
type AIGatewayTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewAIGatewayTimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIGatewayTimeseriesGroupService(opts ...option.RequestOption) (r *AIGatewayTimeseriesGroupService) {
	r = &AIGatewayTimeseriesGroupService{}
	r.Options = opts
	return
}