// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package diagnostics_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/diagnostics"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestTracerouteNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Diagnostics.Traceroutes.New(context.TODO(), diagnostics.TracerouteNewParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Targets:   khulnasoft.F([]string{"203.0.113.1", "khulnasoft.com"}),
		Colos:     khulnasoft.F([]string{"den", "sin"}),
		Options: khulnasoft.F(diagnostics.TracerouteNewParamsOptions{
			MaxTTL:        khulnasoft.F(int64(15)),
			PacketType:    khulnasoft.F(diagnostics.TracerouteNewParamsOptionsPacketTypeIcmp),
			PacketsPerTTL: khulnasoft.F(int64(0)),
			Port:          khulnasoft.F(int64(0)),
			WaitTime:      khulnasoft.F(int64(1)),
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
