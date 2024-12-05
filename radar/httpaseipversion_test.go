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

func TestHTTPAseIPVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.IPVersion.Get(
		context.TODO(),
		radar.HTTPAseIPVersionGetParamsIPVersionIPv4,
		radar.HTTPAseIPVersionGetParams{
			ASN:           khulnasoft.F([]string{"string", "string", "string"}),
			BotClass:      khulnasoft.F([]radar.HTTPAseIPVersionGetParamsBotClass{radar.HTTPAseIPVersionGetParamsBotClassLikelyAutomated, radar.HTTPAseIPVersionGetParamsBotClassLikelyHuman}),
			BrowserFamily: khulnasoft.F([]radar.HTTPAseIPVersionGetParamsBrowserFamily{radar.HTTPAseIPVersionGetParamsBrowserFamilyChrome, radar.HTTPAseIPVersionGetParamsBrowserFamilyEdge, radar.HTTPAseIPVersionGetParamsBrowserFamilyFirefox}),
			Continent:     khulnasoft.F([]string{"string", "string", "string"}),
			DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    khulnasoft.F([]radar.HTTPAseIPVersionGetParamsDeviceType{radar.HTTPAseIPVersionGetParamsDeviceTypeDesktop, radar.HTTPAseIPVersionGetParamsDeviceTypeMobile, radar.HTTPAseIPVersionGetParamsDeviceTypeOther}),
			Format:        khulnasoft.F(radar.HTTPAseIPVersionGetParamsFormatJson),
			HTTPProtocol:  khulnasoft.F([]radar.HTTPAseIPVersionGetParamsHTTPProtocol{radar.HTTPAseIPVersionGetParamsHTTPProtocolHTTP, radar.HTTPAseIPVersionGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   khulnasoft.F([]radar.HTTPAseIPVersionGetParamsHTTPVersion{radar.HTTPAseIPVersionGetParamsHTTPVersionHttPv1, radar.HTTPAseIPVersionGetParamsHTTPVersionHttPv2, radar.HTTPAseIPVersionGetParamsHTTPVersionHttPv3}),
			Limit:         khulnasoft.F(int64(5)),
			Location:      khulnasoft.F([]string{"string", "string", "string"}),
			Name:          khulnasoft.F([]string{"string", "string", "string"}),
			OS:            khulnasoft.F([]radar.HTTPAseIPVersionGetParamsOS{radar.HTTPAseIPVersionGetParamsOSWindows, radar.HTTPAseIPVersionGetParamsOSMacosx, radar.HTTPAseIPVersionGetParamsOSIos}),
			TLSVersion:    khulnasoft.F([]radar.HTTPAseIPVersionGetParamsTLSVersion{radar.HTTPAseIPVersionGetParamsTLSVersionTlSv1_0, radar.HTTPAseIPVersionGetParamsTLSVersionTlSv1_1, radar.HTTPAseIPVersionGetParamsTLSVersionTlSv1_2}),
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
