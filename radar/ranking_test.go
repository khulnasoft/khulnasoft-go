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

func TestRankingTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Ranking.TimeseriesGroups(context.TODO(), radar.RankingTimeseriesGroupsParams{
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Domains:     khulnasoft.F([]string{"string", "string", "string"}),
		Format:      khulnasoft.F(radar.RankingTimeseriesGroupsParamsFormatJson),
		Limit:       khulnasoft.F(int64(5)),
		Location:    khulnasoft.F([]string{"string", "string", "string"}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		RankingType: khulnasoft.F(radar.RankingTimeseriesGroupsParamsRankingTypePopular),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRankingTopWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Ranking.Top(context.TODO(), radar.RankingTopParams{
		Date:        khulnasoft.F([]string{"string", "string", "string"}),
		Format:      khulnasoft.F(radar.RankingTopParamsFormatJson),
		Limit:       khulnasoft.F(int64(5)),
		Location:    khulnasoft.F([]string{"string", "string", "string"}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		RankingType: khulnasoft.F(radar.RankingTopParamsRankingTypePopular),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
