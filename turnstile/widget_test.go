// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/turnstile"
)

func TestWidgetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Turnstile.Widgets.New(context.TODO(), turnstile.WidgetNewParams{
		AccountID:      khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Domains:        khulnasoft.F([]turnstile.WidgetDomainParam{"203.0.113.1", "khulnasoft.com", "blog.example.com"}),
		Mode:           khulnasoft.F(turnstile.WidgetNewParamsModeNonInteractive),
		Name:           khulnasoft.F("blog.khulnasoft.com login form"),
		Direction:      khulnasoft.F(turnstile.WidgetNewParamsDirectionAsc),
		Order:          khulnasoft.F(turnstile.WidgetNewParamsOrderID),
		Page:           khulnasoft.F(1.000000),
		PerPage:        khulnasoft.F(5.000000),
		BotFightMode:   khulnasoft.F(false),
		ClearanceLevel: khulnasoft.F(turnstile.WidgetNewParamsClearanceLevelNoClearance),
		Offlabel:       khulnasoft.F(false),
		Region:         khulnasoft.F(turnstile.WidgetNewParamsRegionWorld),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWidgetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Turnstile.Widgets.Update(
		context.TODO(),
		"0x4AAF00AAAABn0R22HWm-YUc",
		turnstile.WidgetUpdateParams{
			AccountID:      khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Domains:        khulnasoft.F([]turnstile.WidgetDomainParam{"203.0.113.1", "khulnasoft.com", "blog.example.com"}),
			Mode:           khulnasoft.F(turnstile.WidgetUpdateParamsModeNonInteractive),
			Name:           khulnasoft.F("blog.khulnasoft.com login form"),
			BotFightMode:   khulnasoft.F(false),
			ClearanceLevel: khulnasoft.F(turnstile.WidgetUpdateParamsClearanceLevelNoClearance),
			Offlabel:       khulnasoft.F(false),
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

func TestWidgetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Turnstile.Widgets.List(context.TODO(), turnstile.WidgetListParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: khulnasoft.F(turnstile.WidgetListParamsDirectionAsc),
		Order:     khulnasoft.F(turnstile.WidgetListParamsOrderID),
		Page:      khulnasoft.F(1.000000),
		PerPage:   khulnasoft.F(5.000000),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWidgetDelete(t *testing.T) {
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
	_, err := client.Turnstile.Widgets.Delete(
		context.TODO(),
		"0x4AAF00AAAABn0R22HWm-YUc",
		turnstile.WidgetDeleteParams{
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

func TestWidgetGet(t *testing.T) {
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
	_, err := client.Turnstile.Widgets.Get(
		context.TODO(),
		"0x4AAF00AAAABn0R22HWm-YUc",
		turnstile.WidgetGetParams{
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

func TestWidgetRotateSecretWithOptionalParams(t *testing.T) {
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
	_, err := client.Turnstile.Widgets.RotateSecret(
		context.TODO(),
		"0x4AAF00AAAABn0R22HWm-YUc",
		turnstile.WidgetRotateSecretParams{
			AccountID:             khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			InvalidateImmediately: khulnasoft.F(true),
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
