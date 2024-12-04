// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/magic_transit"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestIPSECTunnelNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.New(context.TODO(), magic_transit.IPSECTunnelNewParams{
		AccountID:          khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		KhulnasoftEndpoint: khulnasoft.F("203.0.113.1"),
		InterfaceAddress:   khulnasoft.F("192.0.2.0/31"),
		Name:               khulnasoft.F("IPsec_1"),
		CustomerEndpoint:   khulnasoft.F("203.0.113.1"),
		Description:        khulnasoft.F("Tunnel for ISP X"),
		HealthCheck: khulnasoft.F(magic_transit.HealthCheckParam{
			Direction: khulnasoft.F(magic_transit.HealthCheckDirectionUnidirectional),
			Enabled:   khulnasoft.F(true),
			Rate:      khulnasoft.F(magic_transit.HealthCheckRateLow),
			Target:    khulnasoft.F("203.0.113.1"),
			Type:      khulnasoft.F(magic_transit.HealthCheckTypeReply),
		}),
		PSK:              khulnasoft.F("O3bwKSjnaoCxDoUxjcq4Rk8ZKkezQUiy"),
		ReplayProtection: khulnasoft.F(false),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIPSECTunnelUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.IPSECTunnelUpdateParams{
			AccountID:          khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			KhulnasoftEndpoint: khulnasoft.F("203.0.113.1"),
			InterfaceAddress:   khulnasoft.F("192.0.2.0/31"),
			Name:               khulnasoft.F("IPsec_1"),
			CustomerEndpoint:   khulnasoft.F("203.0.113.1"),
			Description:        khulnasoft.F("Tunnel for ISP X"),
			HealthCheck: khulnasoft.F(magic_transit.HealthCheckParam{
				Direction: khulnasoft.F(magic_transit.HealthCheckDirectionUnidirectional),
				Enabled:   khulnasoft.F(true),
				Rate:      khulnasoft.F(magic_transit.HealthCheckRateLow),
				Target:    khulnasoft.F("203.0.113.1"),
				Type:      khulnasoft.F(magic_transit.HealthCheckTypeReply),
			}),
			PSK:              khulnasoft.F("O3bwKSjnaoCxDoUxjcq4Rk8ZKkezQUiy"),
			ReplayProtection: khulnasoft.F(false),
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

func TestIPSECTunnelList(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.List(context.TODO(), magic_transit.IPSECTunnelListParams{
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

func TestIPSECTunnelDelete(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.IPSECTunnelDeleteParams{
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

func TestIPSECTunnelGet(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.IPSECTunnelGetParams{
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

func TestIPSECTunnelPSKGenerate(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.PSKGenerate(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.IPSECTunnelPSKGenerateParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body:      map[string]interface{}{},
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
