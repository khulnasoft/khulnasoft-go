// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/r2"
)

func TestTemporaryCredentialNewWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.TemporaryCredentials.New(context.TODO(), r2.TemporaryCredentialNewParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		TemporaryCredential: r2.TemporaryCredentialParam{
			Bucket:            khulnasoft.F("example-bucket"),
			ParentAccessKeyID: khulnasoft.F("example-access-key-id"),
			Permission:        khulnasoft.F(r2.TemporaryCredentialPermissionAdminReadWrite),
			TTLSeconds:        khulnasoft.F(3600.000000),
			Objects:           khulnasoft.F([]string{"example-object"}),
			Prefixes:          khulnasoft.F([]string{"example-prefix/"}),
		},
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
