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

func TestAccessGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Groups.New(context.TODO(), zero_trust.AccessGroupNewParams{
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
		IsDefault: khulnasoft.F(true),
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
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Groups.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessGroupUpdateParams{
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
			IsDefault: khulnasoft.F(true),
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

func TestAccessGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Groups.List(context.TODO(), zero_trust.AccessGroupListParams{
		AccountID: khulnasoft.F("account_id"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessGroupDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Groups.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessGroupDeleteParams{
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

func TestAccessGroupGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Groups.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessGroupGetParams{
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
