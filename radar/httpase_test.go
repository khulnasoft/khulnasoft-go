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

func TestHTTPAseGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.Get(context.TODO(), radar.HTTPAseGetParams{
		ASN:           khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:      khulnasoft.F([]radar.HTTPAseGetParamsBotClass{radar.HTTPAseGetParamsBotClassLikelyAutomated, radar.HTTPAseGetParamsBotClassLikelyHuman}),
		BrowserFamily: khulnasoft.F([]radar.HTTPAseGetParamsBrowserFamily{radar.HTTPAseGetParamsBrowserFamilyChrome, radar.HTTPAseGetParamsBrowserFamilyEdge, radar.HTTPAseGetParamsBrowserFamilyFirefox}),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    khulnasoft.F([]radar.HTTPAseGetParamsDeviceType{radar.HTTPAseGetParamsDeviceTypeDesktop, radar.HTTPAseGetParamsDeviceTypeMobile, radar.HTTPAseGetParamsDeviceTypeOther}),
		Format:        khulnasoft.F(radar.HTTPAseGetParamsFormatJson),
		HTTPProtocol:  khulnasoft.F([]radar.HTTPAseGetParamsHTTPProtocol{radar.HTTPAseGetParamsHTTPProtocolHTTP, radar.HTTPAseGetParamsHTTPProtocolHTTPS}),
		HTTPVersion:   khulnasoft.F([]radar.HTTPAseGetParamsHTTPVersion{radar.HTTPAseGetParamsHTTPVersionHttPv1, radar.HTTPAseGetParamsHTTPVersionHttPv2, radar.HTTPAseGetParamsHTTPVersionHttPv3}),
		IPVersion:     khulnasoft.F([]radar.HTTPAseGetParamsIPVersion{radar.HTTPAseGetParamsIPVersionIPv4, radar.HTTPAseGetParamsIPVersionIPv6}),
		Limit:         khulnasoft.F(int64(5)),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		OS:            khulnasoft.F([]radar.HTTPAseGetParamsOS{radar.HTTPAseGetParamsOSWindows, radar.HTTPAseGetParamsOSMacosx, radar.HTTPAseGetParamsOSIos}),
		TLSVersion:    khulnasoft.F([]radar.HTTPAseGetParamsTLSVersion{radar.HTTPAseGetParamsTLSVersionTlSv1_0, radar.HTTPAseGetParamsTLSVersionTlSv1_1, radar.HTTPAseGetParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
