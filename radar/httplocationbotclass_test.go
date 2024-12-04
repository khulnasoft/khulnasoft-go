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

func TestHTTPLocationBotClassGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.BotClass.Get(
		context.TODO(),
		radar.HTTPLocationBotClassGetParamsBotClassLikelyAutomated,
		radar.HTTPLocationBotClassGetParams{
			ASN:           khulnasoft.F([]string{"string", "string", "string"}),
			BrowserFamily: khulnasoft.F([]radar.HTTPLocationBotClassGetParamsBrowserFamily{radar.HTTPLocationBotClassGetParamsBrowserFamilyChrome, radar.HTTPLocationBotClassGetParamsBrowserFamilyEdge, radar.HTTPLocationBotClassGetParamsBrowserFamilyFirefox}),
			Continent:     khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    khulnasoft.F([]radar.HTTPLocationBotClassGetParamsDeviceType{radar.HTTPLocationBotClassGetParamsDeviceTypeDesktop, radar.HTTPLocationBotClassGetParamsDeviceTypeMobile, radar.HTTPLocationBotClassGetParamsDeviceTypeOther}),
			Format:        khulnasoft.F(radar.HTTPLocationBotClassGetParamsFormatJson),
			HTTPProtocol:  khulnasoft.F([]radar.HTTPLocationBotClassGetParamsHTTPProtocol{radar.HTTPLocationBotClassGetParamsHTTPProtocolHTTP, radar.HTTPLocationBotClassGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   khulnasoft.F([]radar.HTTPLocationBotClassGetParamsHTTPVersion{radar.HTTPLocationBotClassGetParamsHTTPVersionHttPv1, radar.HTTPLocationBotClassGetParamsHTTPVersionHttPv2, radar.HTTPLocationBotClassGetParamsHTTPVersionHttPv3}),
			IPVersion:     khulnasoft.F([]radar.HTTPLocationBotClassGetParamsIPVersion{radar.HTTPLocationBotClassGetParamsIPVersionIPv4, radar.HTTPLocationBotClassGetParamsIPVersionIPv6}),
			Limit:         khulnasoft.F(int64(5)),
			Location:      khulnasoft.F([]string{"string", "string", "string"}),
			Name:          khulnasoft.F([]string{"string", "string", "string"}),
			OS:            khulnasoft.F([]radar.HTTPLocationBotClassGetParamsOS{radar.HTTPLocationBotClassGetParamsOSWindows, radar.HTTPLocationBotClassGetParamsOSMacosx, radar.HTTPLocationBotClassGetParamsOSIos}),
			TLSVersion:    khulnasoft.F([]radar.HTTPLocationBotClassGetParamsTLSVersion{radar.HTTPLocationBotClassGetParamsTLSVersionTlSv1_0, radar.HTTPLocationBotClassGetParamsTLSVersionTlSv1_1, radar.HTTPLocationBotClassGetParamsTLSVersionTlSv1_2}),
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
