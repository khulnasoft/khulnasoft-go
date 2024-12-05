// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/ai_gateway"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestAIGatewayNew(t *testing.T) {
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
	_, err := client.AIGateway.New(context.TODO(), ai_gateway.AIGatewayNewParams{
		AccountID:               khulnasoft.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
		ID:                      khulnasoft.F("my-gateway"),
		CacheInvalidateOnUpdate: khulnasoft.F(true),
		CacheTTL:                khulnasoft.F(int64(0)),
		CollectLogs:             khulnasoft.F(true),
		RateLimitingInterval:    khulnasoft.F(int64(0)),
		RateLimitingLimit:       khulnasoft.F(int64(0)),
		RateLimitingTechnique:   khulnasoft.F(ai_gateway.AIGatewayNewParamsRateLimitingTechniqueFixed),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIGatewayUpdate(t *testing.T) {
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
	_, err := client.AIGateway.Update(
		context.TODO(),
		"my-gateway",
		ai_gateway.AIGatewayUpdateParams{
			AccountID:               khulnasoft.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
			CacheInvalidateOnUpdate: khulnasoft.F(true),
			CacheTTL:                khulnasoft.F(int64(0)),
			CollectLogs:             khulnasoft.F(true),
			RateLimitingInterval:    khulnasoft.F(int64(0)),
			RateLimitingLimit:       khulnasoft.F(int64(0)),
			RateLimitingTechnique:   khulnasoft.F(ai_gateway.AIGatewayUpdateParamsRateLimitingTechniqueFixed),
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

func TestAIGatewayListWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.List(context.TODO(), ai_gateway.AIGatewayListParams{
		AccountID:        khulnasoft.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
		ID:               khulnasoft.F("my-gateway"),
		OrderBy:          khulnasoft.F("order_by"),
		OrderByDirection: khulnasoft.F(ai_gateway.AIGatewayListParamsOrderByDirectionAsc),
		Page:             khulnasoft.F(int64(1)),
		PerPage:          khulnasoft.F(int64(5)),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIGatewayDelete(t *testing.T) {
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
	_, err := client.AIGateway.Delete(
		context.TODO(),
		"id",
		ai_gateway.AIGatewayDeleteParams{
			AccountID: khulnasoft.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
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

func TestAIGatewayGet(t *testing.T) {
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
	_, err := client.AIGateway.Get(
		context.TODO(),
		"my-gateway",
		ai_gateway.AIGatewayGetParams{
			AccountID: khulnasoft.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
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
