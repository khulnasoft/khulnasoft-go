// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_tls_client_auth_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/origin_tls_client_auth"
)

func TestHostnameUpdate(t *testing.T) {
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
	_, err := client.OriginTLSClientAuth.Hostnames.Update(context.TODO(), origin_tls_client_auth.HostnameUpdateParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Config: khulnasoft.F([]origin_tls_client_auth.HostnameUpdateParamsConfig{{
			CERTID:   khulnasoft.F("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60"),
			Enabled:  khulnasoft.F(true),
			Hostname: khulnasoft.F("app.example.com"),
		}, {
			CERTID:   khulnasoft.F("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60"),
			Enabled:  khulnasoft.F(true),
			Hostname: khulnasoft.F("app.example.com"),
		}, {
			CERTID:   khulnasoft.F("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60"),
			Enabled:  khulnasoft.F(true),
			Hostname: khulnasoft.F("app.example.com"),
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

func TestHostnameGet(t *testing.T) {
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
	_, err := client.OriginTLSClientAuth.Hostnames.Get(
		context.TODO(),
		"app.example.com",
		origin_tls_client_auth.HostnameGetParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
