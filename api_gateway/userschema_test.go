// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/api_gateway"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestUserSchemaNewWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.UserSchemas.New(context.TODO(), api_gateway.UserSchemaNewParams{
		ZoneID:            khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		File:              khulnasoft.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		Kind:              khulnasoft.F(api_gateway.UserSchemaNewParamsKindOpenAPIV3),
		Name:              khulnasoft.F("petstore schema"),
		ValidationEnabled: khulnasoft.F(api_gateway.UserSchemaNewParamsValidationEnabledTrue),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserSchemaListWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.UserSchemas.List(context.TODO(), api_gateway.UserSchemaListParams{
		ZoneID:            khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		OmitSource:        khulnasoft.F(true),
		Page:              khulnasoft.F(int64(1)),
		PerPage:           khulnasoft.F(int64(5)),
		ValidationEnabled: khulnasoft.F(true),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserSchemaDelete(t *testing.T) {
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
	_, err := client.APIGateway.UserSchemas.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		api_gateway.UserSchemaDeleteParams{
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

func TestUserSchemaEditWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.UserSchemas.Edit(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		api_gateway.UserSchemaEditParams{
			ZoneID:            khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ValidationEnabled: khulnasoft.F(api_gateway.UserSchemaEditParamsValidationEnabledTrue),
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

func TestUserSchemaGetWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.UserSchemas.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		api_gateway.UserSchemaGetParams{
			ZoneID:     khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			OmitSource: khulnasoft.F(true),
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
