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

func TestEmailSecurityTimeseriesGroupARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.ARC(context.TODO(), radar.EmailSecurityTimeseriesGroupARCParams{
		AggInterval: khulnasoft.F(radar.EmailSecurityTimeseriesGroupARCParamsAggInterval15m),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailSecurityTimeseriesGroupARCParamsDKIM{radar.EmailSecurityTimeseriesGroupARCParamsDKIMPass, radar.EmailSecurityTimeseriesGroupARCParamsDKIMNone, radar.EmailSecurityTimeseriesGroupARCParamsDKIMFail}),
		DMARC:       khulnasoft.F([]radar.EmailSecurityTimeseriesGroupARCParamsDMARC{radar.EmailSecurityTimeseriesGroupARCParamsDMARCPass, radar.EmailSecurityTimeseriesGroupARCParamsDMARCNone, radar.EmailSecurityTimeseriesGroupARCParamsDMARCFail}),
		Format:      khulnasoft.F(radar.EmailSecurityTimeseriesGroupARCParamsFormatJson),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupARCParamsSPF{radar.EmailSecurityTimeseriesGroupARCParamsSPFPass, radar.EmailSecurityTimeseriesGroupARCParamsSPFNone, radar.EmailSecurityTimeseriesGroupARCParamsSPFFail}),
		TLSVersion:  khulnasoft.F([]radar.EmailSecurityTimeseriesGroupARCParamsTLSVersion{radar.EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupDKIMWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.DKIM(context.TODO(), radar.EmailSecurityTimeseriesGroupDKIMParams{
		AggInterval: khulnasoft.F(radar.EmailSecurityTimeseriesGroupDKIMParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsARC{radar.EmailSecurityTimeseriesGroupDKIMParamsARCPass, radar.EmailSecurityTimeseriesGroupDKIMParamsARCNone, radar.EmailSecurityTimeseriesGroupDKIMParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DMARC:       khulnasoft.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsDMARC{radar.EmailSecurityTimeseriesGroupDKIMParamsDMARCPass, radar.EmailSecurityTimeseriesGroupDKIMParamsDMARCNone, radar.EmailSecurityTimeseriesGroupDKIMParamsDMARCFail}),
		Format:      khulnasoft.F(radar.EmailSecurityTimeseriesGroupDKIMParamsFormatJson),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsSPF{radar.EmailSecurityTimeseriesGroupDKIMParamsSPFPass, radar.EmailSecurityTimeseriesGroupDKIMParamsSPFNone, radar.EmailSecurityTimeseriesGroupDKIMParamsSPFFail}),
		TLSVersion:  khulnasoft.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsTLSVersion{radar.EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupDMARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.DMARC(context.TODO(), radar.EmailSecurityTimeseriesGroupDMARCParams{
		AggInterval: khulnasoft.F(radar.EmailSecurityTimeseriesGroupDMARCParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsARC{radar.EmailSecurityTimeseriesGroupDMARCParamsARCPass, radar.EmailSecurityTimeseriesGroupDMARCParamsARCNone, radar.EmailSecurityTimeseriesGroupDMARCParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsDKIM{radar.EmailSecurityTimeseriesGroupDMARCParamsDKIMPass, radar.EmailSecurityTimeseriesGroupDMARCParamsDKIMNone, radar.EmailSecurityTimeseriesGroupDMARCParamsDKIMFail}),
		Format:      khulnasoft.F(radar.EmailSecurityTimeseriesGroupDMARCParamsFormatJson),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsSPF{radar.EmailSecurityTimeseriesGroupDMARCParamsSPFPass, radar.EmailSecurityTimeseriesGroupDMARCParamsSPFNone, radar.EmailSecurityTimeseriesGroupDMARCParamsSPFFail}),
		TLSVersion:  khulnasoft.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsTLSVersion{radar.EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupMaliciousWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.Malicious(context.TODO(), radar.EmailSecurityTimeseriesGroupMaliciousParams{
		AggInterval: khulnasoft.F(radar.EmailSecurityTimeseriesGroupMaliciousParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsARC{radar.EmailSecurityTimeseriesGroupMaliciousParamsARCPass, radar.EmailSecurityTimeseriesGroupMaliciousParamsARCNone, radar.EmailSecurityTimeseriesGroupMaliciousParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsDKIM{radar.EmailSecurityTimeseriesGroupMaliciousParamsDKIMPass, radar.EmailSecurityTimeseriesGroupMaliciousParamsDKIMNone, radar.EmailSecurityTimeseriesGroupMaliciousParamsDKIMFail}),
		DMARC:       khulnasoft.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsDMARC{radar.EmailSecurityTimeseriesGroupMaliciousParamsDMARCPass, radar.EmailSecurityTimeseriesGroupMaliciousParamsDMARCNone, radar.EmailSecurityTimeseriesGroupMaliciousParamsDMARCFail}),
		Format:      khulnasoft.F(radar.EmailSecurityTimeseriesGroupMaliciousParamsFormatJson),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsSPF{radar.EmailSecurityTimeseriesGroupMaliciousParamsSPFPass, radar.EmailSecurityTimeseriesGroupMaliciousParamsSPFNone, radar.EmailSecurityTimeseriesGroupMaliciousParamsSPFFail}),
		TLSVersion:  khulnasoft.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion{radar.EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupSpamWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.Spam(context.TODO(), radar.EmailSecurityTimeseriesGroupSpamParams{
		AggInterval: khulnasoft.F(radar.EmailSecurityTimeseriesGroupSpamParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSpamParamsARC{radar.EmailSecurityTimeseriesGroupSpamParamsARCPass, radar.EmailSecurityTimeseriesGroupSpamParamsARCNone, radar.EmailSecurityTimeseriesGroupSpamParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSpamParamsDKIM{radar.EmailSecurityTimeseriesGroupSpamParamsDKIMPass, radar.EmailSecurityTimeseriesGroupSpamParamsDKIMNone, radar.EmailSecurityTimeseriesGroupSpamParamsDKIMFail}),
		DMARC:       khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSpamParamsDMARC{radar.EmailSecurityTimeseriesGroupSpamParamsDMARCPass, radar.EmailSecurityTimeseriesGroupSpamParamsDMARCNone, radar.EmailSecurityTimeseriesGroupSpamParamsDMARCFail}),
		Format:      khulnasoft.F(radar.EmailSecurityTimeseriesGroupSpamParamsFormatJson),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSpamParamsSPF{radar.EmailSecurityTimeseriesGroupSpamParamsSPFPass, radar.EmailSecurityTimeseriesGroupSpamParamsSPFNone, radar.EmailSecurityTimeseriesGroupSpamParamsSPFFail}),
		TLSVersion:  khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSpamParamsTLSVersion{radar.EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupSPFWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.SPF(context.TODO(), radar.EmailSecurityTimeseriesGroupSPFParams{
		AggInterval: khulnasoft.F(radar.EmailSecurityTimeseriesGroupSPFParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSPFParamsARC{radar.EmailSecurityTimeseriesGroupSPFParamsARCPass, radar.EmailSecurityTimeseriesGroupSPFParamsARCNone, radar.EmailSecurityTimeseriesGroupSPFParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSPFParamsDKIM{radar.EmailSecurityTimeseriesGroupSPFParamsDKIMPass, radar.EmailSecurityTimeseriesGroupSPFParamsDKIMNone, radar.EmailSecurityTimeseriesGroupSPFParamsDKIMFail}),
		DMARC:       khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSPFParamsDMARC{radar.EmailSecurityTimeseriesGroupSPFParamsDMARCPass, radar.EmailSecurityTimeseriesGroupSPFParamsDMARCNone, radar.EmailSecurityTimeseriesGroupSPFParamsDMARCFail}),
		Format:      khulnasoft.F(radar.EmailSecurityTimeseriesGroupSPFParamsFormatJson),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		TLSVersion:  khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSPFParamsTLSVersion{radar.EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupSpoofWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.Spoof(context.TODO(), radar.EmailSecurityTimeseriesGroupSpoofParams{
		AggInterval: khulnasoft.F(radar.EmailSecurityTimeseriesGroupSpoofParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsARC{radar.EmailSecurityTimeseriesGroupSpoofParamsARCPass, radar.EmailSecurityTimeseriesGroupSpoofParamsARCNone, radar.EmailSecurityTimeseriesGroupSpoofParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsDKIM{radar.EmailSecurityTimeseriesGroupSpoofParamsDKIMPass, radar.EmailSecurityTimeseriesGroupSpoofParamsDKIMNone, radar.EmailSecurityTimeseriesGroupSpoofParamsDKIMFail}),
		DMARC:       khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsDMARC{radar.EmailSecurityTimeseriesGroupSpoofParamsDMARCPass, radar.EmailSecurityTimeseriesGroupSpoofParamsDMARCNone, radar.EmailSecurityTimeseriesGroupSpoofParamsDMARCFail}),
		Format:      khulnasoft.F(radar.EmailSecurityTimeseriesGroupSpoofParamsFormatJson),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsSPF{radar.EmailSecurityTimeseriesGroupSpoofParamsSPFPass, radar.EmailSecurityTimeseriesGroupSpoofParamsSPFNone, radar.EmailSecurityTimeseriesGroupSpoofParamsSPFFail}),
		TLSVersion:  khulnasoft.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsTLSVersion{radar.EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupThreatCategoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.ThreatCategory(context.TODO(), radar.EmailSecurityTimeseriesGroupThreatCategoryParams{
		AggInterval: khulnasoft.F(radar.EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsARC{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsARCPass, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsARCNone, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMPass, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMNone, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMFail}),
		DMARC:       khulnasoft.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCPass, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCNone, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCFail}),
		Format:      khulnasoft.F(radar.EmailSecurityTimeseriesGroupThreatCategoryParamsFormatJson),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsSPF{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsSPFPass, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsSPFNone, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsSPFFail}),
		TLSVersion:  khulnasoft.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupTLSVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.TLSVersion(context.TODO(), radar.EmailSecurityTimeseriesGroupTLSVersionParams{
		AggInterval: khulnasoft.F(radar.EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval15m),
		ARC:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsARC{radar.EmailSecurityTimeseriesGroupTLSVersionParamsARCPass, radar.EmailSecurityTimeseriesGroupTLSVersionParamsARCNone, radar.EmailSecurityTimeseriesGroupTLSVersionParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsDKIM{radar.EmailSecurityTimeseriesGroupTLSVersionParamsDKIMPass, radar.EmailSecurityTimeseriesGroupTLSVersionParamsDKIMNone, radar.EmailSecurityTimeseriesGroupTLSVersionParamsDKIMFail}),
		DMARC:       khulnasoft.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsDMARC{radar.EmailSecurityTimeseriesGroupTLSVersionParamsDMARCPass, radar.EmailSecurityTimeseriesGroupTLSVersionParamsDMARCNone, radar.EmailSecurityTimeseriesGroupTLSVersionParamsDMARCFail}),
		Format:      khulnasoft.F(radar.EmailSecurityTimeseriesGroupTLSVersionParamsFormatJson),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsSPF{radar.EmailSecurityTimeseriesGroupTLSVersionParamsSPFPass, radar.EmailSecurityTimeseriesGroupTLSVersionParamsSPFNone, radar.EmailSecurityTimeseriesGroupTLSVersionParamsSPFFail}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
