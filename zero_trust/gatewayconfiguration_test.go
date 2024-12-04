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

func TestGatewayConfigurationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Configurations.Update(context.TODO(), zero_trust.GatewayConfigurationUpdateParams{
		AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		Settings: khulnasoft.F(zero_trust.GatewayConfigurationSettingsParam{
			ActivityLog: khulnasoft.F(zero_trust.ActivityLogSettingsParam{
				Enabled: khulnasoft.F(true),
			}),
			Antivirus: khulnasoft.F(zero_trust.AntiVirusSettingsParam{
				EnabledDownloadPhase: khulnasoft.F(false),
				EnabledUploadPhase:   khulnasoft.F(false),
				FailClosed:           khulnasoft.F(false),
				NotificationSettings: khulnasoft.F(zero_trust.NotificationSettingsParam{
					Enabled:    khulnasoft.F(true),
					Msg:        khulnasoft.F("msg"),
					SupportURL: khulnasoft.F("support_url"),
				}),
			}),
			BlockPage: khulnasoft.F(zero_trust.BlockPageSettingsParam{
				BackgroundColor: khulnasoft.F("background_color"),
				Enabled:         khulnasoft.F(true),
				FooterText:      khulnasoft.F("--footer--"),
				HeaderText:      khulnasoft.F("--header--"),
				LogoPath:        khulnasoft.F("https://logos.com/a.png"),
				MailtoAddress:   khulnasoft.F("admin@example.com"),
				MailtoSubject:   khulnasoft.F("Blocked User Inquiry"),
				Name:            khulnasoft.F("Khulnasoft"),
				SuppressFooter:  khulnasoft.F(false),
			}),
			BodyScanning: khulnasoft.F(zero_trust.BodyScanningSettingsParam{
				InspectionMode: khulnasoft.F("deep"),
			}),
			BrowserIsolation: khulnasoft.F(zero_trust.BrowserIsolationSettingsParam{
				NonIdentityEnabled:         khulnasoft.F(true),
				URLBrowserIsolationEnabled: khulnasoft.F(true),
			}),
			Certificate: khulnasoft.F(zero_trust.GatewayConfigurationSettingsCertificateParam{
				ID: khulnasoft.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
			}),
			CustomCertificate: khulnasoft.F(zero_trust.CustomCertificateSettingsParam{
				Enabled: khulnasoft.F(true),
				ID:      khulnasoft.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
			}),
			ExtendedEmailMatching: khulnasoft.F(zero_trust.ExtendedEmailMatchingParam{
				Enabled: khulnasoft.F(true),
			}),
			Fips: khulnasoft.F(zero_trust.FipsSettingsParam{
				TLS: khulnasoft.F(true),
			}),
			ProtocolDetection: khulnasoft.F(zero_trust.ProtocolDetectionParam{
				Enabled: khulnasoft.F(true),
			}),
			TLSDecrypt: khulnasoft.F(zero_trust.TLSSettingsParam{
				Enabled: khulnasoft.F(true),
			}),
		}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayConfigurationEditWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Configurations.Edit(context.TODO(), zero_trust.GatewayConfigurationEditParams{
		AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		Settings: khulnasoft.F(zero_trust.GatewayConfigurationSettingsParam{
			ActivityLog: khulnasoft.F(zero_trust.ActivityLogSettingsParam{
				Enabled: khulnasoft.F(true),
			}),
			Antivirus: khulnasoft.F(zero_trust.AntiVirusSettingsParam{
				EnabledDownloadPhase: khulnasoft.F(false),
				EnabledUploadPhase:   khulnasoft.F(false),
				FailClosed:           khulnasoft.F(false),
				NotificationSettings: khulnasoft.F(zero_trust.NotificationSettingsParam{
					Enabled:    khulnasoft.F(true),
					Msg:        khulnasoft.F("msg"),
					SupportURL: khulnasoft.F("support_url"),
				}),
			}),
			BlockPage: khulnasoft.F(zero_trust.BlockPageSettingsParam{
				BackgroundColor: khulnasoft.F("background_color"),
				Enabled:         khulnasoft.F(true),
				FooterText:      khulnasoft.F("--footer--"),
				HeaderText:      khulnasoft.F("--header--"),
				LogoPath:        khulnasoft.F("https://logos.com/a.png"),
				MailtoAddress:   khulnasoft.F("admin@example.com"),
				MailtoSubject:   khulnasoft.F("Blocked User Inquiry"),
				Name:            khulnasoft.F("Khulnasoft"),
				SuppressFooter:  khulnasoft.F(false),
			}),
			BodyScanning: khulnasoft.F(zero_trust.BodyScanningSettingsParam{
				InspectionMode: khulnasoft.F("deep"),
			}),
			BrowserIsolation: khulnasoft.F(zero_trust.BrowserIsolationSettingsParam{
				NonIdentityEnabled:         khulnasoft.F(true),
				URLBrowserIsolationEnabled: khulnasoft.F(true),
			}),
			Certificate: khulnasoft.F(zero_trust.GatewayConfigurationSettingsCertificateParam{
				ID: khulnasoft.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
			}),
			CustomCertificate: khulnasoft.F(zero_trust.CustomCertificateSettingsParam{
				Enabled: khulnasoft.F(true),
				ID:      khulnasoft.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
			}),
			ExtendedEmailMatching: khulnasoft.F(zero_trust.ExtendedEmailMatchingParam{
				Enabled: khulnasoft.F(true),
			}),
			Fips: khulnasoft.F(zero_trust.FipsSettingsParam{
				TLS: khulnasoft.F(true),
			}),
			ProtocolDetection: khulnasoft.F(zero_trust.ProtocolDetectionParam{
				Enabled: khulnasoft.F(true),
			}),
			TLSDecrypt: khulnasoft.F(zero_trust.TLSSettingsParam{
				Enabled: khulnasoft.F(true),
			}),
		}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayConfigurationGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Configurations.Get(context.TODO(), zero_trust.GatewayConfigurationGetParams{
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
