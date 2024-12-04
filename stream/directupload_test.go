// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/stream"
)

func TestDirectUploadNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Stream.DirectUpload.New(context.TODO(), stream.DirectUploadNewParams{
		AccountID:          khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		MaxDurationSeconds: khulnasoft.F(int64(1)),
		AllowedOrigins:     khulnasoft.F([]stream.AllowedOriginsParam{"example.com"}),
		Creator:            khulnasoft.F("creator-id_abcde12345"),
		Expiry:             khulnasoft.F(time.Now()),
		Meta: khulnasoft.F[any](map[string]interface{}{
			"name": "video12345.mp4",
		}),
		RequireSignedURLs:     khulnasoft.F(true),
		ScheduledDeletion:     khulnasoft.F(time.Now()),
		ThumbnailTimestampPct: khulnasoft.F(0.529241),
		Watermark: khulnasoft.F(stream.DirectUploadNewParamsWatermark{
			UID: khulnasoft.F("ea95132c15732412d22c1476fa83f27a"),
		}),
		UploadCreator: khulnasoft.F("creator-id_abcde12345"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
