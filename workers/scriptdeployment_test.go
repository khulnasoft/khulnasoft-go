// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/workers"
)

func TestScriptDeploymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Scripts.Deployments.New(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptDeploymentNewParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Strategy:  khulnasoft.F(workers.ScriptDeploymentNewParamsStrategyPercentage),
			Versions: khulnasoft.F([]workers.ScriptDeploymentNewParamsVersion{{
				Percentage: khulnasoft.F(100.000000),
				VersionID:  khulnasoft.F("bcf48806-b317-4351-9ee7-36e7d557d4de"),
			}, {
				Percentage: khulnasoft.F(100.000000),
				VersionID:  khulnasoft.F("bcf48806-b317-4351-9ee7-36e7d557d4de"),
			}, {
				Percentage: khulnasoft.F(100.000000),
				VersionID:  khulnasoft.F("bcf48806-b317-4351-9ee7-36e7d557d4de"),
			}}),
			Force: khulnasoft.F(true),
			Annotations: khulnasoft.F(workers.DeploymentParam{
				WorkersMessage: khulnasoft.F("Deploy bug fix."),
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

func TestScriptDeploymentGet(t *testing.T) {
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
	_, err := client.Workers.Scripts.Deployments.Get(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptDeploymentGetParams{
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
