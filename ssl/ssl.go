// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

import (
	"github.com/khulnasoft/khulnasoft-go/option"
)

// SSLService contains methods and other services that help with interacting with
// the khulnasoft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSSLService] method instead.
type SSLService struct {
	Options          []option.RequestOption
	Analyze          *AnalyzeService
	CertificatePacks *CertificatePackService
	Recommendations  *RecommendationService
	Universal        *UniversalService
	Verification     *VerificationService
}

// NewSSLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSSLService(opts ...option.RequestOption) (r *SSLService) {
	r = &SSLService{}
	r.Options = opts
	r.Analyze = NewAnalyzeService(opts...)
	r.CertificatePacks = NewCertificatePackService(opts...)
	r.Recommendations = NewRecommendationService(opts...)
	r.Universal = NewUniversalService(opts...)
	r.Verification = NewVerificationService(opts...)
	return
}