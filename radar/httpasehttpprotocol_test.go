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

func TestHTTPAseHTTPProtocolGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.HTTPProtocol.Get(
		context.TODO(),
		radar.HTTPAseHTTPProtocolGetParamsHTTPProtocolHTTP,
		radar.HTTPAseHTTPProtocolGetParams{
			ASN:           khulnasoft.F([]string{"string", "string", "string"}),
			BotClass:      khulnasoft.F([]radar.HTTPAseHTTPProtocolGetParamsBotClass{radar.HTTPAseHTTPProtocolGetParamsBotClassLikelyAutomated, radar.HTTPAseHTTPProtocolGetParamsBotClassLikelyHuman}),
			BrowserFamily: khulnasoft.F([]radar.HTTPAseHTTPProtocolGetParamsBrowserFamily{radar.HTTPAseHTTPProtocolGetParamsBrowserFamilyChrome, radar.HTTPAseHTTPProtocolGetParamsBrowserFamilyEdge, radar.HTTPAseHTTPProtocolGetParamsBrowserFamilyFirefox}),
			Continent:     khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    khulnasoft.F([]radar.HTTPAseHTTPProtocolGetParamsDeviceType{radar.HTTPAseHTTPProtocolGetParamsDeviceTypeDesktop, radar.HTTPAseHTTPProtocolGetParamsDeviceTypeMobile, radar.HTTPAseHTTPProtocolGetParamsDeviceTypeOther}),
			Format:        khulnasoft.F(radar.HTTPAseHTTPProtocolGetParamsFormatJson),
			IPVersion:     khulnasoft.F([]radar.HTTPAseHTTPProtocolGetParamsIPVersion{radar.HTTPAseHTTPProtocolGetParamsIPVersionIPv4, radar.HTTPAseHTTPProtocolGetParamsIPVersionIPv6}),
			Limit:         khulnasoft.F(int64(5)),
			Location:      khulnasoft.F([]string{"string", "string", "string"}),
			Name:          khulnasoft.F([]string{"string", "string", "string"}),
			OS:            khulnasoft.F([]radar.HTTPAseHTTPProtocolGetParamsOS{radar.HTTPAseHTTPProtocolGetParamsOSWindows, radar.HTTPAseHTTPProtocolGetParamsOSMacosx, radar.HTTPAseHTTPProtocolGetParamsOSIos}),
			TLSVersion:    khulnasoft.F([]radar.HTTPAseHTTPProtocolGetParamsTLSVersion{radar.HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_0, radar.HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_1, radar.HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_2}),
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