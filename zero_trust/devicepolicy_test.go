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

func TestDevicePolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.New(context.TODO(), zero_trust.DevicePolicyNewParams{
		AccountID:           khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		Match:               khulnasoft.F("user.identity == \"test@khulnasoft.com\""),
		Name:                khulnasoft.F("Allow Developers"),
		Precedence:          khulnasoft.F(100.000000),
		AllowModeSwitch:     khulnasoft.F(true),
		AllowUpdates:        khulnasoft.F(true),
		AllowedToLeave:      khulnasoft.F(true),
		AutoConnect:         khulnasoft.F(0.000000),
		CaptivePortal:       khulnasoft.F(180.000000),
		Description:         khulnasoft.F("Policy for test teams."),
		DisableAutoFallback: khulnasoft.F(true),
		Enabled:             khulnasoft.F(true),
		ExcludeOfficeIPs:    khulnasoft.F(true),
		LANAllowMinutes:     khulnasoft.F(30.000000),
		LANAllowSubnetSize:  khulnasoft.F(24.000000),
		ServiceModeV2: khulnasoft.F(zero_trust.DevicePolicyNewParamsServiceModeV2{
			Mode: khulnasoft.F("proxy"),
			Port: khulnasoft.F(3000.000000),
		}),
		SupportURL:     khulnasoft.F("https://1.1.1.1/help"),
		SwitchLocked:   khulnasoft.F(true),
		TunnelProtocol: khulnasoft.F("wireguard"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDevicePolicyList(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.List(context.TODO(), zero_trust.DevicePolicyListParams{
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

func TestDevicePolicyDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePolicyDeleteParams{
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

func TestDevicePolicyEditWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.Edit(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePolicyEditParams{
			AccountID:           khulnasoft.F("699d98642c564d2e855e9661899b7252"),
			AllowModeSwitch:     khulnasoft.F(true),
			AllowUpdates:        khulnasoft.F(true),
			AllowedToLeave:      khulnasoft.F(true),
			AutoConnect:         khulnasoft.F(0.000000),
			CaptivePortal:       khulnasoft.F(180.000000),
			Description:         khulnasoft.F("Policy for test teams."),
			DisableAutoFallback: khulnasoft.F(true),
			Enabled:             khulnasoft.F(true),
			ExcludeOfficeIPs:    khulnasoft.F(true),
			Match:               khulnasoft.F("user.identity == \"test@khulnasoft.com\""),
			Name:                khulnasoft.F("Allow Developers"),
			Precedence:          khulnasoft.F(100.000000),
			ServiceModeV2: khulnasoft.F(zero_trust.DevicePolicyEditParamsServiceModeV2{
				Mode: khulnasoft.F("proxy"),
				Port: khulnasoft.F(3000.000000),
			}),
			SupportURL:     khulnasoft.F("https://1.1.1.1/help"),
			SwitchLocked:   khulnasoft.F(true),
			TunnelProtocol: khulnasoft.F("wireguard"),
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

func TestDevicePolicyGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePolicyGetParams{
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
