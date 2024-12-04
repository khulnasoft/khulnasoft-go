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

func TestHTTPLocationOSGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.OS.Get(
		context.TODO(),
		radar.HTTPLocationOSGetParamsOSWindows,
		radar.HTTPLocationOSGetParams{
			ASN:           khulnasoft.F([]string{"string", "string", "string"}),
			BotClass:      khulnasoft.F([]radar.HTTPLocationOSGetParamsBotClass{radar.HTTPLocationOSGetParamsBotClassLikelyAutomated, radar.HTTPLocationOSGetParamsBotClassLikelyHuman}),
			BrowserFamily: khulnasoft.F([]radar.HTTPLocationOSGetParamsBrowserFamily{radar.HTTPLocationOSGetParamsBrowserFamilyChrome, radar.HTTPLocationOSGetParamsBrowserFamilyEdge, radar.HTTPLocationOSGetParamsBrowserFamilyFirefox}),
			Continent:     khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    khulnasoft.F([]radar.HTTPLocationOSGetParamsDeviceType{radar.HTTPLocationOSGetParamsDeviceTypeDesktop, radar.HTTPLocationOSGetParamsDeviceTypeMobile, radar.HTTPLocationOSGetParamsDeviceTypeOther}),
			Format:        khulnasoft.F(radar.HTTPLocationOSGetParamsFormatJson),
			HTTPProtocol:  khulnasoft.F([]radar.HTTPLocationOSGetParamsHTTPProtocol{radar.HTTPLocationOSGetParamsHTTPProtocolHTTP, radar.HTTPLocationOSGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   khulnasoft.F([]radar.HTTPLocationOSGetParamsHTTPVersion{radar.HTTPLocationOSGetParamsHTTPVersionHttPv1, radar.HTTPLocationOSGetParamsHTTPVersionHttPv2, radar.HTTPLocationOSGetParamsHTTPVersionHttPv3}),
			IPVersion:     khulnasoft.F([]radar.HTTPLocationOSGetParamsIPVersion{radar.HTTPLocationOSGetParamsIPVersionIPv4, radar.HTTPLocationOSGetParamsIPVersionIPv6}),
			Limit:         khulnasoft.F(int64(5)),
			Location:      khulnasoft.F([]string{"string", "string", "string"}),
			Name:          khulnasoft.F([]string{"string", "string", "string"}),
			TLSVersion:    khulnasoft.F([]radar.HTTPLocationOSGetParamsTLSVersion{radar.HTTPLocationOSGetParamsTLSVersionTlSv1_0, radar.HTTPLocationOSGetParamsTLSVersionTlSv1_1, radar.HTTPLocationOSGetParamsTLSVersionTlSv1_2}),
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
