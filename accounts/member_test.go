// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/accounts"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/shared"
)

func TestMemberNewWithOptionalParams(t *testing.T) {
	t.Skip("HTTP 422 error from prism")
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
	_, err := client.Accounts.Members.New(context.TODO(), accounts.MemberNewParams{
		AccountID: khulnasoft.F("eb78d65290b24279ba6f44721b3ea3c4"),
		Body: accounts.MemberNewParamsBodyIAMCreateMemberWithRoles{
			Email:  khulnasoft.F("user@example.com"),
			Roles:  khulnasoft.F([]string{"3536bcfad5faccb999b47003c79917fb", "3536bcfad5faccb999b47003c79917fb", "3536bcfad5faccb999b47003c79917fb"}),
			Status: khulnasoft.F(accounts.MemberNewParamsBodyIAMCreateMemberWithRolesStatusAccepted),
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

func TestMemberUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Accounts.Members.Update(
		context.TODO(),
		"4536bcfad5faccb111b47003c79917fa",
		accounts.MemberUpdateParams{
			AccountID: khulnasoft.F("eb78d65290b24279ba6f44721b3ea3c4"),
			Body: shared.MemberParam{
				Roles: khulnasoft.F([]shared.MemberRoleParam{{
					ID: khulnasoft.F("3536bcfad5faccb999b47003c79917fb"),
				}, {
					ID: khulnasoft.F("3536bcfad5faccb999b47003c79917fb"),
				}, {
					ID: khulnasoft.F("3536bcfad5faccb999b47003c79917fb"),
				}}),
			},
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

func TestMemberListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Members.List(context.TODO(), accounts.MemberListParams{
		AccountID: khulnasoft.F("eb78d65290b24279ba6f44721b3ea3c4"),
		Direction: khulnasoft.F(accounts.MemberListParamsDirectionAsc),
		Order:     khulnasoft.F(accounts.MemberListParamsOrderUserFirstName),
		Page:      khulnasoft.F(1.000000),
		PerPage:   khulnasoft.F(5.000000),
		Status:    khulnasoft.F(accounts.MemberListParamsStatusAccepted),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemberDelete(t *testing.T) {
	t.Skip("HTTP 422 error from prism")
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
	_, err := client.Accounts.Members.Delete(
		context.TODO(),
		"4536bcfad5faccb111b47003c79917fa",
		accounts.MemberDeleteParams{
			AccountID: khulnasoft.F("eb78d65290b24279ba6f44721b3ea3c4"),
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

func TestMemberGet(t *testing.T) {
	t.Skip("HTTP 422 error from prism")
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
	_, err := client.Accounts.Members.Get(
		context.TODO(),
		"4536bcfad5faccb111b47003c79917fa",
		accounts.MemberGetParams{
			AccountID: khulnasoft.F("eb78d65290b24279ba6f44721b3ea3c4"),
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
