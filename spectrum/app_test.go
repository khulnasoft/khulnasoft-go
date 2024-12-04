// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/shared"
	"github.com/khulnasoft/khulnasoft-go/spectrum"
)

func TestAppNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Spectrum.Apps.New(context.TODO(), spectrum.AppNewParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body: spectrum.AppNewParamsBodySpectrumConfigAppConfig{
			DNS: khulnasoft.F(spectrum.DNSParam{
				Name: khulnasoft.F("ssh.example.com"),
				Type: khulnasoft.F(spectrum.DNSTypeCNAME),
			}),
			IPFirewall:       khulnasoft.F(true),
			Protocol:         khulnasoft.F("tcp/22"),
			ProxyProtocol:    khulnasoft.F(spectrum.AppNewParamsBodySpectrumConfigAppConfigProxyProtocolOff),
			TLS:              khulnasoft.F(spectrum.AppNewParamsBodySpectrumConfigAppConfigTLSOff),
			TrafficType:      khulnasoft.F(spectrum.AppNewParamsBodySpectrumConfigAppConfigTrafficTypeDirect),
			ArgoSmartRouting: khulnasoft.F(true),
			EdgeIPs: khulnasoft.F[spectrum.EdgeIPsUnionParam](spectrum.EdgeIPsObjectParam{
				Connectivity: khulnasoft.F(spectrum.EdgeIPsObjectConnectivityAll),
				Type:         khulnasoft.F(spectrum.EdgeIPsObjectTypeDynamic),
			}),
			OriginDirect: khulnasoft.F([]string{"tcp://127.0.0.1:8080", "tcp://127.0.0.1:8080", "tcp://127.0.0.1:8080"}),
			OriginDNS: khulnasoft.F(spectrum.OriginDNSParam{
				Name: khulnasoft.F("origin.example.com"),
				TTL:  khulnasoft.F(int64(600)),
				Type: khulnasoft.F(spectrum.OriginDNSTypeEmpty),
			}),
			OriginPort: khulnasoft.F[spectrum.OriginPortUnionParam](shared.UnionInt(int64(22))),
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

func TestAppUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Spectrum.Apps.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		spectrum.AppUpdateParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: spectrum.AppUpdateParamsBodySpectrumConfigAppConfig{
				DNS: khulnasoft.F(spectrum.DNSParam{
					Name: khulnasoft.F("ssh.example.com"),
					Type: khulnasoft.F(spectrum.DNSTypeCNAME),
				}),
				IPFirewall:       khulnasoft.F(true),
				Protocol:         khulnasoft.F("tcp/22"),
				ProxyProtocol:    khulnasoft.F(spectrum.AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocolOff),
				TLS:              khulnasoft.F(spectrum.AppUpdateParamsBodySpectrumConfigAppConfigTLSOff),
				TrafficType:      khulnasoft.F(spectrum.AppUpdateParamsBodySpectrumConfigAppConfigTrafficTypeDirect),
				ArgoSmartRouting: khulnasoft.F(true),
				EdgeIPs: khulnasoft.F[spectrum.EdgeIPsUnionParam](spectrum.EdgeIPsObjectParam{
					Connectivity: khulnasoft.F(spectrum.EdgeIPsObjectConnectivityAll),
					Type:         khulnasoft.F(spectrum.EdgeIPsObjectTypeDynamic),
				}),
				OriginDirect: khulnasoft.F([]string{"tcp://127.0.0.1:8080", "tcp://127.0.0.1:8080", "tcp://127.0.0.1:8080"}),
				OriginDNS: khulnasoft.F(spectrum.OriginDNSParam{
					Name: khulnasoft.F("origin.example.com"),
					TTL:  khulnasoft.F(int64(600)),
					Type: khulnasoft.F(spectrum.OriginDNSTypeEmpty),
				}),
				OriginPort: khulnasoft.F[spectrum.OriginPortUnionParam](shared.UnionInt(int64(22))),
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

func TestAppListWithOptionalParams(t *testing.T) {
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
	_, err := client.Spectrum.Apps.List(context.TODO(), spectrum.AppListParams{
		ZoneID:    khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: khulnasoft.F(spectrum.AppListParamsDirectionAsc),
		Order:     khulnasoft.F(spectrum.AppListParamsOrderProtocol),
		Page:      khulnasoft.F(1.000000),
		PerPage:   khulnasoft.F(1.000000),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppDelete(t *testing.T) {
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
	_, err := client.Spectrum.Apps.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		spectrum.AppDeleteParams{
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

func TestAppGet(t *testing.T) {
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
	_, err := client.Spectrum.Apps.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		spectrum.AppGetParams{
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
