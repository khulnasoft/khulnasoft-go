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

func TestEmailSecurityTopTldSpamGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Tlds.Spam.Get(
		context.TODO(),
		radar.EmailSecurityTopTldSpamGetParamsSpamSpam,
		radar.EmailSecurityTopTldSpamGetParams{
			ARC:         khulnasoft.F([]radar.EmailSecurityTopTldSpamGetParamsARC{radar.EmailSecurityTopTldSpamGetParamsARCPass, radar.EmailSecurityTopTldSpamGetParamsARCNone, radar.EmailSecurityTopTldSpamGetParamsARCFail}),
			DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DKIM:        khulnasoft.F([]radar.EmailSecurityTopTldSpamGetParamsDKIM{radar.EmailSecurityTopTldSpamGetParamsDKIMPass, radar.EmailSecurityTopTldSpamGetParamsDKIMNone, radar.EmailSecurityTopTldSpamGetParamsDKIMFail}),
			DMARC:       khulnasoft.F([]radar.EmailSecurityTopTldSpamGetParamsDMARC{radar.EmailSecurityTopTldSpamGetParamsDMARCPass, radar.EmailSecurityTopTldSpamGetParamsDMARCNone, radar.EmailSecurityTopTldSpamGetParamsDMARCFail}),
			Format:      khulnasoft.F(radar.EmailSecurityTopTldSpamGetParamsFormatJson),
			Limit:       khulnasoft.F(int64(5)),
			Name:        khulnasoft.F([]string{"string", "string", "string"}),
			SPF:         khulnasoft.F([]radar.EmailSecurityTopTldSpamGetParamsSPF{radar.EmailSecurityTopTldSpamGetParamsSPFPass, radar.EmailSecurityTopTldSpamGetParamsSPFNone, radar.EmailSecurityTopTldSpamGetParamsSPFFail}),
			TldCategory: khulnasoft.F(radar.EmailSecurityTopTldSpamGetParamsTldCategoryClassic),
			TLSVersion:  khulnasoft.F([]radar.EmailSecurityTopTldSpamGetParamsTLSVersion{radar.EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_0, radar.EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_1, radar.EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_2}),
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
