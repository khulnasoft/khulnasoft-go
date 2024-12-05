// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package request_tracers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/request_tracers"
)

func TestTraceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.RequestTracers.Traces.New(context.TODO(), request_tracers.TraceNewParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Method:    khulnasoft.F("PUT"),
		URL:       khulnasoft.F("https://some.zone/some_path"),
		Body: khulnasoft.F(request_tracers.TraceNewParamsBody{
			Base64:    khulnasoft.F("c29tZV9yZXF1ZXN0X2JvZHk="),
			Json:      khulnasoft.F[any](map[string]interface{}{}),
			PlainText: khulnasoft.F("plain_text"),
		}),
		Context: khulnasoft.F(request_tracers.TraceNewParamsContext{
			BotScore: khulnasoft.F(int64(0)),
			Geoloc: khulnasoft.F(request_tracers.TraceNewParamsContextGeoloc{
				City:                khulnasoft.F("London"),
				Continent:           khulnasoft.F("continent"),
				IsEuCountry:         khulnasoft.F(true),
				ISOCode:             khulnasoft.F("iso_code"),
				Latitude:            khulnasoft.F(0.000000),
				Longitude:           khulnasoft.F(0.000000),
				PostalCode:          khulnasoft.F("postal_code"),
				RegionCode:          khulnasoft.F("region_code"),
				Subdivision2ISOCode: khulnasoft.F("subdivision_2_iso_code"),
				Timezone:            khulnasoft.F("timezone"),
			}),
			SkipChallenge: khulnasoft.F(true),
			ThreatScore:   khulnasoft.F(int64(0)),
		}),
		Cookies: khulnasoft.F(map[string]string{
			"cookie_name_1": "cookie_value_1",
			"cookie_name_2": "cookie_value_2",
		}),
		Headers: khulnasoft.F(map[string]string{
			"header_name_1": "header_value_1",
			"header_name_2": "header_value_2",
		}),
		Protocol:     khulnasoft.F("HTTP/1.1"),
		SkipResponse: khulnasoft.F(true),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
