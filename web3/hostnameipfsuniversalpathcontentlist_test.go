// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/web3"
)

func TestHostnameIPFSUniversalPathContentListUpdate(t *testing.T) {
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
	_, err := client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		web3.HostnameIPFSUniversalPathContentListUpdateParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Action: khulnasoft.F(web3.HostnameIPFSUniversalPathContentListUpdateParamsActionBlock),
			Entries: khulnasoft.F([]web3.HostnameIPFSUniversalPathContentListUpdateParamsEntry{{
				Content:     khulnasoft.F("QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB"),
				Description: khulnasoft.F("this is my content list entry"),
				Type:        khulnasoft.F(web3.HostnameIPFSUniversalPathContentListUpdateParamsEntriesTypeCid),
			}, {
				Content:     khulnasoft.F("QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB"),
				Description: khulnasoft.F("this is my content list entry"),
				Type:        khulnasoft.F(web3.HostnameIPFSUniversalPathContentListUpdateParamsEntriesTypeCid),
			}, {
				Content:     khulnasoft.F("QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB"),
				Description: khulnasoft.F("this is my content list entry"),
				Type:        khulnasoft.F(web3.HostnameIPFSUniversalPathContentListUpdateParamsEntriesTypeCid),
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

func TestHostnameIPFSUniversalPathContentListGet(t *testing.T) {
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
	_, err := client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		web3.HostnameIPFSUniversalPathContentListGetParams{
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
