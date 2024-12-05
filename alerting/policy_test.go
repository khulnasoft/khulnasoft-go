// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/alerting"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestPolicyNewWithOptionalParams(t *testing.T) {
	t.Skip("prism errors - https://github.com/khulnasoft/khulnasoft-python/actions/runs/9327225061/job/25676826349?pr=482#step:5:4274")
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
	_, err := client.Alerting.Policies.New(context.TODO(), alerting.PolicyNewParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		AlertType: khulnasoft.F(alerting.PolicyNewParamsAlertTypeAccessCustomCertificateExpirationType),
		Enabled:   khulnasoft.F(true),
		Mechanisms: khulnasoft.F(alerting.MechanismParam{
			"email": []alerting.MechanismItemParam{{
				ID: khulnasoft.F("test@example.com"),
			}},
			"pagerduty": []alerting.MechanismItemParam{{
				ID: khulnasoft.F("e8133a15-00a4-4d69-aec1-32f70c51f6e5"),
			}},
			"webhooks": []alerting.MechanismItemParam{{
				ID: khulnasoft.F("14cc1190-5d2b-4b98-a696-c424cb2ad05f"),
			}},
		}),
		Name:          khulnasoft.F("SSL Notification Event Policy"),
		AlertInterval: khulnasoft.F("30m"),
		Description:   khulnasoft.F("Something describing the policy."),
		Filters: khulnasoft.F(alerting.PolicyFilterParam{
			Actions:                      khulnasoft.F([]string{"string", "string", "string"}),
			AffectedASNs:                 khulnasoft.F([]string{"string", "string", "string"}),
			AffectedComponents:           khulnasoft.F([]string{"string", "string", "string"}),
			AffectedLocations:            khulnasoft.F([]string{"string", "string", "string"}),
			AirportCode:                  khulnasoft.F([]string{"string", "string", "string"}),
			AlertTriggerPreferences:      khulnasoft.F([]string{"string", "string", "string"}),
			AlertTriggerPreferencesValue: khulnasoft.F([]string{"string", "string", "string"}),
			Enabled:                      khulnasoft.F([]string{"string", "string", "string"}),
			Environment:                  khulnasoft.F([]string{"string", "string", "string"}),
			Event:                        khulnasoft.F([]string{"string", "string", "string"}),
			EventSource:                  khulnasoft.F([]string{"string", "string", "string"}),
			EventType:                    khulnasoft.F([]string{"string", "string", "string"}),
			GroupBy:                      khulnasoft.F([]string{"string", "string", "string"}),
			HealthCheckID:                khulnasoft.F([]string{"string", "string", "string"}),
			IncidentImpact:               khulnasoft.F([]alerting.PolicyFilterIncidentImpact{alerting.PolicyFilterIncidentImpactIncidentImpactNone, alerting.PolicyFilterIncidentImpactIncidentImpactMinor, alerting.PolicyFilterIncidentImpactIncidentImpactMajor}),
			InputID:                      khulnasoft.F([]string{"string", "string", "string"}),
			Limit:                        khulnasoft.F([]string{"string", "string", "string"}),
			LogoTag:                      khulnasoft.F([]string{"string", "string", "string"}),
			MegabitsPerSecond:            khulnasoft.F([]string{"string", "string", "string"}),
			NewHealth:                    khulnasoft.F([]string{"string", "string", "string"}),
			NewStatus:                    khulnasoft.F([]string{"string", "string", "string"}),
			PacketsPerSecond:             khulnasoft.F([]string{"string", "string", "string"}),
			PoolID:                       khulnasoft.F([]string{"string", "string", "string"}),
			PopName:                      khulnasoft.F([]string{"string", "string", "string"}),
			Product:                      khulnasoft.F([]string{"string", "string", "string"}),
			ProjectID:                    khulnasoft.F([]string{"string", "string", "string"}),
			Protocol:                     khulnasoft.F([]string{"string", "string", "string"}),
			QueryTag:                     khulnasoft.F([]string{"string", "string", "string"}),
			RequestsPerSecond:            khulnasoft.F([]string{"string", "string", "string"}),
			Selectors:                    khulnasoft.F([]string{"string", "string", "string"}),
			Services:                     khulnasoft.F([]string{"string", "string", "string"}),
			Slo:                          khulnasoft.F([]string{"99.9"}),
			Status:                       khulnasoft.F([]string{"string", "string", "string"}),
			TargetHostname:               khulnasoft.F([]string{"string", "string", "string"}),
			TargetIP:                     khulnasoft.F([]string{"string", "string", "string"}),
			TargetZoneName:               khulnasoft.F([]string{"string", "string", "string"}),
			TrafficExclusions:            khulnasoft.F([]alerting.PolicyFilterTrafficExclusion{alerting.PolicyFilterTrafficExclusionSecurityEvents}),
			TunnelID:                     khulnasoft.F([]string{"string", "string", "string"}),
			TunnelName:                   khulnasoft.F([]string{"string", "string", "string"}),
			Where:                        khulnasoft.F([]string{"string", "string", "string"}),
			Zones:                        khulnasoft.F([]string{"string", "string", "string"}),
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

func TestPolicyUpdateWithOptionalParams(t *testing.T) {
	t.Skip("prism errors - https://github.com/khulnasoft/khulnasoft-python/actions/runs/9327225061/job/25676826349?pr=482#step:5:4274")
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
	_, err := client.Alerting.Policies.Update(
		context.TODO(),
		"0da2b59e-f118-439d-8097-bdfb215203c9",
		alerting.PolicyUpdateParams{
			AccountID:     khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AlertInterval: khulnasoft.F("30m"),
			AlertType:     khulnasoft.F(alerting.PolicyUpdateParamsAlertTypeAccessCustomCertificateExpirationType),
			Description:   khulnasoft.F("Something describing the policy."),
			Enabled:       khulnasoft.F(true),
			Filters: khulnasoft.F(alerting.PolicyFilterParam{
				Actions:                      khulnasoft.F([]string{"string", "string", "string"}),
				AffectedASNs:                 khulnasoft.F([]string{"string", "string", "string"}),
				AffectedComponents:           khulnasoft.F([]string{"string", "string", "string"}),
				AffectedLocations:            khulnasoft.F([]string{"string", "string", "string"}),
				AirportCode:                  khulnasoft.F([]string{"string", "string", "string"}),
				AlertTriggerPreferences:      khulnasoft.F([]string{"string", "string", "string"}),
				AlertTriggerPreferencesValue: khulnasoft.F([]string{"string", "string", "string"}),
				Enabled:                      khulnasoft.F([]string{"string", "string", "string"}),
				Environment:                  khulnasoft.F([]string{"string", "string", "string"}),
				Event:                        khulnasoft.F([]string{"string", "string", "string"}),
				EventSource:                  khulnasoft.F([]string{"string", "string", "string"}),
				EventType:                    khulnasoft.F([]string{"string", "string", "string"}),
				GroupBy:                      khulnasoft.F([]string{"string", "string", "string"}),
				HealthCheckID:                khulnasoft.F([]string{"string", "string", "string"}),
				IncidentImpact:               khulnasoft.F([]alerting.PolicyFilterIncidentImpact{alerting.PolicyFilterIncidentImpactIncidentImpactNone, alerting.PolicyFilterIncidentImpactIncidentImpactMinor, alerting.PolicyFilterIncidentImpactIncidentImpactMajor}),
				InputID:                      khulnasoft.F([]string{"string", "string", "string"}),
				Limit:                        khulnasoft.F([]string{"string", "string", "string"}),
				LogoTag:                      khulnasoft.F([]string{"string", "string", "string"}),
				MegabitsPerSecond:            khulnasoft.F([]string{"string", "string", "string"}),
				NewHealth:                    khulnasoft.F([]string{"string", "string", "string"}),
				NewStatus:                    khulnasoft.F([]string{"string", "string", "string"}),
				PacketsPerSecond:             khulnasoft.F([]string{"string", "string", "string"}),
				PoolID:                       khulnasoft.F([]string{"string", "string", "string"}),
				PopName:                      khulnasoft.F([]string{"string", "string", "string"}),
				Product:                      khulnasoft.F([]string{"string", "string", "string"}),
				ProjectID:                    khulnasoft.F([]string{"string", "string", "string"}),
				Protocol:                     khulnasoft.F([]string{"string", "string", "string"}),
				QueryTag:                     khulnasoft.F([]string{"string", "string", "string"}),
				RequestsPerSecond:            khulnasoft.F([]string{"string", "string", "string"}),
				Selectors:                    khulnasoft.F([]string{"string", "string", "string"}),
				Services:                     khulnasoft.F([]string{"string", "string", "string"}),
				Slo:                          khulnasoft.F([]string{"99.9"}),
				Status:                       khulnasoft.F([]string{"string", "string", "string"}),
				TargetHostname:               khulnasoft.F([]string{"string", "string", "string"}),
				TargetIP:                     khulnasoft.F([]string{"string", "string", "string"}),
				TargetZoneName:               khulnasoft.F([]string{"string", "string", "string"}),
				TrafficExclusions:            khulnasoft.F([]alerting.PolicyFilterTrafficExclusion{alerting.PolicyFilterTrafficExclusionSecurityEvents}),
				TunnelID:                     khulnasoft.F([]string{"string", "string", "string"}),
				TunnelName:                   khulnasoft.F([]string{"string", "string", "string"}),
				Where:                        khulnasoft.F([]string{"string", "string", "string"}),
				Zones:                        khulnasoft.F([]string{"string", "string", "string"}),
			}),
			Mechanisms: khulnasoft.F(alerting.MechanismParam{
				"email": []alerting.MechanismItemParam{{
					ID: khulnasoft.F("test@example.com"),
				}},
				"pagerduty": []alerting.MechanismItemParam{{
					ID: khulnasoft.F("e8133a15-00a4-4d69-aec1-32f70c51f6e5"),
				}},
				"webhooks": []alerting.MechanismItemParam{{
					ID: khulnasoft.F("14cc1190-5d2b-4b98-a696-c424cb2ad05f"),
				}},
			}),
			Name: khulnasoft.F("SSL Notification Event Policy"),
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

func TestPolicyList(t *testing.T) {
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
	_, err := client.Alerting.Policies.List(context.TODO(), alerting.PolicyListParams{
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

func TestPolicyDelete(t *testing.T) {
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
	_, err := client.Alerting.Policies.Delete(
		context.TODO(),
		"0da2b59e-f118-439d-8097-bdfb215203c9",
		alerting.PolicyDeleteParams{
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

func TestPolicyGet(t *testing.T) {
	t.Skip("prism errors - https://github.com/khulnasoft/khulnasoft-python/actions/runs/9327225061/job/25676826349?pr=482#step:5:4274")
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
	_, err := client.Alerting.Policies.Get(
		context.TODO(),
		"0da2b59e-f118-439d-8097-bdfb215203c9",
		alerting.PolicyGetParams{
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
