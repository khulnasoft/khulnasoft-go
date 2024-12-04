// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/stream"
)

func TestClipNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Stream.Clip.New(context.TODO(), stream.ClipNewParams{
		AccountID:             khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ClippedFromVideoUID:   khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		EndTimeSeconds:        khulnasoft.F(int64(0)),
		StartTimeSeconds:      khulnasoft.F(int64(0)),
		AllowedOrigins:        khulnasoft.F([]stream.AllowedOriginsParam{"example.com"}),
		Creator:               khulnasoft.F("creator-id_abcde12345"),
		MaxDurationSeconds:    khulnasoft.F(int64(1)),
		RequireSignedURLs:     khulnasoft.F(true),
		ThumbnailTimestampPct: khulnasoft.F(0.529241),
		Watermark: khulnasoft.F(stream.ClipNewParamsWatermark{
			UID: khulnasoft.F("ea95132c15732412d22c1476fa83f27a"),
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
