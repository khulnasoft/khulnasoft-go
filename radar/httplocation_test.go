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

func TestHTTPLocationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.Get(context.TODO(), radar.HTTPLocationGetParams{
		ASN:           khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:      khulnasoft.F([]radar.HTTPLocationGetParamsBotClass{radar.HTTPLocationGetParamsBotClassLikelyAutomated, radar.HTTPLocationGetParamsBotClassLikelyHuman}),
		BrowserFamily: khulnasoft.F([]radar.HTTPLocationGetParamsBrowserFamily{radar.HTTPLocationGetParamsBrowserFamilyChrome, radar.HTTPLocationGetParamsBrowserFamilyEdge, radar.HTTPLocationGetParamsBrowserFamilyFirefox}),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    khulnasoft.F([]radar.HTTPLocationGetParamsDeviceType{radar.HTTPLocationGetParamsDeviceTypeDesktop, radar.HTTPLocationGetParamsDeviceTypeMobile, radar.HTTPLocationGetParamsDeviceTypeOther}),
		Format:        khulnasoft.F(radar.HTTPLocationGetParamsFormatJson),
		HTTPProtocol:  khulnasoft.F([]radar.HTTPLocationGetParamsHTTPProtocol{radar.HTTPLocationGetParamsHTTPProtocolHTTP, radar.HTTPLocationGetParamsHTTPProtocolHTTPS}),
		HTTPVersion:   khulnasoft.F([]radar.HTTPLocationGetParamsHTTPVersion{radar.HTTPLocationGetParamsHTTPVersionHttPv1, radar.HTTPLocationGetParamsHTTPVersionHttPv2, radar.HTTPLocationGetParamsHTTPVersionHttPv3}),
		IPVersion:     khulnasoft.F([]radar.HTTPLocationGetParamsIPVersion{radar.HTTPLocationGetParamsIPVersionIPv4, radar.HTTPLocationGetParamsIPVersionIPv6}),
		Limit:         khulnasoft.F(int64(5)),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		OS:            khulnasoft.F([]radar.HTTPLocationGetParamsOS{radar.HTTPLocationGetParamsOSWindows, radar.HTTPLocationGetParamsOSMacosx, radar.HTTPLocationGetParamsOSIos}),
		TLSVersion:    khulnasoft.F([]radar.HTTPLocationGetParamsTLSVersion{radar.HTTPLocationGetParamsTLSVersionTlSv1_0, radar.HTTPLocationGetParamsTLSVersionTlSv1_1, radar.HTTPLocationGetParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
