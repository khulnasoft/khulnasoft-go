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

func TestAttackLayer3TimeseriesGroupBitrateWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Bitrate(context.TODO(), radar.AttackLayer3TimeseriesGroupBitrateParams{
		AggInterval:   khulnasoft.F(radar.AttackLayer3TimeseriesGroupBitrateParamsAggInterval15m),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     khulnasoft.F(radar.AttackLayer3TimeseriesGroupBitrateParamsDirectionOrigin),
		Format:        khulnasoft.F(radar.AttackLayer3TimeseriesGroupBitrateParamsFormatJson),
		IPVersion:     khulnasoft.F([]radar.AttackLayer3TimeseriesGroupBitrateParamsIPVersion{radar.AttackLayer3TimeseriesGroupBitrateParamsIPVersionIPv4, radar.AttackLayer3TimeseriesGroupBitrateParamsIPVersionIPv6}),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		Normalization: khulnasoft.F(radar.AttackLayer3TimeseriesGroupBitrateParamsNormalizationPercentage),
		Protocol:      khulnasoft.F([]radar.AttackLayer3TimeseriesGroupBitrateParamsProtocol{radar.AttackLayer3TimeseriesGroupBitrateParamsProtocolUdp, radar.AttackLayer3TimeseriesGroupBitrateParamsProtocolTCP, radar.AttackLayer3TimeseriesGroupBitrateParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupDurationWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Duration(context.TODO(), radar.AttackLayer3TimeseriesGroupDurationParams{
		AggInterval:   khulnasoft.F(radar.AttackLayer3TimeseriesGroupDurationParamsAggInterval15m),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     khulnasoft.F(radar.AttackLayer3TimeseriesGroupDurationParamsDirectionOrigin),
		Format:        khulnasoft.F(radar.AttackLayer3TimeseriesGroupDurationParamsFormatJson),
		IPVersion:     khulnasoft.F([]radar.AttackLayer3TimeseriesGroupDurationParamsIPVersion{radar.AttackLayer3TimeseriesGroupDurationParamsIPVersionIPv4, radar.AttackLayer3TimeseriesGroupDurationParamsIPVersionIPv6}),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		Normalization: khulnasoft.F(radar.AttackLayer3TimeseriesGroupDurationParamsNormalizationPercentage),
		Protocol:      khulnasoft.F([]radar.AttackLayer3TimeseriesGroupDurationParamsProtocol{radar.AttackLayer3TimeseriesGroupDurationParamsProtocolUdp, radar.AttackLayer3TimeseriesGroupDurationParamsProtocolTCP, radar.AttackLayer3TimeseriesGroupDurationParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Get(context.TODO(), radar.AttackLayer3TimeseriesGroupGetParams{
		AggInterval: khulnasoft.F(radar.AttackLayer3TimeseriesGroupGetParamsAggInterval15m),
		ASN:         khulnasoft.F([]string{"string", "string", "string"}),
		Continent:   khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      khulnasoft.F(radar.AttackLayer3TimeseriesGroupGetParamsFormatJson),
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

func TestAttackLayer3TimeseriesGroupIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Industry(context.TODO(), radar.AttackLayer3TimeseriesGroupIndustryParams{
		AggInterval:   khulnasoft.F(radar.AttackLayer3TimeseriesGroupIndustryParamsAggInterval15m),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     khulnasoft.F(radar.AttackLayer3TimeseriesGroupIndustryParamsDirectionOrigin),
		Format:        khulnasoft.F(radar.AttackLayer3TimeseriesGroupIndustryParamsFormatJson),
		IPVersion:     khulnasoft.F([]radar.AttackLayer3TimeseriesGroupIndustryParamsIPVersion{radar.AttackLayer3TimeseriesGroupIndustryParamsIPVersionIPv4, radar.AttackLayer3TimeseriesGroupIndustryParamsIPVersionIPv6}),
		LimitPerGroup: khulnasoft.F(int64(4)),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		Normalization: khulnasoft.F(radar.AttackLayer3TimeseriesGroupIndustryParamsNormalizationPercentage),
		Protocol:      khulnasoft.F([]radar.AttackLayer3TimeseriesGroupIndustryParamsProtocol{radar.AttackLayer3TimeseriesGroupIndustryParamsProtocolUdp, radar.AttackLayer3TimeseriesGroupIndustryParamsProtocolTCP, radar.AttackLayer3TimeseriesGroupIndustryParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.IPVersion(context.TODO(), radar.AttackLayer3TimeseriesGroupIPVersionParams{
		AggInterval:   khulnasoft.F(radar.AttackLayer3TimeseriesGroupIPVersionParamsAggInterval15m),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     khulnasoft.F(radar.AttackLayer3TimeseriesGroupIPVersionParamsDirectionOrigin),
		Format:        khulnasoft.F(radar.AttackLayer3TimeseriesGroupIPVersionParamsFormatJson),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		Normalization: khulnasoft.F(radar.AttackLayer3TimeseriesGroupIPVersionParamsNormalizationPercentage),
		Protocol:      khulnasoft.F([]radar.AttackLayer3TimeseriesGroupIPVersionParamsProtocol{radar.AttackLayer3TimeseriesGroupIPVersionParamsProtocolUdp, radar.AttackLayer3TimeseriesGroupIPVersionParamsProtocolTCP, radar.AttackLayer3TimeseriesGroupIPVersionParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Protocol(context.TODO(), radar.AttackLayer3TimeseriesGroupProtocolParams{
		AggInterval:   khulnasoft.F(radar.AttackLayer3TimeseriesGroupProtocolParamsAggInterval15m),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     khulnasoft.F(radar.AttackLayer3TimeseriesGroupProtocolParamsDirectionOrigin),
		Format:        khulnasoft.F(radar.AttackLayer3TimeseriesGroupProtocolParamsFormatJson),
		IPVersion:     khulnasoft.F([]radar.AttackLayer3TimeseriesGroupProtocolParamsIPVersion{radar.AttackLayer3TimeseriesGroupProtocolParamsIPVersionIPv4, radar.AttackLayer3TimeseriesGroupProtocolParamsIPVersionIPv6}),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		Normalization: khulnasoft.F(radar.AttackLayer3TimeseriesGroupProtocolParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupVectorWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Vector(context.TODO(), radar.AttackLayer3TimeseriesGroupVectorParams{
		AggInterval:   khulnasoft.F(radar.AttackLayer3TimeseriesGroupVectorParamsAggInterval15m),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     khulnasoft.F(radar.AttackLayer3TimeseriesGroupVectorParamsDirectionOrigin),
		Format:        khulnasoft.F(radar.AttackLayer3TimeseriesGroupVectorParamsFormatJson),
		IPVersion:     khulnasoft.F([]radar.AttackLayer3TimeseriesGroupVectorParamsIPVersion{radar.AttackLayer3TimeseriesGroupVectorParamsIPVersionIPv4, radar.AttackLayer3TimeseriesGroupVectorParamsIPVersionIPv6}),
		LimitPerGroup: khulnasoft.F(int64(4)),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		Normalization: khulnasoft.F(radar.AttackLayer3TimeseriesGroupVectorParamsNormalizationPercentage),
		Protocol:      khulnasoft.F([]radar.AttackLayer3TimeseriesGroupVectorParamsProtocol{radar.AttackLayer3TimeseriesGroupVectorParamsProtocolUdp, radar.AttackLayer3TimeseriesGroupVectorParamsProtocolTCP, radar.AttackLayer3TimeseriesGroupVectorParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TimeseriesGroupVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.Vertical(context.TODO(), radar.AttackLayer3TimeseriesGroupVerticalParams{
		AggInterval:   khulnasoft.F(radar.AttackLayer3TimeseriesGroupVerticalParamsAggInterval15m),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     khulnasoft.F(radar.AttackLayer3TimeseriesGroupVerticalParamsDirectionOrigin),
		Format:        khulnasoft.F(radar.AttackLayer3TimeseriesGroupVerticalParamsFormatJson),
		IPVersion:     khulnasoft.F([]radar.AttackLayer3TimeseriesGroupVerticalParamsIPVersion{radar.AttackLayer3TimeseriesGroupVerticalParamsIPVersionIPv4, radar.AttackLayer3TimeseriesGroupVerticalParamsIPVersionIPv6}),
		LimitPerGroup: khulnasoft.F(int64(4)),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		Normalization: khulnasoft.F(radar.AttackLayer3TimeseriesGroupVerticalParamsNormalizationPercentage),
		Protocol:      khulnasoft.F([]radar.AttackLayer3TimeseriesGroupVerticalParamsProtocol{radar.AttackLayer3TimeseriesGroupVerticalParamsProtocolUdp, radar.AttackLayer3TimeseriesGroupVerticalParamsProtocolTCP, radar.AttackLayer3TimeseriesGroupVerticalParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
