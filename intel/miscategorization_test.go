// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/intel"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestMiscategorizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.Miscategorizations.New(context.TODO(), intel.MiscategorizationNewParams{
		AccountID:       khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ContentAdds:     khulnasoft.F([]int64{int64(82)}),
		ContentRemoves:  khulnasoft.F([]int64{int64(155)}),
		IndicatorType:   khulnasoft.F(intel.MiscategorizationNewParamsIndicatorTypeDomain),
		IP:              khulnasoft.F("ip"),
		SecurityAdds:    khulnasoft.F([]int64{int64(117), int64(131)}),
		SecurityRemoves: khulnasoft.F([]int64{int64(83)}),
		URL:             khulnasoft.F("url"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
