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

func TestHTTPAseTLSVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.TLSVersion.Get(
		context.TODO(),
		radar.HTTPAseTLSVersionGetParamsTLSVersionTlSv1_0,
		radar.HTTPAseTLSVersionGetParams{
			ASN:           khulnasoft.F([]string{"string", "string", "string"}),
			BotClass:      khulnasoft.F([]radar.HTTPAseTLSVersionGetParamsBotClass{radar.HTTPAseTLSVersionGetParamsBotClassLikelyAutomated, radar.HTTPAseTLSVersionGetParamsBotClassLikelyHuman}),
			BrowserFamily: khulnasoft.F([]radar.HTTPAseTLSVersionGetParamsBrowserFamily{radar.HTTPAseTLSVersionGetParamsBrowserFamilyChrome, radar.HTTPAseTLSVersionGetParamsBrowserFamilyEdge, radar.HTTPAseTLSVersionGetParamsBrowserFamilyFirefox}),
			Continent:     khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    khulnasoft.F([]radar.HTTPAseTLSVersionGetParamsDeviceType{radar.HTTPAseTLSVersionGetParamsDeviceTypeDesktop, radar.HTTPAseTLSVersionGetParamsDeviceTypeMobile, radar.HTTPAseTLSVersionGetParamsDeviceTypeOther}),
			Format:        khulnasoft.F(radar.HTTPAseTLSVersionGetParamsFormatJson),
			HTTPProtocol:  khulnasoft.F([]radar.HTTPAseTLSVersionGetParamsHTTPProtocol{radar.HTTPAseTLSVersionGetParamsHTTPProtocolHTTP, radar.HTTPAseTLSVersionGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   khulnasoft.F([]radar.HTTPAseTLSVersionGetParamsHTTPVersion{radar.HTTPAseTLSVersionGetParamsHTTPVersionHttPv1, radar.HTTPAseTLSVersionGetParamsHTTPVersionHttPv2, radar.HTTPAseTLSVersionGetParamsHTTPVersionHttPv3}),
			IPVersion:     khulnasoft.F([]radar.HTTPAseTLSVersionGetParamsIPVersion{radar.HTTPAseTLSVersionGetParamsIPVersionIPv4, radar.HTTPAseTLSVersionGetParamsIPVersionIPv6}),
			Limit:         khulnasoft.F(int64(5)),
			Location:      khulnasoft.F([]string{"string", "string", "string"}),
			Name:          khulnasoft.F([]string{"string", "string", "string"}),
			OS:            khulnasoft.F([]radar.HTTPAseTLSVersionGetParamsOS{radar.HTTPAseTLSVersionGetParamsOSWindows, radar.HTTPAseTLSVersionGetParamsOSMacosx, radar.HTTPAseTLSVersionGetParamsOSIos}),
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
