// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/zero_trust"
)

func TestSeatEdit(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.ZeroTrust.Seats.Edit(context.TODO(), zero_trust.SeatEditParams{
		AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		Body: []zero_trust.SeatEditParamsBody{{
			AccessSeat:  khulnasoft.F(false),
			GatewaySeat: khulnasoft.F(false),
			SeatUID:     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
		}, {
			AccessSeat:  khulnasoft.F(false),
			GatewaySeat: khulnasoft.F(false),
			SeatUID:     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
		}, {
			AccessSeat:  khulnasoft.F(false),
			GatewaySeat: khulnasoft.F(false),
			SeatUID:     khulnasoft.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
		}},
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
