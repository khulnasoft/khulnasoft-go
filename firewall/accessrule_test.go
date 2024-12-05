// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/firewall"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestAccessRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.AccessRules.New(context.TODO(), firewall.AccessRuleNewParams{
		Configuration: khulnasoft.F[firewall.AccessRuleNewParamsConfigurationUnion](firewall.AccessRuleIPConfigurationParam{
			Target: khulnasoft.F(firewall.AccessRuleIPConfigurationTargetIP),
			Value:  khulnasoft.F("198.51.100.4"),
		}),
		Mode:      khulnasoft.F(firewall.AccessRuleNewParamsModeBlock),
		AccountID: khulnasoft.F("account_id"),
		Notes:     khulnasoft.F("This rule is enabled because of an event that occurred on date X."),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.AccessRules.List(context.TODO(), firewall.AccessRuleListParams{
		AccountID: khulnasoft.F("account_id"),
		Configuration: khulnasoft.F(firewall.AccessRuleListParamsConfiguration{
			Target: khulnasoft.F(firewall.AccessRuleListParamsConfigurationTargetIP),
			Value:  khulnasoft.F("198.51.100.4"),
		}),
		Direction: khulnasoft.F(firewall.AccessRuleListParamsDirectionAsc),
		Match:     khulnasoft.F(firewall.AccessRuleListParamsMatchAny),
		Mode:      khulnasoft.F(firewall.AccessRuleListParamsModeBlock),
		Notes:     khulnasoft.F("my note"),
		Order:     khulnasoft.F(firewall.AccessRuleListParamsOrderConfigurationTarget),
		Page:      khulnasoft.F(1.000000),
		PerPage:   khulnasoft.F(20.000000),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessRuleDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.AccessRules.Delete(
		context.TODO(),
		"de677e5818985db1285d0e80225f06e5",
		firewall.AccessRuleDeleteParams{
			AccountID: khulnasoft.F("account_id"),
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

func TestAccessRuleEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.AccessRules.Edit(
		context.TODO(),
		"de677e5818985db1285d0e80225f06e5",
		firewall.AccessRuleEditParams{
			Configuration: khulnasoft.F[firewall.AccessRuleEditParamsConfigurationUnion](firewall.AccessRuleIPConfigurationParam{
				Target: khulnasoft.F(firewall.AccessRuleIPConfigurationTargetIP),
				Value:  khulnasoft.F("198.51.100.4"),
			}),
			Mode:      khulnasoft.F(firewall.AccessRuleEditParamsModeBlock),
			AccountID: khulnasoft.F("account_id"),
			Notes:     khulnasoft.F("This rule is enabled because of an event that occurred on date X."),
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

func TestAccessRuleGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.AccessRules.Get(
		context.TODO(),
		"de677e5818985db1285d0e80225f06e5",
		firewall.AccessRuleGetParams{
			AccountID: khulnasoft.F("account_id"),
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
