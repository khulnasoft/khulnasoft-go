// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_scanner_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/url_scanner"
)

func TestURLScannerScanWithOptionalParams(t *testing.T) {
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
	_, err := client.URLScanner.Scan(
		context.TODO(),
		"accountId",
		url_scanner.URLScannerScanParams{
			AccountScans: khulnasoft.F(true),
			ASN:          khulnasoft.F("13335"),
			DateEnd:      khulnasoft.F(time.Now()),
			DateStart:    khulnasoft.F(time.Now()),
			Hash:         khulnasoft.F("hash"),
			Hostname:     khulnasoft.F("example.com"),
			IP:           khulnasoft.F("1.1.1.1"),
			IsMalicious:  khulnasoft.F(true),
			Limit:        khulnasoft.F(int64(100)),
			NextCursor:   khulnasoft.F("next_cursor"),
			PageASN:      khulnasoft.F("page_asn"),
			PageHostname: khulnasoft.F("page_hostname"),
			PageIP:       khulnasoft.F("page_ip"),
			PagePath:     khulnasoft.F("page_path"),
			PageURL:      khulnasoft.F("page_url"),
			Path:         khulnasoft.F("/samples/subresource-integrity/"),
			ScanID:       khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			URL:          khulnasoft.F("https://example.com/?hello"),
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
