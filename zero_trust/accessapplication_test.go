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

func TestAccessApplicationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.New(context.TODO(), zero_trust.AccessApplicationNewParams{
		Body: zero_trust.AccessApplicationNewParamsBodySelfHostedApplication{
			Domain:                   khulnasoft.F("test.example.com/admin"),
			Type:                     khulnasoft.F("self_hosted"),
			AllowAuthenticateViaWARP: khulnasoft.F(true),
			AllowedIdPs:              khulnasoft.F([]zero_trust.AllowedIdPsParam{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
			AppLauncherVisible:       khulnasoft.F(true),
			AutoRedirectToIdentity:   khulnasoft.F(true),
			CORSHeaders: khulnasoft.F(zero_trust.CORSHeadersParam{
				AllowAllHeaders:  khulnasoft.F(true),
				AllowAllMethods:  khulnasoft.F(true),
				AllowAllOrigins:  khulnasoft.F(true),
				AllowCredentials: khulnasoft.F(true),
				AllowedHeaders:   khulnasoft.F([]zero_trust.AllowedHeadersParam{"string", "string", "string"}),
				AllowedMethods:   khulnasoft.F([]zero_trust.AllowedMethods{zero_trust.AllowedMethodsGet}),
				AllowedOrigins:   khulnasoft.F([]zero_trust.AllowedOriginsParam{"https://example.com"}),
				MaxAge:           khulnasoft.F(-1.000000),
			}),
			CustomDenyMessage:        khulnasoft.F("custom_deny_message"),
			CustomDenyURL:            khulnasoft.F("custom_deny_url"),
			CustomNonIdentityDenyURL: khulnasoft.F("custom_non_identity_deny_url"),
			CustomPages:              khulnasoft.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
			EnableBindingCookie:      khulnasoft.F(true),
			HTTPOnlyCookieAttribute:  khulnasoft.F(true),
			LogoURL:                  khulnasoft.F("https://www.khulnasoft.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
			Name:                     khulnasoft.F("Admin Site"),
			OptionsPreflightBypass:   khulnasoft.F(true),
			PathCookieAttribute:      khulnasoft.F(true),
			Policies: khulnasoft.F([]zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationPolicyUnion{zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
				ID:         khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				Precedence: khulnasoft.F(int64(0)),
			}, zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
				ID:         khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				Precedence: khulnasoft.F(int64(0)),
			}, zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
				ID:         khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				Precedence: khulnasoft.F(int64(0)),
			}}),
			SameSiteCookieAttribute: khulnasoft.F("strict"),
			SCIMConfig: khulnasoft.F(zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfig{
				IdPUID:    khulnasoft.F("idp_uid"),
				RemoteURI: khulnasoft.F("remote_uri"),
				Authentication: khulnasoft.F[zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion](zero_trust.SCIMConfigAuthenticationHTTPBasicParam{
					Password: khulnasoft.F("password"),
					Scheme:   khulnasoft.F(zero_trust.SCIMConfigAuthenticationHTTPBasicSchemeHttpbasic),
					User:     khulnasoft.F("user"),
				}),
				DeactivateOnDelete: khulnasoft.F(true),
				Enabled:            khulnasoft.F(true),
				Mappings: khulnasoft.F([]zero_trust.SCIMConfigMappingParam{{
					Schema:  khulnasoft.F("urn:ietf:params:scim:schemas:core:2.0:User"),
					Enabled: khulnasoft.F(true),
					Filter:  khulnasoft.F("title pr or userType eq \"Intern\""),
					Operations: khulnasoft.F(zero_trust.SCIMConfigMappingOperationsParam{
						Create: khulnasoft.F(true),
						Delete: khulnasoft.F(true),
						Update: khulnasoft.F(true),
					}),
					TransformJsonata: khulnasoft.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
				}, {
					Schema:  khulnasoft.F("urn:ietf:params:scim:schemas:core:2.0:User"),
					Enabled: khulnasoft.F(true),
					Filter:  khulnasoft.F("title pr or userType eq \"Intern\""),
					Operations: khulnasoft.F(zero_trust.SCIMConfigMappingOperationsParam{
						Create: khulnasoft.F(true),
						Delete: khulnasoft.F(true),
						Update: khulnasoft.F(true),
					}),
					TransformJsonata: khulnasoft.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
				}, {
					Schema:  khulnasoft.F("urn:ietf:params:scim:schemas:core:2.0:User"),
					Enabled: khulnasoft.F(true),
					Filter:  khulnasoft.F("title pr or userType eq \"Intern\""),
					Operations: khulnasoft.F(zero_trust.SCIMConfigMappingOperationsParam{
						Create: khulnasoft.F(true),
						Delete: khulnasoft.F(true),
						Update: khulnasoft.F(true),
					}),
					TransformJsonata: khulnasoft.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
				}}),
			}),
			SelfHostedDomains:      khulnasoft.F([]zero_trust.SelfHostedDomainsParam{"test.example.com/admin", "test.anotherexample.com/staff"}),
			ServiceAuth401Redirect: khulnasoft.F(true),
			SessionDuration:        khulnasoft.F("24h"),
			SkipInterstitial:       khulnasoft.F(true),
			Tags:                   khulnasoft.F([]string{"engineers", "engineers", "engineers"}),
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

func TestAccessApplicationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.AccessApplicationUpdateParams{
			Body: zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplication{
				Domain:                   khulnasoft.F("test.example.com/admin"),
				Type:                     khulnasoft.F("self_hosted"),
				AllowAuthenticateViaWARP: khulnasoft.F(true),
				AllowedIdPs:              khulnasoft.F([]zero_trust.AllowedIdPsParam{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
				AppLauncherVisible:       khulnasoft.F(true),
				AutoRedirectToIdentity:   khulnasoft.F(true),
				CORSHeaders: khulnasoft.F(zero_trust.CORSHeadersParam{
					AllowAllHeaders:  khulnasoft.F(true),
					AllowAllMethods:  khulnasoft.F(true),
					AllowAllOrigins:  khulnasoft.F(true),
					AllowCredentials: khulnasoft.F(true),
					AllowedHeaders:   khulnasoft.F([]zero_trust.AllowedHeadersParam{"string", "string", "string"}),
					AllowedMethods:   khulnasoft.F([]zero_trust.AllowedMethods{zero_trust.AllowedMethodsGet}),
					AllowedOrigins:   khulnasoft.F([]zero_trust.AllowedOriginsParam{"https://example.com"}),
					MaxAge:           khulnasoft.F(-1.000000),
				}),
				CustomDenyMessage:        khulnasoft.F("custom_deny_message"),
				CustomDenyURL:            khulnasoft.F("custom_deny_url"),
				CustomNonIdentityDenyURL: khulnasoft.F("custom_non_identity_deny_url"),
				CustomPages:              khulnasoft.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
				EnableBindingCookie:      khulnasoft.F(true),
				HTTPOnlyCookieAttribute:  khulnasoft.F(true),
				LogoURL:                  khulnasoft.F("https://www.khulnasoft.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
				Name:                     khulnasoft.F("Admin Site"),
				OptionsPreflightBypass:   khulnasoft.F(true),
				PathCookieAttribute:      khulnasoft.F(true),
				Policies: khulnasoft.F([]zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationPolicyUnion{zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
					ID:         khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					Precedence: khulnasoft.F(int64(0)),
				}, zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
					ID:         khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					Precedence: khulnasoft.F(int64(0)),
				}, zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
					ID:         khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					Precedence: khulnasoft.F(int64(0)),
				}}),
				SameSiteCookieAttribute: khulnasoft.F("strict"),
				SCIMConfig: khulnasoft.F(zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfig{
					IdPUID:    khulnasoft.F("idp_uid"),
					RemoteURI: khulnasoft.F("remote_uri"),
					Authentication: khulnasoft.F[zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion](zero_trust.SCIMConfigAuthenticationHTTPBasicParam{
						Password: khulnasoft.F("password"),
						Scheme:   khulnasoft.F(zero_trust.SCIMConfigAuthenticationHTTPBasicSchemeHttpbasic),
						User:     khulnasoft.F("user"),
					}),
					DeactivateOnDelete: khulnasoft.F(true),
					Enabled:            khulnasoft.F(true),
					Mappings: khulnasoft.F([]zero_trust.SCIMConfigMappingParam{{
						Schema:  khulnasoft.F("urn:ietf:params:scim:schemas:core:2.0:User"),
						Enabled: khulnasoft.F(true),
						Filter:  khulnasoft.F("title pr or userType eq \"Intern\""),
						Operations: khulnasoft.F(zero_trust.SCIMConfigMappingOperationsParam{
							Create: khulnasoft.F(true),
							Delete: khulnasoft.F(true),
							Update: khulnasoft.F(true),
						}),
						TransformJsonata: khulnasoft.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					}, {
						Schema:  khulnasoft.F("urn:ietf:params:scim:schemas:core:2.0:User"),
						Enabled: khulnasoft.F(true),
						Filter:  khulnasoft.F("title pr or userType eq \"Intern\""),
						Operations: khulnasoft.F(zero_trust.SCIMConfigMappingOperationsParam{
							Create: khulnasoft.F(true),
							Delete: khulnasoft.F(true),
							Update: khulnasoft.F(true),
						}),
						TransformJsonata: khulnasoft.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					}, {
						Schema:  khulnasoft.F("urn:ietf:params:scim:schemas:core:2.0:User"),
						Enabled: khulnasoft.F(true),
						Filter:  khulnasoft.F("title pr or userType eq \"Intern\""),
						Operations: khulnasoft.F(zero_trust.SCIMConfigMappingOperationsParam{
							Create: khulnasoft.F(true),
							Delete: khulnasoft.F(true),
							Update: khulnasoft.F(true),
						}),
						TransformJsonata: khulnasoft.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					}}),
				}),
				SelfHostedDomains:      khulnasoft.F([]zero_trust.SelfHostedDomainsParam{"test.example.com/admin", "test.anotherexample.com/staff"}),
				ServiceAuth401Redirect: khulnasoft.F(true),
				SessionDuration:        khulnasoft.F("24h"),
				SkipInterstitial:       khulnasoft.F(true),
				Tags:                   khulnasoft.F([]string{"engineers", "engineers", "engineers"}),
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

func TestAccessApplicationListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.List(context.TODO(), zero_trust.AccessApplicationListParams{
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

func TestAccessApplicationDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.AccessApplicationDeleteParams{
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

func TestAccessApplicationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.AccessApplicationGetParams{
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

func TestAccessApplicationRevokeTokensWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.RevokeTokens(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.AccessApplicationRevokeTokensParams{
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
