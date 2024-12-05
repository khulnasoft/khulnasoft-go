// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/rulesets"
)

func TestPhaseUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Rulesets.Phases.Update(
		context.TODO(),
		rulesets.PhaseDDoSL4,
		rulesets.PhaseUpdateParams{
			Rules: khulnasoft.F([]rulesets.PhaseUpdateParamsRuleUnion{rulesets.BlockRuleParam{
				ID:     khulnasoft.F("3a03d665bac047339bb530ecb439a90d"),
				Action: khulnasoft.F(rulesets.BlockRuleActionBlock),
				ActionParameters: khulnasoft.F(rulesets.BlockRuleActionParametersParam{
					Response: khulnasoft.F(rulesets.BlockRuleActionParametersResponseParam{
						Content:     khulnasoft.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: khulnasoft.F("application/json"),
						StatusCode:  khulnasoft.F(int64(400)),
					}),
				}),
				Description: khulnasoft.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     khulnasoft.F(true),
				Expression:  khulnasoft.F("ip.src ne 1.1.1.1"),
				Logging: khulnasoft.F(rulesets.LoggingParam{
					Enabled: khulnasoft.F(true),
				}),
				Ref: khulnasoft.F("my_ref"),
			}, rulesets.BlockRuleParam{
				ID:     khulnasoft.F("3a03d665bac047339bb530ecb439a90d"),
				Action: khulnasoft.F(rulesets.BlockRuleActionBlock),
				ActionParameters: khulnasoft.F(rulesets.BlockRuleActionParametersParam{
					Response: khulnasoft.F(rulesets.BlockRuleActionParametersResponseParam{
						Content:     khulnasoft.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: khulnasoft.F("application/json"),
						StatusCode:  khulnasoft.F(int64(400)),
					}),
				}),
				Description: khulnasoft.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     khulnasoft.F(true),
				Expression:  khulnasoft.F("ip.src ne 1.1.1.1"),
				Logging: khulnasoft.F(rulesets.LoggingParam{
					Enabled: khulnasoft.F(true),
				}),
				Ref: khulnasoft.F("my_ref"),
			}, rulesets.BlockRuleParam{
				ID:     khulnasoft.F("3a03d665bac047339bb530ecb439a90d"),
				Action: khulnasoft.F(rulesets.BlockRuleActionBlock),
				ActionParameters: khulnasoft.F(rulesets.BlockRuleActionParametersParam{
					Response: khulnasoft.F(rulesets.BlockRuleActionParametersResponseParam{
						Content:     khulnasoft.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: khulnasoft.F("application/json"),
						StatusCode:  khulnasoft.F(int64(400)),
					}),
				}),
				Description: khulnasoft.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     khulnasoft.F(true),
				Expression:  khulnasoft.F("ip.src ne 1.1.1.1"),
				Logging: khulnasoft.F(rulesets.LoggingParam{
					Enabled: khulnasoft.F(true),
				}),
				Ref: khulnasoft.F("my_ref"),
			}}),
			AccountID:   khulnasoft.F("account_id"),
			Description: khulnasoft.F("My ruleset to execute managed rulesets"),
			Name:        khulnasoft.F("My ruleset"),
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

func TestPhaseGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Rulesets.Phases.Get(
		context.TODO(),
		rulesets.PhaseDDoSL4,
		rulesets.PhaseGetParams{
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
