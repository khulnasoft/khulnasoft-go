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

func TestHTTPAseHTTPMethodGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.HTTPMethod.Get(
		context.TODO(),
		radar.HTTPAseHTTPMethodGetParamsHTTPVersionHttPv1,
		radar.HTTPAseHTTPMethodGetParams{
			ASN:           khulnasoft.F([]string{"string", "string", "string"}),
			BotClass:      khulnasoft.F([]radar.HTTPAseHTTPMethodGetParamsBotClass{radar.HTTPAseHTTPMethodGetParamsBotClassLikelyAutomated, radar.HTTPAseHTTPMethodGetParamsBotClassLikelyHuman}),
			BrowserFamily: khulnasoft.F([]radar.HTTPAseHTTPMethodGetParamsBrowserFamily{radar.HTTPAseHTTPMethodGetParamsBrowserFamilyChrome, radar.HTTPAseHTTPMethodGetParamsBrowserFamilyEdge, radar.HTTPAseHTTPMethodGetParamsBrowserFamilyFirefox}),
			Continent:     khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    khulnasoft.F([]radar.HTTPAseHTTPMethodGetParamsDeviceType{radar.HTTPAseHTTPMethodGetParamsDeviceTypeDesktop, radar.HTTPAseHTTPMethodGetParamsDeviceTypeMobile, radar.HTTPAseHTTPMethodGetParamsDeviceTypeOther}),
			Format:        khulnasoft.F(radar.HTTPAseHTTPMethodGetParamsFormatJson),
			HTTPProtocol:  khulnasoft.F([]radar.HTTPAseHTTPMethodGetParamsHTTPProtocol{radar.HTTPAseHTTPMethodGetParamsHTTPProtocolHTTP, radar.HTTPAseHTTPMethodGetParamsHTTPProtocolHTTPS}),
			IPVersion:     khulnasoft.F([]radar.HTTPAseHTTPMethodGetParamsIPVersion{radar.HTTPAseHTTPMethodGetParamsIPVersionIPv4, radar.HTTPAseHTTPMethodGetParamsIPVersionIPv6}),
			Limit:         khulnasoft.F(int64(5)),
			Location:      khulnasoft.F([]string{"string", "string", "string"}),
			Name:          khulnasoft.F([]string{"string", "string", "string"}),
			OS:            khulnasoft.F([]radar.HTTPAseHTTPMethodGetParamsOS{radar.HTTPAseHTTPMethodGetParamsOSWindows, radar.HTTPAseHTTPMethodGetParamsOSMacosx, radar.HTTPAseHTTPMethodGetParamsOSIos}),
			TLSVersion:    khulnasoft.F([]radar.HTTPAseHTTPMethodGetParamsTLSVersion{radar.HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_0, radar.HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_1, radar.HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_2}),
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
