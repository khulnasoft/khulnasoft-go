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

func TestAttackLayer7TimeseriesGroupGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Get(context.TODO(), radar.AttackLayer7TimeseriesGroupGetParams{
		AggInterval: khulnasoft.F(radar.AttackLayer7TimeseriesGroupGetParamsAggInterval15m),
		ASN:         khulnasoft.F([]string{"string", "string", "string"}),
		Continent:   khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      khulnasoft.F(radar.AttackLayer7TimeseriesGroupGetParamsFormatJson),
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

func TestAttackLayer7TimeseriesGroupHTTPMethodWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.HTTPMethod(context.TODO(), radar.AttackLayer7TimeseriesGroupHTTPMethodParams{
		AggInterval:       khulnasoft.F(radar.AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval15m),
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TimeseriesGroupHTTPMethodParamsFormatJson),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv1, radar.AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv2, radar.AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion{radar.AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv4, radar.AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv6}),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDDoS, radar.AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductWAF, radar.AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
		Normalization:     khulnasoft.F(radar.AttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.HTTPVersion(context.TODO(), radar.AttackLayer7TimeseriesGroupHTTPVersionParams{
		AggInterval:       khulnasoft.F(radar.AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval15m),
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TimeseriesGroupHTTPVersionParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodGet, radar.AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPost, radar.AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodDelete}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion{radar.AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv4, radar.AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv6}),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDDoS, radar.AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductWAF, radar.AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
		Normalization:     khulnasoft.F(radar.AttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Industry(context.TODO(), radar.AttackLayer7TimeseriesGroupIndustryParams{
		AggInterval:       khulnasoft.F(radar.AttackLayer7TimeseriesGroupIndustryParamsAggInterval15m),
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TimeseriesGroupIndustryParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodGet, radar.AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPost, radar.AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv1, radar.AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv2, radar.AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7TimeseriesGroupIndustryParamsIPVersion{radar.AttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv4, radar.AttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv6}),
		LimitPerGroup:     khulnasoft.F(int64(4)),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupIndustryParamsMitigationProductDDoS, radar.AttackLayer7TimeseriesGroupIndustryParamsMitigationProductWAF, radar.AttackLayer7TimeseriesGroupIndustryParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
		Normalization:     khulnasoft.F(radar.AttackLayer7TimeseriesGroupIndustryParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.IPVersion(context.TODO(), radar.AttackLayer7TimeseriesGroupIPVersionParams{
		AggInterval:       khulnasoft.F(radar.AttackLayer7TimeseriesGroupIPVersionParamsAggInterval15m),
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TimeseriesGroupIPVersionParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodGet, radar.AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPost, radar.AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv1, radar.AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv2, radar.AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv3}),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDDoS, radar.AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductWAF, radar.AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
		Normalization:     khulnasoft.F(radar.AttackLayer7TimeseriesGroupIPVersionParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupManagedRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.ManagedRules(context.TODO(), radar.AttackLayer7TimeseriesGroupManagedRulesParams{
		AggInterval:       khulnasoft.F(radar.AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval15m),
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TimeseriesGroupManagedRulesParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodGet, radar.AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPost, radar.AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv1, radar.AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv2, radar.AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7TimeseriesGroupManagedRulesParamsIPVersion{radar.AttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv4, radar.AttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv6}),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDDoS, radar.AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductWAF, radar.AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
		Normalization:     khulnasoft.F(radar.AttackLayer7TimeseriesGroupManagedRulesParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupMitigationProductWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.MitigationProduct(context.TODO(), radar.AttackLayer7TimeseriesGroupMitigationProductParams{
		AggInterval:   khulnasoft.F(radar.AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval15m),
		ASN:           khulnasoft.F([]string{"string", "string", "string"}),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:        khulnasoft.F(radar.AttackLayer7TimeseriesGroupMitigationProductParamsFormatJson),
		HTTPMethod:    khulnasoft.F([]radar.AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodGet, radar.AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPost, radar.AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodDelete}),
		HTTPVersion:   khulnasoft.F([]radar.AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv1, radar.AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv2, radar.AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv3}),
		IPVersion:     khulnasoft.F([]radar.AttackLayer7TimeseriesGroupMitigationProductParamsIPVersion{radar.AttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv4, radar.AttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv6}),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		Normalization: khulnasoft.F(radar.AttackLayer7TimeseriesGroupMitigationProductParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TimeseriesGroupVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Vertical(context.TODO(), radar.AttackLayer7TimeseriesGroupVerticalParams{
		AggInterval:       khulnasoft.F(radar.AttackLayer7TimeseriesGroupVerticalParamsAggInterval15m),
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TimeseriesGroupVerticalParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod{radar.AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodGet, radar.AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPost, radar.AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion{radar.AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv1, radar.AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv2, radar.AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7TimeseriesGroupVerticalParamsIPVersion{radar.AttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv4, radar.AttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv6}),
		LimitPerGroup:     khulnasoft.F(int64(4)),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct{radar.AttackLayer7TimeseriesGroupVerticalParamsMitigationProductDDoS, radar.AttackLayer7TimeseriesGroupVerticalParamsMitigationProductWAF, radar.AttackLayer7TimeseriesGroupVerticalParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
		Normalization:     khulnasoft.F(radar.AttackLayer7TimeseriesGroupVerticalParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
