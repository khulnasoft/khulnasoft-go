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

func TestHTTPLocationTLSVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.TLSVersion.Get(
		context.TODO(),
		radar.HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_0,
		radar.HTTPLocationTLSVersionGetParams{
			ASN:           khulnasoft.F([]string{"string", "string", "string"}),
			BotClass:      khulnasoft.F([]radar.HTTPLocationTLSVersionGetParamsBotClass{radar.HTTPLocationTLSVersionGetParamsBotClassLikelyAutomated, radar.HTTPLocationTLSVersionGetParamsBotClassLikelyHuman}),
			BrowserFamily: khulnasoft.F([]radar.HTTPLocationTLSVersionGetParamsBrowserFamily{radar.HTTPLocationTLSVersionGetParamsBrowserFamilyChrome, radar.HTTPLocationTLSVersionGetParamsBrowserFamilyEdge, radar.HTTPLocationTLSVersionGetParamsBrowserFamilyFirefox}),
			Continent:     khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    khulnasoft.F([]radar.HTTPLocationTLSVersionGetParamsDeviceType{radar.HTTPLocationTLSVersionGetParamsDeviceTypeDesktop, radar.HTTPLocationTLSVersionGetParamsDeviceTypeMobile, radar.HTTPLocationTLSVersionGetParamsDeviceTypeOther}),
			Format:        khulnasoft.F(radar.HTTPLocationTLSVersionGetParamsFormatJson),
			HTTPProtocol:  khulnasoft.F([]radar.HTTPLocationTLSVersionGetParamsHTTPProtocol{radar.HTTPLocationTLSVersionGetParamsHTTPProtocolHTTP, radar.HTTPLocationTLSVersionGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   khulnasoft.F([]radar.HTTPLocationTLSVersionGetParamsHTTPVersion{radar.HTTPLocationTLSVersionGetParamsHTTPVersionHttPv1, radar.HTTPLocationTLSVersionGetParamsHTTPVersionHttPv2, radar.HTTPLocationTLSVersionGetParamsHTTPVersionHttPv3}),
			IPVersion:     khulnasoft.F([]radar.HTTPLocationTLSVersionGetParamsIPVersion{radar.HTTPLocationTLSVersionGetParamsIPVersionIPv4, radar.HTTPLocationTLSVersionGetParamsIPVersionIPv6}),
			Limit:         khulnasoft.F(int64(5)),
			Location:      khulnasoft.F([]string{"string", "string", "string"}),
			Name:          khulnasoft.F([]string{"string", "string", "string"}),
			OS:            khulnasoft.F([]radar.HTTPLocationTLSVersionGetParamsOS{radar.HTTPLocationTLSVersionGetParamsOSWindows, radar.HTTPLocationTLSVersionGetParamsOSMacosx, radar.HTTPLocationTLSVersionGetParamsOSIos}),
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
