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

func TestEmailRoutingTimeseriesGroupARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.ARC(context.TODO(), radar.EmailRoutingTimeseriesGroupARCParams{
		AggInterval: khulnasoft.F(radar.EmailRoutingTimeseriesGroupARCParamsAggInterval15m),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailRoutingTimeseriesGroupARCParamsDKIM{radar.EmailRoutingTimeseriesGroupARCParamsDKIMPass, radar.EmailRoutingTimeseriesGroupARCParamsDKIMNone, radar.EmailRoutingTimeseriesGroupARCParamsDKIMFail}),
		DMARC:       khulnasoft.F([]radar.EmailRoutingTimeseriesGroupARCParamsDMARC{radar.EmailRoutingTimeseriesGroupARCParamsDMARCPass, radar.EmailRoutingTimeseriesGroupARCParamsDMARCNone, radar.EmailRoutingTimeseriesGroupARCParamsDMARCFail}),
		Encrypted:   khulnasoft.F([]radar.EmailRoutingTimeseriesGroupARCParamsEncrypted{radar.EmailRoutingTimeseriesGroupARCParamsEncryptedEncrypted, radar.EmailRoutingTimeseriesGroupARCParamsEncryptedNotEncrypted}),
		Format:      khulnasoft.F(radar.EmailRoutingTimeseriesGroupARCParamsFormatJson),
		IPVersion:   khulnasoft.F([]radar.EmailRoutingTimeseriesGroupARCParamsIPVersion{radar.EmailRoutingTimeseriesGroupARCParamsIPVersionIPv4, radar.EmailRoutingTimeseriesGroupARCParamsIPVersionIPv6}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailRoutingTimeseriesGroupARCParamsSPF{radar.EmailRoutingTimeseriesGroupARCParamsSPFPass, radar.EmailRoutingTimeseriesGroupARCParamsSPFNone, radar.EmailRoutingTimeseriesGroupARCParamsSPFFail}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingTimeseriesGroupDKIMWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.DKIM(context.TODO(), radar.EmailRoutingTimeseriesGroupDKIMParams{
		AggInterval: khulnasoft.F(radar.EmailRoutingTimeseriesGroupDKIMParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsARC{radar.EmailRoutingTimeseriesGroupDKIMParamsARCPass, radar.EmailRoutingTimeseriesGroupDKIMParamsARCNone, radar.EmailRoutingTimeseriesGroupDKIMParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DMARC:       khulnasoft.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsDMARC{radar.EmailRoutingTimeseriesGroupDKIMParamsDMARCPass, radar.EmailRoutingTimeseriesGroupDKIMParamsDMARCNone, radar.EmailRoutingTimeseriesGroupDKIMParamsDMARCFail}),
		Encrypted:   khulnasoft.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsEncrypted{radar.EmailRoutingTimeseriesGroupDKIMParamsEncryptedEncrypted, radar.EmailRoutingTimeseriesGroupDKIMParamsEncryptedNotEncrypted}),
		Format:      khulnasoft.F(radar.EmailRoutingTimeseriesGroupDKIMParamsFormatJson),
		IPVersion:   khulnasoft.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsIPVersion{radar.EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv4, radar.EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv6}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsSPF{radar.EmailRoutingTimeseriesGroupDKIMParamsSPFPass, radar.EmailRoutingTimeseriesGroupDKIMParamsSPFNone, radar.EmailRoutingTimeseriesGroupDKIMParamsSPFFail}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingTimeseriesGroupDMARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.DMARC(context.TODO(), radar.EmailRoutingTimeseriesGroupDMARCParams{
		AggInterval: khulnasoft.F(radar.EmailRoutingTimeseriesGroupDMARCParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsARC{radar.EmailRoutingTimeseriesGroupDMARCParamsARCPass, radar.EmailRoutingTimeseriesGroupDMARCParamsARCNone, radar.EmailRoutingTimeseriesGroupDMARCParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsDKIM{radar.EmailRoutingTimeseriesGroupDMARCParamsDKIMPass, radar.EmailRoutingTimeseriesGroupDMARCParamsDKIMNone, radar.EmailRoutingTimeseriesGroupDMARCParamsDKIMFail}),
		Encrypted:   khulnasoft.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsEncrypted{radar.EmailRoutingTimeseriesGroupDMARCParamsEncryptedEncrypted, radar.EmailRoutingTimeseriesGroupDMARCParamsEncryptedNotEncrypted}),
		Format:      khulnasoft.F(radar.EmailRoutingTimeseriesGroupDMARCParamsFormatJson),
		IPVersion:   khulnasoft.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsIPVersion{radar.EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv4, radar.EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv6}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsSPF{radar.EmailRoutingTimeseriesGroupDMARCParamsSPFPass, radar.EmailRoutingTimeseriesGroupDMARCParamsSPFNone, radar.EmailRoutingTimeseriesGroupDMARCParamsSPFFail}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingTimeseriesGroupEncryptedWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.Encrypted(context.TODO(), radar.EmailRoutingTimeseriesGroupEncryptedParams{
		AggInterval: khulnasoft.F(radar.EmailRoutingTimeseriesGroupEncryptedParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsARC{radar.EmailRoutingTimeseriesGroupEncryptedParamsARCPass, radar.EmailRoutingTimeseriesGroupEncryptedParamsARCNone, radar.EmailRoutingTimeseriesGroupEncryptedParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsDKIM{radar.EmailRoutingTimeseriesGroupEncryptedParamsDKIMPass, radar.EmailRoutingTimeseriesGroupEncryptedParamsDKIMNone, radar.EmailRoutingTimeseriesGroupEncryptedParamsDKIMFail}),
		DMARC:       khulnasoft.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsDMARC{radar.EmailRoutingTimeseriesGroupEncryptedParamsDMARCPass, radar.EmailRoutingTimeseriesGroupEncryptedParamsDMARCNone, radar.EmailRoutingTimeseriesGroupEncryptedParamsDMARCFail}),
		Format:      khulnasoft.F(radar.EmailRoutingTimeseriesGroupEncryptedParamsFormatJson),
		IPVersion:   khulnasoft.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsIPVersion{radar.EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv4, radar.EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv6}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsSPF{radar.EmailRoutingTimeseriesGroupEncryptedParamsSPFPass, radar.EmailRoutingTimeseriesGroupEncryptedParamsSPFNone, radar.EmailRoutingTimeseriesGroupEncryptedParamsSPFFail}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingTimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.IPVersion(context.TODO(), radar.EmailRoutingTimeseriesGroupIPVersionParams{
		AggInterval: khulnasoft.F(radar.EmailRoutingTimeseriesGroupIPVersionParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsARC{radar.EmailRoutingTimeseriesGroupIPVersionParamsARCPass, radar.EmailRoutingTimeseriesGroupIPVersionParamsARCNone, radar.EmailRoutingTimeseriesGroupIPVersionParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsDKIM{radar.EmailRoutingTimeseriesGroupIPVersionParamsDKIMPass, radar.EmailRoutingTimeseriesGroupIPVersionParamsDKIMNone, radar.EmailRoutingTimeseriesGroupIPVersionParamsDKIMFail}),
		DMARC:       khulnasoft.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsDMARC{radar.EmailRoutingTimeseriesGroupIPVersionParamsDMARCPass, radar.EmailRoutingTimeseriesGroupIPVersionParamsDMARCNone, radar.EmailRoutingTimeseriesGroupIPVersionParamsDMARCFail}),
		Encrypted:   khulnasoft.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsEncrypted{radar.EmailRoutingTimeseriesGroupIPVersionParamsEncryptedEncrypted, radar.EmailRoutingTimeseriesGroupIPVersionParamsEncryptedNotEncrypted}),
		Format:      khulnasoft.F(radar.EmailRoutingTimeseriesGroupIPVersionParamsFormatJson),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsSPF{radar.EmailRoutingTimeseriesGroupIPVersionParamsSPFPass, radar.EmailRoutingTimeseriesGroupIPVersionParamsSPFNone, radar.EmailRoutingTimeseriesGroupIPVersionParamsSPFFail}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingTimeseriesGroupSPFWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.SPF(context.TODO(), radar.EmailRoutingTimeseriesGroupSPFParams{
		AggInterval: khulnasoft.F(radar.EmailRoutingTimeseriesGroupSPFParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailRoutingTimeseriesGroupSPFParamsARC{radar.EmailRoutingTimeseriesGroupSPFParamsARCPass, radar.EmailRoutingTimeseriesGroupSPFParamsARCNone, radar.EmailRoutingTimeseriesGroupSPFParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailRoutingTimeseriesGroupSPFParamsDKIM{radar.EmailRoutingTimeseriesGroupSPFParamsDKIMPass, radar.EmailRoutingTimeseriesGroupSPFParamsDKIMNone, radar.EmailRoutingTimeseriesGroupSPFParamsDKIMFail}),
		DMARC:       khulnasoft.F([]radar.EmailRoutingTimeseriesGroupSPFParamsDMARC{radar.EmailRoutingTimeseriesGroupSPFParamsDMARCPass, radar.EmailRoutingTimeseriesGroupSPFParamsDMARCNone, radar.EmailRoutingTimeseriesGroupSPFParamsDMARCFail}),
		Encrypted:   khulnasoft.F([]radar.EmailRoutingTimeseriesGroupSPFParamsEncrypted{radar.EmailRoutingTimeseriesGroupSPFParamsEncryptedEncrypted, radar.EmailRoutingTimeseriesGroupSPFParamsEncryptedNotEncrypted}),
		Format:      khulnasoft.F(radar.EmailRoutingTimeseriesGroupSPFParamsFormatJson),
		IPVersion:   khulnasoft.F([]radar.EmailRoutingTimeseriesGroupSPFParamsIPVersion{radar.EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv4, radar.EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv6}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
