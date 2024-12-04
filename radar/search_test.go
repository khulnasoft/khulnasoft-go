// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/radar"
)

func TestSearchGlobalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Search.Global(context.TODO(), radar.SearchGlobalParams{
		Query:         khulnasoft.F("United"),
		Exclude:       khulnasoft.F([]radar.SearchGlobalParamsExclude{radar.SearchGlobalParamsExcludeSpecialEvents, radar.SearchGlobalParamsExcludeNotebooks, radar.SearchGlobalParamsExcludeLocations}),
		Format:        khulnasoft.F(radar.SearchGlobalParamsFormatJson),
		Include:       khulnasoft.F([]radar.SearchGlobalParamsInclude{radar.SearchGlobalParamsIncludeSpecialEvents, radar.SearchGlobalParamsIncludeNotebooks, radar.SearchGlobalParamsIncludeLocations}),
		Limit:         khulnasoft.F(int64(5)),
		LimitPerGroup: khulnasoft.F(0.000000),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
