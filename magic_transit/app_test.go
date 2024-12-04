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

func TestAppNewWithOptionalParams(t *testing.T) {
	t.Skip("prism errors - https://github.com/khulnasoft/khulnasoft-python/actions/runs/9360388260/job/25765690361?pr=482#step:5:7212")
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
	_, err := client.MagicTransit.Apps.New(context.TODO(), magic_transit.AppNewParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Name:      khulnasoft.F("Khulnasoft Dashboard"),
		Type:      khulnasoft.F("Development"),
		Hostnames: khulnasoft.F([]string{"auth.khulnasoft.com", "auth.khulnasoft.com", "auth.khulnasoft.com"}),
		IPSubnets: khulnasoft.F([]string{"1.1.1.1/32", "1.1.1.1/32", "1.1.1.1/32"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppUpdateWithOptionalParams(t *testing.T) {
	t.Skip("prism errors - https://github.com/khulnasoft/khulnasoft-python/actions/runs/9360388260/job/25765690361?pr=482#step:5:7212")
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
	_, err := client.MagicTransit.Apps.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.AppUpdateParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Hostnames: khulnasoft.F([]string{"auth.khulnasoft.com", "auth.khulnasoft.com", "auth.khulnasoft.com"}),
			IPSubnets: khulnasoft.F([]string{"1.1.1.1/32", "1.1.1.1/32", "1.1.1.1/32"}),
			Name:      khulnasoft.F("Khulnasoft Dashboard"),
			Type:      khulnasoft.F("Development"),
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

func TestAppList(t *testing.T) {
	t.Skip("prism errors - https://github.com/khulnasoft/khulnasoft-python/actions/runs/9360388260/job/25765690361?pr=482#step:5:7212")
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
	_, err := client.MagicTransit.Apps.List(context.TODO(), magic_transit.AppListParams{
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

func TestAppDelete(t *testing.T) {
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
	_, err := client.MagicTransit.Apps.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.AppDeleteParams{
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
