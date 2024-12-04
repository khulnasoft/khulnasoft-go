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

func TestEmailRoutingSummaryARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.ARC(context.TODO(), radar.EmailRoutingSummaryARCParams{
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      khulnasoft.F([]radar.EmailRoutingSummaryARCParamsDKIM{radar.EmailRoutingSummaryARCParamsDKIMPass, radar.EmailRoutingSummaryARCParamsDKIMNone, radar.EmailRoutingSummaryARCParamsDKIMFail}),
		DMARC:     khulnasoft.F([]radar.EmailRoutingSummaryARCParamsDMARC{radar.EmailRoutingSummaryARCParamsDMARCPass, radar.EmailRoutingSummaryARCParamsDMARCNone, radar.EmailRoutingSummaryARCParamsDMARCFail}),
		Encrypted: khulnasoft.F([]radar.EmailRoutingSummaryARCParamsEncrypted{radar.EmailRoutingSummaryARCParamsEncryptedEncrypted, radar.EmailRoutingSummaryARCParamsEncryptedNotEncrypted}),
		Format:    khulnasoft.F(radar.EmailRoutingSummaryARCParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.EmailRoutingSummaryARCParamsIPVersion{radar.EmailRoutingSummaryARCParamsIPVersionIPv4, radar.EmailRoutingSummaryARCParamsIPVersionIPv6}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		SPF:       khulnasoft.F([]radar.EmailRoutingSummaryARCParamsSPF{radar.EmailRoutingSummaryARCParamsSPFPass, radar.EmailRoutingSummaryARCParamsSPFNone, radar.EmailRoutingSummaryARCParamsSPFFail}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingSummaryDKIMWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.DKIM(context.TODO(), radar.EmailRoutingSummaryDKIMParams{
		ARC:       khulnasoft.F([]radar.EmailRoutingSummaryDKIMParamsARC{radar.EmailRoutingSummaryDKIMParamsARCPass, radar.EmailRoutingSummaryDKIMParamsARCNone, radar.EmailRoutingSummaryDKIMParamsARCFail}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DMARC:     khulnasoft.F([]radar.EmailRoutingSummaryDKIMParamsDMARC{radar.EmailRoutingSummaryDKIMParamsDMARCPass, radar.EmailRoutingSummaryDKIMParamsDMARCNone, radar.EmailRoutingSummaryDKIMParamsDMARCFail}),
		Encrypted: khulnasoft.F([]radar.EmailRoutingSummaryDKIMParamsEncrypted{radar.EmailRoutingSummaryDKIMParamsEncryptedEncrypted, radar.EmailRoutingSummaryDKIMParamsEncryptedNotEncrypted}),
		Format:    khulnasoft.F(radar.EmailRoutingSummaryDKIMParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.EmailRoutingSummaryDKIMParamsIPVersion{radar.EmailRoutingSummaryDKIMParamsIPVersionIPv4, radar.EmailRoutingSummaryDKIMParamsIPVersionIPv6}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		SPF:       khulnasoft.F([]radar.EmailRoutingSummaryDKIMParamsSPF{radar.EmailRoutingSummaryDKIMParamsSPFPass, radar.EmailRoutingSummaryDKIMParamsSPFNone, radar.EmailRoutingSummaryDKIMParamsSPFFail}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingSummaryDMARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.DMARC(context.TODO(), radar.EmailRoutingSummaryDMARCParams{
		ARC:       khulnasoft.F([]radar.EmailRoutingSummaryDMARCParamsARC{radar.EmailRoutingSummaryDMARCParamsARCPass, radar.EmailRoutingSummaryDMARCParamsARCNone, radar.EmailRoutingSummaryDMARCParamsARCFail}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      khulnasoft.F([]radar.EmailRoutingSummaryDMARCParamsDKIM{radar.EmailRoutingSummaryDMARCParamsDKIMPass, radar.EmailRoutingSummaryDMARCParamsDKIMNone, radar.EmailRoutingSummaryDMARCParamsDKIMFail}),
		Encrypted: khulnasoft.F([]radar.EmailRoutingSummaryDMARCParamsEncrypted{radar.EmailRoutingSummaryDMARCParamsEncryptedEncrypted, radar.EmailRoutingSummaryDMARCParamsEncryptedNotEncrypted}),
		Format:    khulnasoft.F(radar.EmailRoutingSummaryDMARCParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.EmailRoutingSummaryDMARCParamsIPVersion{radar.EmailRoutingSummaryDMARCParamsIPVersionIPv4, radar.EmailRoutingSummaryDMARCParamsIPVersionIPv6}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		SPF:       khulnasoft.F([]radar.EmailRoutingSummaryDMARCParamsSPF{radar.EmailRoutingSummaryDMARCParamsSPFPass, radar.EmailRoutingSummaryDMARCParamsSPFNone, radar.EmailRoutingSummaryDMARCParamsSPFFail}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingSummaryEncryptedWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.Encrypted(context.TODO(), radar.EmailRoutingSummaryEncryptedParams{
		ARC:       khulnasoft.F([]radar.EmailRoutingSummaryEncryptedParamsARC{radar.EmailRoutingSummaryEncryptedParamsARCPass, radar.EmailRoutingSummaryEncryptedParamsARCNone, radar.EmailRoutingSummaryEncryptedParamsARCFail}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      khulnasoft.F([]radar.EmailRoutingSummaryEncryptedParamsDKIM{radar.EmailRoutingSummaryEncryptedParamsDKIMPass, radar.EmailRoutingSummaryEncryptedParamsDKIMNone, radar.EmailRoutingSummaryEncryptedParamsDKIMFail}),
		DMARC:     khulnasoft.F([]radar.EmailRoutingSummaryEncryptedParamsDMARC{radar.EmailRoutingSummaryEncryptedParamsDMARCPass, radar.EmailRoutingSummaryEncryptedParamsDMARCNone, radar.EmailRoutingSummaryEncryptedParamsDMARCFail}),
		Format:    khulnasoft.F(radar.EmailRoutingSummaryEncryptedParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.EmailRoutingSummaryEncryptedParamsIPVersion{radar.EmailRoutingSummaryEncryptedParamsIPVersionIPv4, radar.EmailRoutingSummaryEncryptedParamsIPVersionIPv6}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		SPF:       khulnasoft.F([]radar.EmailRoutingSummaryEncryptedParamsSPF{radar.EmailRoutingSummaryEncryptedParamsSPFPass, radar.EmailRoutingSummaryEncryptedParamsSPFNone, radar.EmailRoutingSummaryEncryptedParamsSPFFail}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingSummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.IPVersion(context.TODO(), radar.EmailRoutingSummaryIPVersionParams{
		ARC:       khulnasoft.F([]radar.EmailRoutingSummaryIPVersionParamsARC{radar.EmailRoutingSummaryIPVersionParamsARCPass, radar.EmailRoutingSummaryIPVersionParamsARCNone, radar.EmailRoutingSummaryIPVersionParamsARCFail}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      khulnasoft.F([]radar.EmailRoutingSummaryIPVersionParamsDKIM{radar.EmailRoutingSummaryIPVersionParamsDKIMPass, radar.EmailRoutingSummaryIPVersionParamsDKIMNone, radar.EmailRoutingSummaryIPVersionParamsDKIMFail}),
		DMARC:     khulnasoft.F([]radar.EmailRoutingSummaryIPVersionParamsDMARC{radar.EmailRoutingSummaryIPVersionParamsDMARCPass, radar.EmailRoutingSummaryIPVersionParamsDMARCNone, radar.EmailRoutingSummaryIPVersionParamsDMARCFail}),
		Encrypted: khulnasoft.F([]radar.EmailRoutingSummaryIPVersionParamsEncrypted{radar.EmailRoutingSummaryIPVersionParamsEncryptedEncrypted, radar.EmailRoutingSummaryIPVersionParamsEncryptedNotEncrypted}),
		Format:    khulnasoft.F(radar.EmailRoutingSummaryIPVersionParamsFormatJson),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		SPF:       khulnasoft.F([]radar.EmailRoutingSummaryIPVersionParamsSPF{radar.EmailRoutingSummaryIPVersionParamsSPFPass, radar.EmailRoutingSummaryIPVersionParamsSPFNone, radar.EmailRoutingSummaryIPVersionParamsSPFFail}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingSummarySPFWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.SPF(context.TODO(), radar.EmailRoutingSummarySPFParams{
		ARC:       khulnasoft.F([]radar.EmailRoutingSummarySPFParamsARC{radar.EmailRoutingSummarySPFParamsARCPass, radar.EmailRoutingSummarySPFParamsARCNone, radar.EmailRoutingSummarySPFParamsARCFail}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      khulnasoft.F([]radar.EmailRoutingSummarySPFParamsDKIM{radar.EmailRoutingSummarySPFParamsDKIMPass, radar.EmailRoutingSummarySPFParamsDKIMNone, radar.EmailRoutingSummarySPFParamsDKIMFail}),
		DMARC:     khulnasoft.F([]radar.EmailRoutingSummarySPFParamsDMARC{radar.EmailRoutingSummarySPFParamsDMARCPass, radar.EmailRoutingSummarySPFParamsDMARCNone, radar.EmailRoutingSummarySPFParamsDMARCFail}),
		Encrypted: khulnasoft.F([]radar.EmailRoutingSummarySPFParamsEncrypted{radar.EmailRoutingSummarySPFParamsEncryptedEncrypted, radar.EmailRoutingSummarySPFParamsEncryptedNotEncrypted}),
		Format:    khulnasoft.F(radar.EmailRoutingSummarySPFParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.EmailRoutingSummarySPFParamsIPVersion{radar.EmailRoutingSummarySPFParamsIPVersionIPv4, radar.EmailRoutingSummarySPFParamsIPVersionIPv6}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
