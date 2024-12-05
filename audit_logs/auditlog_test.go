// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audit_logs_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/audit_logs"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/shared"
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
	_, err := client.AuditLogs.List(context.TODO(), audit_logs.AuditLogListParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ID:        khulnasoft.F("f174be97-19b1-40d6-954d-70cd5fbd52db"),
		Action: khulnasoft.F(audit_logs.AuditLogListParamsAction{
			Type: khulnasoft.F("add"),
		}),
		Actor: khulnasoft.F(audit_logs.AuditLogListParamsActor{
			Email: khulnasoft.F("alice@example.com"),
			IP:    khulnasoft.F("17.168.228.63"),
		}),
		Before:       khulnasoft.F[audit_logs.AuditLogListParamsBeforeUnion](shared.UnionTime(time.Now())),
		Direction:    khulnasoft.F(audit_logs.AuditLogListParamsDirectionDesc),
		Export:       khulnasoft.F(true),
		HideUserLogs: khulnasoft.F(true),
		Page:         khulnasoft.F(50.000000),
		PerPage:      khulnasoft.F(25.000000),
		Since:        khulnasoft.F[audit_logs.AuditLogListParamsSinceUnion](shared.UnionTime(time.Now())),
		Zone: khulnasoft.F(audit_logs.AuditLogListParamsZone{
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
