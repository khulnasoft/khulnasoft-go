// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/logs"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/shared"
)

func TestReceivedGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Logs.Received.Get(context.TODO(), logs.ReceivedGetParams{
		ZoneID:     khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		End:        khulnasoft.F[logs.ReceivedGetParamsEndUnion](shared.UnionString("2018-05-20T10:01:00Z")),
		Count:      khulnasoft.F(int64(1)),
		Fields:     khulnasoft.F("ClientIP,RayID,EdgeStartTimestamp"),
		Sample:     khulnasoft.F(0.100000),
		Start:      khulnasoft.F[logs.ReceivedGetParamsStartUnion](shared.UnionString("2018-05-20T10:00:00Z")),
		Timestamps: khulnasoft.F(logs.ReceivedGetParamsTimestampsUnix),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
