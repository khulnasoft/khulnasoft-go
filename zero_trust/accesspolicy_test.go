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

func TestAccessPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Policies.New(context.TODO(), zero_trust.AccessPolicyNewParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Decision:  khulnasoft.F(zero_trust.DecisionAllow),
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
		Name: khulnasoft.F("Allow devs"),
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
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessPolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Policies.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessPolicyUpdateParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Decision:  khulnasoft.F(zero_trust.DecisionAllow),
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
			Name: khulnasoft.F("Allow devs"),
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

func TestAccessPolicyList(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Policies.List(context.TODO(), zero_trust.AccessPolicyListParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessPolicyDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Policies.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessPolicyDeleteParams{
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

func TestAccessPolicyGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Policies.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessPolicyGetParams{
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
