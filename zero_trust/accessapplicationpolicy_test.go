// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/zero_trust"
)

func TestAccessApplicationPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.Policies.New(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyNewParams{
			Decision: khulnasoft.F(zero_trust.DecisionAllow),
			Include: khulnasoft.F([]zero_trust.AccessRuleUnionParam{zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}}),
			Name:      khulnasoft.F("Allow devs"),
			AccountID: khulnasoft.F("account_id"),
			ApprovalGroups: khulnasoft.F([]zero_trust.ApprovalGroupParam{{
				ApprovalsNeeded: khulnasoft.F(1.000000),
				EmailAddresses:  khulnasoft.F([]string{"test1@khulnasoft.com", "test2@khulnasoft.com"}),
				EmailListUUID:   khulnasoft.F("email_list_uuid"),
			}, {
				ApprovalsNeeded: khulnasoft.F(3.000000),
				EmailAddresses:  khulnasoft.F([]string{"test@khulnasoft.com", "test2@khulnasoft.com"}),
				EmailListUUID:   khulnasoft.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: khulnasoft.F(true),
			Exclude: khulnasoft.F([]zero_trust.AccessRuleUnionParam{zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}}),
			IsolationRequired:            khulnasoft.F(false),
			Precedence:                   khulnasoft.F(int64(0)),
			PurposeJustificationPrompt:   khulnasoft.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: khulnasoft.F(true),
			Require: khulnasoft.F([]zero_trust.AccessRuleUnionParam{zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}}),
			SessionDuration: khulnasoft.F("24h"),
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

func TestAccessApplicationPolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.Policies.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyUpdateParams{
			Decision: khulnasoft.F(zero_trust.DecisionAllow),
			Include: khulnasoft.F([]zero_trust.AccessRuleUnionParam{zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}}),
			Name:      khulnasoft.F("Allow devs"),
			AccountID: khulnasoft.F("account_id"),
			ApprovalGroups: khulnasoft.F([]zero_trust.ApprovalGroupParam{{
				ApprovalsNeeded: khulnasoft.F(1.000000),
				EmailAddresses:  khulnasoft.F([]string{"test1@khulnasoft.com", "test2@khulnasoft.com"}),
				EmailListUUID:   khulnasoft.F("email_list_uuid"),
			}, {
				ApprovalsNeeded: khulnasoft.F(3.000000),
				EmailAddresses:  khulnasoft.F([]string{"test@khulnasoft.com", "test2@khulnasoft.com"}),
				EmailListUUID:   khulnasoft.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: khulnasoft.F(true),
			Exclude: khulnasoft.F([]zero_trust.AccessRuleUnionParam{zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}}),
			IsolationRequired:            khulnasoft.F(false),
			Precedence:                   khulnasoft.F(int64(0)),
			PurposeJustificationPrompt:   khulnasoft.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: khulnasoft.F(true),
			Require: khulnasoft.F([]zero_trust.AccessRuleUnionParam{zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: khulnasoft.F(zero_trust.EmailRuleEmailParam{
					Email: khulnasoft.F("test@example.com"),
				}),
			}}),
			SessionDuration: khulnasoft.F("24h"),
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

func TestAccessApplicationPolicyListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.Policies.List(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyListParams{
			AccountID: khulnasoft.F("account_id"),
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

func TestAccessApplicationPolicyDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.Policies.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyDeleteParams{
			AccountID: khulnasoft.F("account_id"),
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

func TestAccessApplicationPolicyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.Policies.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyGetParams{
			AccountID: khulnasoft.F("account_id"),
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
