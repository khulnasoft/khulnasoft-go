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

func TestEmailSecurityTopTldMaliciousGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Tlds.Malicious.Get(
		context.TODO(),
		radar.EmailSecurityTopTldMaliciousGetParamsMaliciousMalicious,
		radar.EmailSecurityTopTldMaliciousGetParams{
			ARC:         khulnasoft.F([]radar.EmailSecurityTopTldMaliciousGetParamsARC{radar.EmailSecurityTopTldMaliciousGetParamsARCPass, radar.EmailSecurityTopTldMaliciousGetParamsARCNone, radar.EmailSecurityTopTldMaliciousGetParamsARCFail}),
			DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
			DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DKIM:        khulnasoft.F([]radar.EmailSecurityTopTldMaliciousGetParamsDKIM{radar.EmailSecurityTopTldMaliciousGetParamsDKIMPass, radar.EmailSecurityTopTldMaliciousGetParamsDKIMNone, radar.EmailSecurityTopTldMaliciousGetParamsDKIMFail}),
			DMARC:       khulnasoft.F([]radar.EmailSecurityTopTldMaliciousGetParamsDMARC{radar.EmailSecurityTopTldMaliciousGetParamsDMARCPass, radar.EmailSecurityTopTldMaliciousGetParamsDMARCNone, radar.EmailSecurityTopTldMaliciousGetParamsDMARCFail}),
			Format:      khulnasoft.F(radar.EmailSecurityTopTldMaliciousGetParamsFormatJson),
			Limit:       khulnasoft.F(int64(5)),
			Name:        khulnasoft.F([]string{"string", "string", "string"}),
			SPF:         khulnasoft.F([]radar.EmailSecurityTopTldMaliciousGetParamsSPF{radar.EmailSecurityTopTldMaliciousGetParamsSPFPass, radar.EmailSecurityTopTldMaliciousGetParamsSPFNone, radar.EmailSecurityTopTldMaliciousGetParamsSPFFail}),
			TldCategory: khulnasoft.F(radar.EmailSecurityTopTldMaliciousGetParamsTldCategoryClassic),
			TLSVersion:  khulnasoft.F([]radar.EmailSecurityTopTldMaliciousGetParamsTLSVersion{radar.EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_0, radar.EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_1, radar.EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_2}),
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
