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

func TestEmailSecuritySummaryARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.ARC(context.TODO(), radar.EmailSecuritySummaryARCParams{
		DateEnd:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:  khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       khulnasoft.F([]radar.EmailSecuritySummaryARCParamsDKIM{radar.EmailSecuritySummaryARCParamsDKIMPass, radar.EmailSecuritySummaryARCParamsDKIMNone, radar.EmailSecuritySummaryARCParamsDKIMFail}),
		DMARC:      khulnasoft.F([]radar.EmailSecuritySummaryARCParamsDMARC{radar.EmailSecuritySummaryARCParamsDMARCPass, radar.EmailSecuritySummaryARCParamsDMARCNone, radar.EmailSecuritySummaryARCParamsDMARCFail}),
		Format:     khulnasoft.F(radar.EmailSecuritySummaryARCParamsFormatJson),
		Name:       khulnasoft.F([]string{"string", "string", "string"}),
		SPF:        khulnasoft.F([]radar.EmailSecuritySummaryARCParamsSPF{radar.EmailSecuritySummaryARCParamsSPFPass, radar.EmailSecuritySummaryARCParamsSPFNone, radar.EmailSecuritySummaryARCParamsSPFFail}),
		TLSVersion: khulnasoft.F([]radar.EmailSecuritySummaryARCParamsTLSVersion{radar.EmailSecuritySummaryARCParamsTLSVersionTlSv1_0, radar.EmailSecuritySummaryARCParamsTLSVersionTlSv1_1, radar.EmailSecuritySummaryARCParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummaryDKIMWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.DKIM(context.TODO(), radar.EmailSecuritySummaryDKIMParams{
		ARC:        khulnasoft.F([]radar.EmailSecuritySummaryDKIMParamsARC{radar.EmailSecuritySummaryDKIMParamsARCPass, radar.EmailSecuritySummaryDKIMParamsARCNone, radar.EmailSecuritySummaryDKIMParamsARCFail}),
		DateEnd:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:  khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DMARC:      khulnasoft.F([]radar.EmailSecuritySummaryDKIMParamsDMARC{radar.EmailSecuritySummaryDKIMParamsDMARCPass, radar.EmailSecuritySummaryDKIMParamsDMARCNone, radar.EmailSecuritySummaryDKIMParamsDMARCFail}),
		Format:     khulnasoft.F(radar.EmailSecuritySummaryDKIMParamsFormatJson),
		Name:       khulnasoft.F([]string{"string", "string", "string"}),
		SPF:        khulnasoft.F([]radar.EmailSecuritySummaryDKIMParamsSPF{radar.EmailSecuritySummaryDKIMParamsSPFPass, radar.EmailSecuritySummaryDKIMParamsSPFNone, radar.EmailSecuritySummaryDKIMParamsSPFFail}),
		TLSVersion: khulnasoft.F([]radar.EmailSecuritySummaryDKIMParamsTLSVersion{radar.EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_0, radar.EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_1, radar.EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummaryDMARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.DMARC(context.TODO(), radar.EmailSecuritySummaryDMARCParams{
		ARC:        khulnasoft.F([]radar.EmailSecuritySummaryDMARCParamsARC{radar.EmailSecuritySummaryDMARCParamsARCPass, radar.EmailSecuritySummaryDMARCParamsARCNone, radar.EmailSecuritySummaryDMARCParamsARCFail}),
		DateEnd:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:  khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       khulnasoft.F([]radar.EmailSecuritySummaryDMARCParamsDKIM{radar.EmailSecuritySummaryDMARCParamsDKIMPass, radar.EmailSecuritySummaryDMARCParamsDKIMNone, radar.EmailSecuritySummaryDMARCParamsDKIMFail}),
		Format:     khulnasoft.F(radar.EmailSecuritySummaryDMARCParamsFormatJson),
		Name:       khulnasoft.F([]string{"string", "string", "string"}),
		SPF:        khulnasoft.F([]radar.EmailSecuritySummaryDMARCParamsSPF{radar.EmailSecuritySummaryDMARCParamsSPFPass, radar.EmailSecuritySummaryDMARCParamsSPFNone, radar.EmailSecuritySummaryDMARCParamsSPFFail}),
		TLSVersion: khulnasoft.F([]radar.EmailSecuritySummaryDMARCParamsTLSVersion{radar.EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_0, radar.EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_1, radar.EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummaryMaliciousWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.Malicious(context.TODO(), radar.EmailSecuritySummaryMaliciousParams{
		ARC:        khulnasoft.F([]radar.EmailSecuritySummaryMaliciousParamsARC{radar.EmailSecuritySummaryMaliciousParamsARCPass, radar.EmailSecuritySummaryMaliciousParamsARCNone, radar.EmailSecuritySummaryMaliciousParamsARCFail}),
		DateEnd:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:  khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       khulnasoft.F([]radar.EmailSecuritySummaryMaliciousParamsDKIM{radar.EmailSecuritySummaryMaliciousParamsDKIMPass, radar.EmailSecuritySummaryMaliciousParamsDKIMNone, radar.EmailSecuritySummaryMaliciousParamsDKIMFail}),
		DMARC:      khulnasoft.F([]radar.EmailSecuritySummaryMaliciousParamsDMARC{radar.EmailSecuritySummaryMaliciousParamsDMARCPass, radar.EmailSecuritySummaryMaliciousParamsDMARCNone, radar.EmailSecuritySummaryMaliciousParamsDMARCFail}),
		Format:     khulnasoft.F(radar.EmailSecuritySummaryMaliciousParamsFormatJson),
		Name:       khulnasoft.F([]string{"string", "string", "string"}),
		SPF:        khulnasoft.F([]radar.EmailSecuritySummaryMaliciousParamsSPF{radar.EmailSecuritySummaryMaliciousParamsSPFPass, radar.EmailSecuritySummaryMaliciousParamsSPFNone, radar.EmailSecuritySummaryMaliciousParamsSPFFail}),
		TLSVersion: khulnasoft.F([]radar.EmailSecuritySummaryMaliciousParamsTLSVersion{radar.EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_0, radar.EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_1, radar.EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummarySpamWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.Spam(context.TODO(), radar.EmailSecuritySummarySpamParams{
		ARC:        khulnasoft.F([]radar.EmailSecuritySummarySpamParamsARC{radar.EmailSecuritySummarySpamParamsARCPass, radar.EmailSecuritySummarySpamParamsARCNone, radar.EmailSecuritySummarySpamParamsARCFail}),
		DateEnd:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:  khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       khulnasoft.F([]radar.EmailSecuritySummarySpamParamsDKIM{radar.EmailSecuritySummarySpamParamsDKIMPass, radar.EmailSecuritySummarySpamParamsDKIMNone, radar.EmailSecuritySummarySpamParamsDKIMFail}),
		DMARC:      khulnasoft.F([]radar.EmailSecuritySummarySpamParamsDMARC{radar.EmailSecuritySummarySpamParamsDMARCPass, radar.EmailSecuritySummarySpamParamsDMARCNone, radar.EmailSecuritySummarySpamParamsDMARCFail}),
		Format:     khulnasoft.F(radar.EmailSecuritySummarySpamParamsFormatJson),
		Name:       khulnasoft.F([]string{"string", "string", "string"}),
		SPF:        khulnasoft.F([]radar.EmailSecuritySummarySpamParamsSPF{radar.EmailSecuritySummarySpamParamsSPFPass, radar.EmailSecuritySummarySpamParamsSPFNone, radar.EmailSecuritySummarySpamParamsSPFFail}),
		TLSVersion: khulnasoft.F([]radar.EmailSecuritySummarySpamParamsTLSVersion{radar.EmailSecuritySummarySpamParamsTLSVersionTlSv1_0, radar.EmailSecuritySummarySpamParamsTLSVersionTlSv1_1, radar.EmailSecuritySummarySpamParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummarySPFWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.SPF(context.TODO(), radar.EmailSecuritySummarySPFParams{
		ARC:        khulnasoft.F([]radar.EmailSecuritySummarySPFParamsARC{radar.EmailSecuritySummarySPFParamsARCPass, radar.EmailSecuritySummarySPFParamsARCNone, radar.EmailSecuritySummarySPFParamsARCFail}),
		DateEnd:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:  khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       khulnasoft.F([]radar.EmailSecuritySummarySPFParamsDKIM{radar.EmailSecuritySummarySPFParamsDKIMPass, radar.EmailSecuritySummarySPFParamsDKIMNone, radar.EmailSecuritySummarySPFParamsDKIMFail}),
		DMARC:      khulnasoft.F([]radar.EmailSecuritySummarySPFParamsDMARC{radar.EmailSecuritySummarySPFParamsDMARCPass, radar.EmailSecuritySummarySPFParamsDMARCNone, radar.EmailSecuritySummarySPFParamsDMARCFail}),
		Format:     khulnasoft.F(radar.EmailSecuritySummarySPFParamsFormatJson),
		Name:       khulnasoft.F([]string{"string", "string", "string"}),
		TLSVersion: khulnasoft.F([]radar.EmailSecuritySummarySPFParamsTLSVersion{radar.EmailSecuritySummarySPFParamsTLSVersionTlSv1_0, radar.EmailSecuritySummarySPFParamsTLSVersionTlSv1_1, radar.EmailSecuritySummarySPFParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummarySpoofWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.Spoof(context.TODO(), radar.EmailSecuritySummarySpoofParams{
		ARC:        khulnasoft.F([]radar.EmailSecuritySummarySpoofParamsARC{radar.EmailSecuritySummarySpoofParamsARCPass, radar.EmailSecuritySummarySpoofParamsARCNone, radar.EmailSecuritySummarySpoofParamsARCFail}),
		DateEnd:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:  khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       khulnasoft.F([]radar.EmailSecuritySummarySpoofParamsDKIM{radar.EmailSecuritySummarySpoofParamsDKIMPass, radar.EmailSecuritySummarySpoofParamsDKIMNone, radar.EmailSecuritySummarySpoofParamsDKIMFail}),
		DMARC:      khulnasoft.F([]radar.EmailSecuritySummarySpoofParamsDMARC{radar.EmailSecuritySummarySpoofParamsDMARCPass, radar.EmailSecuritySummarySpoofParamsDMARCNone, radar.EmailSecuritySummarySpoofParamsDMARCFail}),
		Format:     khulnasoft.F(radar.EmailSecuritySummarySpoofParamsFormatJson),
		Name:       khulnasoft.F([]string{"string", "string", "string"}),
		SPF:        khulnasoft.F([]radar.EmailSecuritySummarySpoofParamsSPF{radar.EmailSecuritySummarySpoofParamsSPFPass, radar.EmailSecuritySummarySpoofParamsSPFNone, radar.EmailSecuritySummarySpoofParamsSPFFail}),
		TLSVersion: khulnasoft.F([]radar.EmailSecuritySummarySpoofParamsTLSVersion{radar.EmailSecuritySummarySpoofParamsTLSVersionTlSv1_0, radar.EmailSecuritySummarySpoofParamsTLSVersionTlSv1_1, radar.EmailSecuritySummarySpoofParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummaryThreatCategoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.ThreatCategory(context.TODO(), radar.EmailSecuritySummaryThreatCategoryParams{
		ARC:        khulnasoft.F([]radar.EmailSecuritySummaryThreatCategoryParamsARC{radar.EmailSecuritySummaryThreatCategoryParamsARCPass, radar.EmailSecuritySummaryThreatCategoryParamsARCNone, radar.EmailSecuritySummaryThreatCategoryParamsARCFail}),
		DateEnd:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:  khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       khulnasoft.F([]radar.EmailSecuritySummaryThreatCategoryParamsDKIM{radar.EmailSecuritySummaryThreatCategoryParamsDKIMPass, radar.EmailSecuritySummaryThreatCategoryParamsDKIMNone, radar.EmailSecuritySummaryThreatCategoryParamsDKIMFail}),
		DMARC:      khulnasoft.F([]radar.EmailSecuritySummaryThreatCategoryParamsDMARC{radar.EmailSecuritySummaryThreatCategoryParamsDMARCPass, radar.EmailSecuritySummaryThreatCategoryParamsDMARCNone, radar.EmailSecuritySummaryThreatCategoryParamsDMARCFail}),
		Format:     khulnasoft.F(radar.EmailSecuritySummaryThreatCategoryParamsFormatJson),
		Name:       khulnasoft.F([]string{"string", "string", "string"}),
		SPF:        khulnasoft.F([]radar.EmailSecuritySummaryThreatCategoryParamsSPF{radar.EmailSecuritySummaryThreatCategoryParamsSPFPass, radar.EmailSecuritySummaryThreatCategoryParamsSPFNone, radar.EmailSecuritySummaryThreatCategoryParamsSPFFail}),
		TLSVersion: khulnasoft.F([]radar.EmailSecuritySummaryThreatCategoryParamsTLSVersion{radar.EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_0, radar.EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_1, radar.EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummaryTLSVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.TLSVersion(context.TODO(), radar.EmailSecuritySummaryTLSVersionParams{
		ARC:       khulnasoft.F([]radar.EmailSecuritySummaryTLSVersionParamsARC{radar.EmailSecuritySummaryTLSVersionParamsARCPass, radar.EmailSecuritySummaryTLSVersionParamsARCNone, radar.EmailSecuritySummaryTLSVersionParamsARCFail}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      khulnasoft.F([]radar.EmailSecuritySummaryTLSVersionParamsDKIM{radar.EmailSecuritySummaryTLSVersionParamsDKIMPass, radar.EmailSecuritySummaryTLSVersionParamsDKIMNone, radar.EmailSecuritySummaryTLSVersionParamsDKIMFail}),
		DMARC:     khulnasoft.F([]radar.EmailSecuritySummaryTLSVersionParamsDMARC{radar.EmailSecuritySummaryTLSVersionParamsDMARCPass, radar.EmailSecuritySummaryTLSVersionParamsDMARCNone, radar.EmailSecuritySummaryTLSVersionParamsDMARCFail}),
		Format:    khulnasoft.F(radar.EmailSecuritySummaryTLSVersionParamsFormatJson),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		SPF:       khulnasoft.F([]radar.EmailSecuritySummaryTLSVersionParamsSPF{radar.EmailSecuritySummaryTLSVersionParamsSPFPass, radar.EmailSecuritySummaryTLSVersionParamsSPFNone, radar.EmailSecuritySummaryTLSVersionParamsSPFFail}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
