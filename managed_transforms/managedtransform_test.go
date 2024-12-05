// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/managed_transforms"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestManagedTransformList(t *testing.T) {
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
	_, err := client.ManagedTransforms.List(context.TODO(), managed_transforms.ManagedTransformListParams{
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

func TestManagedTransformEdit(t *testing.T) {
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
	_, err := client.ManagedTransforms.Edit(context.TODO(), managed_transforms.ManagedTransformEditParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ManagedRequestHeaders: khulnasoft.F([]managed_transforms.RequestModelParam{{
			ID:      khulnasoft.F("add_cf-bot-score_header"),
			Enabled: khulnasoft.F(true),
		}, {
			ID:      khulnasoft.F("add_cf-bot-score_header"),
			Enabled: khulnasoft.F(true),
		}, {
			ID:      khulnasoft.F("add_cf-bot-score_header"),
			Enabled: khulnasoft.F(true),
		}}),
		ManagedResponseHeaders: khulnasoft.F([]managed_transforms.RequestModelParam{{
			ID:      khulnasoft.F("add_cf-bot-score_header"),
			Enabled: khulnasoft.F(true),
		}, {
			ID:      khulnasoft.F("add_cf-bot-score_header"),
			Enabled: khulnasoft.F(true),
		}, {
			ID:      khulnasoft.F("add_cf-bot-score_header"),
			Enabled: khulnasoft.F(true),
		}}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
