// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/magic_network_monitoring"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestConfigNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Configs.New(context.TODO(), magic_network_monitoring.ConfigNewParams{
		AccountID:       khulnasoft.F("6f91088a406011ed95aed352566e8d4c"),
		DefaultSampling: khulnasoft.F(1.000000),
		Name:            khulnasoft.F("khulnasoft user's account"),
		RouterIPs:       khulnasoft.F([]string{"203.0.113.1", "203.0.113.1", "203.0.113.1"}),
		WARPDevices: khulnasoft.F([]magic_network_monitoring.ConfigNewParamsWARPDevice{{
			ID:       khulnasoft.F("5360368d-b351-4791-abe1-93550dabd351"),
			Name:     khulnasoft.F("My warp device"),
			RouterIP: khulnasoft.F("203.0.113.1"),
		}, {
			ID:       khulnasoft.F("5360368d-b351-4791-abe1-93550dabd351"),
			Name:     khulnasoft.F("My warp device"),
			RouterIP: khulnasoft.F("203.0.113.1"),
		}, {
			ID:       khulnasoft.F("5360368d-b351-4791-abe1-93550dabd351"),
			Name:     khulnasoft.F("My warp device"),
			RouterIP: khulnasoft.F("203.0.113.1"),
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

func TestConfigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Configs.Update(context.TODO(), magic_network_monitoring.ConfigUpdateParams{
		AccountID:       khulnasoft.F("6f91088a406011ed95aed352566e8d4c"),
		DefaultSampling: khulnasoft.F(1.000000),
		Name:            khulnasoft.F("khulnasoft user's account"),
		RouterIPs:       khulnasoft.F([]string{"203.0.113.1", "203.0.113.1", "203.0.113.1"}),
		WARPDevices: khulnasoft.F([]magic_network_monitoring.ConfigUpdateParamsWARPDevice{{
			ID:       khulnasoft.F("5360368d-b351-4791-abe1-93550dabd351"),
			Name:     khulnasoft.F("My warp device"),
			RouterIP: khulnasoft.F("203.0.113.1"),
		}, {
			ID:       khulnasoft.F("5360368d-b351-4791-abe1-93550dabd351"),
			Name:     khulnasoft.F("My warp device"),
			RouterIP: khulnasoft.F("203.0.113.1"),
		}, {
			ID:       khulnasoft.F("5360368d-b351-4791-abe1-93550dabd351"),
			Name:     khulnasoft.F("My warp device"),
			RouterIP: khulnasoft.F("203.0.113.1"),
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

func TestConfigDelete(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Configs.Delete(context.TODO(), magic_network_monitoring.ConfigDeleteParams{
		AccountID: khulnasoft.F("6f91088a406011ed95aed352566e8d4c"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigEditWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Configs.Edit(context.TODO(), magic_network_monitoring.ConfigEditParams{
		AccountID:       khulnasoft.F("6f91088a406011ed95aed352566e8d4c"),
		DefaultSampling: khulnasoft.F(1.000000),
		Name:            khulnasoft.F("khulnasoft user's account"),
		RouterIPs:       khulnasoft.F([]string{"203.0.113.1", "203.0.113.1", "203.0.113.1"}),
		WARPDevices: khulnasoft.F([]magic_network_monitoring.ConfigEditParamsWARPDevice{{
			ID:       khulnasoft.F("5360368d-b351-4791-abe1-93550dabd351"),
			Name:     khulnasoft.F("My warp device"),
			RouterIP: khulnasoft.F("203.0.113.1"),
		}, {
			ID:       khulnasoft.F("5360368d-b351-4791-abe1-93550dabd351"),
			Name:     khulnasoft.F("My warp device"),
			RouterIP: khulnasoft.F("203.0.113.1"),
		}, {
			ID:       khulnasoft.F("5360368d-b351-4791-abe1-93550dabd351"),
			Name:     khulnasoft.F("My warp device"),
			RouterIP: khulnasoft.F("203.0.113.1"),
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

func TestConfigGet(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Configs.Get(context.TODO(), magic_network_monitoring.ConfigGetParams{
		AccountID: khulnasoft.F("6f91088a406011ed95aed352566e8d4c"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
