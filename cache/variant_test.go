// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cache_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/cache"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestVariantDelete(t *testing.T) {
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
	_, err := client.Cache.Variants.Delete(context.TODO(), cache.VariantDeleteParams{
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

func TestVariantEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Cache.Variants.Edit(context.TODO(), cache.VariantEditParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Value: khulnasoft.F(cache.VariantEditParamsValue{
			AVIF: khulnasoft.F([]string{"image/webp", "image/jpeg"}),
			BMP:  khulnasoft.F([]string{"image/webp", "image/jpeg"}),
			GIF:  khulnasoft.F([]string{"image/webp", "image/jpeg"}),
			JP2:  khulnasoft.F([]string{"image/webp", "image/avif"}),
			JPEG: khulnasoft.F([]string{"image/webp", "image/avif"}),
			JPG:  khulnasoft.F([]string{"image/webp", "image/avif"}),
			JPG2: khulnasoft.F([]string{"image/webp", "image/avif"}),
			PNG:  khulnasoft.F([]string{"image/webp", "image/avif"}),
			TIF:  khulnasoft.F([]string{"image/webp", "image/avif"}),
			TIFF: khulnasoft.F([]string{"image/webp", "image/avif"}),
			WebP: khulnasoft.F([]string{"image/jpeg", "image/avif"}),
		}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVariantGet(t *testing.T) {
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
	_, err := client.Cache.Variants.Get(context.TODO(), cache.VariantGetParams{
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
