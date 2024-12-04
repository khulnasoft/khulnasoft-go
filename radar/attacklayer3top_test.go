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

func TestAttackLayer3TopAttacksWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Attacks(context.TODO(), radar.AttackLayer3TopAttacksParams{
		Continent:        khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:          khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:        khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:        khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:           khulnasoft.F(radar.AttackLayer3TopAttacksParamsFormatJson),
		IPVersion:        khulnasoft.F([]radar.AttackLayer3TopAttacksParamsIPVersion{radar.AttackLayer3TopAttacksParamsIPVersionIPv4, radar.AttackLayer3TopAttacksParamsIPVersionIPv6}),
		Limit:            khulnasoft.F(int64(5)),
		LimitDirection:   khulnasoft.F(radar.AttackLayer3TopAttacksParamsLimitDirectionOrigin),
		LimitPerLocation: khulnasoft.F(int64(10)),
		Location:         khulnasoft.F([]string{"string", "string", "string"}),
		Name:             khulnasoft.F([]string{"string", "string", "string"}),
		Normalization:    khulnasoft.F(radar.AttackLayer3TopAttacksParamsNormalizationPercentage),
		Protocol:         khulnasoft.F([]radar.AttackLayer3TopAttacksParamsProtocol{radar.AttackLayer3TopAttacksParamsProtocolUdp, radar.AttackLayer3TopAttacksParamsProtocolTCP, radar.AttackLayer3TopAttacksParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TopIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Industry(context.TODO(), radar.AttackLayer3TopIndustryParams{
		Continent: khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    khulnasoft.F(radar.AttackLayer3TopIndustryParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.AttackLayer3TopIndustryParamsIPVersion{radar.AttackLayer3TopIndustryParamsIPVersionIPv4, radar.AttackLayer3TopIndustryParamsIPVersionIPv6}),
		Limit:     khulnasoft.F(int64(5)),
		Location:  khulnasoft.F([]string{"string", "string", "string"}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		Protocol:  khulnasoft.F([]radar.AttackLayer3TopIndustryParamsProtocol{radar.AttackLayer3TopIndustryParamsProtocolUdp, radar.AttackLayer3TopIndustryParamsProtocolTCP, radar.AttackLayer3TopIndustryParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TopVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Vertical(context.TODO(), radar.AttackLayer3TopVerticalParams{
		Continent: khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    khulnasoft.F(radar.AttackLayer3TopVerticalParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.AttackLayer3TopVerticalParamsIPVersion{radar.AttackLayer3TopVerticalParamsIPVersionIPv4, radar.AttackLayer3TopVerticalParamsIPVersionIPv6}),
		Limit:     khulnasoft.F(int64(5)),
		Location:  khulnasoft.F([]string{"string", "string", "string"}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		Protocol:  khulnasoft.F([]radar.AttackLayer3TopVerticalParamsProtocol{radar.AttackLayer3TopVerticalParamsProtocolUdp, radar.AttackLayer3TopVerticalParamsProtocolTCP, radar.AttackLayer3TopVerticalParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
