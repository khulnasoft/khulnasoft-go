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

func TestMonitorNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Monitors.New(context.TODO(), load_balancers.MonitorNewParams{
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
		Type:      khulnasoft.F(load_balancers.MonitorNewParamsTypeHTTP),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Monitors.Update(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		load_balancers.MonitorUpdateParams{
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
			Type:      khulnasoft.F(load_balancers.MonitorUpdateParamsTypeHTTP),
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

func TestMonitorList(t *testing.T) {
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
	_, err := client.LoadBalancers.Monitors.List(context.TODO(), load_balancers.MonitorListParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorDelete(t *testing.T) {
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
	_, err := client.LoadBalancers.Monitors.Delete(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		load_balancers.MonitorDeleteParams{
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

func TestMonitorEditWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Monitors.Edit(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		load_balancers.MonitorEditParams{
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
			Type:      khulnasoft.F(load_balancers.MonitorEditParamsTypeHTTP),
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

func TestMonitorGet(t *testing.T) {
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
	_, err := client.LoadBalancers.Monitors.Get(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		load_balancers.MonitorGetParams{
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
