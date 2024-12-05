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

func TestGatewayLocationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Locations.New(context.TODO(), zero_trust.GatewayLocationNewParams{
		AccountID:           khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		Name:                khulnasoft.F("Austin Office Location"),
		ClientDefault:       khulnasoft.F(false),
		DNSDestinationIPsID: khulnasoft.F("0e4a32c6-6fb8-4858-9296-98f51631e8e6"),
		ECSSupport:          khulnasoft.F(false),
		Endpoints: khulnasoft.F(zero_trust.EndpointParam{
			DOH: khulnasoft.F(zero_trust.DOHEndpointParam{
				Enabled: khulnasoft.F(true),
				Networks: khulnasoft.F([]zero_trust.IPNetworkParam{{
					Network: khulnasoft.F("2001:85a3::/64"),
				}, {
					Network: khulnasoft.F("2001:85a3::/64"),
				}, {
					Network: khulnasoft.F("2001:85a3::/64"),
				}}),
				RequireToken: khulnasoft.F(true),
			}),
			DOT: khulnasoft.F(zero_trust.DOTEndpointParam{
				Enabled: khulnasoft.F(true),
				Networks: khulnasoft.F([]zero_trust.IPNetworkParam{{
					Network: khulnasoft.F("2001:85a3::/64"),
				}, {
					Network: khulnasoft.F("2001:85a3::/64"),
				}, {
					Network: khulnasoft.F("2001:85a3::/64"),
				}}),
			}),
			IPV4: khulnasoft.F(zero_trust.IPV4EndpointParam{
				Enabled: khulnasoft.F(true),
			}),
			IPV6: khulnasoft.F(zero_trust.IPV6EndpointParam{
				Enabled: khulnasoft.F(true),
				Networks: khulnasoft.F([]zero_trust.IPV6NetworkParam{{
					Network: khulnasoft.F("2001:85a3::/64"),
				}, {
					Network: khulnasoft.F("2001:85a3::/64"),
				}, {
					Network: khulnasoft.F("2001:85a3::/64"),
				}}),
			}),
		}),
		Networks: khulnasoft.F([]zero_trust.GatewayLocationNewParamsNetwork{{
			Network: khulnasoft.F("192.0.2.1/32"),
		}, {
			Network: khulnasoft.F("192.0.2.1/32"),
		}, {
			Network: khulnasoft.F("192.0.2.1/32"),
		}}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayLocationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Locations.Update(
		context.TODO(),
		"ed35569b41ce4d1facfe683550f54086",
		zero_trust.GatewayLocationUpdateParams{
			AccountID:           khulnasoft.F("699d98642c564d2e855e9661899b7252"),
			Name:                khulnasoft.F("Austin Office Location"),
			ClientDefault:       khulnasoft.F(false),
			DNSDestinationIPsID: khulnasoft.F("0e4a32c6-6fb8-4858-9296-98f51631e8e6"),
			ECSSupport:          khulnasoft.F(false),
			Endpoints: khulnasoft.F(zero_trust.EndpointParam{
				DOH: khulnasoft.F(zero_trust.DOHEndpointParam{
					Enabled: khulnasoft.F(true),
					Networks: khulnasoft.F([]zero_trust.IPNetworkParam{{
						Network: khulnasoft.F("2001:85a3::/64"),
					}, {
						Network: khulnasoft.F("2001:85a3::/64"),
					}, {
						Network: khulnasoft.F("2001:85a3::/64"),
					}}),
					RequireToken: khulnasoft.F(true),
				}),
				DOT: khulnasoft.F(zero_trust.DOTEndpointParam{
					Enabled: khulnasoft.F(true),
					Networks: khulnasoft.F([]zero_trust.IPNetworkParam{{
						Network: khulnasoft.F("2001:85a3::/64"),
					}, {
						Network: khulnasoft.F("2001:85a3::/64"),
					}, {
						Network: khulnasoft.F("2001:85a3::/64"),
					}}),
				}),
				IPV4: khulnasoft.F(zero_trust.IPV4EndpointParam{
					Enabled: khulnasoft.F(true),
				}),
				IPV6: khulnasoft.F(zero_trust.IPV6EndpointParam{
					Enabled: khulnasoft.F(true),
					Networks: khulnasoft.F([]zero_trust.IPV6NetworkParam{{
						Network: khulnasoft.F("2001:85a3::/64"),
					}, {
						Network: khulnasoft.F("2001:85a3::/64"),
					}, {
						Network: khulnasoft.F("2001:85a3::/64"),
					}}),
				}),
			}),
			Networks: khulnasoft.F([]zero_trust.GatewayLocationUpdateParamsNetwork{{
				Network: khulnasoft.F("192.0.2.1/32"),
			}, {
				Network: khulnasoft.F("192.0.2.1/32"),
			}, {
				Network: khulnasoft.F("192.0.2.1/32"),
			}}),
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

func TestGatewayLocationList(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Locations.List(context.TODO(), zero_trust.GatewayLocationListParams{
		AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayLocationDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Locations.Delete(
		context.TODO(),
		"ed35569b41ce4d1facfe683550f54086",
		zero_trust.GatewayLocationDeleteParams{
			AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
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

func TestGatewayLocationGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Locations.Get(
		context.TODO(),
		"ed35569b41ce4d1facfe683550f54086",
		zero_trust.GatewayLocationGetParams{
			AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
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
