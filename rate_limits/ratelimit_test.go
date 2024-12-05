// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limits_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/rate_limits"
)

func TestRateLimitNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.RateLimits.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		rate_limits.RateLimitNewParams{
			Action: khulnasoft.F(rate_limits.RateLimitNewParamsAction{
				Mode: khulnasoft.F(rate_limits.RateLimitNewParamsActionModeSimulate),
				Response: khulnasoft.F(rate_limits.RateLimitNewParamsActionResponse{
					Body:        khulnasoft.F("<error>This request has been rate-limited.</error>"),
					ContentType: khulnasoft.F("text/xml"),
				}),
				Timeout: khulnasoft.F(86400.000000),
			}),
			Match: khulnasoft.F(rate_limits.RateLimitNewParamsMatch{
				Headers: khulnasoft.F([]rate_limits.RateLimitNewParamsMatchHeader{{
					Name:  khulnasoft.F("Cf-Cache-Status"),
					Op:    khulnasoft.F(rate_limits.RateLimitNewParamsMatchHeadersOpEq),
					Value: khulnasoft.F("HIT"),
				}, {
					Name:  khulnasoft.F("Cf-Cache-Status"),
					Op:    khulnasoft.F(rate_limits.RateLimitNewParamsMatchHeadersOpEq),
					Value: khulnasoft.F("HIT"),
				}, {
					Name:  khulnasoft.F("Cf-Cache-Status"),
					Op:    khulnasoft.F(rate_limits.RateLimitNewParamsMatchHeadersOpEq),
					Value: khulnasoft.F("HIT"),
				}}),
				Request: khulnasoft.F(rate_limits.RateLimitNewParamsMatchRequest{
					Methods: khulnasoft.F([]rate_limits.Methods{rate_limits.MethodsGet, rate_limits.MethodsPost}),
					Schemes: khulnasoft.F([]string{"HTTP", "HTTPS"}),
					URL:     khulnasoft.F("*.example.org/path*"),
				}),
				Response: khulnasoft.F(rate_limits.RateLimitNewParamsMatchResponse{
					OriginTraffic: khulnasoft.F(true),
				}),
			}),
			Period:    khulnasoft.F(900.000000),
			Threshold: khulnasoft.F(60.000000),
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

func TestRateLimitListWithOptionalParams(t *testing.T) {
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
	_, err := client.RateLimits.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		rate_limits.RateLimitListParams{
			Page:    khulnasoft.F(1.000000),
			PerPage: khulnasoft.F(1.000000),
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

func TestRateLimitDelete(t *testing.T) {
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
	_, err := client.RateLimits.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"372e67954025e0ba6aaa6d586b9e0b59",
	)
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRateLimitEditWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.RateLimits.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"372e67954025e0ba6aaa6d586b9e0b59",
		rate_limits.RateLimitEditParams{
			Action: khulnasoft.F(rate_limits.RateLimitEditParamsAction{
				Mode: khulnasoft.F(rate_limits.RateLimitEditParamsActionModeSimulate),
				Response: khulnasoft.F(rate_limits.RateLimitEditParamsActionResponse{
					Body:        khulnasoft.F("<error>This request has been rate-limited.</error>"),
					ContentType: khulnasoft.F("text/xml"),
				}),
				Timeout: khulnasoft.F(86400.000000),
			}),
			Match: khulnasoft.F(rate_limits.RateLimitEditParamsMatch{
				Headers: khulnasoft.F([]rate_limits.RateLimitEditParamsMatchHeader{{
					Name:  khulnasoft.F("Cf-Cache-Status"),
					Op:    khulnasoft.F(rate_limits.RateLimitEditParamsMatchHeadersOpEq),
					Value: khulnasoft.F("HIT"),
				}, {
					Name:  khulnasoft.F("Cf-Cache-Status"),
					Op:    khulnasoft.F(rate_limits.RateLimitEditParamsMatchHeadersOpEq),
					Value: khulnasoft.F("HIT"),
				}, {
					Name:  khulnasoft.F("Cf-Cache-Status"),
					Op:    khulnasoft.F(rate_limits.RateLimitEditParamsMatchHeadersOpEq),
					Value: khulnasoft.F("HIT"),
				}}),
				Request: khulnasoft.F(rate_limits.RateLimitEditParamsMatchRequest{
					Methods: khulnasoft.F([]rate_limits.Methods{rate_limits.MethodsGet, rate_limits.MethodsPost}),
					Schemes: khulnasoft.F([]string{"HTTP", "HTTPS"}),
					URL:     khulnasoft.F("*.example.org/path*"),
				}),
				Response: khulnasoft.F(rate_limits.RateLimitEditParamsMatchResponse{
					OriginTraffic: khulnasoft.F(true),
				}),
			}),
			Period:    khulnasoft.F(900.000000),
			Threshold: khulnasoft.F(60.000000),
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

func TestRateLimitGet(t *testing.T) {
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
	_, err := client.RateLimits.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"372e67954025e0ba6aaa6d586b9e0b59",
	)
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
