// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/intel"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestDNSListWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.DNS.List(context.TODO(), intel.DNSListParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		IPV4:      khulnasoft.F("ipv4"),
		Page:      khulnasoft.F(1.000000),
		PerPage:   khulnasoft.F(20.000000),
		StartEndParams: khulnasoft.F(intel.DNSListParamsStartEndParams{
			End:   khulnasoft.F(time.Now()),
			Start: khulnasoft.F(time.Now()),
		}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
