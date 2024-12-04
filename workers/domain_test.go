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

func TestDomainUpdate(t *testing.T) {
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
	_, err := client.Workers.Domains.Update(context.TODO(), workers.DomainUpdateParams{
		AccountID:   khulnasoft.F("9a7806061c88ada191ed06f989cc3dac"),
		Environment: khulnasoft.F("production"),
		Hostname:    khulnasoft.F("foo.example.com"),
		Service:     khulnasoft.F("foo"),
		ZoneID:      khulnasoft.F("593c9c94de529bbbfaac7c53ced0447d"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDomainListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Domains.List(context.TODO(), workers.DomainListParams{
		AccountID:   khulnasoft.F("9a7806061c88ada191ed06f989cc3dac"),
		Environment: khulnasoft.F("production"),
		Hostname:    khulnasoft.F("foo.example.com"),
		Service:     khulnasoft.F("foo"),
		ZoneID:      khulnasoft.F("593c9c94de529bbbfaac7c53ced0447d"),
		ZoneName:    khulnasoft.F("example.com"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDomainDelete(t *testing.T) {
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
	err := client.Workers.Domains.Delete(
		context.TODO(),
		"dbe10b4bc17c295377eabd600e1787fd",
		workers.DomainDeleteParams{
			AccountID: khulnasoft.F("9a7806061c88ada191ed06f989cc3dac"),
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

func TestDomainGet(t *testing.T) {
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
	_, err := client.Workers.Domains.Get(
		context.TODO(),
		"dbe10b4bc17c295377eabd600e1787fd",
		workers.DomainGetParams{
			AccountID: khulnasoft.F("9a7806061c88ada191ed06f989cc3dac"),
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
