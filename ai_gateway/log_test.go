// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/ai_gateway"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.Logs.List(
		context.TODO(),
		"my-gateway",
		ai_gateway.LogListParams{
			AccountID:           khulnasoft.F("0d37909e38d3e99c29fa2cd343ac421a"),
			Cached:              khulnasoft.F(true),
			Direction:           khulnasoft.F(ai_gateway.LogListParamsDirectionAsc),
			EndDate:             khulnasoft.F(time.Now()),
			Feedback:            khulnasoft.F(ai_gateway.LogListParamsFeedback0),
			MaxCost:             khulnasoft.F(0.000000),
			MaxDuration:         khulnasoft.F(0.000000),
			MaxTokensIn:         khulnasoft.F(0.000000),
			MaxTokensOut:        khulnasoft.F(0.000000),
			MaxTotalTokens:      khulnasoft.F(0.000000),
			MetaInfo:            khulnasoft.F(true),
			MinCost:             khulnasoft.F(0.000000),
			MinDuration:         khulnasoft.F(0.000000),
			MinTokensIn:         khulnasoft.F(0.000000),
			MinTokensOut:        khulnasoft.F(0.000000),
			MinTotalTokens:      khulnasoft.F(0.000000),
			Model:               khulnasoft.F("model"),
			ModelType:           khulnasoft.F("model_type"),
			OrderBy:             khulnasoft.F(ai_gateway.LogListParamsOrderByCreatedAt),
			OrderByDirection:    khulnasoft.F(ai_gateway.LogListParamsOrderByDirectionAsc),
			Page:                khulnasoft.F(int64(1)),
			PerPage:             khulnasoft.F(int64(5)),
			Provider:            khulnasoft.F("provider"),
			RequestContentType:  khulnasoft.F("request_content_type"),
			ResponseContentType: khulnasoft.F("response_content_type"),
			Search:              khulnasoft.F("search"),
			StartDate:           khulnasoft.F(time.Now()),
			Success:             khulnasoft.F(true),
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
