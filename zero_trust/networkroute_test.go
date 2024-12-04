// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/zero_trust"
)

func TestNetworkRouteNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.Routes.New(context.TODO(), zero_trust.NetworkRouteNewParams{
		AccountID:        khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		Network:          khulnasoft.F("172.16.0.0/16"),
		TunnelID:         khulnasoft.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
		Comment:          khulnasoft.F("Example comment for this route."),
		VirtualNetworkID: khulnasoft.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkRouteListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.Routes.List(context.TODO(), zero_trust.NetworkRouteListParams{
		AccountID:        khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		Comment:          khulnasoft.F("Example comment for this route."),
		ExistedAt:        khulnasoft.F(time.Now()),
		IsDeleted:        khulnasoft.F(true),
		NetworkSubset:    khulnasoft.F("172.16.0.0/16"),
		NetworkSuperset:  khulnasoft.F("172.16.0.0/16"),
		Page:             khulnasoft.F(1.000000),
		PerPage:          khulnasoft.F(1.000000),
		RouteID:          khulnasoft.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
		TunTypes:         khulnasoft.F("cfd_tunnel,warp_connector"),
		TunnelID:         khulnasoft.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
		VirtualNetworkID: khulnasoft.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkRouteDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.Routes.Delete(
		context.TODO(),
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		zero_trust.NetworkRouteDeleteParams{
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

func TestNetworkRouteEditWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.Routes.Edit(
		context.TODO(),
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		zero_trust.NetworkRouteEditParams{
			AccountID:        khulnasoft.F("699d98642c564d2e855e9661899b7252"),
			Comment:          khulnasoft.F("Example comment for this route."),
			Network:          khulnasoft.F("172.16.0.0/16"),
			TunnelID:         khulnasoft.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
			VirtualNetworkID: khulnasoft.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
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

func TestNetworkRouteGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.Routes.Get(
		context.TODO(),
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		zero_trust.NetworkRouteGetParams{
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
