// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/page_shield"
)

func TestCookieListWithOptionalParams(t *testing.T) {
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
	_, err := client.PageShield.Cookies.List(context.TODO(), page_shield.CookieListParams{
		ZoneID:    khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: khulnasoft.F(page_shield.CookieListParamsDirectionAsc),
		Domain:    khulnasoft.F("example.com"),
		Export:    khulnasoft.F(page_shield.CookieListParamsExportCsv),
		Hosts:     khulnasoft.F("blog.khulnasoft.com,www.example*,*khulnasoft.com"),
		HTTPOnly:  khulnasoft.F(true),
		Name:      khulnasoft.F("session_id"),
		OrderBy:   khulnasoft.F(page_shield.CookieListParamsOrderByFirstSeenAt),
		Page:      khulnasoft.F("2"),
		PageURL:   khulnasoft.F("example.com/page,*/checkout,example.com/*,*checkout*"),
		Path:      khulnasoft.F("/"),
		PerPage:   khulnasoft.F(100.000000),
		SameSite:  khulnasoft.F(page_shield.CookieListParamsSameSiteLax),
		Secure:    khulnasoft.F(true),
		Type:      khulnasoft.F(page_shield.CookieListParamsTypeFirstParty),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCookieGet(t *testing.T) {
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
	_, err := client.PageShield.Cookies.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		page_shield.CookieGetParams{
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
