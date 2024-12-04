// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_certificates_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/custom_certificates"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestPrioritizeUpdate(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := khulnasoft.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.CustomCertificates.Prioritize.Update(context.TODO(), custom_certificates.PrioritizeUpdateParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Certificates: khulnasoft.F([]custom_certificates.PrioritizeUpdateParamsCertificate{{
			ID:       khulnasoft.F("5a7805061c76ada191ed06f989cc3dac"),
			Priority: khulnasoft.F(2.000000),
		}, {
			ID:       khulnasoft.F("9a7806061c88ada191ed06f989cc3dac"),
			Priority: khulnasoft.F(1.000000),
		}}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
