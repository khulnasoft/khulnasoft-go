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

func TestUserSchemaOperationListWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.UserSchemas.Operations.List(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		api_gateway.UserSchemaOperationListParams{
			ZoneID:          khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Endpoint:        khulnasoft.F("/api/v1"),
			Feature:         khulnasoft.F([]api_gateway.UserSchemaOperationListParamsFeature{api_gateway.UserSchemaOperationListParamsFeatureThresholds}),
			Host:            khulnasoft.F([]string{"api.khulnasoft.com"}),
			Method:          khulnasoft.F([]string{"GET"}),
			OperationStatus: khulnasoft.F(api_gateway.UserSchemaOperationListParamsOperationStatusNew),
			Page:            khulnasoft.F(int64(1)),
			PerPage:         khulnasoft.F(int64(5)),
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
