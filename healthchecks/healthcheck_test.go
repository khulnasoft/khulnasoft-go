// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthchecks_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/healthchecks"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestHealthcheckNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Healthchecks.New(context.TODO(), healthchecks.HealthcheckNewParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		QueryHealthcheck: healthchecks.QueryHealthcheckParam{
			Address:              khulnasoft.F("www.example.com"),
			Name:                 khulnasoft.F("server-1"),
			CheckRegions:         khulnasoft.F([]healthchecks.CheckRegion{healthchecks.CheckRegionWnam, healthchecks.CheckRegionEnam}),
			ConsecutiveFails:     khulnasoft.F(int64(0)),
			ConsecutiveSuccesses: khulnasoft.F(int64(0)),
			Description:          khulnasoft.F("Health check for www.example.com"),
			HTTPConfig: khulnasoft.F(healthchecks.HTTPConfigurationParam{
				AllowInsecure:   khulnasoft.F(true),
				ExpectedBody:    khulnasoft.F("success"),
				ExpectedCodes:   khulnasoft.F([]string{"2xx", "302"}),
				FollowRedirects: khulnasoft.F(true),
				Header: khulnasoft.F(map[string][]string{
					"Host":     {"example.com"},
					"X-App-ID": {"abc123"},
				}),
				Method: khulnasoft.F(healthchecks.HTTPConfigurationMethodGet),
				Path:   khulnasoft.F("/health"),
				Port:   khulnasoft.F(int64(0)),
			}),
			Interval:  khulnasoft.F(int64(0)),
			Retries:   khulnasoft.F(int64(0)),
			Suspended: khulnasoft.F(true),
			TCPConfig: khulnasoft.F(healthchecks.TCPConfigurationParam{
				Method: khulnasoft.F(healthchecks.TCPConfigurationMethodConnectionEstablished),
				Port:   khulnasoft.F(int64(0)),
			}),
			Timeout: khulnasoft.F(int64(0)),
			Type:    khulnasoft.F("HTTPS"),
		},
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHealthcheckUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Healthchecks.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		healthchecks.HealthcheckUpdateParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			QueryHealthcheck: healthchecks.QueryHealthcheckParam{
				Address:              khulnasoft.F("www.example.com"),
				Name:                 khulnasoft.F("server-1"),
				CheckRegions:         khulnasoft.F([]healthchecks.CheckRegion{healthchecks.CheckRegionWnam, healthchecks.CheckRegionEnam}),
				ConsecutiveFails:     khulnasoft.F(int64(0)),
				ConsecutiveSuccesses: khulnasoft.F(int64(0)),
				Description:          khulnasoft.F("Health check for www.example.com"),
				HTTPConfig: khulnasoft.F(healthchecks.HTTPConfigurationParam{
					AllowInsecure:   khulnasoft.F(true),
					ExpectedBody:    khulnasoft.F("success"),
					ExpectedCodes:   khulnasoft.F([]string{"2xx", "302"}),
					FollowRedirects: khulnasoft.F(true),
					Header: khulnasoft.F(map[string][]string{
						"Host":     {"example.com"},
						"X-App-ID": {"abc123"},
					}),
					Method: khulnasoft.F(healthchecks.HTTPConfigurationMethodGet),
					Path:   khulnasoft.F("/health"),
					Port:   khulnasoft.F(int64(0)),
				}),
				Interval:  khulnasoft.F(int64(0)),
				Retries:   khulnasoft.F(int64(0)),
				Suspended: khulnasoft.F(true),
				TCPConfig: khulnasoft.F(healthchecks.TCPConfigurationParam{
					Method: khulnasoft.F(healthchecks.TCPConfigurationMethodConnectionEstablished),
					Port:   khulnasoft.F(int64(0)),
				}),
				Timeout: khulnasoft.F(int64(0)),
				Type:    khulnasoft.F("HTTPS"),
			},
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

func TestHealthcheckListWithOptionalParams(t *testing.T) {
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
	_, err := client.Healthchecks.List(context.TODO(), healthchecks.HealthcheckListParams{
		ZoneID:  khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Page:    khulnasoft.F(1.000000),
		PerPage: khulnasoft.F(5.000000),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHealthcheckDelete(t *testing.T) {
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
	_, err := client.Healthchecks.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		healthchecks.HealthcheckDeleteParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestHealthcheckEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Healthchecks.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		healthchecks.HealthcheckEditParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			QueryHealthcheck: healthchecks.QueryHealthcheckParam{
				Address:              khulnasoft.F("www.example.com"),
				Name:                 khulnasoft.F("server-1"),
				CheckRegions:         khulnasoft.F([]healthchecks.CheckRegion{healthchecks.CheckRegionWnam, healthchecks.CheckRegionEnam}),
				ConsecutiveFails:     khulnasoft.F(int64(0)),
				ConsecutiveSuccesses: khulnasoft.F(int64(0)),
				Description:          khulnasoft.F("Health check for www.example.com"),
				HTTPConfig: khulnasoft.F(healthchecks.HTTPConfigurationParam{
					AllowInsecure:   khulnasoft.F(true),
					ExpectedBody:    khulnasoft.F("success"),
					ExpectedCodes:   khulnasoft.F([]string{"2xx", "302"}),
					FollowRedirects: khulnasoft.F(true),
					Header: khulnasoft.F(map[string][]string{
						"Host":     {"example.com"},
						"X-App-ID": {"abc123"},
					}),
					Method: khulnasoft.F(healthchecks.HTTPConfigurationMethodGet),
					Path:   khulnasoft.F("/health"),
					Port:   khulnasoft.F(int64(0)),
				}),
				Interval:  khulnasoft.F(int64(0)),
				Retries:   khulnasoft.F(int64(0)),
				Suspended: khulnasoft.F(true),
				TCPConfig: khulnasoft.F(healthchecks.TCPConfigurationParam{
					Method: khulnasoft.F(healthchecks.TCPConfigurationMethodConnectionEstablished),
					Port:   khulnasoft.F(int64(0)),
				}),
				Timeout: khulnasoft.F(int64(0)),
				Type:    khulnasoft.F("HTTPS"),
			},
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

func TestHealthcheckGet(t *testing.T) {
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
	_, err := client.Healthchecks.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		healthchecks.HealthcheckGetParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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