// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rules_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/rules"
)

func TestListItemNew(t *testing.T) {
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
	_, err := client.Rules.Lists.Items.New(
		context.TODO(),
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		rules.ListItemNewParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: []rules.ListItemNewParamsBody{{
				ASN:     khulnasoft.F(int64(5567)),
				Comment: khulnasoft.F("Private IP address"),
				Hostname: khulnasoft.F(rules.HostnameParam{
					URLHostname: khulnasoft.F("example.com"),
				}),
				IP: khulnasoft.F("10.0.0.1"),
				Redirect: khulnasoft.F(rules.RedirectParam{
					SourceURL:           khulnasoft.F("example.com/arch"),
					TargetURL:           khulnasoft.F("https://archlinux.org/"),
					IncludeSubdomains:   khulnasoft.F(true),
					PreservePathSuffix:  khulnasoft.F(true),
					PreserveQueryString: khulnasoft.F(true),
					StatusCode:          khulnasoft.F(rules.RedirectStatusCode301),
					SubpathMatching:     khulnasoft.F(true),
				}),
			}, {
				ASN:     khulnasoft.F(int64(5567)),
				Comment: khulnasoft.F("Private IP address"),
				Hostname: khulnasoft.F(rules.HostnameParam{
					URLHostname: khulnasoft.F("example.com"),
				}),
				IP: khulnasoft.F("10.0.0.1"),
				Redirect: khulnasoft.F(rules.RedirectParam{
					SourceURL:           khulnasoft.F("example.com/arch"),
					TargetURL:           khulnasoft.F("https://archlinux.org/"),
					IncludeSubdomains:   khulnasoft.F(true),
					PreservePathSuffix:  khulnasoft.F(true),
					PreserveQueryString: khulnasoft.F(true),
					StatusCode:          khulnasoft.F(rules.RedirectStatusCode301),
					SubpathMatching:     khulnasoft.F(true),
				}),
			}, {
				ASN:     khulnasoft.F(int64(5567)),
				Comment: khulnasoft.F("Private IP address"),
				Hostname: khulnasoft.F(rules.HostnameParam{
					URLHostname: khulnasoft.F("example.com"),
				}),
				IP: khulnasoft.F("10.0.0.1"),
				Redirect: khulnasoft.F(rules.RedirectParam{
					SourceURL:           khulnasoft.F("example.com/arch"),
					TargetURL:           khulnasoft.F("https://archlinux.org/"),
					IncludeSubdomains:   khulnasoft.F(true),
					PreservePathSuffix:  khulnasoft.F(true),
					PreserveQueryString: khulnasoft.F(true),
					StatusCode:          khulnasoft.F(rules.RedirectStatusCode301),
					SubpathMatching:     khulnasoft.F(true),
				}),
			}},
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

func TestListItemUpdate(t *testing.T) {
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
	_, err := client.Rules.Lists.Items.Update(
		context.TODO(),
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		rules.ListItemUpdateParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: []rules.ListItemUpdateParamsBody{{
				ASN:     khulnasoft.F(int64(5567)),
				Comment: khulnasoft.F("Private IP address"),
				Hostname: khulnasoft.F(rules.HostnameParam{
					URLHostname: khulnasoft.F("example.com"),
				}),
				IP: khulnasoft.F("10.0.0.1"),
				Redirect: khulnasoft.F(rules.RedirectParam{
					SourceURL:           khulnasoft.F("example.com/arch"),
					TargetURL:           khulnasoft.F("https://archlinux.org/"),
					IncludeSubdomains:   khulnasoft.F(true),
					PreservePathSuffix:  khulnasoft.F(true),
					PreserveQueryString: khulnasoft.F(true),
					StatusCode:          khulnasoft.F(rules.RedirectStatusCode301),
					SubpathMatching:     khulnasoft.F(true),
				}),
			}, {
				ASN:     khulnasoft.F(int64(5567)),
				Comment: khulnasoft.F("Private IP address"),
				Hostname: khulnasoft.F(rules.HostnameParam{
					URLHostname: khulnasoft.F("example.com"),
				}),
				IP: khulnasoft.F("10.0.0.1"),
				Redirect: khulnasoft.F(rules.RedirectParam{
					SourceURL:           khulnasoft.F("example.com/arch"),
					TargetURL:           khulnasoft.F("https://archlinux.org/"),
					IncludeSubdomains:   khulnasoft.F(true),
					PreservePathSuffix:  khulnasoft.F(true),
					PreserveQueryString: khulnasoft.F(true),
					StatusCode:          khulnasoft.F(rules.RedirectStatusCode301),
					SubpathMatching:     khulnasoft.F(true),
				}),
			}, {
				ASN:     khulnasoft.F(int64(5567)),
				Comment: khulnasoft.F("Private IP address"),
				Hostname: khulnasoft.F(rules.HostnameParam{
					URLHostname: khulnasoft.F("example.com"),
				}),
				IP: khulnasoft.F("10.0.0.1"),
				Redirect: khulnasoft.F(rules.RedirectParam{
					SourceURL:           khulnasoft.F("example.com/arch"),
					TargetURL:           khulnasoft.F("https://archlinux.org/"),
					IncludeSubdomains:   khulnasoft.F(true),
					PreservePathSuffix:  khulnasoft.F(true),
					PreserveQueryString: khulnasoft.F(true),
					StatusCode:          khulnasoft.F(rules.RedirectStatusCode301),
					SubpathMatching:     khulnasoft.F(true),
				}),
			}},
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

func TestListItemListWithOptionalParams(t *testing.T) {
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
	_, err := client.Rules.Lists.Items.List(
		context.TODO(),
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		rules.ListItemListParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Cursor:    khulnasoft.F("zzz"),
			PerPage:   khulnasoft.F(int64(1)),
			Search:    khulnasoft.F("1.1.1."),
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

func TestListItemDelete(t *testing.T) {
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
	_, err := client.Rules.Lists.Items.Delete(
		context.TODO(),
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		rules.ListItemDeleteParams{
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

func TestListItemGet(t *testing.T) {
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
	_, err := client.Rules.Lists.Items.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"2c0fc9fa937b11eaa1b71c4d701ab86e",
		"34b12448945f11eaa1b71c4d701ab86e",
	)
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
