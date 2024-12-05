// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/load_balancers"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestPoolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Pools.New(context.TODO(), load_balancers.PoolNewParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Name:      khulnasoft.F("primary-dc-1"),
		Origins: khulnasoft.F([]load_balancers.OriginParam{{
			Address: khulnasoft.F("0.0.0.0"),
			Enabled: khulnasoft.F(true),
			Header: khulnasoft.F(load_balancers.HeaderParam{
				Host: khulnasoft.F([]load_balancers.HostParam{"example.com", "example.com", "example.com"}),
			}),
			Name:             khulnasoft.F("app-server-1"),
			VirtualNetworkID: khulnasoft.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           khulnasoft.F(0.600000),
		}, {
			Address: khulnasoft.F("0.0.0.0"),
			Enabled: khulnasoft.F(true),
			Header: khulnasoft.F(load_balancers.HeaderParam{
				Host: khulnasoft.F([]load_balancers.HostParam{"example.com", "example.com", "example.com"}),
			}),
			Name:             khulnasoft.F("app-server-1"),
			VirtualNetworkID: khulnasoft.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           khulnasoft.F(0.600000),
		}, {
			Address: khulnasoft.F("0.0.0.0"),
			Enabled: khulnasoft.F(true),
			Header: khulnasoft.F(load_balancers.HeaderParam{
				Host: khulnasoft.F([]load_balancers.HostParam{"example.com", "example.com", "example.com"}),
			}),
			Name:             khulnasoft.F("app-server-1"),
			VirtualNetworkID: khulnasoft.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           khulnasoft.F(0.600000),
		}}),
		Description: khulnasoft.F("Primary data center - Provider XYZ"),
		Enabled:     khulnasoft.F(false),
		Latitude:    khulnasoft.F(0.000000),
		LoadShedding: khulnasoft.F(load_balancers.LoadSheddingParam{
			DefaultPercent: khulnasoft.F(0.000000),
			DefaultPolicy:  khulnasoft.F(load_balancers.LoadSheddingDefaultPolicyRandom),
			SessionPercent: khulnasoft.F(0.000000),
			SessionPolicy:  khulnasoft.F(load_balancers.LoadSheddingSessionPolicyHash),
		}),
		Longitude:         khulnasoft.F(0.000000),
		MinimumOrigins:    khulnasoft.F(int64(0)),
		Monitor:           khulnasoft.F("monitor"),
		NotificationEmail: khulnasoft.F("someone@example.com,sometwo@example.com"),
		NotificationFilter: khulnasoft.F(load_balancers.NotificationFilterParam{
			Origin: khulnasoft.F(load_balancers.FilterOptionsParam{
				Disable: khulnasoft.F(true),
				Healthy: khulnasoft.F(true),
			}),
			Pool: khulnasoft.F(load_balancers.FilterOptionsParam{
				Disable: khulnasoft.F(true),
				Healthy: khulnasoft.F(false),
			}),
		}),
		OriginSteering: khulnasoft.F(load_balancers.OriginSteeringParam{
			Policy: khulnasoft.F(load_balancers.OriginSteeringPolicyRandom),
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

func TestPoolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Pools.Update(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		load_balancers.PoolUpdateParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:      khulnasoft.F("primary-dc-1"),
			Origins: khulnasoft.F([]load_balancers.OriginParam{{
				Address: khulnasoft.F("0.0.0.0"),
				Enabled: khulnasoft.F(true),
				Header: khulnasoft.F(load_balancers.HeaderParam{
					Host: khulnasoft.F([]load_balancers.HostParam{"example.com", "example.com", "example.com"}),
				}),
				Name:             khulnasoft.F("app-server-1"),
				VirtualNetworkID: khulnasoft.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           khulnasoft.F(0.600000),
			}, {
				Address: khulnasoft.F("0.0.0.0"),
				Enabled: khulnasoft.F(true),
				Header: khulnasoft.F(load_balancers.HeaderParam{
					Host: khulnasoft.F([]load_balancers.HostParam{"example.com", "example.com", "example.com"}),
				}),
				Name:             khulnasoft.F("app-server-1"),
				VirtualNetworkID: khulnasoft.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           khulnasoft.F(0.600000),
			}, {
				Address: khulnasoft.F("0.0.0.0"),
				Enabled: khulnasoft.F(true),
				Header: khulnasoft.F(load_balancers.HeaderParam{
					Host: khulnasoft.F([]load_balancers.HostParam{"example.com", "example.com", "example.com"}),
				}),
				Name:             khulnasoft.F("app-server-1"),
				VirtualNetworkID: khulnasoft.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           khulnasoft.F(0.600000),
			}}),
			CheckRegions: khulnasoft.F([]load_balancers.CheckRegion{load_balancers.CheckRegionWnam, load_balancers.CheckRegionEnam}),
			Description:  khulnasoft.F("Primary data center - Provider XYZ"),
			Enabled:      khulnasoft.F(false),
			Latitude:     khulnasoft.F(0.000000),
			LoadShedding: khulnasoft.F(load_balancers.LoadSheddingParam{
				DefaultPercent: khulnasoft.F(0.000000),
				DefaultPolicy:  khulnasoft.F(load_balancers.LoadSheddingDefaultPolicyRandom),
				SessionPercent: khulnasoft.F(0.000000),
				SessionPolicy:  khulnasoft.F(load_balancers.LoadSheddingSessionPolicyHash),
			}),
			Longitude:         khulnasoft.F(0.000000),
			MinimumOrigins:    khulnasoft.F(int64(0)),
			Monitor:           khulnasoft.F("monitor"),
			NotificationEmail: khulnasoft.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: khulnasoft.F(load_balancers.NotificationFilterParam{
				Origin: khulnasoft.F(load_balancers.FilterOptionsParam{
					Disable: khulnasoft.F(true),
					Healthy: khulnasoft.F(true),
				}),
				Pool: khulnasoft.F(load_balancers.FilterOptionsParam{
					Disable: khulnasoft.F(true),
					Healthy: khulnasoft.F(false),
				}),
			}),
			OriginSteering: khulnasoft.F(load_balancers.OriginSteeringParam{
				Policy: khulnasoft.F(load_balancers.OriginSteeringPolicyRandom),
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

func TestPoolListWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Pools.List(context.TODO(), load_balancers.PoolListParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Monitor:   khulnasoft.F("monitor"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPoolDelete(t *testing.T) {
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
	_, err := client.LoadBalancers.Pools.Delete(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		load_balancers.PoolDeleteParams{
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

func TestPoolEditWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Pools.Edit(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		load_balancers.PoolEditParams{
			AccountID:    khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CheckRegions: khulnasoft.F([]load_balancers.CheckRegion{load_balancers.CheckRegionWnam, load_balancers.CheckRegionEnam}),
			Description:  khulnasoft.F("Primary data center - Provider XYZ"),
			Enabled:      khulnasoft.F(false),
			Latitude:     khulnasoft.F(0.000000),
			LoadShedding: khulnasoft.F(load_balancers.LoadSheddingParam{
				DefaultPercent: khulnasoft.F(0.000000),
				DefaultPolicy:  khulnasoft.F(load_balancers.LoadSheddingDefaultPolicyRandom),
				SessionPercent: khulnasoft.F(0.000000),
				SessionPolicy:  khulnasoft.F(load_balancers.LoadSheddingSessionPolicyHash),
			}),
			Longitude:         khulnasoft.F(0.000000),
			MinimumOrigins:    khulnasoft.F(int64(0)),
			Monitor:           khulnasoft.F("monitor"),
			Name:              khulnasoft.F("primary-dc-1"),
			NotificationEmail: khulnasoft.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: khulnasoft.F(load_balancers.NotificationFilterParam{
				Origin: khulnasoft.F(load_balancers.FilterOptionsParam{
					Disable: khulnasoft.F(true),
					Healthy: khulnasoft.F(true),
				}),
				Pool: khulnasoft.F(load_balancers.FilterOptionsParam{
					Disable: khulnasoft.F(true),
					Healthy: khulnasoft.F(false),
				}),
			}),
			OriginSteering: khulnasoft.F(load_balancers.OriginSteeringParam{
				Policy: khulnasoft.F(load_balancers.OriginSteeringPolicyRandom),
			}),
			Origins: khulnasoft.F([]load_balancers.OriginParam{{
				Address: khulnasoft.F("0.0.0.0"),
				Enabled: khulnasoft.F(true),
				Header: khulnasoft.F(load_balancers.HeaderParam{
					Host: khulnasoft.F([]load_balancers.HostParam{"example.com", "example.com", "example.com"}),
				}),
				Name:             khulnasoft.F("app-server-1"),
				VirtualNetworkID: khulnasoft.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           khulnasoft.F(0.600000),
			}, {
				Address: khulnasoft.F("0.0.0.0"),
				Enabled: khulnasoft.F(true),
				Header: khulnasoft.F(load_balancers.HeaderParam{
					Host: khulnasoft.F([]load_balancers.HostParam{"example.com", "example.com", "example.com"}),
				}),
				Name:             khulnasoft.F("app-server-1"),
				VirtualNetworkID: khulnasoft.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           khulnasoft.F(0.600000),
			}, {
				Address: khulnasoft.F("0.0.0.0"),
				Enabled: khulnasoft.F(true),
				Header: khulnasoft.F(load_balancers.HeaderParam{
					Host: khulnasoft.F([]load_balancers.HostParam{"example.com", "example.com", "example.com"}),
				}),
				Name:             khulnasoft.F("app-server-1"),
				VirtualNetworkID: khulnasoft.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           khulnasoft.F(0.600000),
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

func TestPoolGet(t *testing.T) {
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
	_, err := client.LoadBalancers.Pools.Get(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		load_balancers.PoolGetParams{
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
