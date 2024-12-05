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

func TestLoadBalancerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.New(context.TODO(), load_balancers.LoadBalancerNewParams{
		ZoneID:       khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		DefaultPools: khulnasoft.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
		FallbackPool: khulnasoft.F("fallback_pool"),
		Name:         khulnasoft.F("www.example.com"),
		AdaptiveRouting: khulnasoft.F(load_balancers.AdaptiveRoutingParam{
			FailoverAcrossPools: khulnasoft.F(true),
		}),
		CountryPools: khulnasoft.F(map[string][]string{
			"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
			"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
		}),
		Description: khulnasoft.F("Load Balancer for www.example.com"),
		LocationStrategy: khulnasoft.F(load_balancers.LocationStrategyParam{
			Mode:      khulnasoft.F(load_balancers.LocationStrategyModePop),
			PreferECS: khulnasoft.F(load_balancers.LocationStrategyPreferECSAlways),
		}),
		Networks: khulnasoft.F([]string{"string", "string", "string"}),
		PopPools: khulnasoft.F(map[string][]string{
			"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
			"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
			"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
		}),
		Proxied: khulnasoft.F(true),
		RandomSteering: khulnasoft.F(load_balancers.RandomSteeringParam{
			DefaultWeight: khulnasoft.F(0.200000),
			PoolWeights: khulnasoft.F(load_balancers.RandomSteeringPoolWeightsParam{
				Key:   khulnasoft.F("key"),
				Value: khulnasoft.F(0.000000),
			}),
		}),
		RegionPools: khulnasoft.F(map[string][]string{
			"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
			"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
		}),
		Rules: khulnasoft.F([]load_balancers.RulesParam{{
			Condition: khulnasoft.F("http.request.uri.path contains \"/testing\""),
			Disabled:  khulnasoft.F(true),
			FixedResponse: khulnasoft.F(load_balancers.RulesFixedResponseParam{
				ContentType: khulnasoft.F("application/json"),
				Location:    khulnasoft.F("www.example.com"),
				MessageBody: khulnasoft.F("Testing Hello"),
				StatusCode:  khulnasoft.F(int64(0)),
			}),
			Name: khulnasoft.F("route the path /testing to testing datacenter."),
			Overrides: khulnasoft.F(load_balancers.RulesOverridesParam{
				AdaptiveRouting: khulnasoft.F(load_balancers.AdaptiveRoutingParam{
					FailoverAcrossPools: khulnasoft.F(true),
				}),
				CountryPools: khulnasoft.F(map[string][]string{
					"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
					"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
				}),
				DefaultPools: khulnasoft.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
				FallbackPool: khulnasoft.F("fallback_pool"),
				LocationStrategy: khulnasoft.F(load_balancers.LocationStrategyParam{
					Mode:      khulnasoft.F(load_balancers.LocationStrategyModePop),
					PreferECS: khulnasoft.F(load_balancers.LocationStrategyPreferECSAlways),
				}),
				PopPools: khulnasoft.F(map[string][]string{
					"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
					"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
				}),
				RandomSteering: khulnasoft.F(load_balancers.RandomSteeringParam{
					DefaultWeight: khulnasoft.F(0.200000),
					PoolWeights: khulnasoft.F(load_balancers.RandomSteeringPoolWeightsParam{
						Key:   khulnasoft.F("key"),
						Value: khulnasoft.F(0.000000),
					}),
				}),
				RegionPools: khulnasoft.F(map[string][]string{
					"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
					"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
				}),
				SessionAffinity: khulnasoft.F(load_balancers.SessionAffinityNone),
				SessionAffinityAttributes: khulnasoft.F(load_balancers.SessionAffinityAttributesParam{
					DrainDuration:        khulnasoft.F(100.000000),
					Headers:              khulnasoft.F([]string{"x"}),
					RequireAllHeaders:    khulnasoft.F(true),
					Samesite:             khulnasoft.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
					Secure:               khulnasoft.F(load_balancers.SessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: khulnasoft.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
				}),
				SessionAffinityTTL: khulnasoft.F(1800.000000),
				SteeringPolicy:     khulnasoft.F(load_balancers.SteeringPolicyOff),
				TTL:                khulnasoft.F(30.000000),
			}),
			Priority:   khulnasoft.F(int64(0)),
			Terminates: khulnasoft.F(true),
		}, {
			Condition: khulnasoft.F("http.request.uri.path contains \"/testing\""),
			Disabled:  khulnasoft.F(true),
			FixedResponse: khulnasoft.F(load_balancers.RulesFixedResponseParam{
				ContentType: khulnasoft.F("application/json"),
				Location:    khulnasoft.F("www.example.com"),
				MessageBody: khulnasoft.F("Testing Hello"),
				StatusCode:  khulnasoft.F(int64(0)),
			}),
			Name: khulnasoft.F("route the path /testing to testing datacenter."),
			Overrides: khulnasoft.F(load_balancers.RulesOverridesParam{
				AdaptiveRouting: khulnasoft.F(load_balancers.AdaptiveRoutingParam{
					FailoverAcrossPools: khulnasoft.F(true),
				}),
				CountryPools: khulnasoft.F(map[string][]string{
					"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
					"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
				}),
				DefaultPools: khulnasoft.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
				FallbackPool: khulnasoft.F("fallback_pool"),
				LocationStrategy: khulnasoft.F(load_balancers.LocationStrategyParam{
					Mode:      khulnasoft.F(load_balancers.LocationStrategyModePop),
					PreferECS: khulnasoft.F(load_balancers.LocationStrategyPreferECSAlways),
				}),
				PopPools: khulnasoft.F(map[string][]string{
					"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
					"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
				}),
				RandomSteering: khulnasoft.F(load_balancers.RandomSteeringParam{
					DefaultWeight: khulnasoft.F(0.200000),
					PoolWeights: khulnasoft.F(load_balancers.RandomSteeringPoolWeightsParam{
						Key:   khulnasoft.F("key"),
						Value: khulnasoft.F(0.000000),
					}),
				}),
				RegionPools: khulnasoft.F(map[string][]string{
					"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
					"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
				}),
				SessionAffinity: khulnasoft.F(load_balancers.SessionAffinityNone),
				SessionAffinityAttributes: khulnasoft.F(load_balancers.SessionAffinityAttributesParam{
					DrainDuration:        khulnasoft.F(100.000000),
					Headers:              khulnasoft.F([]string{"x"}),
					RequireAllHeaders:    khulnasoft.F(true),
					Samesite:             khulnasoft.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
					Secure:               khulnasoft.F(load_balancers.SessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: khulnasoft.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
				}),
				SessionAffinityTTL: khulnasoft.F(1800.000000),
				SteeringPolicy:     khulnasoft.F(load_balancers.SteeringPolicyOff),
				TTL:                khulnasoft.F(30.000000),
			}),
			Priority:   khulnasoft.F(int64(0)),
			Terminates: khulnasoft.F(true),
		}, {
			Condition: khulnasoft.F("http.request.uri.path contains \"/testing\""),
			Disabled:  khulnasoft.F(true),
			FixedResponse: khulnasoft.F(load_balancers.RulesFixedResponseParam{
				ContentType: khulnasoft.F("application/json"),
				Location:    khulnasoft.F("www.example.com"),
				MessageBody: khulnasoft.F("Testing Hello"),
				StatusCode:  khulnasoft.F(int64(0)),
			}),
			Name: khulnasoft.F("route the path /testing to testing datacenter."),
			Overrides: khulnasoft.F(load_balancers.RulesOverridesParam{
				AdaptiveRouting: khulnasoft.F(load_balancers.AdaptiveRoutingParam{
					FailoverAcrossPools: khulnasoft.F(true),
				}),
				CountryPools: khulnasoft.F(map[string][]string{
					"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
					"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
				}),
				DefaultPools: khulnasoft.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
				FallbackPool: khulnasoft.F("fallback_pool"),
				LocationStrategy: khulnasoft.F(load_balancers.LocationStrategyParam{
					Mode:      khulnasoft.F(load_balancers.LocationStrategyModePop),
					PreferECS: khulnasoft.F(load_balancers.LocationStrategyPreferECSAlways),
				}),
				PopPools: khulnasoft.F(map[string][]string{
					"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
					"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
				}),
				RandomSteering: khulnasoft.F(load_balancers.RandomSteeringParam{
					DefaultWeight: khulnasoft.F(0.200000),
					PoolWeights: khulnasoft.F(load_balancers.RandomSteeringPoolWeightsParam{
						Key:   khulnasoft.F("key"),
						Value: khulnasoft.F(0.000000),
					}),
				}),
				RegionPools: khulnasoft.F(map[string][]string{
					"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
					"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
				}),
				SessionAffinity: khulnasoft.F(load_balancers.SessionAffinityNone),
				SessionAffinityAttributes: khulnasoft.F(load_balancers.SessionAffinityAttributesParam{
					DrainDuration:        khulnasoft.F(100.000000),
					Headers:              khulnasoft.F([]string{"x"}),
					RequireAllHeaders:    khulnasoft.F(true),
					Samesite:             khulnasoft.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
					Secure:               khulnasoft.F(load_balancers.SessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: khulnasoft.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
				}),
				SessionAffinityTTL: khulnasoft.F(1800.000000),
				SteeringPolicy:     khulnasoft.F(load_balancers.SteeringPolicyOff),
				TTL:                khulnasoft.F(30.000000),
			}),
			Priority:   khulnasoft.F(int64(0)),
			Terminates: khulnasoft.F(true),
		}}),
		SessionAffinity: khulnasoft.F(load_balancers.SessionAffinityNone),
		SessionAffinityAttributes: khulnasoft.F(load_balancers.SessionAffinityAttributesParam{
			DrainDuration:        khulnasoft.F(100.000000),
			Headers:              khulnasoft.F([]string{"x"}),
			RequireAllHeaders:    khulnasoft.F(true),
			Samesite:             khulnasoft.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
			Secure:               khulnasoft.F(load_balancers.SessionAffinityAttributesSecureAuto),
			ZeroDowntimeFailover: khulnasoft.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
		}),
		SessionAffinityTTL: khulnasoft.F(1800.000000),
		SteeringPolicy:     khulnasoft.F(load_balancers.SteeringPolicyOff),
		TTL:                khulnasoft.F(30.000000),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		load_balancers.LoadBalancerUpdateParams{
			ZoneID:       khulnasoft.F("699d98642c564d2e855e9661899b7252"),
			DefaultPools: khulnasoft.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			FallbackPool: khulnasoft.F("fallback_pool"),
			Name:         khulnasoft.F("www.example.com"),
			AdaptiveRouting: khulnasoft.F(load_balancers.AdaptiveRoutingParam{
				FailoverAcrossPools: khulnasoft.F(true),
			}),
			CountryPools: khulnasoft.F(map[string][]string{
				"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
				"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			Description: khulnasoft.F("Load Balancer for www.example.com"),
			Enabled:     khulnasoft.F(true),
			LocationStrategy: khulnasoft.F(load_balancers.LocationStrategyParam{
				Mode:      khulnasoft.F(load_balancers.LocationStrategyModePop),
				PreferECS: khulnasoft.F(load_balancers.LocationStrategyPreferECSAlways),
			}),
			Networks: khulnasoft.F([]string{"string", "string", "string"}),
			PopPools: khulnasoft.F(map[string][]string{
				"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
				"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
				"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			Proxied: khulnasoft.F(true),
			RandomSteering: khulnasoft.F(load_balancers.RandomSteeringParam{
				DefaultWeight: khulnasoft.F(0.200000),
				PoolWeights: khulnasoft.F(load_balancers.RandomSteeringPoolWeightsParam{
					Key:   khulnasoft.F("key"),
					Value: khulnasoft.F(0.000000),
				}),
			}),
			RegionPools: khulnasoft.F(map[string][]string{
				"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
				"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
			}),
			Rules: khulnasoft.F([]load_balancers.RulesParam{{
				Condition: khulnasoft.F("http.request.uri.path contains \"/testing\""),
				Disabled:  khulnasoft.F(true),
				FixedResponse: khulnasoft.F(load_balancers.RulesFixedResponseParam{
					ContentType: khulnasoft.F("application/json"),
					Location:    khulnasoft.F("www.example.com"),
					MessageBody: khulnasoft.F("Testing Hello"),
					StatusCode:  khulnasoft.F(int64(0)),
				}),
				Name: khulnasoft.F("route the path /testing to testing datacenter."),
				Overrides: khulnasoft.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: khulnasoft.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: khulnasoft.F(true),
					}),
					CountryPools: khulnasoft.F(map[string][]string{
						"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
						"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					DefaultPools: khulnasoft.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: khulnasoft.F("fallback_pool"),
					LocationStrategy: khulnasoft.F(load_balancers.LocationStrategyParam{
						Mode:      khulnasoft.F(load_balancers.LocationStrategyModePop),
						PreferECS: khulnasoft.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					PopPools: khulnasoft.F(map[string][]string{
						"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
						"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
						"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					RandomSteering: khulnasoft.F(load_balancers.RandomSteeringParam{
						DefaultWeight: khulnasoft.F(0.200000),
						PoolWeights: khulnasoft.F(load_balancers.RandomSteeringPoolWeightsParam{
							Key:   khulnasoft.F("key"),
							Value: khulnasoft.F(0.000000),
						}),
					}),
					RegionPools: khulnasoft.F(map[string][]string{
						"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
						"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					}),
					SessionAffinity: khulnasoft.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: khulnasoft.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        khulnasoft.F(100.000000),
						Headers:              khulnasoft.F([]string{"x"}),
						RequireAllHeaders:    khulnasoft.F(true),
						Samesite:             khulnasoft.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               khulnasoft.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: khulnasoft.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: khulnasoft.F(1800.000000),
					SteeringPolicy:     khulnasoft.F(load_balancers.SteeringPolicyOff),
					TTL:                khulnasoft.F(30.000000),
				}),
				Priority:   khulnasoft.F(int64(0)),
				Terminates: khulnasoft.F(true),
			}, {
				Condition: khulnasoft.F("http.request.uri.path contains \"/testing\""),
				Disabled:  khulnasoft.F(true),
				FixedResponse: khulnasoft.F(load_balancers.RulesFixedResponseParam{
					ContentType: khulnasoft.F("application/json"),
					Location:    khulnasoft.F("www.example.com"),
					MessageBody: khulnasoft.F("Testing Hello"),
					StatusCode:  khulnasoft.F(int64(0)),
				}),
				Name: khulnasoft.F("route the path /testing to testing datacenter."),
				Overrides: khulnasoft.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: khulnasoft.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: khulnasoft.F(true),
					}),
					CountryPools: khulnasoft.F(map[string][]string{
						"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
						"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					DefaultPools: khulnasoft.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: khulnasoft.F("fallback_pool"),
					LocationStrategy: khulnasoft.F(load_balancers.LocationStrategyParam{
						Mode:      khulnasoft.F(load_balancers.LocationStrategyModePop),
						PreferECS: khulnasoft.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					PopPools: khulnasoft.F(map[string][]string{
						"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
						"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
						"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					RandomSteering: khulnasoft.F(load_balancers.RandomSteeringParam{
						DefaultWeight: khulnasoft.F(0.200000),
						PoolWeights: khulnasoft.F(load_balancers.RandomSteeringPoolWeightsParam{
							Key:   khulnasoft.F("key"),
							Value: khulnasoft.F(0.000000),
						}),
					}),
					RegionPools: khulnasoft.F(map[string][]string{
						"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
						"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					}),
					SessionAffinity: khulnasoft.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: khulnasoft.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        khulnasoft.F(100.000000),
						Headers:              khulnasoft.F([]string{"x"}),
						RequireAllHeaders:    khulnasoft.F(true),
						Samesite:             khulnasoft.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               khulnasoft.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: khulnasoft.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: khulnasoft.F(1800.000000),
					SteeringPolicy:     khulnasoft.F(load_balancers.SteeringPolicyOff),
					TTL:                khulnasoft.F(30.000000),
				}),
				Priority:   khulnasoft.F(int64(0)),
				Terminates: khulnasoft.F(true),
			}, {
				Condition: khulnasoft.F("http.request.uri.path contains \"/testing\""),
				Disabled:  khulnasoft.F(true),
				FixedResponse: khulnasoft.F(load_balancers.RulesFixedResponseParam{
					ContentType: khulnasoft.F("application/json"),
					Location:    khulnasoft.F("www.example.com"),
					MessageBody: khulnasoft.F("Testing Hello"),
					StatusCode:  khulnasoft.F(int64(0)),
				}),
				Name: khulnasoft.F("route the path /testing to testing datacenter."),
				Overrides: khulnasoft.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: khulnasoft.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: khulnasoft.F(true),
					}),
					CountryPools: khulnasoft.F(map[string][]string{
						"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
						"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					DefaultPools: khulnasoft.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: khulnasoft.F("fallback_pool"),
					LocationStrategy: khulnasoft.F(load_balancers.LocationStrategyParam{
						Mode:      khulnasoft.F(load_balancers.LocationStrategyModePop),
						PreferECS: khulnasoft.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					PopPools: khulnasoft.F(map[string][]string{
						"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
						"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
						"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					RandomSteering: khulnasoft.F(load_balancers.RandomSteeringParam{
						DefaultWeight: khulnasoft.F(0.200000),
						PoolWeights: khulnasoft.F(load_balancers.RandomSteeringPoolWeightsParam{
							Key:   khulnasoft.F("key"),
							Value: khulnasoft.F(0.000000),
						}),
					}),
					RegionPools: khulnasoft.F(map[string][]string{
						"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
						"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					}),
					SessionAffinity: khulnasoft.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: khulnasoft.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        khulnasoft.F(100.000000),
						Headers:              khulnasoft.F([]string{"x"}),
						RequireAllHeaders:    khulnasoft.F(true),
						Samesite:             khulnasoft.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               khulnasoft.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: khulnasoft.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: khulnasoft.F(1800.000000),
					SteeringPolicy:     khulnasoft.F(load_balancers.SteeringPolicyOff),
					TTL:                khulnasoft.F(30.000000),
				}),
				Priority:   khulnasoft.F(int64(0)),
				Terminates: khulnasoft.F(true),
			}}),
			SessionAffinity: khulnasoft.F(load_balancers.SessionAffinityNone),
			SessionAffinityAttributes: khulnasoft.F(load_balancers.SessionAffinityAttributesParam{
				DrainDuration:        khulnasoft.F(100.000000),
				Headers:              khulnasoft.F([]string{"x"}),
				RequireAllHeaders:    khulnasoft.F(true),
				Samesite:             khulnasoft.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
				Secure:               khulnasoft.F(load_balancers.SessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: khulnasoft.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
			}),
			SessionAffinityTTL: khulnasoft.F(1800.000000),
			SteeringPolicy:     khulnasoft.F(load_balancers.SteeringPolicyOff),
			TTL:                khulnasoft.F(30.000000),
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

func TestLoadBalancerList(t *testing.T) {
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
	_, err := client.LoadBalancers.List(context.TODO(), load_balancers.LoadBalancerListParams{
		ZoneID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerDelete(t *testing.T) {
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
	_, err := client.LoadBalancers.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		load_balancers.LoadBalancerDeleteParams{
			ZoneID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
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

func TestLoadBalancerEditWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Edit(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		load_balancers.LoadBalancerEditParams{
			ZoneID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
			AdaptiveRouting: khulnasoft.F(load_balancers.AdaptiveRoutingParam{
				FailoverAcrossPools: khulnasoft.F(true),
			}),
			CountryPools: khulnasoft.F(map[string][]string{
				"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
				"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			DefaultPools: khulnasoft.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			Description:  khulnasoft.F("Load Balancer for www.example.com"),
			Enabled:      khulnasoft.F(true),
			FallbackPool: khulnasoft.F("fallback_pool"),
			LocationStrategy: khulnasoft.F(load_balancers.LocationStrategyParam{
				Mode:      khulnasoft.F(load_balancers.LocationStrategyModePop),
				PreferECS: khulnasoft.F(load_balancers.LocationStrategyPreferECSAlways),
			}),
			Name: khulnasoft.F("www.example.com"),
			PopPools: khulnasoft.F(map[string][]string{
				"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
				"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
				"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			Proxied: khulnasoft.F(true),
			RandomSteering: khulnasoft.F(load_balancers.RandomSteeringParam{
				DefaultWeight: khulnasoft.F(0.200000),
				PoolWeights: khulnasoft.F(load_balancers.RandomSteeringPoolWeightsParam{
					Key:   khulnasoft.F("key"),
					Value: khulnasoft.F(0.000000),
				}),
			}),
			RegionPools: khulnasoft.F(map[string][]string{
				"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
				"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
			}),
			Rules: khulnasoft.F([]load_balancers.RulesParam{{
				Condition: khulnasoft.F("http.request.uri.path contains \"/testing\""),
				Disabled:  khulnasoft.F(true),
				FixedResponse: khulnasoft.F(load_balancers.RulesFixedResponseParam{
					ContentType: khulnasoft.F("application/json"),
					Location:    khulnasoft.F("www.example.com"),
					MessageBody: khulnasoft.F("Testing Hello"),
					StatusCode:  khulnasoft.F(int64(0)),
				}),
				Name: khulnasoft.F("route the path /testing to testing datacenter."),
				Overrides: khulnasoft.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: khulnasoft.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: khulnasoft.F(true),
					}),
					CountryPools: khulnasoft.F(map[string][]string{
						"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
						"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					DefaultPools: khulnasoft.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: khulnasoft.F("fallback_pool"),
					LocationStrategy: khulnasoft.F(load_balancers.LocationStrategyParam{
						Mode:      khulnasoft.F(load_balancers.LocationStrategyModePop),
						PreferECS: khulnasoft.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					PopPools: khulnasoft.F(map[string][]string{
						"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
						"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
						"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					RandomSteering: khulnasoft.F(load_balancers.RandomSteeringParam{
						DefaultWeight: khulnasoft.F(0.200000),
						PoolWeights: khulnasoft.F(load_balancers.RandomSteeringPoolWeightsParam{
							Key:   khulnasoft.F("key"),
							Value: khulnasoft.F(0.000000),
						}),
					}),
					RegionPools: khulnasoft.F(map[string][]string{
						"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
						"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					}),
					SessionAffinity: khulnasoft.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: khulnasoft.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        khulnasoft.F(100.000000),
						Headers:              khulnasoft.F([]string{"x"}),
						RequireAllHeaders:    khulnasoft.F(true),
						Samesite:             khulnasoft.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               khulnasoft.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: khulnasoft.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: khulnasoft.F(1800.000000),
					SteeringPolicy:     khulnasoft.F(load_balancers.SteeringPolicyOff),
					TTL:                khulnasoft.F(30.000000),
				}),
				Priority:   khulnasoft.F(int64(0)),
				Terminates: khulnasoft.F(true),
			}, {
				Condition: khulnasoft.F("http.request.uri.path contains \"/testing\""),
				Disabled:  khulnasoft.F(true),
				FixedResponse: khulnasoft.F(load_balancers.RulesFixedResponseParam{
					ContentType: khulnasoft.F("application/json"),
					Location:    khulnasoft.F("www.example.com"),
					MessageBody: khulnasoft.F("Testing Hello"),
					StatusCode:  khulnasoft.F(int64(0)),
				}),
				Name: khulnasoft.F("route the path /testing to testing datacenter."),
				Overrides: khulnasoft.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: khulnasoft.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: khulnasoft.F(true),
					}),
					CountryPools: khulnasoft.F(map[string][]string{
						"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
						"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					DefaultPools: khulnasoft.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: khulnasoft.F("fallback_pool"),
					LocationStrategy: khulnasoft.F(load_balancers.LocationStrategyParam{
						Mode:      khulnasoft.F(load_balancers.LocationStrategyModePop),
						PreferECS: khulnasoft.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					PopPools: khulnasoft.F(map[string][]string{
						"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
						"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
						"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					RandomSteering: khulnasoft.F(load_balancers.RandomSteeringParam{
						DefaultWeight: khulnasoft.F(0.200000),
						PoolWeights: khulnasoft.F(load_balancers.RandomSteeringPoolWeightsParam{
							Key:   khulnasoft.F("key"),
							Value: khulnasoft.F(0.000000),
						}),
					}),
					RegionPools: khulnasoft.F(map[string][]string{
						"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
						"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					}),
					SessionAffinity: khulnasoft.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: khulnasoft.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        khulnasoft.F(100.000000),
						Headers:              khulnasoft.F([]string{"x"}),
						RequireAllHeaders:    khulnasoft.F(true),
						Samesite:             khulnasoft.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               khulnasoft.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: khulnasoft.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: khulnasoft.F(1800.000000),
					SteeringPolicy:     khulnasoft.F(load_balancers.SteeringPolicyOff),
					TTL:                khulnasoft.F(30.000000),
				}),
				Priority:   khulnasoft.F(int64(0)),
				Terminates: khulnasoft.F(true),
			}, {
				Condition: khulnasoft.F("http.request.uri.path contains \"/testing\""),
				Disabled:  khulnasoft.F(true),
				FixedResponse: khulnasoft.F(load_balancers.RulesFixedResponseParam{
					ContentType: khulnasoft.F("application/json"),
					Location:    khulnasoft.F("www.example.com"),
					MessageBody: khulnasoft.F("Testing Hello"),
					StatusCode:  khulnasoft.F(int64(0)),
				}),
				Name: khulnasoft.F("route the path /testing to testing datacenter."),
				Overrides: khulnasoft.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: khulnasoft.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: khulnasoft.F(true),
					}),
					CountryPools: khulnasoft.F(map[string][]string{
						"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
						"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					DefaultPools: khulnasoft.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: khulnasoft.F("fallback_pool"),
					LocationStrategy: khulnasoft.F(load_balancers.LocationStrategyParam{
						Mode:      khulnasoft.F(load_balancers.LocationStrategyModePop),
						PreferECS: khulnasoft.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					PopPools: khulnasoft.F(map[string][]string{
						"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
						"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
						"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					RandomSteering: khulnasoft.F(load_balancers.RandomSteeringParam{
						DefaultWeight: khulnasoft.F(0.200000),
						PoolWeights: khulnasoft.F(load_balancers.RandomSteeringPoolWeightsParam{
							Key:   khulnasoft.F("key"),
							Value: khulnasoft.F(0.000000),
						}),
					}),
					RegionPools: khulnasoft.F(map[string][]string{
						"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
						"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					}),
					SessionAffinity: khulnasoft.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: khulnasoft.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        khulnasoft.F(100.000000),
						Headers:              khulnasoft.F([]string{"x"}),
						RequireAllHeaders:    khulnasoft.F(true),
						Samesite:             khulnasoft.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               khulnasoft.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: khulnasoft.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: khulnasoft.F(1800.000000),
					SteeringPolicy:     khulnasoft.F(load_balancers.SteeringPolicyOff),
					TTL:                khulnasoft.F(30.000000),
				}),
				Priority:   khulnasoft.F(int64(0)),
				Terminates: khulnasoft.F(true),
			}}),
			SessionAffinity: khulnasoft.F(load_balancers.SessionAffinityNone),
			SessionAffinityAttributes: khulnasoft.F(load_balancers.SessionAffinityAttributesParam{
				DrainDuration:        khulnasoft.F(100.000000),
				Headers:              khulnasoft.F([]string{"x"}),
				RequireAllHeaders:    khulnasoft.F(true),
				Samesite:             khulnasoft.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
				Secure:               khulnasoft.F(load_balancers.SessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: khulnasoft.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
			}),
			SessionAffinityTTL: khulnasoft.F(1800.000000),
			SteeringPolicy:     khulnasoft.F(load_balancers.SteeringPolicyOff),
			TTL:                khulnasoft.F(30.000000),
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

func TestLoadBalancerGet(t *testing.T) {
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
	_, err := client.LoadBalancers.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		load_balancers.LoadBalancerGetParams{
			ZoneID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
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
