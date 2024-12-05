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

func TestIdentityProviderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.New(context.TODO(), zero_trust.IdentityProviderNewParams{
		IdentityProvider: zero_trust.AzureADParam{
			Config: khulnasoft.F(zero_trust.AzureADConfigParam{
				Claims:                   khulnasoft.F([]string{"email_verified", "preferred_username", "custom_claim_name"}),
				ClientID:                 khulnasoft.F("<your client id>"),
				ClientSecret:             khulnasoft.F("<your client secret>"),
				ConditionalAccessEnabled: khulnasoft.F(true),
				DirectoryID:              khulnasoft.F("<your azure directory uuid>"),
				EmailClaimName:           khulnasoft.F("custom_claim_name"),
				Prompt:                   khulnasoft.F(zero_trust.AzureADConfigPromptLogin),
				SupportGroups:            khulnasoft.F(true),
			}),
			Name: khulnasoft.F("Widget Corps IDP"),
			Type: khulnasoft.F(zero_trust.IdentityProviderTypeOnetimepin),
			SCIMConfig: khulnasoft.F(zero_trust.IdentityProviderSCIMConfigParam{
				Enabled:                khulnasoft.F(true),
				GroupMemberDeprovision: khulnasoft.F(true),
				SeatDeprovision:        khulnasoft.F(true),
				Secret:                 khulnasoft.F("secret"),
				UserDeprovision:        khulnasoft.F(true),
			}),
		},
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

func TestIdentityProviderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.IdentityProviderUpdateParams{
			IdentityProvider: zero_trust.AzureADParam{
				Config: khulnasoft.F(zero_trust.AzureADConfigParam{
					Claims:                   khulnasoft.F([]string{"email_verified", "preferred_username", "custom_claim_name"}),
					ClientID:                 khulnasoft.F("<your client id>"),
					ClientSecret:             khulnasoft.F("<your client secret>"),
					ConditionalAccessEnabled: khulnasoft.F(true),
					DirectoryID:              khulnasoft.F("<your azure directory uuid>"),
					EmailClaimName:           khulnasoft.F("custom_claim_name"),
					Prompt:                   khulnasoft.F(zero_trust.AzureADConfigPromptLogin),
					SupportGroups:            khulnasoft.F(true),
				}),
				Name: khulnasoft.F("Widget Corps IDP"),
				Type: khulnasoft.F(zero_trust.IdentityProviderTypeOnetimepin),
				SCIMConfig: khulnasoft.F(zero_trust.IdentityProviderSCIMConfigParam{
					Enabled:                khulnasoft.F(true),
					GroupMemberDeprovision: khulnasoft.F(true),
					SeatDeprovision:        khulnasoft.F(true),
					Secret:                 khulnasoft.F("secret"),
					UserDeprovision:        khulnasoft.F(true),
				}),
			},
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

func TestIdentityProviderListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.List(context.TODO(), zero_trust.IdentityProviderListParams{
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

func TestIdentityProviderDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.IdentityProviderDeleteParams{
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

func TestIdentityProviderGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.IdentityProviderGetParams{
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
