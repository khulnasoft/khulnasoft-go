// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/workers"
)

func TestScriptVersionNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Workers.Scripts.Versions.New(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptVersionNewParams{
			AccountID:   khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AnyPartName: khulnasoft.F([]io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents"))), io.Reader(bytes.NewBuffer([]byte("some file contents"))), io.Reader(bytes.NewBuffer([]byte("some file contents")))}),
			Metadata: khulnasoft.F(workers.ScriptVersionNewParamsMetadata{
				Annotations: khulnasoft.F(workers.ScriptVersionNewParamsMetadataAnnotations{
					WorkersMessage: khulnasoft.F("Fixed worker code."),
					WorkersTag:     khulnasoft.F("workers/tag"),
				}),
				Bindings: khulnasoft.F([]interface{}{map[string]interface{}{
					"name": "MY_ENV_VAR",
					"text": "my_data",
					"type": "plain_text",
				}}),
				CompatibilityDate:  khulnasoft.F("2023-07-25"),
				CompatibilityFlags: khulnasoft.F([]string{"string", "string", "string"}),
				KeepBindings:       khulnasoft.F([]string{"string", "string", "string"}),
				MainModule:         khulnasoft.F("worker.js"),
				UsageModel:         khulnasoft.F(workers.ScriptVersionNewParamsMetadataUsageModelStandard),
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

func TestScriptVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Scripts.Versions.List(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptVersionListParams{
			AccountID:  khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Deployable: khulnasoft.F(true),
			Page:       khulnasoft.F(int64(0)),
			PerPage:    khulnasoft.F(int64(0)),
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

func TestScriptVersionGet(t *testing.T) {
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
	_, err := client.Workers.Scripts.Versions.Get(
		context.TODO(),
		"this-is_my_script-01",
		"bcf48806-b317-4351-9ee7-36e7d557d4de",
		workers.ScriptVersionGetParams{
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
