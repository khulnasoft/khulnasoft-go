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

func TestAttackLayer7TopAttacksWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.Attacks(context.TODO(), radar.AttackLayer7TopAttacksParams{
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TopAttacksParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7TopAttacksParamsHTTPMethod{radar.AttackLayer7TopAttacksParamsHTTPMethodGet, radar.AttackLayer7TopAttacksParamsHTTPMethodPost, radar.AttackLayer7TopAttacksParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7TopAttacksParamsHTTPVersion{radar.AttackLayer7TopAttacksParamsHTTPVersionHttPv1, radar.AttackLayer7TopAttacksParamsHTTPVersionHttPv2, radar.AttackLayer7TopAttacksParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7TopAttacksParamsIPVersion{radar.AttackLayer7TopAttacksParamsIPVersionIPv4, radar.AttackLayer7TopAttacksParamsIPVersionIPv6}),
		Limit:             khulnasoft.F(int64(5)),
		LimitDirection:    khulnasoft.F(radar.AttackLayer7TopAttacksParamsLimitDirectionOrigin),
		LimitPerLocation:  khulnasoft.F(int64(10)),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		Magnitude:         khulnasoft.F(radar.AttackLayer7TopAttacksParamsMagnitudeAffectedZones),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TopAttacksParamsMitigationProduct{radar.AttackLayer7TopAttacksParamsMitigationProductDDoS, radar.AttackLayer7TopAttacksParamsMitigationProductWAF, radar.AttackLayer7TopAttacksParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
		Normalization:     khulnasoft.F(radar.AttackLayer7TopAttacksParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TopIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.Industry(context.TODO(), radar.AttackLayer7TopIndustryParams{
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TopIndustryParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7TopIndustryParamsHTTPMethod{radar.AttackLayer7TopIndustryParamsHTTPMethodGet, radar.AttackLayer7TopIndustryParamsHTTPMethodPost, radar.AttackLayer7TopIndustryParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7TopIndustryParamsHTTPVersion{radar.AttackLayer7TopIndustryParamsHTTPVersionHttPv1, radar.AttackLayer7TopIndustryParamsHTTPVersionHttPv2, radar.AttackLayer7TopIndustryParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7TopIndustryParamsIPVersion{radar.AttackLayer7TopIndustryParamsIPVersionIPv4, radar.AttackLayer7TopIndustryParamsIPVersionIPv6}),
		Limit:             khulnasoft.F(int64(5)),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TopIndustryParamsMitigationProduct{radar.AttackLayer7TopIndustryParamsMitigationProductDDoS, radar.AttackLayer7TopIndustryParamsMitigationProductWAF, radar.AttackLayer7TopIndustryParamsMitigationProductBotManagement}),
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

func TestAttackLayer7TopVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.Vertical(context.TODO(), radar.AttackLayer7TopVerticalParams{
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TopVerticalParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7TopVerticalParamsHTTPMethod{radar.AttackLayer7TopVerticalParamsHTTPMethodGet, radar.AttackLayer7TopVerticalParamsHTTPMethodPost, radar.AttackLayer7TopVerticalParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7TopVerticalParamsHTTPVersion{radar.AttackLayer7TopVerticalParamsHTTPVersionHttPv1, radar.AttackLayer7TopVerticalParamsHTTPVersionHttPv2, radar.AttackLayer7TopVerticalParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7TopVerticalParamsIPVersion{radar.AttackLayer7TopVerticalParamsIPVersionIPv4, radar.AttackLayer7TopVerticalParamsIPVersionIPv6}),
		Limit:             khulnasoft.F(int64(5)),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TopVerticalParamsMitigationProduct{radar.AttackLayer7TopVerticalParamsMitigationProductDDoS, radar.AttackLayer7TopVerticalParamsMitigationProductWAF, radar.AttackLayer7TopVerticalParamsMitigationProductBotManagement}),
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
