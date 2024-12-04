// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/storage"
)

func TestAnalyticsListWithOptionalParams(t *testing.T) {
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
	_, err := client.Storage.Analytics.List(context.TODO(), storage.AnalyticsListParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Query: khulnasoft.F(storage.AnalyticsListParamsQuery{
			Dimensions: khulnasoft.F([]storage.AnalyticsListParamsQueryDimension{storage.AnalyticsListParamsQueryDimensionAccountID, storage.AnalyticsListParamsQueryDimensionResponseCode, storage.AnalyticsListParamsQueryDimensionRequestType}),
			Filters:    khulnasoft.F("requestType==read AND responseCode!=200"),
			Limit:      khulnasoft.F(int64(0)),
			Metrics:    khulnasoft.F([]storage.AnalyticsListParamsQueryMetric{storage.AnalyticsListParamsQueryMetricRequests, storage.AnalyticsListParamsQueryMetricWriteKiB, storage.AnalyticsListParamsQueryMetricReadKiB}),
			Since:      khulnasoft.F(time.Now()),
			Sort:       khulnasoft.F([]string{"+requests", "-responseCode"}),
			Until:      khulnasoft.F(time.Now()),
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

func TestAnalyticsStoredWithOptionalParams(t *testing.T) {
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
	_, err := client.Storage.Analytics.Stored(context.TODO(), storage.AnalyticsStoredParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Query: khulnasoft.F(storage.AnalyticsStoredParamsQuery{
			Dimensions: khulnasoft.F([]storage.AnalyticsStoredParamsQueryDimension{storage.AnalyticsStoredParamsQueryDimensionNamespaceID}),
			Filters:    khulnasoft.F("namespaceId==a4e8cbb7-1b58-4990-925e-e026d40c4c64"),
			Limit:      khulnasoft.F(int64(0)),
			Metrics:    khulnasoft.F([]storage.AnalyticsStoredParamsQueryMetric{storage.AnalyticsStoredParamsQueryMetricStoredBytes, storage.AnalyticsStoredParamsQueryMetricStoredKeys}),
			Since:      khulnasoft.F(time.Now()),
			Sort:       khulnasoft.F([]string{"+storedBytes", "-namespaceId"}),
			Until:      khulnasoft.F(time.Now()),
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
