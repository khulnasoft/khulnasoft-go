// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/radar"
)

func TestAttackLayer7SummaryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.Get(context.TODO(), radar.AttackLayer7SummaryGetParams{
		ASN:       khulnasoft.F([]string{"string", "string", "string"}),
		Continent: khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    khulnasoft.F(radar.AttackLayer7SummaryGetParamsFormatJson),
		Location:  khulnasoft.F([]string{"string", "string", "string"}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7SummaryHTTPMethodWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.HTTPMethod(context.TODO(), radar.AttackLayer7SummaryHTTPMethodParams{
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7SummaryHTTPMethodParamsFormatJson),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7SummaryHTTPMethodParamsHTTPVersion{radar.AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv1, radar.AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv2, radar.AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7SummaryHTTPMethodParamsIPVersion{radar.AttackLayer7SummaryHTTPMethodParamsIPVersionIPv4, radar.AttackLayer7SummaryHTTPMethodParamsIPVersionIPv6}),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7SummaryHTTPMethodParamsMitigationProduct{radar.AttackLayer7SummaryHTTPMethodParamsMitigationProductDDoS, radar.AttackLayer7SummaryHTTPMethodParamsMitigationProductWAF, radar.AttackLayer7SummaryHTTPMethodParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7SummaryHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.HTTPVersion(context.TODO(), radar.AttackLayer7SummaryHTTPVersionParams{
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7SummaryHTTPVersionParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7SummaryHTTPVersionParamsHTTPMethod{radar.AttackLayer7SummaryHTTPVersionParamsHTTPMethodGet, radar.AttackLayer7SummaryHTTPVersionParamsHTTPMethodPost, radar.AttackLayer7SummaryHTTPVersionParamsHTTPMethodDelete}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7SummaryHTTPVersionParamsIPVersion{radar.AttackLayer7SummaryHTTPVersionParamsIPVersionIPv4, radar.AttackLayer7SummaryHTTPVersionParamsIPVersionIPv6}),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7SummaryHTTPVersionParamsMitigationProduct{radar.AttackLayer7SummaryHTTPVersionParamsMitigationProductDDoS, radar.AttackLayer7SummaryHTTPVersionParamsMitigationProductWAF, radar.AttackLayer7SummaryHTTPVersionParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7SummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.IPVersion(context.TODO(), radar.AttackLayer7SummaryIPVersionParams{
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7SummaryIPVersionParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7SummaryIPVersionParamsHTTPMethod{radar.AttackLayer7SummaryIPVersionParamsHTTPMethodGet, radar.AttackLayer7SummaryIPVersionParamsHTTPMethodPost, radar.AttackLayer7SummaryIPVersionParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7SummaryIPVersionParamsHTTPVersion{radar.AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv1, radar.AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv2, radar.AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv3}),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7SummaryIPVersionParamsMitigationProduct{radar.AttackLayer7SummaryIPVersionParamsMitigationProductDDoS, radar.AttackLayer7SummaryIPVersionParamsMitigationProductWAF, radar.AttackLayer7SummaryIPVersionParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7SummaryManagedRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.ManagedRules(context.TODO(), radar.AttackLayer7SummaryManagedRulesParams{
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7SummaryManagedRulesParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7SummaryManagedRulesParamsHTTPMethod{radar.AttackLayer7SummaryManagedRulesParamsHTTPMethodGet, radar.AttackLayer7SummaryManagedRulesParamsHTTPMethodPost, radar.AttackLayer7SummaryManagedRulesParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7SummaryManagedRulesParamsHTTPVersion{radar.AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv1, radar.AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv2, radar.AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7SummaryManagedRulesParamsIPVersion{radar.AttackLayer7SummaryManagedRulesParamsIPVersionIPv4, radar.AttackLayer7SummaryManagedRulesParamsIPVersionIPv6}),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7SummaryManagedRulesParamsMitigationProduct{radar.AttackLayer7SummaryManagedRulesParamsMitigationProductDDoS, radar.AttackLayer7SummaryManagedRulesParamsMitigationProductWAF, radar.AttackLayer7SummaryManagedRulesParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7SummaryMitigationProductWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.MitigationProduct(context.TODO(), radar.AttackLayer7SummaryMitigationProductParams{
		ASN:         khulnasoft.F([]string{"string", "string", "string"}),
		Continent:   khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      khulnasoft.F(radar.AttackLayer7SummaryMitigationProductParamsFormatJson),
		HTTPMethod:  khulnasoft.F([]radar.AttackLayer7SummaryMitigationProductParamsHTTPMethod{radar.AttackLayer7SummaryMitigationProductParamsHTTPMethodGet, radar.AttackLayer7SummaryMitigationProductParamsHTTPMethodPost, radar.AttackLayer7SummaryMitigationProductParamsHTTPMethodDelete}),
		HTTPVersion: khulnasoft.F([]radar.AttackLayer7SummaryMitigationProductParamsHTTPVersion{radar.AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv1, radar.AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv2, radar.AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv3}),
		IPVersion:   khulnasoft.F([]radar.AttackLayer7SummaryMitigationProductParamsIPVersion{radar.AttackLayer7SummaryMitigationProductParamsIPVersionIPv4, radar.AttackLayer7SummaryMitigationProductParamsIPVersionIPv6}),
		Location:    khulnasoft.F([]string{"string", "string", "string"}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
