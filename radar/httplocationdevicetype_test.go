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

func TestHTTPLocationDeviceTypeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.DeviceType.Get(
		context.TODO(),
		radar.HTTPLocationDeviceTypeGetParamsDeviceTypeDesktop,
		radar.HTTPLocationDeviceTypeGetParams{
			ASN:           khulnasoft.F([]string{"string", "string", "string"}),
			BotClass:      khulnasoft.F([]radar.HTTPLocationDeviceTypeGetParamsBotClass{radar.HTTPLocationDeviceTypeGetParamsBotClassLikelyAutomated, radar.HTTPLocationDeviceTypeGetParamsBotClassLikelyHuman}),
			BrowserFamily: khulnasoft.F([]radar.HTTPLocationDeviceTypeGetParamsBrowserFamily{radar.HTTPLocationDeviceTypeGetParamsBrowserFamilyChrome, radar.HTTPLocationDeviceTypeGetParamsBrowserFamilyEdge, radar.HTTPLocationDeviceTypeGetParamsBrowserFamilyFirefox}),
			Continent:     khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Format:        khulnasoft.F(radar.HTTPLocationDeviceTypeGetParamsFormatJson),
			HTTPProtocol:  khulnasoft.F([]radar.HTTPLocationDeviceTypeGetParamsHTTPProtocol{radar.HTTPLocationDeviceTypeGetParamsHTTPProtocolHTTP, radar.HTTPLocationDeviceTypeGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   khulnasoft.F([]radar.HTTPLocationDeviceTypeGetParamsHTTPVersion{radar.HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv1, radar.HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv2, radar.HTTPLocationDeviceTypeGetParamsHTTPVersionHttPv3}),
			IPVersion:     khulnasoft.F([]radar.HTTPLocationDeviceTypeGetParamsIPVersion{radar.HTTPLocationDeviceTypeGetParamsIPVersionIPv4, radar.HTTPLocationDeviceTypeGetParamsIPVersionIPv6}),
			Limit:         khulnasoft.F(int64(5)),
			Location:      khulnasoft.F([]string{"string", "string", "string"}),
			Name:          khulnasoft.F([]string{"string", "string", "string"}),
			OS:            khulnasoft.F([]radar.HTTPLocationDeviceTypeGetParamsOS{radar.HTTPLocationDeviceTypeGetParamsOSWindows, radar.HTTPLocationDeviceTypeGetParamsOSMacosx, radar.HTTPLocationDeviceTypeGetParamsOSIos}),
			TLSVersion:    khulnasoft.F([]radar.HTTPLocationDeviceTypeGetParamsTLSVersion{radar.HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_0, radar.HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_1, radar.HTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_2}),
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
