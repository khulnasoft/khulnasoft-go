// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/workers"
	"github.com/khulnasoft/khulnasoft-go/workers_for_platforms"
)

func TestDispatchNamespaceScriptSettingEditWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Settings.Edit(
		context.TODO(),
		"my-dispatch-namespace",
		"this-is_my_script-01",
		workers_for_platforms.DispatchNamespaceScriptSettingEditParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Settings: khulnasoft.F(workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettings{
				Bindings: khulnasoft.F([]workers.BindingUnionParam{workers.KVNamespaceBindingParam{
					Type: khulnasoft.F(workers.KVNamespaceBindingTypeKVNamespace),
				}, workers.KVNamespaceBindingParam{
					Type: khulnasoft.F(workers.KVNamespaceBindingTypeKVNamespace),
				}, workers.KVNamespaceBindingParam{
					Type: khulnasoft.F(workers.KVNamespaceBindingTypeKVNamespace),
				}}),
				CompatibilityDate:  khulnasoft.F("2022-04-05"),
				CompatibilityFlags: khulnasoft.F([]string{"formdata_parser_supports_files", "formdata_parser_supports_files", "formdata_parser_supports_files"}),
				Limits: khulnasoft.F(workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsLimits{
					CPUMs: khulnasoft.F(int64(50)),
				}),
				Logpush: khulnasoft.F(false),
				Migrations: khulnasoft.F[workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion](workers.SingleStepMigrationParam{
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
				Tags: khulnasoft.F([]string{"my-tag", "my-tag", "my-tag"}),
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
				UsageModel: khulnasoft.F("unbound"),
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

func TestDispatchNamespaceScriptSettingGet(t *testing.T) {
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
	_, err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Settings.Get(
		context.TODO(),
		"my-dispatch-namespace",
		"this-is_my_script-01",
		workers_for_platforms.DispatchNamespaceScriptSettingGetParams{
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
