// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/shared"
	"github.com/khulnasoft/khulnasoft-go/user"
)

func TestAuditLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.User.AuditLogs.List(context.TODO(), user.AuditLogListParams{
		ID: khulnasoft.F("f174be97-19b1-40d6-954d-70cd5fbd52db"),
		Action: khulnasoft.F(user.AuditLogListParamsAction{
			Type: khulnasoft.F("add"),
		}),
		Actor: khulnasoft.F(user.AuditLogListParamsActor{
			Email: khulnasoft.F("alice@example.com"),
			IP:    khulnasoft.F("17.168.228.63"),
		}),
		Before:       khulnasoft.F[user.AuditLogListParamsBeforeUnion](shared.UnionTime(time.Now())),
		Direction:    khulnasoft.F(user.AuditLogListParamsDirectionDesc),
		Export:       khulnasoft.F(true),
		HideUserLogs: khulnasoft.F(true),
		Page:         khulnasoft.F(50.000000),
		PerPage:      khulnasoft.F(25.000000),
		Since:        khulnasoft.F[user.AuditLogListParamsSinceUnion](shared.UnionTime(time.Now())),
		Zone: khulnasoft.F(user.AuditLogListParamsZone{
			Name: khulnasoft.F("example.com"),
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
