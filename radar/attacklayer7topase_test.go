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

func TestAttackLayer7TopAseOriginWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.Ases.Origin(context.TODO(), radar.AttackLayer7TopAseOriginParams{
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TopAseOriginParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7TopAseOriginParamsHTTPMethod{radar.AttackLayer7TopAseOriginParamsHTTPMethodGet, radar.AttackLayer7TopAseOriginParamsHTTPMethodPost, radar.AttackLayer7TopAseOriginParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7TopAseOriginParamsHTTPVersion{radar.AttackLayer7TopAseOriginParamsHTTPVersionHttPv1, radar.AttackLayer7TopAseOriginParamsHTTPVersionHttPv2, radar.AttackLayer7TopAseOriginParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7TopAseOriginParamsIPVersion{radar.AttackLayer7TopAseOriginParamsIPVersionIPv4, radar.AttackLayer7TopAseOriginParamsIPVersionIPv6}),
		Limit:             khulnasoft.F(int64(5)),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TopAseOriginParamsMitigationProduct{radar.AttackLayer7TopAseOriginParamsMitigationProductDDoS, radar.AttackLayer7TopAseOriginParamsMitigationProductWAF, radar.AttackLayer7TopAseOriginParamsMitigationProductBotManagement}),
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
