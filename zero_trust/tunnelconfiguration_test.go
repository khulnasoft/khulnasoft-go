// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/zero_trust"
)

func TestTunnelConfigurationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Tunnels.Configurations.Update(
		context.TODO(),
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		zero_trust.TunnelConfigurationUpdateParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Config: khulnasoft.F(zero_trust.TunnelConfigurationUpdateParamsConfig{
				Ingress: khulnasoft.F([]zero_trust.TunnelConfigurationUpdateParamsConfigIngress{{
					Hostname: khulnasoft.F("tunnel.example.com"),
					Service:  khulnasoft.F("https://localhost:8001"),
					OriginRequest: khulnasoft.F(zero_trust.TunnelConfigurationUpdateParamsConfigIngressOriginRequest{
						Access: khulnasoft.F(zero_trust.TunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess{
							AUDTag:   khulnasoft.F([]string{"string", "string", "string"}),
							TeamName: khulnasoft.F("teamName"),
							Required: khulnasoft.F(true),
						}),
						CAPool:                 khulnasoft.F("caPool"),
						ConnectTimeout:         khulnasoft.F(int64(0)),
						DisableChunkedEncoding: khulnasoft.F(true),
						HTTP2Origin:            khulnasoft.F(true),
						HTTPHostHeader:         khulnasoft.F("httpHostHeader"),
						KeepAliveConnections:   khulnasoft.F(int64(0)),
						KeepAliveTimeout:       khulnasoft.F(int64(0)),
						NoHappyEyeballs:        khulnasoft.F(true),
						NoTLSVerify:            khulnasoft.F(true),
						OriginServerName:       khulnasoft.F("originServerName"),
						ProxyType:              khulnasoft.F("proxyType"),
						TCPKeepAlive:           khulnasoft.F(int64(0)),
						TLSTimeout:             khulnasoft.F(int64(0)),
					}),
					Path: khulnasoft.F("subpath"),
				}}),
				OriginRequest: khulnasoft.F(zero_trust.TunnelConfigurationUpdateParamsConfigOriginRequest{
					Access: khulnasoft.F(zero_trust.TunnelConfigurationUpdateParamsConfigOriginRequestAccess{
						AUDTag:   khulnasoft.F([]string{"string", "string", "string"}),
						TeamName: khulnasoft.F("teamName"),
						Required: khulnasoft.F(true),
					}),
					CAPool:                 khulnasoft.F("caPool"),
					ConnectTimeout:         khulnasoft.F(int64(0)),
					DisableChunkedEncoding: khulnasoft.F(true),
					HTTP2Origin:            khulnasoft.F(true),
					HTTPHostHeader:         khulnasoft.F("httpHostHeader"),
					KeepAliveConnections:   khulnasoft.F(int64(0)),
					KeepAliveTimeout:       khulnasoft.F(int64(0)),
					NoHappyEyeballs:        khulnasoft.F(true),
					NoTLSVerify:            khulnasoft.F(true),
					OriginServerName:       khulnasoft.F("originServerName"),
					ProxyType:              khulnasoft.F("proxyType"),
					TCPKeepAlive:           khulnasoft.F(int64(0)),
					TLSTimeout:             khulnasoft.F(int64(0)),
				}),
			}),
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

func TestTunnelConfigurationGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Tunnels.Configurations.Get(
		context.TODO(),
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		zero_trust.TunnelConfigurationGetParams{
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
