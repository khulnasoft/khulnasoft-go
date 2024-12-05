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

func TestEmailSecurityTopTldGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Tlds.Get(context.TODO(), radar.EmailSecurityTopTldGetParams{
		ARC:         khulnasoft.F([]radar.EmailSecurityTopTldGetParamsARC{radar.EmailSecurityTopTldGetParamsARCPass, radar.EmailSecurityTopTldGetParamsARCNone, radar.EmailSecurityTopTldGetParamsARCFail}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        khulnasoft.F([]radar.EmailSecurityTopTldGetParamsDKIM{radar.EmailSecurityTopTldGetParamsDKIMPass, radar.EmailSecurityTopTldGetParamsDKIMNone, radar.EmailSecurityTopTldGetParamsDKIMFail}),
		DMARC:       khulnasoft.F([]radar.EmailSecurityTopTldGetParamsDMARC{radar.EmailSecurityTopTldGetParamsDMARCPass, radar.EmailSecurityTopTldGetParamsDMARCNone, radar.EmailSecurityTopTldGetParamsDMARCFail}),
		Format:      khulnasoft.F(radar.EmailSecurityTopTldGetParamsFormatJson),
		Limit:       khulnasoft.F(int64(5)),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
		SPF:         khulnasoft.F([]radar.EmailSecurityTopTldGetParamsSPF{radar.EmailSecurityTopTldGetParamsSPFPass, radar.EmailSecurityTopTldGetParamsSPFNone, radar.EmailSecurityTopTldGetParamsSPFFail}),
		TldCategory: khulnasoft.F(radar.EmailSecurityTopTldGetParamsTldCategoryClassic),
		TLSVersion:  khulnasoft.F([]radar.EmailSecurityTopTldGetParamsTLSVersion{radar.EmailSecurityTopTldGetParamsTLSVersionTlSv1_0, radar.EmailSecurityTopTldGetParamsTLSVersionTlSv1_1, radar.EmailSecurityTopTldGetParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
