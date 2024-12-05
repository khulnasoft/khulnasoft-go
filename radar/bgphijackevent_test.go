// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/radar"
)

func TestBGPHijackEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.BGP.Hijacks.Events.List(context.TODO(), radar.BGPHijackEventListParams{
		DateEnd:         khulnasoft.F(time.Now()),
		DateRange:       khulnasoft.F("7d"),
		DateStart:       khulnasoft.F(time.Now()),
		EventID:         khulnasoft.F(int64(0)),
		Format:          khulnasoft.F(radar.BGPHijackEventListParamsFormatJson),
		HijackerASN:     khulnasoft.F(int64(0)),
		InvolvedASN:     khulnasoft.F(int64(0)),
		InvolvedCountry: khulnasoft.F("involvedCountry"),
		MaxConfidence:   khulnasoft.F(int64(0)),
		MinConfidence:   khulnasoft.F(int64(0)),
		Page:            khulnasoft.F(int64(0)),
		PerPage:         khulnasoft.F(int64(0)),
		Prefix:          khulnasoft.F("1.1.1.0/24"),
		SortBy:          khulnasoft.F(radar.BGPHijackEventListParamsSortByID),
		SortOrder:       khulnasoft.F(radar.BGPHijackEventListParamsSortOrderAsc),
		VictimASN:       khulnasoft.F(int64(0)),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
