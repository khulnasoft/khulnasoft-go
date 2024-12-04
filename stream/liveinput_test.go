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

func TestLiveInputNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Stream.LiveInputs.New(context.TODO(), stream.LiveInputNewParams{
		AccountID:                khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		DefaultCreator:           khulnasoft.F("defaultCreator"),
		DeleteRecordingAfterDays: khulnasoft.F(45.000000),
		Meta: khulnasoft.F[any](map[string]interface{}{
			"name": "test stream 1",
		}),
		Recording: khulnasoft.F(stream.LiveInputNewParamsRecording{
			AllowedOrigins:      khulnasoft.F([]string{"example.com"}),
			HideLiveViewerCount: khulnasoft.F(false),
			Mode:                khulnasoft.F(stream.LiveInputNewParamsRecordingModeOff),
			RequireSignedURLs:   khulnasoft.F(false),
			TimeoutSeconds:      khulnasoft.F(int64(0)),
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

func TestLiveInputUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Stream.LiveInputs.Update(
		context.TODO(),
		"66be4bf738797e01e1fca35a7bdecdcd",
		stream.LiveInputUpdateParams{
			AccountID:                khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			DefaultCreator:           khulnasoft.F("defaultCreator"),
			DeleteRecordingAfterDays: khulnasoft.F(45.000000),
			Meta: khulnasoft.F[any](map[string]interface{}{
				"name": "test stream 1",
			}),
			Recording: khulnasoft.F(stream.LiveInputUpdateParamsRecording{
				AllowedOrigins:      khulnasoft.F([]string{"example.com"}),
				HideLiveViewerCount: khulnasoft.F(false),
				Mode:                khulnasoft.F(stream.LiveInputUpdateParamsRecordingModeOff),
				RequireSignedURLs:   khulnasoft.F(false),
				TimeoutSeconds:      khulnasoft.F(int64(0)),
			}),
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

func TestLiveInputListWithOptionalParams(t *testing.T) {
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
	_, err := client.Stream.LiveInputs.List(context.TODO(), stream.LiveInputListParams{
		AccountID:     khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		IncludeCounts: khulnasoft.F(true),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLiveInputDelete(t *testing.T) {
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
	err := client.Stream.LiveInputs.Delete(
		context.TODO(),
		"66be4bf738797e01e1fca35a7bdecdcd",
		stream.LiveInputDeleteParams{
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

func TestLiveInputGet(t *testing.T) {
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
	_, err := client.Stream.LiveInputs.Get(
		context.TODO(),
		"66be4bf738797e01e1fca35a7bdecdcd",
		stream.LiveInputGetParams{
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
