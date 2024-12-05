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

func TestHTTPTopBrowserWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Browser(context.TODO(), radar.HTTPTopBrowserParams{
		ASN:           khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:      khulnasoft.F([]radar.HTTPTopBrowserParamsBotClass{radar.HTTPTopBrowserParamsBotClassLikelyAutomated, radar.HTTPTopBrowserParamsBotClassLikelyHuman}),
		BrowserFamily: khulnasoft.F([]radar.HTTPTopBrowserParamsBrowserFamily{radar.HTTPTopBrowserParamsBrowserFamilyChrome, radar.HTTPTopBrowserParamsBrowserFamilyEdge, radar.HTTPTopBrowserParamsBrowserFamilyFirefox}),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    khulnasoft.F([]radar.HTTPTopBrowserParamsDeviceType{radar.HTTPTopBrowserParamsDeviceTypeDesktop, radar.HTTPTopBrowserParamsDeviceTypeMobile, radar.HTTPTopBrowserParamsDeviceTypeOther}),
		Format:        khulnasoft.F(radar.HTTPTopBrowserParamsFormatJson),
		HTTPProtocol:  khulnasoft.F([]radar.HTTPTopBrowserParamsHTTPProtocol{radar.HTTPTopBrowserParamsHTTPProtocolHTTP, radar.HTTPTopBrowserParamsHTTPProtocolHTTPS}),
		HTTPVersion:   khulnasoft.F([]radar.HTTPTopBrowserParamsHTTPVersion{radar.HTTPTopBrowserParamsHTTPVersionHttPv1, radar.HTTPTopBrowserParamsHTTPVersionHttPv2, radar.HTTPTopBrowserParamsHTTPVersionHttPv3}),
		IPVersion:     khulnasoft.F([]radar.HTTPTopBrowserParamsIPVersion{radar.HTTPTopBrowserParamsIPVersionIPv4, radar.HTTPTopBrowserParamsIPVersionIPv6}),
		Limit:         khulnasoft.F(int64(5)),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		OS:            khulnasoft.F([]radar.HTTPTopBrowserParamsOS{radar.HTTPTopBrowserParamsOSWindows, radar.HTTPTopBrowserParamsOSMacosx, radar.HTTPTopBrowserParamsOSIos}),
		TLSVersion:    khulnasoft.F([]radar.HTTPTopBrowserParamsTLSVersion{radar.HTTPTopBrowserParamsTLSVersionTlSv1_0, radar.HTTPTopBrowserParamsTLSVersionTlSv1_1, radar.HTTPTopBrowserParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTopBrowserFamilyWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.BrowserFamily(context.TODO(), radar.HTTPTopBrowserFamilyParams{
		ASN:           khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:      khulnasoft.F([]radar.HTTPTopBrowserFamilyParamsBotClass{radar.HTTPTopBrowserFamilyParamsBotClassLikelyAutomated, radar.HTTPTopBrowserFamilyParamsBotClassLikelyHuman}),
		BrowserFamily: khulnasoft.F([]radar.HTTPTopBrowserFamilyParamsBrowserFamily{radar.HTTPTopBrowserFamilyParamsBrowserFamilyChrome, radar.HTTPTopBrowserFamilyParamsBrowserFamilyEdge, radar.HTTPTopBrowserFamilyParamsBrowserFamilyFirefox}),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    khulnasoft.F([]radar.HTTPTopBrowserFamilyParamsDeviceType{radar.HTTPTopBrowserFamilyParamsDeviceTypeDesktop, radar.HTTPTopBrowserFamilyParamsDeviceTypeMobile, radar.HTTPTopBrowserFamilyParamsDeviceTypeOther}),
		Format:        khulnasoft.F(radar.HTTPTopBrowserFamilyParamsFormatJson),
		HTTPProtocol:  khulnasoft.F([]radar.HTTPTopBrowserFamilyParamsHTTPProtocol{radar.HTTPTopBrowserFamilyParamsHTTPProtocolHTTP, radar.HTTPTopBrowserFamilyParamsHTTPProtocolHTTPS}),
		HTTPVersion:   khulnasoft.F([]radar.HTTPTopBrowserFamilyParamsHTTPVersion{radar.HTTPTopBrowserFamilyParamsHTTPVersionHttPv1, radar.HTTPTopBrowserFamilyParamsHTTPVersionHttPv2, radar.HTTPTopBrowserFamilyParamsHTTPVersionHttPv3}),
		IPVersion:     khulnasoft.F([]radar.HTTPTopBrowserFamilyParamsIPVersion{radar.HTTPTopBrowserFamilyParamsIPVersionIPv4, radar.HTTPTopBrowserFamilyParamsIPVersionIPv6}),
		Limit:         khulnasoft.F(int64(5)),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		OS:            khulnasoft.F([]radar.HTTPTopBrowserFamilyParamsOS{radar.HTTPTopBrowserFamilyParamsOSWindows, radar.HTTPTopBrowserFamilyParamsOSMacosx, radar.HTTPTopBrowserFamilyParamsOSIos}),
		TLSVersion:    khulnasoft.F([]radar.HTTPTopBrowserFamilyParamsTLSVersion{radar.HTTPTopBrowserFamilyParamsTLSVersionTlSv1_0, radar.HTTPTopBrowserFamilyParamsTLSVersionTlSv1_1, radar.HTTPTopBrowserFamilyParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
