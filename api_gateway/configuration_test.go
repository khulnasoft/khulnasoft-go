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

func TestConfigurationUpdate(t *testing.T) {
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
	_, err := client.APIGateway.Configurations.Update(context.TODO(), api_gateway.ConfigurationUpdateParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Configuration: api_gateway.ConfigurationParam{
			AuthIDCharacteristics: khulnasoft.F([]api_gateway.ConfigurationAuthIDCharacteristicsUnionParam{api_gateway.ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicParam{
				Name: khulnasoft.F("authorization"),
				Type: khulnasoft.F(api_gateway.ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicTypeHeader),
			}}),
		},
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigurationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.Configurations.Get(context.TODO(), api_gateway.ConfigurationGetParams{
		ZoneID:     khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Properties: khulnasoft.F([]api_gateway.ConfigurationGetParamsProperty{api_gateway.ConfigurationGetParamsPropertyAuthIDCharacteristics}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
