// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/intel"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestAttackSurfaceReportIssueListWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.List(context.TODO(), intel.AttackSurfaceReportIssueListParams{
		AccountID:     khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Dismissed:     khulnasoft.F(false),
		IssueClass:    khulnasoft.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueClassNeq: khulnasoft.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueType:     khulnasoft.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
		IssueTypeNeq:  khulnasoft.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
		Page:          khulnasoft.F(int64(1)),
		PerPage:       khulnasoft.F(int64(25)),
		Product:       khulnasoft.F([]string{"access", "dns"}),
		ProductNeq:    khulnasoft.F([]string{"access", "dns"}),
		Severity:      khulnasoft.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
		SeverityNeq:   khulnasoft.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
		Subject:       khulnasoft.F([]string{"example.com", "example.com", "example.com"}),
		SubjectNeq:    khulnasoft.F([]string{"example.com", "example.com", "example.com"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackSurfaceReportIssueClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.Class(context.TODO(), intel.AttackSurfaceReportIssueClassParams{
		AccountID:     khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Dismissed:     khulnasoft.F(false),
		IssueClass:    khulnasoft.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueClassNeq: khulnasoft.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueType:     khulnasoft.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
		IssueTypeNeq:  khulnasoft.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
		Product:       khulnasoft.F([]string{"access", "dns"}),
		ProductNeq:    khulnasoft.F([]string{"access", "dns"}),
		Severity:      khulnasoft.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
		SeverityNeq:   khulnasoft.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
		Subject:       khulnasoft.F([]string{"example.com", "example.com", "example.com"}),
		SubjectNeq:    khulnasoft.F([]string{"example.com", "example.com", "example.com"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackSurfaceReportIssueDismissWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.Dismiss(
		context.TODO(),
		"issue_id",
		intel.AttackSurfaceReportIssueDismissParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Dismiss:   khulnasoft.F(true),
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

func TestAttackSurfaceReportIssueSeverityWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.Severity(context.TODO(), intel.AttackSurfaceReportIssueSeverityParams{
		AccountID:     khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Dismissed:     khulnasoft.F(false),
		IssueClass:    khulnasoft.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueClassNeq: khulnasoft.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueType:     khulnasoft.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
		IssueTypeNeq:  khulnasoft.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
		Product:       khulnasoft.F([]string{"access", "dns"}),
		ProductNeq:    khulnasoft.F([]string{"access", "dns"}),
		Severity:      khulnasoft.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
		SeverityNeq:   khulnasoft.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
		Subject:       khulnasoft.F([]string{"example.com", "example.com", "example.com"}),
		SubjectNeq:    khulnasoft.F([]string{"example.com", "example.com", "example.com"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackSurfaceReportIssueTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.Type(context.TODO(), intel.AttackSurfaceReportIssueTypeParams{
		AccountID:     khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Dismissed:     khulnasoft.F(false),
		IssueClass:    khulnasoft.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueClassNeq: khulnasoft.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueType:     khulnasoft.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
		IssueTypeNeq:  khulnasoft.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
		Product:       khulnasoft.F([]string{"access", "dns"}),
		ProductNeq:    khulnasoft.F([]string{"access", "dns"}),
		Severity:      khulnasoft.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
		SeverityNeq:   khulnasoft.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
		Subject:       khulnasoft.F([]string{"example.com", "example.com", "example.com"}),
		SubjectNeq:    khulnasoft.F([]string{"example.com", "example.com", "example.com"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
