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

func TestHTTPAseBotClassGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.BotClass.Get(
		context.TODO(),
		radar.HTTPAseBotClassGetParamsBotClassLikelyAutomated,
		radar.HTTPAseBotClassGetParams{
			ASN:           khulnasoft.F([]string{"string", "string", "string"}),
			BrowserFamily: khulnasoft.F([]radar.HTTPAseBotClassGetParamsBrowserFamily{radar.HTTPAseBotClassGetParamsBrowserFamilyChrome, radar.HTTPAseBotClassGetParamsBrowserFamilyEdge, radar.HTTPAseBotClassGetParamsBrowserFamilyFirefox}),
			Continent:     khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    khulnasoft.F([]radar.HTTPAseBotClassGetParamsDeviceType{radar.HTTPAseBotClassGetParamsDeviceTypeDesktop, radar.HTTPAseBotClassGetParamsDeviceTypeMobile, radar.HTTPAseBotClassGetParamsDeviceTypeOther}),
			Format:        khulnasoft.F(radar.HTTPAseBotClassGetParamsFormatJson),
			HTTPProtocol:  khulnasoft.F([]radar.HTTPAseBotClassGetParamsHTTPProtocol{radar.HTTPAseBotClassGetParamsHTTPProtocolHTTP, radar.HTTPAseBotClassGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   khulnasoft.F([]radar.HTTPAseBotClassGetParamsHTTPVersion{radar.HTTPAseBotClassGetParamsHTTPVersionHttPv1, radar.HTTPAseBotClassGetParamsHTTPVersionHttPv2, radar.HTTPAseBotClassGetParamsHTTPVersionHttPv3}),
			IPVersion:     khulnasoft.F([]radar.HTTPAseBotClassGetParamsIPVersion{radar.HTTPAseBotClassGetParamsIPVersionIPv4, radar.HTTPAseBotClassGetParamsIPVersionIPv6}),
			Limit:         khulnasoft.F(int64(5)),
			Location:      khulnasoft.F([]string{"string", "string", "string"}),
			Name:          khulnasoft.F([]string{"string", "string", "string"}),
			OS:            khulnasoft.F([]radar.HTTPAseBotClassGetParamsOS{radar.HTTPAseBotClassGetParamsOSWindows, radar.HTTPAseBotClassGetParamsOSMacosx, radar.HTTPAseBotClassGetParamsOSIos}),
			TLSVersion:    khulnasoft.F([]radar.HTTPAseBotClassGetParamsTLSVersion{radar.HTTPAseBotClassGetParamsTLSVersionTlSv1_0, radar.HTTPAseBotClassGetParamsTLSVersionTlSv1_1, radar.HTTPAseBotClassGetParamsTLSVersionTlSv1_2}),
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
