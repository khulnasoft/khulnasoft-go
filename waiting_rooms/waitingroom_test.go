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

func TestWaitingRoomNewWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.New(context.TODO(), waiting_rooms.WaitingRoomNewParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Query: waiting_rooms.QueryParam{
			Host:              khulnasoft.F("shop.example.com"),
			Name:              khulnasoft.F("production_webinar"),
			NewUsersPerMinute: khulnasoft.F(int64(200)),
			TotalActiveUsers:  khulnasoft.F(int64(200)),
			AdditionalRoutes: khulnasoft.F([]waiting_rooms.AdditionalRoutesParam{{
				Host: khulnasoft.F("shop2.example.com"),
				Path: khulnasoft.F("/shop2/checkout"),
			}, {
				Host: khulnasoft.F("shop2.example.com"),
				Path: khulnasoft.F("/shop2/checkout"),
			}, {
				Host: khulnasoft.F("shop2.example.com"),
				Path: khulnasoft.F("/shop2/checkout"),
			}}),
			CookieAttributes: khulnasoft.F(waiting_rooms.CookieAttributesParam{
				Samesite: khulnasoft.F(waiting_rooms.CookieAttributesSamesiteAuto),
				Secure:   khulnasoft.F(waiting_rooms.CookieAttributesSecureAuto),
			}),
			CookieSuffix:            khulnasoft.F("abcd"),
			CustomPageHTML:          khulnasoft.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
			DefaultTemplateLanguage: khulnasoft.F(waiting_rooms.QueryDefaultTemplateLanguageEnUs),
			Description:             khulnasoft.F("Production - DO NOT MODIFY"),
			DisableSessionRenewal:   khulnasoft.F(false),
			EnabledOriginCommands:   khulnasoft.F([]waiting_rooms.QueryEnabledOriginCommand{waiting_rooms.QueryEnabledOriginCommandRevoke}),
			JsonResponseEnabled:     khulnasoft.F(false),
			Path:                    khulnasoft.F("/shop/checkout"),
			QueueAll:                khulnasoft.F(true),
			QueueingMethod:          khulnasoft.F(waiting_rooms.QueryQueueingMethodFifo),
			QueueingStatusCode:      khulnasoft.F(waiting_rooms.QueryQueueingStatusCode200),
			SessionDuration:         khulnasoft.F(int64(1)),
			Suspended:               khulnasoft.F(true),
		},
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaitingRoomUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.WaitingRoomUpdateParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Query: waiting_rooms.QueryParam{
				Host:              khulnasoft.F("shop.example.com"),
				Name:              khulnasoft.F("production_webinar"),
				NewUsersPerMinute: khulnasoft.F(int64(200)),
				TotalActiveUsers:  khulnasoft.F(int64(200)),
				AdditionalRoutes: khulnasoft.F([]waiting_rooms.AdditionalRoutesParam{{
					Host: khulnasoft.F("shop2.example.com"),
					Path: khulnasoft.F("/shop2/checkout"),
				}, {
					Host: khulnasoft.F("shop2.example.com"),
					Path: khulnasoft.F("/shop2/checkout"),
				}, {
					Host: khulnasoft.F("shop2.example.com"),
					Path: khulnasoft.F("/shop2/checkout"),
				}}),
				CookieAttributes: khulnasoft.F(waiting_rooms.CookieAttributesParam{
					Samesite: khulnasoft.F(waiting_rooms.CookieAttributesSamesiteAuto),
					Secure:   khulnasoft.F(waiting_rooms.CookieAttributesSecureAuto),
				}),
				CookieSuffix:            khulnasoft.F("abcd"),
				CustomPageHTML:          khulnasoft.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
				DefaultTemplateLanguage: khulnasoft.F(waiting_rooms.QueryDefaultTemplateLanguageEnUs),
				Description:             khulnasoft.F("Production - DO NOT MODIFY"),
				DisableSessionRenewal:   khulnasoft.F(false),
				EnabledOriginCommands:   khulnasoft.F([]waiting_rooms.QueryEnabledOriginCommand{waiting_rooms.QueryEnabledOriginCommandRevoke}),
				JsonResponseEnabled:     khulnasoft.F(false),
				Path:                    khulnasoft.F("/shop/checkout"),
				QueueAll:                khulnasoft.F(true),
				QueueingMethod:          khulnasoft.F(waiting_rooms.QueryQueueingMethodFifo),
				QueueingStatusCode:      khulnasoft.F(waiting_rooms.QueryQueueingStatusCode200),
				SessionDuration:         khulnasoft.F(int64(1)),
				Suspended:               khulnasoft.F(true),
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

func TestWaitingRoomListWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.List(context.TODO(), waiting_rooms.WaitingRoomListParams{
		ZoneID:  khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Page:    khulnasoft.F(1.000000),
		PerPage: khulnasoft.F(5.000000),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaitingRoomDelete(t *testing.T) {
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
	_, err := client.WaitingRooms.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.WaitingRoomDeleteParams{
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

func TestWaitingRoomEditWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.Edit(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.WaitingRoomEditParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Query: waiting_rooms.QueryParam{
				Host:              khulnasoft.F("shop.example.com"),
				Name:              khulnasoft.F("production_webinar"),
				NewUsersPerMinute: khulnasoft.F(int64(200)),
				TotalActiveUsers:  khulnasoft.F(int64(200)),
				AdditionalRoutes: khulnasoft.F([]waiting_rooms.AdditionalRoutesParam{{
					Host: khulnasoft.F("shop2.example.com"),
					Path: khulnasoft.F("/shop2/checkout"),
				}, {
					Host: khulnasoft.F("shop2.example.com"),
					Path: khulnasoft.F("/shop2/checkout"),
				}, {
					Host: khulnasoft.F("shop2.example.com"),
					Path: khulnasoft.F("/shop2/checkout"),
				}}),
				CookieAttributes: khulnasoft.F(waiting_rooms.CookieAttributesParam{
					Samesite: khulnasoft.F(waiting_rooms.CookieAttributesSamesiteAuto),
					Secure:   khulnasoft.F(waiting_rooms.CookieAttributesSecureAuto),
				}),
				CookieSuffix:            khulnasoft.F("abcd"),
				CustomPageHTML:          khulnasoft.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
				DefaultTemplateLanguage: khulnasoft.F(waiting_rooms.QueryDefaultTemplateLanguageEnUs),
				Description:             khulnasoft.F("Production - DO NOT MODIFY"),
				DisableSessionRenewal:   khulnasoft.F(false),
				EnabledOriginCommands:   khulnasoft.F([]waiting_rooms.QueryEnabledOriginCommand{waiting_rooms.QueryEnabledOriginCommandRevoke}),
				JsonResponseEnabled:     khulnasoft.F(false),
				Path:                    khulnasoft.F("/shop/checkout"),
				QueueAll:                khulnasoft.F(true),
				QueueingMethod:          khulnasoft.F(waiting_rooms.QueryQueueingMethodFifo),
				QueueingStatusCode:      khulnasoft.F(waiting_rooms.QueryQueueingStatusCode200),
				SessionDuration:         khulnasoft.F(int64(1)),
				Suspended:               khulnasoft.F(true),
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

func TestWaitingRoomGet(t *testing.T) {
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
	_, err := client.WaitingRooms.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.WaitingRoomGetParams{
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
