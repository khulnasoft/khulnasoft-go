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

func TestTunnelNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Tunnels.New(context.TODO(), zero_trust.TunnelNewParams{
		AccountID:    khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		Name:         khulnasoft.F("blog"),
		ConfigSrc:    khulnasoft.F(zero_trust.TunnelNewParamsConfigSrcLocal),
		TunnelSecret: khulnasoft.F("AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTunnelListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Tunnels.List(context.TODO(), zero_trust.TunnelListParams{
		AccountID:     khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		ExcludePrefix: khulnasoft.F("vpc1-"),
		ExistedAt:     khulnasoft.F(time.Now()),
		IncludePrefix: khulnasoft.F("vpc1-"),
		IsDeleted:     khulnasoft.F(true),
		Name:          khulnasoft.F("blog"),
		Page:          khulnasoft.F(1.000000),
		PerPage:       khulnasoft.F(1.000000),
		Status:        khulnasoft.F(zero_trust.TunnelListParamsStatusInactive),
		UUID:          khulnasoft.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
		WasActiveAt:   khulnasoft.F(time.Now()),
		WasInactiveAt: khulnasoft.F(time.Now()),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTunnelDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Tunnels.Delete(
		context.TODO(),
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		zero_trust.TunnelDeleteParams{
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

func TestTunnelEditWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Tunnels.Edit(
		context.TODO(),
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		zero_trust.TunnelEditParams{
			AccountID:    khulnasoft.F("699d98642c564d2e855e9661899b7252"),
			Name:         khulnasoft.F("blog"),
			TunnelSecret: khulnasoft.F("AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="),
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

func TestTunnelGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Tunnels.Get(
		context.TODO(),
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		zero_trust.TunnelGetParams{
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
