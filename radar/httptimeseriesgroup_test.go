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

func TestHTTPTimeseriesGroupBotClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.BotClass(context.TODO(), radar.HTTPTimeseriesGroupBotClassParams{
		AggInterval:  khulnasoft.F(radar.HTTPTimeseriesGroupBotClassParamsAggInterval15m),
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPTimeseriesGroupBotClassParamsDeviceType{radar.HTTPTimeseriesGroupBotClassParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupBotClassParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupBotClassParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPTimeseriesGroupBotClassParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPTimeseriesGroupBotClassParamsHTTPProtocol{radar.HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPTimeseriesGroupBotClassParamsHTTPVersion{radar.HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv3}),
		IPVersion:    khulnasoft.F([]radar.HTTPTimeseriesGroupBotClassParamsIPVersion{radar.HTTPTimeseriesGroupBotClassParamsIPVersionIPv4, radar.HTTPTimeseriesGroupBotClassParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPTimeseriesGroupBotClassParamsOS{radar.HTTPTimeseriesGroupBotClassParamsOSWindows, radar.HTTPTimeseriesGroupBotClassParamsOSMacosx, radar.HTTPTimeseriesGroupBotClassParamsOSIos}),
		TLSVersion:   khulnasoft.F([]radar.HTTPTimeseriesGroupBotClassParamsTLSVersion{radar.HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupBrowserWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.Browser(context.TODO(), radar.HTTPTimeseriesGroupBrowserParams{
		AggInterval:   khulnasoft.F(radar.HTTPTimeseriesGroupBrowserParamsAggInterval15m),
		ASN:           khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:      khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserParamsBotClass{radar.HTTPTimeseriesGroupBrowserParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupBrowserParamsBotClassLikelyHuman}),
		Continent:     khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:       khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserParamsDeviceType{radar.HTTPTimeseriesGroupBrowserParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupBrowserParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupBrowserParamsDeviceTypeOther}),
		Format:        khulnasoft.F(radar.HTTPTimeseriesGroupBrowserParamsFormatJson),
		HTTPProtocol:  khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserParamsHTTPProtocol{radar.HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTPS}),
		HTTPVersion:   khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserParamsHTTPVersion{radar.HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv3}),
		IPVersion:     khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserParamsIPVersion{radar.HTTPTimeseriesGroupBrowserParamsIPVersionIPv4, radar.HTTPTimeseriesGroupBrowserParamsIPVersionIPv6}),
		LimitPerGroup: khulnasoft.F(int64(4)),
		Location:      khulnasoft.F([]string{"string", "string", "string"}),
		Name:          khulnasoft.F([]string{"string", "string", "string"}),
		OS:            khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserParamsOS{radar.HTTPTimeseriesGroupBrowserParamsOSWindows, radar.HTTPTimeseriesGroupBrowserParamsOSMacosx, radar.HTTPTimeseriesGroupBrowserParamsOSIos}),
		TLSVersion:    khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserParamsTLSVersion{radar.HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupBrowserFamilyWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.BrowserFamily(context.TODO(), radar.HTTPTimeseriesGroupBrowserFamilyParams{
		AggInterval:  khulnasoft.F(radar.HTTPTimeseriesGroupBrowserFamilyParamsAggInterval15m),
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsBotClass{radar.HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsDeviceType{radar.HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPTimeseriesGroupBrowserFamilyParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol{radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion{radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv3}),
		IPVersion:    khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsIPVersion{radar.HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv4, radar.HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsOS{radar.HTTPTimeseriesGroupBrowserFamilyParamsOSWindows, radar.HTTPTimeseriesGroupBrowserFamilyParamsOSMacosx, radar.HTTPTimeseriesGroupBrowserFamilyParamsOSIos}),
		TLSVersion:   khulnasoft.F([]radar.HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion{radar.HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupDeviceTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.DeviceType(context.TODO(), radar.HTTPTimeseriesGroupDeviceTypeParams{
		AggInterval:  khulnasoft.F(radar.HTTPTimeseriesGroupDeviceTypeParamsAggInterval15m),
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsBotClass{radar.HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:       khulnasoft.F(radar.HTTPTimeseriesGroupDeviceTypeParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol{radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion{radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv3}),
		IPVersion:    khulnasoft.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsIPVersion{radar.HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv4, radar.HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsOS{radar.HTTPTimeseriesGroupDeviceTypeParamsOSWindows, radar.HTTPTimeseriesGroupDeviceTypeParamsOSMacosx, radar.HTTPTimeseriesGroupDeviceTypeParamsOSIos}),
		TLSVersion:   khulnasoft.F([]radar.HTTPTimeseriesGroupDeviceTypeParamsTLSVersion{radar.HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupHTTPProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.HTTPProtocol(context.TODO(), radar.HTTPTimeseriesGroupHTTPProtocolParams{
		AggInterval: khulnasoft.F(radar.HTTPTimeseriesGroupHTTPProtocolParamsAggInterval15m),
		ASN:         khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:    khulnasoft.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsBotClass{radar.HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyHuman}),
		Continent:   khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:  khulnasoft.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsDeviceType{radar.HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeOther}),
		Format:      khulnasoft.F(radar.HTTPTimeseriesGroupHTTPProtocolParamsFormatJson),
		HTTPVersion: khulnasoft.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion{radar.HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv3}),
		IPVersion:   khulnasoft.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsIPVersion{radar.HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv4, radar.HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv6}),
		Location:    khulnasoft.F([]string{"string", "string", "string"}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		OS:          khulnasoft.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsOS{radar.HTTPTimeseriesGroupHTTPProtocolParamsOSWindows, radar.HTTPTimeseriesGroupHTTPProtocolParamsOSMacosx, radar.HTTPTimeseriesGroupHTTPProtocolParamsOSIos}),
		TLSVersion:  khulnasoft.F([]radar.HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion{radar.HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.HTTPVersion(context.TODO(), radar.HTTPTimeseriesGroupHTTPVersionParams{
		AggInterval:  khulnasoft.F(radar.HTTPTimeseriesGroupHTTPVersionParamsAggInterval15m),
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsBotClass{radar.HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsDeviceType{radar.HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPTimeseriesGroupHTTPVersionParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol{radar.HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTPS}),
		IPVersion:    khulnasoft.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsIPVersion{radar.HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv4, radar.HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsOS{radar.HTTPTimeseriesGroupHTTPVersionParamsOSWindows, radar.HTTPTimeseriesGroupHTTPVersionParamsOSMacosx, radar.HTTPTimeseriesGroupHTTPVersionParamsOSIos}),
		TLSVersion:   khulnasoft.F([]radar.HTTPTimeseriesGroupHTTPVersionParamsTLSVersion{radar.HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.IPVersion(context.TODO(), radar.HTTPTimeseriesGroupIPVersionParams{
		AggInterval:  khulnasoft.F(radar.HTTPTimeseriesGroupIPVersionParamsAggInterval15m),
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPTimeseriesGroupIPVersionParamsBotClass{radar.HTTPTimeseriesGroupIPVersionParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupIPVersionParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPTimeseriesGroupIPVersionParamsDeviceType{radar.HTTPTimeseriesGroupIPVersionParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupIPVersionParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupIPVersionParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPTimeseriesGroupIPVersionParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPTimeseriesGroupIPVersionParamsHTTPProtocol{radar.HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPTimeseriesGroupIPVersionParamsHTTPVersion{radar.HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv3}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPTimeseriesGroupIPVersionParamsOS{radar.HTTPTimeseriesGroupIPVersionParamsOSWindows, radar.HTTPTimeseriesGroupIPVersionParamsOSMacosx, radar.HTTPTimeseriesGroupIPVersionParamsOSIos}),
		TLSVersion:   khulnasoft.F([]radar.HTTPTimeseriesGroupIPVersionParamsTLSVersion{radar.HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupOSWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.OS(context.TODO(), radar.HTTPTimeseriesGroupOSParams{
		AggInterval:  khulnasoft.F(radar.HTTPTimeseriesGroupOSParamsAggInterval15m),
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPTimeseriesGroupOSParamsBotClass{radar.HTTPTimeseriesGroupOSParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupOSParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPTimeseriesGroupOSParamsDeviceType{radar.HTTPTimeseriesGroupOSParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupOSParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupOSParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPTimeseriesGroupOSParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPTimeseriesGroupOSParamsHTTPProtocol{radar.HTTPTimeseriesGroupOSParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupOSParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPTimeseriesGroupOSParamsHTTPVersion{radar.HTTPTimeseriesGroupOSParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupOSParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupOSParamsHTTPVersionHttPv3}),
		IPVersion:    khulnasoft.F([]radar.HTTPTimeseriesGroupOSParamsIPVersion{radar.HTTPTimeseriesGroupOSParamsIPVersionIPv4, radar.HTTPTimeseriesGroupOSParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		TLSVersion:   khulnasoft.F([]radar.HTTPTimeseriesGroupOSParamsTLSVersion{radar.HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupPostQuantumWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.PostQuantum(context.TODO(), radar.HTTPTimeseriesGroupPostQuantumParams{
		AggInterval:  khulnasoft.F(radar.HTTPTimeseriesGroupPostQuantumParamsAggInterval15m),
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPTimeseriesGroupPostQuantumParamsBotClass{radar.HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPTimeseriesGroupPostQuantumParamsDeviceType{radar.HTTPTimeseriesGroupPostQuantumParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupPostQuantumParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupPostQuantumParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPTimeseriesGroupPostQuantumParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol{radar.HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPTimeseriesGroupPostQuantumParamsHTTPVersion{radar.HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv3}),
		IPVersion:    khulnasoft.F([]radar.HTTPTimeseriesGroupPostQuantumParamsIPVersion{radar.HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv4, radar.HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPTimeseriesGroupPostQuantumParamsOS{radar.HTTPTimeseriesGroupPostQuantumParamsOSWindows, radar.HTTPTimeseriesGroupPostQuantumParamsOSMacosx, radar.HTTPTimeseriesGroupPostQuantumParamsOSIos}),
		TLSVersion:   khulnasoft.F([]radar.HTTPTimeseriesGroupPostQuantumParamsTLSVersion{radar.HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_0, radar.HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_1, radar.HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTimeseriesGroupTLSVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.TLSVersion(context.TODO(), radar.HTTPTimeseriesGroupTLSVersionParams{
		AggInterval:  khulnasoft.F(radar.HTTPTimeseriesGroupTLSVersionParamsAggInterval15m),
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPTimeseriesGroupTLSVersionParamsBotClass{radar.HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyAutomated, radar.HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPTimeseriesGroupTLSVersionParamsDeviceType{radar.HTTPTimeseriesGroupTLSVersionParamsDeviceTypeDesktop, radar.HTTPTimeseriesGroupTLSVersionParamsDeviceTypeMobile, radar.HTTPTimeseriesGroupTLSVersionParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPTimeseriesGroupTLSVersionParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol{radar.HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTP, radar.HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPTimeseriesGroupTLSVersionParamsHTTPVersion{radar.HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv1, radar.HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv2, radar.HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv3}),
		IPVersion:    khulnasoft.F([]radar.HTTPTimeseriesGroupTLSVersionParamsIPVersion{radar.HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv4, radar.HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPTimeseriesGroupTLSVersionParamsOS{radar.HTTPTimeseriesGroupTLSVersionParamsOSWindows, radar.HTTPTimeseriesGroupTLSVersionParamsOSMacosx, radar.HTTPTimeseriesGroupTLSVersionParamsOSIos}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
