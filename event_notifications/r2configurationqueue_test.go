// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event_notifications_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/event_notifications"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestR2ConfigurationQueueUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.EventNotifications.R2.Configuration.Queues.Update(
		context.TODO(),
		"example-bucket",
		"queue_id",
		event_notifications.R2ConfigurationQueueUpdateParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Rules: khulnasoft.F([]event_notifications.R2ConfigurationQueueUpdateParamsRule{{
				Actions: khulnasoft.F([]event_notifications.R2ConfigurationQueueUpdateParamsRulesAction{event_notifications.R2ConfigurationQueueUpdateParamsRulesActionPutObject, event_notifications.R2ConfigurationQueueUpdateParamsRulesActionCopyObject}),
				Prefix:  khulnasoft.F("img/"),
				Suffix:  khulnasoft.F(".jpeg"),
			}, {
				Actions: khulnasoft.F([]event_notifications.R2ConfigurationQueueUpdateParamsRulesAction{event_notifications.R2ConfigurationQueueUpdateParamsRulesActionPutObject, event_notifications.R2ConfigurationQueueUpdateParamsRulesActionCopyObject}),
				Prefix:  khulnasoft.F("img/"),
				Suffix:  khulnasoft.F(".jpeg"),
			}, {
				Actions: khulnasoft.F([]event_notifications.R2ConfigurationQueueUpdateParamsRulesAction{event_notifications.R2ConfigurationQueueUpdateParamsRulesActionPutObject, event_notifications.R2ConfigurationQueueUpdateParamsRulesActionCopyObject}),
				Prefix:  khulnasoft.F("img/"),
				Suffix:  khulnasoft.F(".jpeg"),
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

func TestR2ConfigurationQueueDelete(t *testing.T) {
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
	_, err := client.EventNotifications.R2.Configuration.Queues.Delete(
		context.TODO(),
		"example-bucket",
		"queue_id",
		event_notifications.R2ConfigurationQueueDeleteParams{
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
