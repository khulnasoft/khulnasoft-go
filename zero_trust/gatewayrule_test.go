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

func TestGatewayRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.New(context.TODO(), zero_trust.GatewayRuleNewParams{
		AccountID:     khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		Action:        khulnasoft.F(zero_trust.GatewayRuleNewParamsActionOn),
		Name:          khulnasoft.F("block bad websites"),
		Description:   khulnasoft.F("Block bad websites based on their host name."),
		DevicePosture: khulnasoft.F("any(device_posture.checks.passed[*] in {\"1308749e-fcfb-4ebc-b051-fe022b632644\"})"),
		Enabled:       khulnasoft.F(true),
		Filters:       khulnasoft.F([]zero_trust.GatewayFilter{zero_trust.GatewayFilterHTTP}),
		Identity:      khulnasoft.F("any(identity.groups.name[*] in {\"finance\"})"),
		Precedence:    khulnasoft.F(int64(0)),
		RuleSettings: khulnasoft.F(zero_trust.RuleSettingParam{
			AddHeaders: khulnasoft.F(map[string]string{
				"foo": "string",
			}),
			AllowChildBypass: khulnasoft.F(false),
			AuditSSH: khulnasoft.F(zero_trust.RuleSettingAuditSSHParam{
				CommandLogging: khulnasoft.F(false),
			}),
			BISOAdminControls: khulnasoft.F(zero_trust.RuleSettingBISOAdminControlsParam{
				DCP: khulnasoft.F(false),
				DD:  khulnasoft.F(false),
				DK:  khulnasoft.F(false),
				DP:  khulnasoft.F(false),
				DU:  khulnasoft.F(false),
			}),
			BlockPageEnabled: khulnasoft.F(true),
			BlockReason:      khulnasoft.F("This website is a security risk"),
			BypassParentRule: khulnasoft.F(false),
			CheckSession: khulnasoft.F(zero_trust.RuleSettingCheckSessionParam{
				Duration: khulnasoft.F("300s"),
				Enforce:  khulnasoft.F(true),
			}),
			DNSResolvers: khulnasoft.F(zero_trust.RuleSettingDNSResolversParam{
				IPV4: khulnasoft.F([]zero_trust.DNSResolverSettingsV4Param{{
					IP:                         khulnasoft.F("2.2.2.2"),
					Port:                       khulnasoft.F(int64(5053)),
					RouteThroughPrivateNetwork: khulnasoft.F(true),
					VnetID:                     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         khulnasoft.F("2.2.2.2"),
					Port:                       khulnasoft.F(int64(5053)),
					RouteThroughPrivateNetwork: khulnasoft.F(true),
					VnetID:                     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         khulnasoft.F("2.2.2.2"),
					Port:                       khulnasoft.F(int64(5053)),
					RouteThroughPrivateNetwork: khulnasoft.F(true),
					VnetID:                     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}}),
				IPV6: khulnasoft.F([]zero_trust.DNSResolverSettingsV6Param{{
					IP:                         khulnasoft.F("2001:DB8::"),
					Port:                       khulnasoft.F(int64(5053)),
					RouteThroughPrivateNetwork: khulnasoft.F(true),
					VnetID:                     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         khulnasoft.F("2001:DB8::"),
					Port:                       khulnasoft.F(int64(5053)),
					RouteThroughPrivateNetwork: khulnasoft.F(true),
					VnetID:                     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         khulnasoft.F("2001:DB8::"),
					Port:                       khulnasoft.F(int64(5053)),
					RouteThroughPrivateNetwork: khulnasoft.F(true),
					VnetID:                     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}}),
			}),
			Egress: khulnasoft.F(zero_trust.RuleSettingEgressParam{
				IPV4:         khulnasoft.F("192.0.2.2"),
				IPV4Fallback: khulnasoft.F("192.0.2.3"),
				IPV6:         khulnasoft.F("2001:DB8::/64"),
			}),
			IgnoreCNAMECategoryMatches:      khulnasoft.F(true),
			InsecureDisableDNSSECValidation: khulnasoft.F(false),
			IPCategories:                    khulnasoft.F(true),
			IPIndicatorFeeds:                khulnasoft.F(true),
			L4override: khulnasoft.F(zero_trust.RuleSettingL4overrideParam{
				IP:   khulnasoft.F("1.1.1.1"),
				Port: khulnasoft.F(int64(0)),
			}),
			NotificationSettings: khulnasoft.F(zero_trust.RuleSettingNotificationSettingsParam{
				Enabled:    khulnasoft.F(true),
				Msg:        khulnasoft.F("msg"),
				SupportURL: khulnasoft.F("support_url"),
			}),
			OverrideHost: khulnasoft.F("example.com"),
			OverrideIPs:  khulnasoft.F([]string{"1.1.1.1", "2.2.2.2"}),
			PayloadLog: khulnasoft.F(zero_trust.RuleSettingPayloadLogParam{
				Enabled: khulnasoft.F(true),
			}),
			ResolveDNSThroughCloudflare: khulnasoft.F(true),
			UntrustedCERT: khulnasoft.F(zero_trust.RuleSettingUntrustedCERTParam{
				Action: khulnasoft.F(zero_trust.RuleSettingUntrustedCERTActionPassThrough),
			}),
		}),
		Schedule: khulnasoft.F(zero_trust.ScheduleParam{
			Fri:      khulnasoft.F("08:00-12:30,13:30-17:00"),
			Mon:      khulnasoft.F("08:00-12:30,13:30-17:00"),
			Sat:      khulnasoft.F("08:00-12:30,13:30-17:00"),
			Sun:      khulnasoft.F("08:00-12:30,13:30-17:00"),
			Thu:      khulnasoft.F("08:00-12:30,13:30-17:00"),
			TimeZone: khulnasoft.F("America/New York"),
			Tue:      khulnasoft.F("08:00-12:30,13:30-17:00"),
			Wed:      khulnasoft.F("08:00-12:30,13:30-17:00"),
		}),
		Traffic: khulnasoft.F("http.request.uri matches \".*a/partial/uri.*\" and http.request.host in $01302951-49f9-47c9-a400-0297e60b6a10"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.GatewayRuleUpdateParams{
			AccountID:     khulnasoft.F("699d98642c564d2e855e9661899b7252"),
			Action:        khulnasoft.F(zero_trust.GatewayRuleUpdateParamsActionOn),
			Name:          khulnasoft.F("block bad websites"),
			Description:   khulnasoft.F("Block bad websites based on their host name."),
			DevicePosture: khulnasoft.F("any(device_posture.checks.passed[*] in {\"1308749e-fcfb-4ebc-b051-fe022b632644\"})"),
			Enabled:       khulnasoft.F(true),
			Filters:       khulnasoft.F([]zero_trust.GatewayFilter{zero_trust.GatewayFilterHTTP}),
			Identity:      khulnasoft.F("any(identity.groups.name[*] in {\"finance\"})"),
			Precedence:    khulnasoft.F(int64(0)),
			RuleSettings: khulnasoft.F(zero_trust.RuleSettingParam{
				AddHeaders: khulnasoft.F(map[string]string{
					"foo": "string",
				}),
				AllowChildBypass: khulnasoft.F(false),
				AuditSSH: khulnasoft.F(zero_trust.RuleSettingAuditSSHParam{
					CommandLogging: khulnasoft.F(false),
				}),
				BISOAdminControls: khulnasoft.F(zero_trust.RuleSettingBISOAdminControlsParam{
					DCP: khulnasoft.F(false),
					DD:  khulnasoft.F(false),
					DK:  khulnasoft.F(false),
					DP:  khulnasoft.F(false),
					DU:  khulnasoft.F(false),
				}),
				BlockPageEnabled: khulnasoft.F(true),
				BlockReason:      khulnasoft.F("This website is a security risk"),
				BypassParentRule: khulnasoft.F(false),
				CheckSession: khulnasoft.F(zero_trust.RuleSettingCheckSessionParam{
					Duration: khulnasoft.F("300s"),
					Enforce:  khulnasoft.F(true),
				}),
				DNSResolvers: khulnasoft.F(zero_trust.RuleSettingDNSResolversParam{
					IPV4: khulnasoft.F([]zero_trust.DNSResolverSettingsV4Param{{
						IP:                         khulnasoft.F("2.2.2.2"),
						Port:                       khulnasoft.F(int64(5053)),
						RouteThroughPrivateNetwork: khulnasoft.F(true),
						VnetID:                     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         khulnasoft.F("2.2.2.2"),
						Port:                       khulnasoft.F(int64(5053)),
						RouteThroughPrivateNetwork: khulnasoft.F(true),
						VnetID:                     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         khulnasoft.F("2.2.2.2"),
						Port:                       khulnasoft.F(int64(5053)),
						RouteThroughPrivateNetwork: khulnasoft.F(true),
						VnetID:                     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
					IPV6: khulnasoft.F([]zero_trust.DNSResolverSettingsV6Param{{
						IP:                         khulnasoft.F("2001:DB8::"),
						Port:                       khulnasoft.F(int64(5053)),
						RouteThroughPrivateNetwork: khulnasoft.F(true),
						VnetID:                     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         khulnasoft.F("2001:DB8::"),
						Port:                       khulnasoft.F(int64(5053)),
						RouteThroughPrivateNetwork: khulnasoft.F(true),
						VnetID:                     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         khulnasoft.F("2001:DB8::"),
						Port:                       khulnasoft.F(int64(5053)),
						RouteThroughPrivateNetwork: khulnasoft.F(true),
						VnetID:                     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
				}),
				Egress: khulnasoft.F(zero_trust.RuleSettingEgressParam{
					IPV4:         khulnasoft.F("192.0.2.2"),
					IPV4Fallback: khulnasoft.F("192.0.2.3"),
					IPV6:         khulnasoft.F("2001:DB8::/64"),
				}),
				IgnoreCNAMECategoryMatches:      khulnasoft.F(true),
				InsecureDisableDNSSECValidation: khulnasoft.F(false),
				IPCategories:                    khulnasoft.F(true),
				IPIndicatorFeeds:                khulnasoft.F(true),
				L4override: khulnasoft.F(zero_trust.RuleSettingL4overrideParam{
					IP:   khulnasoft.F("1.1.1.1"),
					Port: khulnasoft.F(int64(0)),
				}),
				NotificationSettings: khulnasoft.F(zero_trust.RuleSettingNotificationSettingsParam{
					Enabled:    khulnasoft.F(true),
					Msg:        khulnasoft.F("msg"),
					SupportURL: khulnasoft.F("support_url"),
				}),
				OverrideHost: khulnasoft.F("example.com"),
				OverrideIPs:  khulnasoft.F([]string{"1.1.1.1", "2.2.2.2"}),
				PayloadLog: khulnasoft.F(zero_trust.RuleSettingPayloadLogParam{
					Enabled: khulnasoft.F(true),
				}),
				ResolveDNSThroughCloudflare: khulnasoft.F(true),
				UntrustedCERT: khulnasoft.F(zero_trust.RuleSettingUntrustedCERTParam{
					Action: khulnasoft.F(zero_trust.RuleSettingUntrustedCERTActionPassThrough),
				}),
			}),
			Schedule: khulnasoft.F(zero_trust.ScheduleParam{
				Fri:      khulnasoft.F("08:00-12:30,13:30-17:00"),
				Mon:      khulnasoft.F("08:00-12:30,13:30-17:00"),
				Sat:      khulnasoft.F("08:00-12:30,13:30-17:00"),
				Sun:      khulnasoft.F("08:00-12:30,13:30-17:00"),
				Thu:      khulnasoft.F("08:00-12:30,13:30-17:00"),
				TimeZone: khulnasoft.F("America/New York"),
				Tue:      khulnasoft.F("08:00-12:30,13:30-17:00"),
				Wed:      khulnasoft.F("08:00-12:30,13:30-17:00"),
			}),
			Traffic: khulnasoft.F("http.request.uri matches \".*a/partial/uri.*\" and http.request.host in $01302951-49f9-47c9-a400-0297e60b6a10"),
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

func TestGatewayRuleList(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.List(context.TODO(), zero_trust.GatewayRuleListParams{
		AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayRuleDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.GatewayRuleDeleteParams{
			AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
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

func TestGatewayRuleGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.GatewayRuleGetParams{
			AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
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
