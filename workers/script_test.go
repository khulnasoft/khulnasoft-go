// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/workers"
)

func TestScriptUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Scripts.Update(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptUpdateParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: workers.ScriptUpdateParamsBodyObject{
				AnyPartName: khulnasoft.F([]io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents"))), io.Reader(bytes.NewBuffer([]byte("some file contents"))), io.Reader(bytes.NewBuffer([]byte("some file contents")))}),
				Metadata: khulnasoft.F(workers.ScriptUpdateParamsBodyObjectMetadata{
					Bindings: khulnasoft.F([]workers.ScriptUpdateParamsBodyObjectMetadataBinding{{
						Name: khulnasoft.F("MY_ENV_VAR"),
						Type: khulnasoft.F("plain_text"),
					}}),
					BodyPart:           khulnasoft.F("worker.js"),
					CompatibilityDate:  khulnasoft.F("2023-07-25"),
					CompatibilityFlags: khulnasoft.F([]string{"string", "string", "string"}),
					KeepBindings:       khulnasoft.F([]string{"string", "string", "string"}),
					Logpush:            khulnasoft.F(false),
					MainModule:         khulnasoft.F("worker.js"),
					Migrations: khulnasoft.F[workers.ScriptUpdateParamsBodyObjectMetadataMigrationsUnion](workers.SingleStepMigrationParam{
						DeletedClasses:   khulnasoft.F([]string{"string", "string", "string"}),
						NewClasses:       khulnasoft.F([]string{"string", "string", "string"}),
						NewSqliteClasses: khulnasoft.F([]string{"string", "string", "string"}),
						NewTag:           khulnasoft.F("v2"),
						OldTag:           khulnasoft.F("v1"),
						RenamedClasses: khulnasoft.F([]workers.SingleStepMigrationRenamedClassParam{{
							From: khulnasoft.F("from"),
							To:   khulnasoft.F("to"),
						}, {
							From: khulnasoft.F("from"),
							To:   khulnasoft.F("to"),
						}, {
							From: khulnasoft.F("from"),
							To:   khulnasoft.F("to"),
						}}),
						TransferredClasses: khulnasoft.F([]workers.SingleStepMigrationTransferredClassParam{{
							From:       khulnasoft.F("from"),
							FromScript: khulnasoft.F("from_script"),
							To:         khulnasoft.F("to"),
						}, {
							From:       khulnasoft.F("from"),
							FromScript: khulnasoft.F("from_script"),
							To:         khulnasoft.F("to"),
						}, {
							From:       khulnasoft.F("from"),
							FromScript: khulnasoft.F("from_script"),
							To:         khulnasoft.F("to"),
						}}),
					}),
					Placement: khulnasoft.F(workers.PlacementConfigurationParam{
						Mode: khulnasoft.F(workers.PlacementConfigurationModeSmart),
					}),
					Tags: khulnasoft.F([]string{"string", "string", "string"}),
					TailConsumers: khulnasoft.F([]workers.ConsumerScriptParam{{
						Service:     khulnasoft.F("my-log-consumer"),
						Environment: khulnasoft.F("production"),
						Namespace:   khulnasoft.F("my-namespace"),
					}, {
						Service:     khulnasoft.F("my-log-consumer"),
						Environment: khulnasoft.F("production"),
						Namespace:   khulnasoft.F("my-namespace"),
					}, {
						Service:     khulnasoft.F("my-log-consumer"),
						Environment: khulnasoft.F("production"),
						Namespace:   khulnasoft.F("my-namespace"),
					}}),
					UsageModel: khulnasoft.F(workers.ScriptUpdateParamsBodyObjectMetadataUsageModelBundled),
					VersionTags: khulnasoft.F(map[string]string{
						"foo": "string",
					}),
				}),
			},
			RollbackTo: khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
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

func TestScriptList(t *testing.T) {
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
	_, err := client.Workers.Scripts.List(context.TODO(), workers.ScriptListParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScriptDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Workers.Scripts.Delete(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptDeleteParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Force:     khulnasoft.F(true),
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

func TestScriptGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := khulnasoft.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	resp, err := client.Workers.Scripts.Get(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptGetParams{
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
