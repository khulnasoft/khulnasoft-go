// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/api_gateway"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestOperationNew(t *testing.T) {
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
	_, err := client.APIGateway.Operations.New(context.TODO(), api_gateway.OperationNewParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body: []api_gateway.OperationNewParamsBody{{
			Endpoint: khulnasoft.F("/api/v1/users/{var1}"),
			Host:     khulnasoft.F("www.example.com"),
			Method:   khulnasoft.F(api_gateway.OperationNewParamsBodyMethodGet),
		}, {
			Endpoint: khulnasoft.F("/api/v1/users/{var1}"),
			Host:     khulnasoft.F("www.example.com"),
			Method:   khulnasoft.F(api_gateway.OperationNewParamsBodyMethodGet),
		}, {
			Endpoint: khulnasoft.F("/api/v1/users/{var1}"),
			Host:     khulnasoft.F("www.example.com"),
			Method:   khulnasoft.F(api_gateway.OperationNewParamsBodyMethodGet),
		}},
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationListWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.Operations.List(context.TODO(), api_gateway.OperationListParams{
		ZoneID:    khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: khulnasoft.F(api_gateway.OperationListParamsDirectionAsc),
		Endpoint:  khulnasoft.F("/api/v1"),
		Feature:   khulnasoft.F([]api_gateway.OperationListParamsFeature{api_gateway.OperationListParamsFeatureThresholds}),
		Host:      khulnasoft.F([]string{"api.khulnasoft.com"}),
		Method:    khulnasoft.F([]string{"GET"}),
		Order:     khulnasoft.F(api_gateway.OperationListParamsOrderMethod),
		Page:      khulnasoft.F(int64(1)),
		PerPage:   khulnasoft.F(int64(5)),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationDelete(t *testing.T) {
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
	_, err := client.APIGateway.Operations.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		api_gateway.OperationDeleteParams{
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

func TestOperationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.Operations.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		api_gateway.OperationGetParams{
			ZoneID:  khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Feature: khulnasoft.F([]api_gateway.OperationGetParamsFeature{api_gateway.OperationGetParamsFeatureThresholds}),
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
