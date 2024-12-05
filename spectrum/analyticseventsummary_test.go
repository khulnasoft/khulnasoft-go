// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/spectrum"
)

func TestAnalyticsEventSummaryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Spectrum.Analytics.Events.Summaries.Get(context.TODO(), spectrum.AnalyticsEventSummaryGetParams{
		ZoneID:     khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Dimensions: khulnasoft.F([]spectrum.Dimension{spectrum.DimensionEvent, spectrum.DimensionAppID}),
		Filters:    khulnasoft.F("event==disconnect%20AND%20coloName!=SFO"),
		Metrics:    khulnasoft.F([]spectrum.AnalyticsEventSummaryGetParamsMetric{spectrum.AnalyticsEventSummaryGetParamsMetricCount, spectrum.AnalyticsEventSummaryGetParamsMetricBytesIngress}),
		Since:      khulnasoft.F(time.Now()),
		Sort:       khulnasoft.F([]string{"+count", "-bytesIngress"}),
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
