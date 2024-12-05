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

func TestHTTPAseDeviceTypeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.DeviceType.Get(
		context.TODO(),
		radar.HTTPAseDeviceTypeGetParamsDeviceTypeDesktop,
		radar.HTTPAseDeviceTypeGetParams{
			ASN:           khulnasoft.F([]string{"string", "string", "string"}),
			BotClass:      khulnasoft.F([]radar.HTTPAseDeviceTypeGetParamsBotClass{radar.HTTPAseDeviceTypeGetParamsBotClassLikelyAutomated, radar.HTTPAseDeviceTypeGetParamsBotClassLikelyHuman}),
			BrowserFamily: khulnasoft.F([]radar.HTTPAseDeviceTypeGetParamsBrowserFamily{radar.HTTPAseDeviceTypeGetParamsBrowserFamilyChrome, radar.HTTPAseDeviceTypeGetParamsBrowserFamilyEdge, radar.HTTPAseDeviceTypeGetParamsBrowserFamilyFirefox}),
			Continent:     khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Format:        khulnasoft.F(radar.HTTPAseDeviceTypeGetParamsFormatJson),
			HTTPProtocol:  khulnasoft.F([]radar.HTTPAseDeviceTypeGetParamsHTTPProtocol{radar.HTTPAseDeviceTypeGetParamsHTTPProtocolHTTP, radar.HTTPAseDeviceTypeGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   khulnasoft.F([]radar.HTTPAseDeviceTypeGetParamsHTTPVersion{radar.HTTPAseDeviceTypeGetParamsHTTPVersionHttPv1, radar.HTTPAseDeviceTypeGetParamsHTTPVersionHttPv2, radar.HTTPAseDeviceTypeGetParamsHTTPVersionHttPv3}),
			IPVersion:     khulnasoft.F([]radar.HTTPAseDeviceTypeGetParamsIPVersion{radar.HTTPAseDeviceTypeGetParamsIPVersionIPv4, radar.HTTPAseDeviceTypeGetParamsIPVersionIPv6}),
			Limit:         khulnasoft.F(int64(5)),
			Location:      khulnasoft.F([]string{"string", "string", "string"}),
			Name:          khulnasoft.F([]string{"string", "string", "string"}),
			OS:            khulnasoft.F([]radar.HTTPAseDeviceTypeGetParamsOS{radar.HTTPAseDeviceTypeGetParamsOSWindows, radar.HTTPAseDeviceTypeGetParamsOSMacosx, radar.HTTPAseDeviceTypeGetParamsOSIos}),
			TLSVersion:    khulnasoft.F([]radar.HTTPAseDeviceTypeGetParamsTLSVersion{radar.HTTPAseDeviceTypeGetParamsTLSVersionTlSv1_0, radar.HTTPAseDeviceTypeGetParamsTLSVersionTlSv1_1, radar.HTTPAseDeviceTypeGetParamsTLSVersionTlSv1_2}),
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
