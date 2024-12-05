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

func TestAttackLayer3TimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Timeseries(context.TODO(), radar.AttackLayer3TimeseriesParams{
		AggInterval:   khulnasoft.F(radar.AttackLayer3TimeseriesParamsAggInterval15m),
		ASN:           khulnasoft.F([]string{"string", "string", "string"}),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     khulnasoft.F(radar.AttackLayer3TimeseriesParamsDirectionOrigin),
		Format:        khulnasoft.F(radar.AttackLayer3TimeseriesParamsFormatJson),
		IPVersion:     khulnasoft.F([]radar.AttackLayer3TimeseriesParamsIPVersion{radar.AttackLayer3TimeseriesParamsIPVersionIPv4, radar.AttackLayer3TimeseriesParamsIPVersionIPv6}),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Metric:        khulnasoft.F(radar.AttackLayer3TimeseriesParamsMetricBytes),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		Normalization: khulnasoft.F(radar.AttackLayer3TimeseriesParamsNormalizationPercentageChange),
		Protocol:      khulnasoft.F([]radar.AttackLayer3TimeseriesParamsProtocol{radar.AttackLayer3TimeseriesParamsProtocolUdp, radar.AttackLayer3TimeseriesParamsProtocolTCP, radar.AttackLayer3TimeseriesParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
