// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/email_routing"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestRuleCatchAllUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailRouting.Rules.CatchAlls.Update(context.TODO(), email_routing.RuleCatchAllUpdateParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Actions: khulnasoft.F([]email_routing.CatchAllActionParam{{
			Type:  khulnasoft.F(email_routing.CatchAllActionTypeDrop),
			Value: khulnasoft.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
		}, {
			Type:  khulnasoft.F(email_routing.CatchAllActionTypeDrop),
			Value: khulnasoft.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
		}, {
			Type:  khulnasoft.F(email_routing.CatchAllActionTypeDrop),
			Value: khulnasoft.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
		}}),
		Matchers: khulnasoft.F([]email_routing.CatchAllMatcherParam{{
			Type: khulnasoft.F(email_routing.CatchAllMatcherTypeAll),
		}, {
			Type: khulnasoft.F(email_routing.CatchAllMatcherTypeAll),
		}, {
			Type: khulnasoft.F(email_routing.CatchAllMatcherTypeAll),
		}}),
		Enabled: khulnasoft.F(email_routing.RuleCatchAllUpdateParamsEnabledTrue),
		Name:    khulnasoft.F("Send to user@example.net rule."),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleCatchAllGet(t *testing.T) {
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
	_, err := client.EmailRouting.Rules.CatchAlls.Get(context.TODO(), email_routing.RuleCatchAllGetParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
