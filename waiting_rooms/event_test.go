// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/waiting_rooms"
)

func TestEventNewWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.Events.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.EventNewParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EventQuery: waiting_rooms.EventQueryParam{
				EventEndTime:          khulnasoft.F("2021-09-28T17:00:00.000Z"),
				EventStartTime:        khulnasoft.F("2021-09-28T15:30:00.000Z"),
				Name:                  khulnasoft.F("production_webinar_event"),
				CustomPageHTML:        khulnasoft.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Event is prequeueing / Queue all enabled {{/waitTimeKnown}}"),
				Description:           khulnasoft.F("Production event - DO NOT MODIFY"),
				DisableSessionRenewal: khulnasoft.F(true),
				NewUsersPerMinute:     khulnasoft.F(int64(200)),
				PrequeueStartTime:     khulnasoft.F("2021-09-28T15:00:00.000Z"),
				QueueingMethod:        khulnasoft.F("random"),
				SessionDuration:       khulnasoft.F(int64(1)),
				ShuffleAtEventStart:   khulnasoft.F(true),
				Suspended:             khulnasoft.F(true),
				TotalActiveUsers:      khulnasoft.F(int64(200)),
			},
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

func TestEventUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.Events.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		waiting_rooms.EventUpdateParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EventQuery: waiting_rooms.EventQueryParam{
				EventEndTime:          khulnasoft.F("2021-09-28T17:00:00.000Z"),
				EventStartTime:        khulnasoft.F("2021-09-28T15:30:00.000Z"),
				Name:                  khulnasoft.F("production_webinar_event"),
				CustomPageHTML:        khulnasoft.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Event is prequeueing / Queue all enabled {{/waitTimeKnown}}"),
				Description:           khulnasoft.F("Production event - DO NOT MODIFY"),
				DisableSessionRenewal: khulnasoft.F(true),
				NewUsersPerMinute:     khulnasoft.F(int64(200)),
				PrequeueStartTime:     khulnasoft.F("2021-09-28T15:00:00.000Z"),
				QueueingMethod:        khulnasoft.F("random"),
				SessionDuration:       khulnasoft.F(int64(1)),
				ShuffleAtEventStart:   khulnasoft.F(true),
				Suspended:             khulnasoft.F(true),
				TotalActiveUsers:      khulnasoft.F(int64(200)),
			},
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

func TestEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.Events.List(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.EventListParams{
			ZoneID:  khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Page:    khulnasoft.F(1.000000),
			PerPage: khulnasoft.F(5.000000),
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

func TestEventDelete(t *testing.T) {
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
	_, err := client.WaitingRooms.Events.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		waiting_rooms.EventDeleteParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestEventEditWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.Events.Edit(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		waiting_rooms.EventEditParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EventQuery: waiting_rooms.EventQueryParam{
				EventEndTime:          khulnasoft.F("2021-09-28T17:00:00.000Z"),
				EventStartTime:        khulnasoft.F("2021-09-28T15:30:00.000Z"),
				Name:                  khulnasoft.F("production_webinar_event"),
				CustomPageHTML:        khulnasoft.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Event is prequeueing / Queue all enabled {{/waitTimeKnown}}"),
				Description:           khulnasoft.F("Production event - DO NOT MODIFY"),
				DisableSessionRenewal: khulnasoft.F(true),
				NewUsersPerMinute:     khulnasoft.F(int64(200)),
				PrequeueStartTime:     khulnasoft.F("2021-09-28T15:00:00.000Z"),
				QueueingMethod:        khulnasoft.F("random"),
				SessionDuration:       khulnasoft.F(int64(1)),
				ShuffleAtEventStart:   khulnasoft.F(true),
				Suspended:             khulnasoft.F(true),
				TotalActiveUsers:      khulnasoft.F(int64(200)),
			},
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

func TestEventGet(t *testing.T) {
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
	_, err := client.WaitingRooms.Events.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		waiting_rooms.EventGetParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
