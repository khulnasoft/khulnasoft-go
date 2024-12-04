// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/iam"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestResourceGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.New(context.TODO(), iam.ResourceGroupNewParams{
		AccountID: khulnasoft.F("eb78d65290b24279ba6f44721b3ea3c4"),
		Scope: khulnasoft.F(iam.ResourceGroupNewParamsScope{
			Key: khulnasoft.F("com.khulnasoft.api.account.eb78d65290b24279ba6f44721b3ea3c4"),
			Objects: khulnasoft.F([]iam.ResourceGroupNewParamsScopeObject{{
				Key: khulnasoft.F("com.khulnasoft.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
			}, {
				Key: khulnasoft.F("com.khulnasoft.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
			}, {
				Key: khulnasoft.F("com.khulnasoft.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
			}}),
		}),
		Meta: khulnasoft.F[any](map[string]interface{}{
			"editable": "false",
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

func TestResourceGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.Update(
		context.TODO(),
		"6d7f2f5f5b1d4a0e9081fdc98d432fd1",
		iam.ResourceGroupUpdateParams{
			AccountID: khulnasoft.F("eb78d65290b24279ba6f44721b3ea3c4"),
			Scope: khulnasoft.F(iam.ResourceGroupUpdateParamsScope{
				Key: khulnasoft.F("com.khulnasoft.api.account.eb78d65290b24279ba6f44721b3ea3c4"),
				Objects: khulnasoft.F([]iam.ResourceGroupUpdateParamsScopeObject{{
					Key: khulnasoft.F("com.khulnasoft.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
				}, {
					Key: khulnasoft.F("com.khulnasoft.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
				}, {
					Key: khulnasoft.F("com.khulnasoft.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
				}}),
			}),
			Meta: khulnasoft.F[any](map[string]interface{}{
				"editable": "false",
			}),
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

func TestResourceGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.List(context.TODO(), iam.ResourceGroupListParams{
		AccountID: khulnasoft.F("eb78d65290b24279ba6f44721b3ea3c4"),
		ID:        khulnasoft.F("6d7f2f5f5b1d4a0e9081fdc98d432fd1"),
		Name:      khulnasoft.F("NameOfTheResourceGroup"),
		Page:      khulnasoft.F(1.000000),
		PerPage:   khulnasoft.F(5.000000),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResourceGroupDelete(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.Delete(
		context.TODO(),
		"6d7f2f5f5b1d4a0e9081fdc98d432fd1",
		iam.ResourceGroupDeleteParams{
			AccountID: khulnasoft.F("eb78d65290b24279ba6f44721b3ea3c4"),
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

func TestResourceGroupGet(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.Get(
		context.TODO(),
		"6d7f2f5f5b1d4a0e9081fdc98d432fd1",
		iam.ResourceGroupGetParams{
			AccountID: khulnasoft.F("eb78d65290b24279ba6f44721b3ea3c4"),
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
