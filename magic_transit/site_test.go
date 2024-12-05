// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/magic_transit"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestSiteNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.New(context.TODO(), magic_transit.SiteNewParams{
		AccountID:   khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Name:        khulnasoft.F("site_1"),
		ConnectorID: khulnasoft.F("ac60d3d0435248289d446cedd870bcf4"),
		Description: khulnasoft.F("description"),
		HaMode:      khulnasoft.F(true),
		Location: khulnasoft.F(magic_transit.SiteLocationParam{
			Lat: khulnasoft.F("37.6192"),
			Lon: khulnasoft.F("122.3816"),
		}),
		SecondaryConnectorID: khulnasoft.F("8d67040d3835dbcf46ce29da440dc482"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSiteUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteUpdateParams{
			AccountID:   khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ConnectorID: khulnasoft.F("ac60d3d0435248289d446cedd870bcf4"),
			Description: khulnasoft.F("description"),
			Location: khulnasoft.F(magic_transit.SiteLocationParam{
				Lat: khulnasoft.F("37.6192"),
				Lon: khulnasoft.F("122.3816"),
			}),
			Name:                 khulnasoft.F("site_1"),
			SecondaryConnectorID: khulnasoft.F("8d67040d3835dbcf46ce29da440dc482"),
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

func TestSiteListWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.List(context.TODO(), magic_transit.SiteListParams{
		AccountID:           khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ConnectorIdentifier: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSiteDelete(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteDeleteParams{
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

func TestSiteEditWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteEditParams{
			AccountID:   khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ConnectorID: khulnasoft.F("ac60d3d0435248289d446cedd870bcf4"),
			Description: khulnasoft.F("description"),
			Location: khulnasoft.F(magic_transit.SiteLocationParam{
				Lat: khulnasoft.F("37.6192"),
				Lon: khulnasoft.F("122.3816"),
			}),
			Name:                 khulnasoft.F("site_1"),
			SecondaryConnectorID: khulnasoft.F("8d67040d3835dbcf46ce29da440dc482"),
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

func TestSiteGet(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteGetParams{
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
