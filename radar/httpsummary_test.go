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

func TestHTTPSummaryBotClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.BotClass(context.TODO(), radar.HTTPSummaryBotClassParams{
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPSummaryBotClassParamsDeviceType{radar.HTTPSummaryBotClassParamsDeviceTypeDesktop, radar.HTTPSummaryBotClassParamsDeviceTypeMobile, radar.HTTPSummaryBotClassParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPSummaryBotClassParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPSummaryBotClassParamsHTTPProtocol{radar.HTTPSummaryBotClassParamsHTTPProtocolHTTP, radar.HTTPSummaryBotClassParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPSummaryBotClassParamsHTTPVersion{radar.HTTPSummaryBotClassParamsHTTPVersionHttPv1, radar.HTTPSummaryBotClassParamsHTTPVersionHttPv2, radar.HTTPSummaryBotClassParamsHTTPVersionHttPv3}),
		IPVersion:    khulnasoft.F([]radar.HTTPSummaryBotClassParamsIPVersion{radar.HTTPSummaryBotClassParamsIPVersionIPv4, radar.HTTPSummaryBotClassParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPSummaryBotClassParamsOS{radar.HTTPSummaryBotClassParamsOSWindows, radar.HTTPSummaryBotClassParamsOSMacosx, radar.HTTPSummaryBotClassParamsOSIos}),
		TLSVersion:   khulnasoft.F([]radar.HTTPSummaryBotClassParamsTLSVersion{radar.HTTPSummaryBotClassParamsTLSVersionTlSv1_0, radar.HTTPSummaryBotClassParamsTLSVersionTlSv1_1, radar.HTTPSummaryBotClassParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryDeviceTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.DeviceType(context.TODO(), radar.HTTPSummaryDeviceTypeParams{
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPSummaryDeviceTypeParamsBotClass{radar.HTTPSummaryDeviceTypeParamsBotClassLikelyAutomated, radar.HTTPSummaryDeviceTypeParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:       khulnasoft.F(radar.HTTPSummaryDeviceTypeParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPSummaryDeviceTypeParamsHTTPProtocol{radar.HTTPSummaryDeviceTypeParamsHTTPProtocolHTTP, radar.HTTPSummaryDeviceTypeParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPSummaryDeviceTypeParamsHTTPVersion{radar.HTTPSummaryDeviceTypeParamsHTTPVersionHttPv1, radar.HTTPSummaryDeviceTypeParamsHTTPVersionHttPv2, radar.HTTPSummaryDeviceTypeParamsHTTPVersionHttPv3}),
		IPVersion:    khulnasoft.F([]radar.HTTPSummaryDeviceTypeParamsIPVersion{radar.HTTPSummaryDeviceTypeParamsIPVersionIPv4, radar.HTTPSummaryDeviceTypeParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPSummaryDeviceTypeParamsOS{radar.HTTPSummaryDeviceTypeParamsOSWindows, radar.HTTPSummaryDeviceTypeParamsOSMacosx, radar.HTTPSummaryDeviceTypeParamsOSIos}),
		TLSVersion:   khulnasoft.F([]radar.HTTPSummaryDeviceTypeParamsTLSVersion{radar.HTTPSummaryDeviceTypeParamsTLSVersionTlSv1_0, radar.HTTPSummaryDeviceTypeParamsTLSVersionTlSv1_1, radar.HTTPSummaryDeviceTypeParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryHTTPProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.HTTPProtocol(context.TODO(), radar.HTTPSummaryHTTPProtocolParams{
		ASN:         khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:    khulnasoft.F([]radar.HTTPSummaryHTTPProtocolParamsBotClass{radar.HTTPSummaryHTTPProtocolParamsBotClassLikelyAutomated, radar.HTTPSummaryHTTPProtocolParamsBotClassLikelyHuman}),
		Continent:   khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:  khulnasoft.F([]radar.HTTPSummaryHTTPProtocolParamsDeviceType{radar.HTTPSummaryHTTPProtocolParamsDeviceTypeDesktop, radar.HTTPSummaryHTTPProtocolParamsDeviceTypeMobile, radar.HTTPSummaryHTTPProtocolParamsDeviceTypeOther}),
		Format:      khulnasoft.F(radar.HTTPSummaryHTTPProtocolParamsFormatJson),
		HTTPVersion: khulnasoft.F([]radar.HTTPSummaryHTTPProtocolParamsHTTPVersion{radar.HTTPSummaryHTTPProtocolParamsHTTPVersionHttPv1, radar.HTTPSummaryHTTPProtocolParamsHTTPVersionHttPv2, radar.HTTPSummaryHTTPProtocolParamsHTTPVersionHttPv3}),
		IPVersion:   khulnasoft.F([]radar.HTTPSummaryHTTPProtocolParamsIPVersion{radar.HTTPSummaryHTTPProtocolParamsIPVersionIPv4, radar.HTTPSummaryHTTPProtocolParamsIPVersionIPv6}),
		Location:    khulnasoft.F([]string{"string", "string", "string"}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		OS:          khulnasoft.F([]radar.HTTPSummaryHTTPProtocolParamsOS{radar.HTTPSummaryHTTPProtocolParamsOSWindows, radar.HTTPSummaryHTTPProtocolParamsOSMacosx, radar.HTTPSummaryHTTPProtocolParamsOSIos}),
		TLSVersion:  khulnasoft.F([]radar.HTTPSummaryHTTPProtocolParamsTLSVersion{radar.HTTPSummaryHTTPProtocolParamsTLSVersionTlSv1_0, radar.HTTPSummaryHTTPProtocolParamsTLSVersionTlSv1_1, radar.HTTPSummaryHTTPProtocolParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.HTTPVersion(context.TODO(), radar.HTTPSummaryHTTPVersionParams{
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPSummaryHTTPVersionParamsBotClass{radar.HTTPSummaryHTTPVersionParamsBotClassLikelyAutomated, radar.HTTPSummaryHTTPVersionParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPSummaryHTTPVersionParamsDeviceType{radar.HTTPSummaryHTTPVersionParamsDeviceTypeDesktop, radar.HTTPSummaryHTTPVersionParamsDeviceTypeMobile, radar.HTTPSummaryHTTPVersionParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPSummaryHTTPVersionParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPSummaryHTTPVersionParamsHTTPProtocol{radar.HTTPSummaryHTTPVersionParamsHTTPProtocolHTTP, radar.HTTPSummaryHTTPVersionParamsHTTPProtocolHTTPS}),
		IPVersion:    khulnasoft.F([]radar.HTTPSummaryHTTPVersionParamsIPVersion{radar.HTTPSummaryHTTPVersionParamsIPVersionIPv4, radar.HTTPSummaryHTTPVersionParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPSummaryHTTPVersionParamsOS{radar.HTTPSummaryHTTPVersionParamsOSWindows, radar.HTTPSummaryHTTPVersionParamsOSMacosx, radar.HTTPSummaryHTTPVersionParamsOSIos}),
		TLSVersion:   khulnasoft.F([]radar.HTTPSummaryHTTPVersionParamsTLSVersion{radar.HTTPSummaryHTTPVersionParamsTLSVersionTlSv1_0, radar.HTTPSummaryHTTPVersionParamsTLSVersionTlSv1_1, radar.HTTPSummaryHTTPVersionParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.IPVersion(context.TODO(), radar.HTTPSummaryIPVersionParams{
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPSummaryIPVersionParamsBotClass{radar.HTTPSummaryIPVersionParamsBotClassLikelyAutomated, radar.HTTPSummaryIPVersionParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPSummaryIPVersionParamsDeviceType{radar.HTTPSummaryIPVersionParamsDeviceTypeDesktop, radar.HTTPSummaryIPVersionParamsDeviceTypeMobile, radar.HTTPSummaryIPVersionParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPSummaryIPVersionParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPSummaryIPVersionParamsHTTPProtocol{radar.HTTPSummaryIPVersionParamsHTTPProtocolHTTP, radar.HTTPSummaryIPVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPSummaryIPVersionParamsHTTPVersion{radar.HTTPSummaryIPVersionParamsHTTPVersionHttPv1, radar.HTTPSummaryIPVersionParamsHTTPVersionHttPv2, radar.HTTPSummaryIPVersionParamsHTTPVersionHttPv3}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPSummaryIPVersionParamsOS{radar.HTTPSummaryIPVersionParamsOSWindows, radar.HTTPSummaryIPVersionParamsOSMacosx, radar.HTTPSummaryIPVersionParamsOSIos}),
		TLSVersion:   khulnasoft.F([]radar.HTTPSummaryIPVersionParamsTLSVersion{radar.HTTPSummaryIPVersionParamsTLSVersionTlSv1_0, radar.HTTPSummaryIPVersionParamsTLSVersionTlSv1_1, radar.HTTPSummaryIPVersionParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryOSWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.OS(context.TODO(), radar.HTTPSummaryOSParams{
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPSummaryOSParamsBotClass{radar.HTTPSummaryOSParamsBotClassLikelyAutomated, radar.HTTPSummaryOSParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPSummaryOSParamsDeviceType{radar.HTTPSummaryOSParamsDeviceTypeDesktop, radar.HTTPSummaryOSParamsDeviceTypeMobile, radar.HTTPSummaryOSParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPSummaryOSParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPSummaryOSParamsHTTPProtocol{radar.HTTPSummaryOSParamsHTTPProtocolHTTP, radar.HTTPSummaryOSParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPSummaryOSParamsHTTPVersion{radar.HTTPSummaryOSParamsHTTPVersionHttPv1, radar.HTTPSummaryOSParamsHTTPVersionHttPv2, radar.HTTPSummaryOSParamsHTTPVersionHttPv3}),
		IPVersion:    khulnasoft.F([]radar.HTTPSummaryOSParamsIPVersion{radar.HTTPSummaryOSParamsIPVersionIPv4, radar.HTTPSummaryOSParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		TLSVersion:   khulnasoft.F([]radar.HTTPSummaryOSParamsTLSVersion{radar.HTTPSummaryOSParamsTLSVersionTlSv1_0, radar.HTTPSummaryOSParamsTLSVersionTlSv1_1, radar.HTTPSummaryOSParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryPostQuantumWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.PostQuantum(context.TODO(), radar.HTTPSummaryPostQuantumParams{
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPSummaryPostQuantumParamsBotClass{radar.HTTPSummaryPostQuantumParamsBotClassLikelyAutomated, radar.HTTPSummaryPostQuantumParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPSummaryPostQuantumParamsDeviceType{radar.HTTPSummaryPostQuantumParamsDeviceTypeDesktop, radar.HTTPSummaryPostQuantumParamsDeviceTypeMobile, radar.HTTPSummaryPostQuantumParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPSummaryPostQuantumParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPSummaryPostQuantumParamsHTTPProtocol{radar.HTTPSummaryPostQuantumParamsHTTPProtocolHTTP, radar.HTTPSummaryPostQuantumParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPSummaryPostQuantumParamsHTTPVersion{radar.HTTPSummaryPostQuantumParamsHTTPVersionHttPv1, radar.HTTPSummaryPostQuantumParamsHTTPVersionHttPv2, radar.HTTPSummaryPostQuantumParamsHTTPVersionHttPv3}),
		IPVersion:    khulnasoft.F([]radar.HTTPSummaryPostQuantumParamsIPVersion{radar.HTTPSummaryPostQuantumParamsIPVersionIPv4, radar.HTTPSummaryPostQuantumParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPSummaryPostQuantumParamsOS{radar.HTTPSummaryPostQuantumParamsOSWindows, radar.HTTPSummaryPostQuantumParamsOSMacosx, radar.HTTPSummaryPostQuantumParamsOSIos}),
		TLSVersion:   khulnasoft.F([]radar.HTTPSummaryPostQuantumParamsTLSVersion{radar.HTTPSummaryPostQuantumParamsTLSVersionTlSv1_0, radar.HTTPSummaryPostQuantumParamsTLSVersionTlSv1_1, radar.HTTPSummaryPostQuantumParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryTLSVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.TLSVersion(context.TODO(), radar.HTTPSummaryTLSVersionParams{
		ASN:          khulnasoft.F([]string{"string", "string", "string"}),
		BotClass:     khulnasoft.F([]radar.HTTPSummaryTLSVersionParamsBotClass{radar.HTTPSummaryTLSVersionParamsBotClassLikelyAutomated, radar.HTTPSummaryTLSVersionParamsBotClassLikelyHuman}),
		Continent:    khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:      khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   khulnasoft.F([]radar.HTTPSummaryTLSVersionParamsDeviceType{radar.HTTPSummaryTLSVersionParamsDeviceTypeDesktop, radar.HTTPSummaryTLSVersionParamsDeviceTypeMobile, radar.HTTPSummaryTLSVersionParamsDeviceTypeOther}),
		Format:       khulnasoft.F(radar.HTTPSummaryTLSVersionParamsFormatJson),
		HTTPProtocol: khulnasoft.F([]radar.HTTPSummaryTLSVersionParamsHTTPProtocol{radar.HTTPSummaryTLSVersionParamsHTTPProtocolHTTP, radar.HTTPSummaryTLSVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  khulnasoft.F([]radar.HTTPSummaryTLSVersionParamsHTTPVersion{radar.HTTPSummaryTLSVersionParamsHTTPVersionHttPv1, radar.HTTPSummaryTLSVersionParamsHTTPVersionHttPv2, radar.HTTPSummaryTLSVersionParamsHTTPVersionHttPv3}),
		IPVersion:    khulnasoft.F([]radar.HTTPSummaryTLSVersionParamsIPVersion{radar.HTTPSummaryTLSVersionParamsIPVersionIPv4, radar.HTTPSummaryTLSVersionParamsIPVersionIPv6}),
		Location:     khulnasoft.F([]string{"string", "string", "string"}),
		Name:         khulnasoft.F([]string{"string", "string", "string"}),
		OS:           khulnasoft.F([]radar.HTTPSummaryTLSVersionParamsOS{radar.HTTPSummaryTLSVersionParamsOSWindows, radar.HTTPSummaryTLSVersionParamsOSMacosx, radar.HTTPSummaryTLSVersionParamsOSIos}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
