// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagerules_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/pagerules"
)

func TestPageruleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pagerules.New(context.TODO(), pagerules.PageruleNewParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Actions: khulnasoft.F([]pagerules.RouteParam{{
			Name: khulnasoft.F(pagerules.RouteNameForwardURL),
			Value: khulnasoft.F(pagerules.RouteValueParam{
				Type: khulnasoft.F(pagerules.RouteValueTypeTemporary),
				URL:  khulnasoft.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
			}),
		}, {
			Name: khulnasoft.F(pagerules.RouteNameForwardURL),
			Value: khulnasoft.F(pagerules.RouteValueParam{
				Type: khulnasoft.F(pagerules.RouteValueTypeTemporary),
				URL:  khulnasoft.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
			}),
		}, {
			Name: khulnasoft.F(pagerules.RouteNameForwardURL),
			Value: khulnasoft.F(pagerules.RouteValueParam{
				Type: khulnasoft.F(pagerules.RouteValueTypeTemporary),
				URL:  khulnasoft.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
			}),
		}}),
		Targets: khulnasoft.F([]pagerules.TargetParam{{
			Constraint: khulnasoft.F(pagerules.TargetConstraintParam{
				Operator: khulnasoft.F(pagerules.TargetConstraintOperatorMatches),
				Value:    khulnasoft.F("*example.com/images/*"),
			}),
			Target: khulnasoft.F(pagerules.TargetTargetURL),
		}}),
		Priority: khulnasoft.F(int64(0)),
		Status:   khulnasoft.F(pagerules.PageruleNewParamsStatusActive),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageruleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pagerules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		pagerules.PageruleUpdateParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Actions: khulnasoft.F([]pagerules.RouteParam{{
				Name: khulnasoft.F(pagerules.RouteNameForwardURL),
				Value: khulnasoft.F(pagerules.RouteValueParam{
					Type: khulnasoft.F(pagerules.RouteValueTypeTemporary),
					URL:  khulnasoft.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}, {
				Name: khulnasoft.F(pagerules.RouteNameForwardURL),
				Value: khulnasoft.F(pagerules.RouteValueParam{
					Type: khulnasoft.F(pagerules.RouteValueTypeTemporary),
					URL:  khulnasoft.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}, {
				Name: khulnasoft.F(pagerules.RouteNameForwardURL),
				Value: khulnasoft.F(pagerules.RouteValueParam{
					Type: khulnasoft.F(pagerules.RouteValueTypeTemporary),
					URL:  khulnasoft.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}}),
			Targets: khulnasoft.F([]pagerules.TargetParam{{
				Constraint: khulnasoft.F(pagerules.TargetConstraintParam{
					Operator: khulnasoft.F(pagerules.TargetConstraintOperatorMatches),
					Value:    khulnasoft.F("*example.com/images/*"),
				}),
				Target: khulnasoft.F(pagerules.TargetTargetURL),
			}}),
			Priority: khulnasoft.F(int64(0)),
			Status:   khulnasoft.F(pagerules.PageruleUpdateParamsStatusActive),
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

func TestPageruleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Pagerules.List(context.TODO(), pagerules.PageruleListParams{
		ZoneID:    khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: khulnasoft.F(pagerules.PageruleListParamsDirectionAsc),
		Match:     khulnasoft.F(pagerules.PageruleListParamsMatchAny),
		Order:     khulnasoft.F(pagerules.PageruleListParamsOrderStatus),
		Status:    khulnasoft.F(pagerules.PageruleListParamsStatusActive),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageruleDelete(t *testing.T) {
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
	_, err := client.Pagerules.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		pagerules.PageruleDeleteParams{
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

func TestPageruleEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Pagerules.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		pagerules.PageruleEditParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Actions: khulnasoft.F([]pagerules.RouteParam{{
				Name: khulnasoft.F(pagerules.RouteNameForwardURL),
				Value: khulnasoft.F(pagerules.RouteValueParam{
					Type: khulnasoft.F(pagerules.RouteValueTypeTemporary),
					URL:  khulnasoft.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}, {
				Name: khulnasoft.F(pagerules.RouteNameForwardURL),
				Value: khulnasoft.F(pagerules.RouteValueParam{
					Type: khulnasoft.F(pagerules.RouteValueTypeTemporary),
					URL:  khulnasoft.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}, {
				Name: khulnasoft.F(pagerules.RouteNameForwardURL),
				Value: khulnasoft.F(pagerules.RouteValueParam{
					Type: khulnasoft.F(pagerules.RouteValueTypeTemporary),
					URL:  khulnasoft.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}}),
			Priority: khulnasoft.F(int64(0)),
			Status:   khulnasoft.F(pagerules.PageruleEditParamsStatusActive),
			Targets: khulnasoft.F([]pagerules.TargetParam{{
				Constraint: khulnasoft.F(pagerules.TargetConstraintParam{
					Operator: khulnasoft.F(pagerules.TargetConstraintOperatorMatches),
					Value:    khulnasoft.F("*example.com/images/*"),
				}),
				Target: khulnasoft.F(pagerules.TargetTargetURL),
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

func TestPageruleGet(t *testing.T) {
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
	_, err := client.Pagerules.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		pagerules.PageruleGetParams{
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
