// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/magic_transit"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestSiteACLNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.ACLs.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteACLNewParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			LAN1: khulnasoft.F(magic_transit.ACLConfigurationParam{
				LANID:   khulnasoft.F("lan_id"),
				LANName: khulnasoft.F("lan_name"),
				Ports:   khulnasoft.F([]int64{int64(1), int64(1), int64(1)}),
				Subnets: khulnasoft.F([]magic_transit.SubnetParam{"192.0.2.1", "192.0.2.1", "192.0.2.1"}),
			}),
			LAN2: khulnasoft.F(magic_transit.ACLConfigurationParam{
				LANID:   khulnasoft.F("lan_id"),
				LANName: khulnasoft.F("lan_name"),
				Ports:   khulnasoft.F([]int64{int64(1), int64(1), int64(1)}),
				Subnets: khulnasoft.F([]magic_transit.SubnetParam{"192.0.2.1", "192.0.2.1", "192.0.2.1"}),
			}),
			Name:           khulnasoft.F("PIN Pad - Cash Register"),
			Description:    khulnasoft.F("Allows local traffic between PIN pads and cash register."),
			ForwardLocally: khulnasoft.F(true),
			Protocols:      khulnasoft.F([]magic_transit.AllowedProtocol{magic_transit.AllowedProtocolTCP, magic_transit.AllowedProtocolUdp, magic_transit.AllowedProtocolIcmp}),
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

func TestSiteACLUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.ACLs.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteACLUpdateParams{
			AccountID:      khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Description:    khulnasoft.F("Allows local traffic between PIN pads and cash register."),
			ForwardLocally: khulnasoft.F(true),
			LAN1: khulnasoft.F(magic_transit.ACLConfigurationParam{
				LANID:   khulnasoft.F("lan_id"),
				LANName: khulnasoft.F("lan_name"),
				Ports:   khulnasoft.F([]int64{int64(1), int64(1), int64(1)}),
				Subnets: khulnasoft.F([]magic_transit.SubnetParam{"192.0.2.1", "192.0.2.1", "192.0.2.1"}),
			}),
			LAN2: khulnasoft.F(magic_transit.ACLConfigurationParam{
				LANID:   khulnasoft.F("lan_id"),
				LANName: khulnasoft.F("lan_name"),
				Ports:   khulnasoft.F([]int64{int64(1), int64(1), int64(1)}),
				Subnets: khulnasoft.F([]magic_transit.SubnetParam{"192.0.2.1", "192.0.2.1", "192.0.2.1"}),
			}),
			Name:      khulnasoft.F("PIN Pad - Cash Register"),
			Protocols: khulnasoft.F([]magic_transit.AllowedProtocol{magic_transit.AllowedProtocolTCP, magic_transit.AllowedProtocolUdp, magic_transit.AllowedProtocolIcmp}),
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

func TestSiteACLList(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.ACLs.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteACLListParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestSiteACLDelete(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.ACLs.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteACLDeleteParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestSiteACLEditWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.ACLs.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteACLEditParams{
			AccountID:      khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Description:    khulnasoft.F("Allows local traffic between PIN pads and cash register."),
			ForwardLocally: khulnasoft.F(true),
			LAN1: khulnasoft.F(magic_transit.ACLConfigurationParam{
				LANID:   khulnasoft.F("lan_id"),
				LANName: khulnasoft.F("lan_name"),
				Ports:   khulnasoft.F([]int64{int64(1), int64(1), int64(1)}),
				Subnets: khulnasoft.F([]magic_transit.SubnetParam{"192.0.2.1", "192.0.2.1", "192.0.2.1"}),
			}),
			LAN2: khulnasoft.F(magic_transit.ACLConfigurationParam{
				LANID:   khulnasoft.F("lan_id"),
				LANName: khulnasoft.F("lan_name"),
				Ports:   khulnasoft.F([]int64{int64(1), int64(1), int64(1)}),
				Subnets: khulnasoft.F([]magic_transit.SubnetParam{"192.0.2.1", "192.0.2.1", "192.0.2.1"}),
			}),
			Name:      khulnasoft.F("PIN Pad - Cash Register"),
			Protocols: khulnasoft.F([]magic_transit.AllowedProtocol{magic_transit.AllowedProtocolTCP, magic_transit.AllowedProtocolUdp, magic_transit.AllowedProtocolIcmp}),
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

func TestSiteACLGet(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.ACLs.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteACLGetParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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