// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/load_balancers"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestPoolHealthNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Pools.Health.New(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		load_balancers.PoolHealthNewParams{
			AccountID:       khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ExpectedCodes:   khulnasoft.F("2xx"),
			AllowInsecure:   khulnasoft.F(true),
			ConsecutiveDown: khulnasoft.F(int64(0)),
			ConsecutiveUp:   khulnasoft.F(int64(0)),
			Description:     khulnasoft.F("Login page monitor"),
			ExpectedBody:    khulnasoft.F("alive"),
			FollowRedirects: khulnasoft.F(true),
			Header: khulnasoft.F(map[string][]string{
				"Host":     {"example.com"},
				"X-App-ID": {"abc123"},
			}),
			Interval:  khulnasoft.F(int64(0)),
			Method:    khulnasoft.F("GET"),
			Path:      khulnasoft.F("/health"),
			Port:      khulnasoft.F(int64(0)),
			ProbeZone: khulnasoft.F("example.com"),
			Retries:   khulnasoft.F(int64(0)),
			Timeout:   khulnasoft.F(int64(0)),
			Type:      khulnasoft.F(load_balancers.PoolHealthNewParamsTypeHTTP),
		},
	)
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPoolHealthGet(t *testing.T) {
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
	_, err := client.LoadBalancers.Pools.Health.Get(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		load_balancers.PoolHealthGetParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		},
	)
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
