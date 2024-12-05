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

func TestOrganizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Organizations.New(context.TODO(), zero_trust.OrganizationNewParams{
		AuthDomain:               khulnasoft.F("test.khulnasoftaccess.com"),
		Name:                     khulnasoft.F("Widget Corps Internal Applications"),
		AccountID:                khulnasoft.F("account_id"),
		AllowAuthenticateViaWARP: khulnasoft.F(true),
		AutoRedirectToIdentity:   khulnasoft.F(true),
		IsUIReadOnly:             khulnasoft.F(true),
		LoginDesign: khulnasoft.F(zero_trust.LoginDesignParam{
			BackgroundColor: khulnasoft.F("#c5ed1b"),
			FooterText:      khulnasoft.F("This is an example description."),
			HeaderText:      khulnasoft.F("This is an example description."),
			LogoPath:        khulnasoft.F("https://example.com/logo.png"),
			TextColor:       khulnasoft.F("#c5ed1b"),
		}),
		SessionDuration:                khulnasoft.F("24h"),
		UIReadOnlyToggleReason:         khulnasoft.F("Temporarily turn off the UI read only lock to make a change via the UI"),
		UserSeatExpirationInactiveTime: khulnasoft.F("730h"),
		WARPAuthSessionDuration:        khulnasoft.F("24h"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Organizations.Update(context.TODO(), zero_trust.OrganizationUpdateParams{
		AccountID:                khulnasoft.F("account_id"),
		AllowAuthenticateViaWARP: khulnasoft.F(true),
		AuthDomain:               khulnasoft.F("test.khulnasoftaccess.com"),
		AutoRedirectToIdentity:   khulnasoft.F(true),
		CustomPages: khulnasoft.F(zero_trust.OrganizationUpdateParamsCustomPages{
			Forbidden:      khulnasoft.F("699d98642c564d2e855e9661899b7252"),
			IdentityDenied: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		}),
		IsUIReadOnly: khulnasoft.F(true),
		LoginDesign: khulnasoft.F(zero_trust.LoginDesignParam{
			BackgroundColor: khulnasoft.F("#c5ed1b"),
			FooterText:      khulnasoft.F("This is an example description."),
			HeaderText:      khulnasoft.F("This is an example description."),
			LogoPath:        khulnasoft.F("https://example.com/logo.png"),
			TextColor:       khulnasoft.F("#c5ed1b"),
		}),
		Name:                           khulnasoft.F("Widget Corps Internal Applications"),
		SessionDuration:                khulnasoft.F("24h"),
		UIReadOnlyToggleReason:         khulnasoft.F("Temporarily turn off the UI read only lock to make a change via the UI"),
		UserSeatExpirationInactiveTime: khulnasoft.F("730h"),
		WARPAuthSessionDuration:        khulnasoft.F("24h"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Organizations.List(context.TODO(), zero_trust.OrganizationListParams{
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

func TestOrganizationRevokeUsersWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Organizations.RevokeUsers(context.TODO(), zero_trust.OrganizationRevokeUsersParams{
		Email:     khulnasoft.F("test@example.com"),
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
