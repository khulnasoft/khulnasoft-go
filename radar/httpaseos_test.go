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

func TestHTTPAseOSGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.OS.Get(
		context.TODO(),
		radar.HTTPAseOSGetParamsOSWindows,
		radar.HTTPAseOSGetParams{
			ASN:           khulnasoft.F([]string{"string", "string", "string"}),
			BotClass:      khulnasoft.F([]radar.HTTPAseOSGetParamsBotClass{radar.HTTPAseOSGetParamsBotClassLikelyAutomated, radar.HTTPAseOSGetParamsBotClassLikelyHuman}),
			BrowserFamily: khulnasoft.F([]radar.HTTPAseOSGetParamsBrowserFamily{radar.HTTPAseOSGetParamsBrowserFamilyChrome, radar.HTTPAseOSGetParamsBrowserFamilyEdge, radar.HTTPAseOSGetParamsBrowserFamilyFirefox}),
			Continent:     khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    khulnasoft.F([]radar.HTTPAseOSGetParamsDeviceType{radar.HTTPAseOSGetParamsDeviceTypeDesktop, radar.HTTPAseOSGetParamsDeviceTypeMobile, radar.HTTPAseOSGetParamsDeviceTypeOther}),
			Format:        khulnasoft.F(radar.HTTPAseOSGetParamsFormatJson),
			HTTPProtocol:  khulnasoft.F([]radar.HTTPAseOSGetParamsHTTPProtocol{radar.HTTPAseOSGetParamsHTTPProtocolHTTP, radar.HTTPAseOSGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   khulnasoft.F([]radar.HTTPAseOSGetParamsHTTPVersion{radar.HTTPAseOSGetParamsHTTPVersionHttPv1, radar.HTTPAseOSGetParamsHTTPVersionHttPv2, radar.HTTPAseOSGetParamsHTTPVersionHttPv3}),
			IPVersion:     khulnasoft.F([]radar.HTTPAseOSGetParamsIPVersion{radar.HTTPAseOSGetParamsIPVersionIPv4, radar.HTTPAseOSGetParamsIPVersionIPv6}),
			Limit:         khulnasoft.F(int64(5)),
			Location:      khulnasoft.F([]string{"string", "string", "string"}),
			Name:          khulnasoft.F([]string{"string", "string", "string"}),
			TLSVersion:    khulnasoft.F([]radar.HTTPAseOSGetParamsTLSVersion{radar.HTTPAseOSGetParamsTLSVersionTlSv1_0, radar.HTTPAseOSGetParamsTLSVersionTlSv1_1, radar.HTTPAseOSGetParamsTLSVersionTlSv1_2}),
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
