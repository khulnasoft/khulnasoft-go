// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/email_security"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestSettingAllowPatternNewWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPatterns.New(context.TODO(), email_security.SettingAllowPatternNewParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body: email_security.SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPattern{
			IsRecipient:  khulnasoft.F(false),
			IsRegex:      khulnasoft.F(false),
			IsSender:     khulnasoft.F(true),
			IsSpoof:      khulnasoft.F(false),
			Pattern:      khulnasoft.F("test@example.com"),
			PatternType:  khulnasoft.F(email_security.SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternTypeEmail),
			VerifySender: khulnasoft.F(true),
			Comments:     khulnasoft.F("Trust all messages send from test@example.com"),
		},
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingAllowPatternListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPatterns.List(context.TODO(), email_security.SettingAllowPatternListParams{
		AccountID:    khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction:    khulnasoft.F(email_security.SettingAllowPatternListParamsDirectionAsc),
		IsRecipient:  khulnasoft.F(true),
		IsSender:     khulnasoft.F(true),
		IsSpoof:      khulnasoft.F(true),
		Order:        khulnasoft.F(email_security.SettingAllowPatternListParamsOrderPattern),
		Page:         khulnasoft.F(int64(1)),
		PatternType:  khulnasoft.F(email_security.SettingAllowPatternListParamsPatternTypeEmail),
		PerPage:      khulnasoft.F(int64(1)),
		Search:       khulnasoft.F("search"),
		VerifySender: khulnasoft.F(true),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingAllowPatternDelete(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPatterns.Delete(
		context.TODO(),
		int64(2401),
		email_security.SettingAllowPatternDeleteParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestSettingAllowPatternEditWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPatterns.Edit(
		context.TODO(),
		int64(2401),
		email_security.SettingAllowPatternEditParams{
			AccountID:    khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Comments:     khulnasoft.F("comments"),
			IsRecipient:  khulnasoft.F(true),
			IsRegex:      khulnasoft.F(true),
			IsSender:     khulnasoft.F(true),
			IsSpoof:      khulnasoft.F(true),
			Pattern:      khulnasoft.F("x"),
			PatternType:  khulnasoft.F(email_security.SettingAllowPatternEditParamsPatternTypeEmail),
			VerifySender: khulnasoft.F(true),
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

func TestSettingAllowPatternGet(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPatterns.Get(
		context.TODO(),
		int64(2401),
		email_security.SettingAllowPatternGetParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
