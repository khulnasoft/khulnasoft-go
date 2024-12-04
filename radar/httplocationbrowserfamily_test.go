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

func TestHTTPLocationBrowserFamilyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.BrowserFamily.Get(
		context.TODO(),
		radar.HTTPLocationBrowserFamilyGetParamsBrowserFamilyChrome,
		radar.HTTPLocationBrowserFamilyGetParams{
			ASN:          khulnasoft.F([]string{"string", "string", "string"}),
			BotClass:     khulnasoft.F([]radar.HTTPLocationBrowserFamilyGetParamsBotClass{radar.HTTPLocationBrowserFamilyGetParamsBotClassLikelyAutomated, radar.HTTPLocationBrowserFamilyGetParamsBotClassLikelyHuman}),
			Continent:    khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   khulnasoft.F([]radar.HTTPLocationBrowserFamilyGetParamsDeviceType{radar.HTTPLocationBrowserFamilyGetParamsDeviceTypeDesktop, radar.HTTPLocationBrowserFamilyGetParamsDeviceTypeMobile, radar.HTTPLocationBrowserFamilyGetParamsDeviceTypeOther}),
			Format:       khulnasoft.F(radar.HTTPLocationBrowserFamilyGetParamsFormatJson),
			HTTPProtocol: khulnasoft.F([]radar.HTTPLocationBrowserFamilyGetParamsHTTPProtocol{radar.HTTPLocationBrowserFamilyGetParamsHTTPProtocolHTTP, radar.HTTPLocationBrowserFamilyGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  khulnasoft.F([]radar.HTTPLocationBrowserFamilyGetParamsHTTPVersion{radar.HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv1, radar.HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv2, radar.HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv3}),
			IPVersion:    khulnasoft.F([]radar.HTTPLocationBrowserFamilyGetParamsIPVersion{radar.HTTPLocationBrowserFamilyGetParamsIPVersionIPv4, radar.HTTPLocationBrowserFamilyGetParamsIPVersionIPv6}),
			Limit:        khulnasoft.F(int64(5)),
			Location:     khulnasoft.F([]string{"string", "string", "string"}),
			Name:         khulnasoft.F([]string{"string", "string", "string"}),
			OS:           khulnasoft.F([]radar.HTTPLocationBrowserFamilyGetParamsOS{radar.HTTPLocationBrowserFamilyGetParamsOSWindows, radar.HTTPLocationBrowserFamilyGetParamsOSMacosx, radar.HTTPLocationBrowserFamilyGetParamsOSIos}),
			TLSVersion:   khulnasoft.F([]radar.HTTPLocationBrowserFamilyGetParamsTLSVersion{radar.HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_0, radar.HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_1, radar.HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_2}),
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
