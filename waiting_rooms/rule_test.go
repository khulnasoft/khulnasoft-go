// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/waiting_rooms"
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
	_, err := client.WaitingRooms.Rules.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.RuleNewParams{
			ZoneID:      khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Action:      khulnasoft.F(waiting_rooms.RuleNewParamsActionBypassWaitingRoom),
			Expression:  khulnasoft.F("ip.src in {10.20.30.40}"),
			Description: khulnasoft.F("allow all traffic from 10.20.30.40"),
			Enabled:     khulnasoft.F(true),
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

func TestRuleUpdate(t *testing.T) {
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
	_, err := client.WaitingRooms.Rules.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.RuleUpdateParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: []waiting_rooms.RuleUpdateParamsBody{{
				Action:      khulnasoft.F(waiting_rooms.RuleUpdateParamsBodyActionBypassWaitingRoom),
				Expression:  khulnasoft.F("ip.src in {10.20.30.40}"),
				Description: khulnasoft.F("allow all traffic from 10.20.30.40"),
				Enabled:     khulnasoft.F(true),
			}, {
				Action:      khulnasoft.F(waiting_rooms.RuleUpdateParamsBodyActionBypassWaitingRoom),
				Expression:  khulnasoft.F("ip.src in {10.20.30.40}"),
				Description: khulnasoft.F("allow all traffic from 10.20.30.40"),
				Enabled:     khulnasoft.F(true),
			}, {
				Action:      khulnasoft.F(waiting_rooms.RuleUpdateParamsBodyActionBypassWaitingRoom),
				Expression:  khulnasoft.F("ip.src in {10.20.30.40}"),
				Description: khulnasoft.F("allow all traffic from 10.20.30.40"),
				Enabled:     khulnasoft.F(true),
			}},
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
	_, err := client.WaitingRooms.Rules.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		waiting_rooms.RuleDeleteParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
	_, err := client.WaitingRooms.Rules.Edit(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		waiting_rooms.RuleEditParams{
			ZoneID:      khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Action:      khulnasoft.F(waiting_rooms.RuleEditParamsActionBypassWaitingRoom),
			Expression:  khulnasoft.F("ip.src in {10.20.30.40}"),
			Description: khulnasoft.F("allow all traffic from 10.20.30.40"),
			Enabled:     khulnasoft.F(true),
			Position: khulnasoft.F[waiting_rooms.RuleEditParamsPositionUnion](waiting_rooms.RuleEditParamsPositionIndex{
				Index: khulnasoft.F(int64(0)),
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
	_, err := client.WaitingRooms.Rules.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.RuleGetParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
