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

func TestHTTPLocationHTTPMethodGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.HTTPMethod.Get(
		context.TODO(),
		radar.HTTPLocationHTTPMethodGetParamsHTTPVersionHttPv1,
		radar.HTTPLocationHTTPMethodGetParams{
			ASN:           khulnasoft.F([]string{"string", "string", "string"}),
			BotClass:      khulnasoft.F([]radar.HTTPLocationHTTPMethodGetParamsBotClass{radar.HTTPLocationHTTPMethodGetParamsBotClassLikelyAutomated, radar.HTTPLocationHTTPMethodGetParamsBotClassLikelyHuman}),
			BrowserFamily: khulnasoft.F([]radar.HTTPLocationHTTPMethodGetParamsBrowserFamily{radar.HTTPLocationHTTPMethodGetParamsBrowserFamilyChrome, radar.HTTPLocationHTTPMethodGetParamsBrowserFamilyEdge, radar.HTTPLocationHTTPMethodGetParamsBrowserFamilyFirefox}),
			Continent:     khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    khulnasoft.F([]radar.HTTPLocationHTTPMethodGetParamsDeviceType{radar.HTTPLocationHTTPMethodGetParamsDeviceTypeDesktop, radar.HTTPLocationHTTPMethodGetParamsDeviceTypeMobile, radar.HTTPLocationHTTPMethodGetParamsDeviceTypeOther}),
			Format:        khulnasoft.F(radar.HTTPLocationHTTPMethodGetParamsFormatJson),
			HTTPProtocol:  khulnasoft.F([]radar.HTTPLocationHTTPMethodGetParamsHTTPProtocol{radar.HTTPLocationHTTPMethodGetParamsHTTPProtocolHTTP, radar.HTTPLocationHTTPMethodGetParamsHTTPProtocolHTTPS}),
			IPVersion:     khulnasoft.F([]radar.HTTPLocationHTTPMethodGetParamsIPVersion{radar.HTTPLocationHTTPMethodGetParamsIPVersionIPv4, radar.HTTPLocationHTTPMethodGetParamsIPVersionIPv6}),
			Limit:         khulnasoft.F(int64(5)),
			Location:      khulnasoft.F([]string{"string", "string", "string"}),
			Name:          khulnasoft.F([]string{"string", "string", "string"}),
			OS:            khulnasoft.F([]radar.HTTPLocationHTTPMethodGetParamsOS{radar.HTTPLocationHTTPMethodGetParamsOSWindows, radar.HTTPLocationHTTPMethodGetParamsOSMacosx, radar.HTTPLocationHTTPMethodGetParamsOSIos}),
			TLSVersion:    khulnasoft.F([]radar.HTTPLocationHTTPMethodGetParamsTLSVersion{radar.HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_0, radar.HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_1, radar.HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_2}),
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
