// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package khulnasoft_test

import (
	"context"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/zones"
)

func TestUsage(t *testing.T) {
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
	zone, err := client.Zones.New(context.TODO(), zones.ZoneNewParams{
		Account: khulnasoft.F(zones.ZoneNewParamsAccount{
			ID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		}),
		Name: khulnasoft.F("example.com"),
		Type: khulnasoft.F(zones.TypeFull),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", zone.ID)
}
