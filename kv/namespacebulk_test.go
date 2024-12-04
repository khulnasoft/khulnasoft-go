// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/kv"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestNamespaceBulkUpdate(t *testing.T) {
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
	_, err := client.KV.Namespaces.Bulk.Update(
		context.TODO(),
		"0f2ac74b498b48028cb68387c421e279",
		kv.NamespaceBulkUpdateParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: []kv.NamespaceBulkUpdateParamsBody{{
				Base64:        khulnasoft.F(true),
				Expiration:    khulnasoft.F(1578435000.000000),
				ExpirationTTL: khulnasoft.F(300.000000),
				Key:           khulnasoft.F("My-Key"),
				Metadata: khulnasoft.F(map[string]interface{}{
					"someMetadataKey": "bar",
				}),
				Value: khulnasoft.F("Some string"),
			}, {
				Base64:        khulnasoft.F(true),
				Expiration:    khulnasoft.F(1578435000.000000),
				ExpirationTTL: khulnasoft.F(300.000000),
				Key:           khulnasoft.F("My-Key"),
				Metadata: khulnasoft.F(map[string]interface{}{
					"someMetadataKey": "bar",
				}),
				Value: khulnasoft.F("Some string"),
			}, {
				Base64:        khulnasoft.F(true),
				Expiration:    khulnasoft.F(1578435000.000000),
				ExpirationTTL: khulnasoft.F(300.000000),
				Key:           khulnasoft.F("My-Key"),
				Metadata: khulnasoft.F(map[string]interface{}{
					"someMetadataKey": "bar",
				}),
				Value: khulnasoft.F("Some string"),
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

func TestNamespaceBulkDelete(t *testing.T) {
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
	_, err := client.KV.Namespaces.Bulk.Delete(
		context.TODO(),
		"0f2ac74b498b48028cb68387c421e279",
		kv.NamespaceBulkDeleteParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
