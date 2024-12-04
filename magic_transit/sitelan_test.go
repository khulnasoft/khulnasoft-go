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

func TestSiteLANNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.LANs.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLANNewParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Physport:  khulnasoft.F(int64(1)),
			VlanTag:   khulnasoft.F(int64(0)),
			HaLink:    khulnasoft.F(true),
			Name:      khulnasoft.F("name"),
			Nat: khulnasoft.F(magic_transit.NatParam{
				StaticPrefix: khulnasoft.F("192.0.2.0/24"),
			}),
			RoutedSubnets: khulnasoft.F([]magic_transit.RoutedSubnetParam{{
				NextHop: khulnasoft.F("192.0.2.1"),
				Prefix:  khulnasoft.F("192.0.2.0/24"),
				Nat: khulnasoft.F(magic_transit.NatParam{
					StaticPrefix: khulnasoft.F("192.0.2.0/24"),
				}),
			}, {
				NextHop: khulnasoft.F("192.0.2.1"),
				Prefix:  khulnasoft.F("192.0.2.0/24"),
				Nat: khulnasoft.F(magic_transit.NatParam{
					StaticPrefix: khulnasoft.F("192.0.2.0/24"),
				}),
			}, {
				NextHop: khulnasoft.F("192.0.2.1"),
				Prefix:  khulnasoft.F("192.0.2.0/24"),
				Nat: khulnasoft.F(magic_transit.NatParam{
					StaticPrefix: khulnasoft.F("192.0.2.0/24"),
				}),
			}}),
			StaticAddressing: khulnasoft.F(magic_transit.LANStaticAddressingParam{
				Address: khulnasoft.F("192.0.2.0/24"),
				DHCPRelay: khulnasoft.F(magic_transit.DHCPRelayParam{
					ServerAddresses: khulnasoft.F([]string{"192.0.2.1", "192.0.2.1", "192.0.2.1"}),
				}),
				DHCPServer: khulnasoft.F(magic_transit.DHCPServerParam{
					DHCPPoolEnd:   khulnasoft.F("192.0.2.1"),
					DHCPPoolStart: khulnasoft.F("192.0.2.1"),
					DNSServer:     khulnasoft.F("192.0.2.1"),
					Reservations: khulnasoft.F(map[string]string{
						"00:11:22:33:44:55": "192.0.2.100",
						"AA:BB:CC:DD:EE:FF": "192.168.1.101",
					}),
				}),
				SecondaryAddress: khulnasoft.F("192.0.2.0/24"),
				VirtualAddress:   khulnasoft.F("192.0.2.0/24"),
			}),
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

func TestSiteLANUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.LANs.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLANUpdateParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:      khulnasoft.F("name"),
			Nat: khulnasoft.F(magic_transit.NatParam{
				StaticPrefix: khulnasoft.F("192.0.2.0/24"),
			}),
			Physport: khulnasoft.F(int64(1)),
			RoutedSubnets: khulnasoft.F([]magic_transit.RoutedSubnetParam{{
				NextHop: khulnasoft.F("192.0.2.1"),
				Prefix:  khulnasoft.F("192.0.2.0/24"),
				Nat: khulnasoft.F(magic_transit.NatParam{
					StaticPrefix: khulnasoft.F("192.0.2.0/24"),
				}),
			}, {
				NextHop: khulnasoft.F("192.0.2.1"),
				Prefix:  khulnasoft.F("192.0.2.0/24"),
				Nat: khulnasoft.F(magic_transit.NatParam{
					StaticPrefix: khulnasoft.F("192.0.2.0/24"),
				}),
			}, {
				NextHop: khulnasoft.F("192.0.2.1"),
				Prefix:  khulnasoft.F("192.0.2.0/24"),
				Nat: khulnasoft.F(magic_transit.NatParam{
					StaticPrefix: khulnasoft.F("192.0.2.0/24"),
				}),
			}}),
			StaticAddressing: khulnasoft.F(magic_transit.LANStaticAddressingParam{
				Address: khulnasoft.F("192.0.2.0/24"),
				DHCPRelay: khulnasoft.F(magic_transit.DHCPRelayParam{
					ServerAddresses: khulnasoft.F([]string{"192.0.2.1", "192.0.2.1", "192.0.2.1"}),
				}),
				DHCPServer: khulnasoft.F(magic_transit.DHCPServerParam{
					DHCPPoolEnd:   khulnasoft.F("192.0.2.1"),
					DHCPPoolStart: khulnasoft.F("192.0.2.1"),
					DNSServer:     khulnasoft.F("192.0.2.1"),
					Reservations: khulnasoft.F(map[string]string{
						"00:11:22:33:44:55": "192.0.2.100",
						"AA:BB:CC:DD:EE:FF": "192.168.1.101",
					}),
				}),
				SecondaryAddress: khulnasoft.F("192.0.2.0/24"),
				VirtualAddress:   khulnasoft.F("192.0.2.0/24"),
			}),
			VlanTag: khulnasoft.F(int64(0)),
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

func TestSiteLANList(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.LANs.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLANListParams{
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

func TestSiteLANDelete(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.LANs.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLANDeleteParams{
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

func TestSiteLANEditWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.LANs.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLANEditParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:      khulnasoft.F("name"),
			Nat: khulnasoft.F(magic_transit.NatParam{
				StaticPrefix: khulnasoft.F("192.0.2.0/24"),
			}),
			Physport: khulnasoft.F(int64(1)),
			RoutedSubnets: khulnasoft.F([]magic_transit.RoutedSubnetParam{{
				NextHop: khulnasoft.F("192.0.2.1"),
				Prefix:  khulnasoft.F("192.0.2.0/24"),
				Nat: khulnasoft.F(magic_transit.NatParam{
					StaticPrefix: khulnasoft.F("192.0.2.0/24"),
				}),
			}, {
				NextHop: khulnasoft.F("192.0.2.1"),
				Prefix:  khulnasoft.F("192.0.2.0/24"),
				Nat: khulnasoft.F(magic_transit.NatParam{
					StaticPrefix: khulnasoft.F("192.0.2.0/24"),
				}),
			}, {
				NextHop: khulnasoft.F("192.0.2.1"),
				Prefix:  khulnasoft.F("192.0.2.0/24"),
				Nat: khulnasoft.F(magic_transit.NatParam{
					StaticPrefix: khulnasoft.F("192.0.2.0/24"),
				}),
			}}),
			StaticAddressing: khulnasoft.F(magic_transit.LANStaticAddressingParam{
				Address: khulnasoft.F("192.0.2.0/24"),
				DHCPRelay: khulnasoft.F(magic_transit.DHCPRelayParam{
					ServerAddresses: khulnasoft.F([]string{"192.0.2.1", "192.0.2.1", "192.0.2.1"}),
				}),
				DHCPServer: khulnasoft.F(magic_transit.DHCPServerParam{
					DHCPPoolEnd:   khulnasoft.F("192.0.2.1"),
					DHCPPoolStart: khulnasoft.F("192.0.2.1"),
					DNSServer:     khulnasoft.F("192.0.2.1"),
					Reservations: khulnasoft.F(map[string]string{
						"00:11:22:33:44:55": "192.0.2.100",
						"AA:BB:CC:DD:EE:FF": "192.168.1.101",
					}),
				}),
				SecondaryAddress: khulnasoft.F("192.0.2.0/24"),
				VirtualAddress:   khulnasoft.F("192.0.2.0/24"),
			}),
			VlanTag: khulnasoft.F(int64(0)),
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

func TestSiteLANGet(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.LANs.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLANGetParams{
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
