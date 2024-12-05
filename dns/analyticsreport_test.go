// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/dns"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestAnalyticsReportGetWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Analytics.Reports.Get(context.TODO(), dns.AnalyticsReportGetParams{
		ZoneID:     khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Dimensions: khulnasoft.F("queryType"),
		Filters:    khulnasoft.F("responseCode==NOERROR,queryType==A"),
		Limit:      khulnasoft.F(int64(100)),
		Metrics:    khulnasoft.F("queryCount,uncachedCount"),
		Since:      khulnasoft.F(time.Now()),
		Sort:       khulnasoft.F("+responseCode,-queryName"),
		Until:      khulnasoft.F(time.Now()),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
