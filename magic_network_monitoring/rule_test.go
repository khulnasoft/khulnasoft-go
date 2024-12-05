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

func TestRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Rules.New(context.TODO(), magic_network_monitoring.RuleNewParams{
		AccountID:              khulnasoft.F("6f91088a406011ed95aed352566e8d4c"),
		Duration:               khulnasoft.F("300s"),
		Name:                   khulnasoft.F("my_rule_1"),
		AutomaticAdvertisement: khulnasoft.F(true),
		Bandwidth:              khulnasoft.F(1000.000000),
		PacketThreshold:        khulnasoft.F(10000.000000),
		Prefixes:               khulnasoft.F([]string{"203.0.113.1/32", "203.0.113.1/32", "203.0.113.1/32"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Rules.Update(context.TODO(), magic_network_monitoring.RuleUpdateParams{
		AccountID:              khulnasoft.F("6f91088a406011ed95aed352566e8d4c"),
		Duration:               khulnasoft.F("300s"),
		Name:                   khulnasoft.F("my_rule_1"),
		ID:                     khulnasoft.F("2890e6fa406311ed9b5a23f70f6fb8cf"),
		AutomaticAdvertisement: khulnasoft.F(true),
		Bandwidth:              khulnasoft.F(1000.000000),
		PacketThreshold:        khulnasoft.F(10000.000000),
		Prefixes:               khulnasoft.F([]string{"203.0.113.1/32", "203.0.113.1/32", "203.0.113.1/32"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleList(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Rules.List(context.TODO(), magic_network_monitoring.RuleListParams{
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

func TestRuleDelete(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Rules.Delete(
		context.TODO(),
		"2890e6fa406311ed9b5a23f70f6fb8cf",
		magic_network_monitoring.RuleDeleteParams{
			AccountID: khulnasoft.F("6f91088a406011ed95aed352566e8d4c"),
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

func TestRuleEditWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Rules.Edit(
		context.TODO(),
		"2890e6fa406311ed9b5a23f70f6fb8cf",
		magic_network_monitoring.RuleEditParams{
			AccountID:              khulnasoft.F("6f91088a406011ed95aed352566e8d4c"),
			AutomaticAdvertisement: khulnasoft.F(true),
			Bandwidth:              khulnasoft.F(1000.000000),
			Duration:               khulnasoft.F("300s"),
			Name:                   khulnasoft.F("my_rule_1"),
			PacketThreshold:        khulnasoft.F(10000.000000),
			Prefixes:               khulnasoft.F([]string{"203.0.113.1/32", "203.0.113.1/32", "203.0.113.1/32"}),
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

func TestRuleGet(t *testing.T) {
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
	_, err := client.MagicNetworkMonitoring.Rules.Get(
		context.TODO(),
		"2890e6fa406311ed9b5a23f70f6fb8cf",
		magic_network_monitoring.RuleGetParams{
			AccountID: khulnasoft.F("6f91088a406011ed95aed352566e8d4c"),
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
