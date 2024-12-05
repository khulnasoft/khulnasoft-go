// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/user"
)

func TestTokenNewWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Tokens.New(context.TODO(), user.TokenNewParams{
		Name: khulnasoft.F("readonly token"),
		Policies: khulnasoft.F([]user.PolicyParam{{
			Effect: khulnasoft.F(user.PolicyEffectAllow),
			PermissionGroups: khulnasoft.F([]user.PolicyPermissionGroupParam{{
				Meta: khulnasoft.F(user.PolicyPermissionGroupsMetaParam{
					Key:   khulnasoft.F("key"),
					Value: khulnasoft.F("value"),
				}),
			}, {
				Meta: khulnasoft.F(user.PolicyPermissionGroupsMetaParam{
					Key:   khulnasoft.F("key"),
					Value: khulnasoft.F("value"),
				}),
			}}),
			Resources: khulnasoft.F(user.PolicyResourcesParam{
				Resource: khulnasoft.F("resource"),
				Scope:    khulnasoft.F("scope"),
			}),
		}, {
			Effect: khulnasoft.F(user.PolicyEffectAllow),
			PermissionGroups: khulnasoft.F([]user.PolicyPermissionGroupParam{{
				Meta: khulnasoft.F(user.PolicyPermissionGroupsMetaParam{
					Key:   khulnasoft.F("key"),
					Value: khulnasoft.F("value"),
				}),
			}, {
				Meta: khulnasoft.F(user.PolicyPermissionGroupsMetaParam{
					Key:   khulnasoft.F("key"),
					Value: khulnasoft.F("value"),
				}),
			}}),
			Resources: khulnasoft.F(user.PolicyResourcesParam{
				Resource: khulnasoft.F("resource"),
				Scope:    khulnasoft.F("scope"),
			}),
		}, {
			Effect: khulnasoft.F(user.PolicyEffectAllow),
			PermissionGroups: khulnasoft.F([]user.PolicyPermissionGroupParam{{
				Meta: khulnasoft.F(user.PolicyPermissionGroupsMetaParam{
					Key:   khulnasoft.F("key"),
					Value: khulnasoft.F("value"),
				}),
			}, {
				Meta: khulnasoft.F(user.PolicyPermissionGroupsMetaParam{
					Key:   khulnasoft.F("key"),
					Value: khulnasoft.F("value"),
				}),
			}}),
			Resources: khulnasoft.F(user.PolicyResourcesParam{
				Resource: khulnasoft.F("resource"),
				Scope:    khulnasoft.F("scope"),
			}),
		}}),
		Condition: khulnasoft.F(user.TokenNewParamsCondition{
			RequestIP: khulnasoft.F(user.TokenNewParamsConditionRequestIP{
				In:    khulnasoft.F([]user.CIDRListParam{"123.123.123.0/24", "2606:4700::/32"}),
				NotIn: khulnasoft.F([]user.CIDRListParam{"123.123.123.100/24", "2606:4700:4700::/48"}),
			}),
		}),
		ExpiresOn: khulnasoft.F(time.Now()),
		NotBefore: khulnasoft.F(time.Now()),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTokenUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Tokens.Update(
		context.TODO(),
		"ed17574386854bf78a67040be0a770b0",
		user.TokenUpdateParams{
			Token: user.TokenParam{
				Name: khulnasoft.F("readonly token"),
				Policies: khulnasoft.F([]user.PolicyParam{{
					Effect: khulnasoft.F(user.PolicyEffectAllow),
					PermissionGroups: khulnasoft.F([]user.PolicyPermissionGroupParam{{
						Meta: khulnasoft.F(user.PolicyPermissionGroupsMetaParam{
							Key:   khulnasoft.F("key"),
							Value: khulnasoft.F("value"),
						}),
					}, {
						Meta: khulnasoft.F(user.PolicyPermissionGroupsMetaParam{
							Key:   khulnasoft.F("key"),
							Value: khulnasoft.F("value"),
						}),
					}}),
					Resources: khulnasoft.F(user.PolicyResourcesParam{
						Resource: khulnasoft.F("resource"),
						Scope:    khulnasoft.F("scope"),
					}),
				}, {
					Effect: khulnasoft.F(user.PolicyEffectAllow),
					PermissionGroups: khulnasoft.F([]user.PolicyPermissionGroupParam{{
						Meta: khulnasoft.F(user.PolicyPermissionGroupsMetaParam{
							Key:   khulnasoft.F("key"),
							Value: khulnasoft.F("value"),
						}),
					}, {
						Meta: khulnasoft.F(user.PolicyPermissionGroupsMetaParam{
							Key:   khulnasoft.F("key"),
							Value: khulnasoft.F("value"),
						}),
					}}),
					Resources: khulnasoft.F(user.PolicyResourcesParam{
						Resource: khulnasoft.F("resource"),
						Scope:    khulnasoft.F("scope"),
					}),
				}, {
					Effect: khulnasoft.F(user.PolicyEffectAllow),
					PermissionGroups: khulnasoft.F([]user.PolicyPermissionGroupParam{{
						Meta: khulnasoft.F(user.PolicyPermissionGroupsMetaParam{
							Key:   khulnasoft.F("key"),
							Value: khulnasoft.F("value"),
						}),
					}, {
						Meta: khulnasoft.F(user.PolicyPermissionGroupsMetaParam{
							Key:   khulnasoft.F("key"),
							Value: khulnasoft.F("value"),
						}),
					}}),
					Resources: khulnasoft.F(user.PolicyResourcesParam{
						Resource: khulnasoft.F("resource"),
						Scope:    khulnasoft.F("scope"),
					}),
				}}),
				Status: khulnasoft.F(user.TokenStatusActive),
				Condition: khulnasoft.F(user.TokenConditionParam{
					RequestIP: khulnasoft.F(user.TokenConditionRequestIPParam{
						In:    khulnasoft.F([]user.CIDRListParam{"123.123.123.0/24", "2606:4700::/32"}),
						NotIn: khulnasoft.F([]user.CIDRListParam{"123.123.123.100/24", "2606:4700:4700::/48"}),
					}),
				}),
				ExpiresOn: khulnasoft.F(time.Now()),
				NotBefore: khulnasoft.F(time.Now()),
			},
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

func TestTokenListWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Tokens.List(context.TODO(), user.TokenListParams{
		Direction: khulnasoft.F(user.TokenListParamsDirectionAsc),
		Page:      khulnasoft.F(1.000000),
		PerPage:   khulnasoft.F(5.000000),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTokenDelete(t *testing.T) {
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
	_, err := client.User.Tokens.Delete(context.TODO(), "ed17574386854bf78a67040be0a770b0")
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTokenGet(t *testing.T) {
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
	_, err := client.User.Tokens.Get(context.TODO(), "ed17574386854bf78a67040be0a770b0")
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTokenVerify(t *testing.T) {
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
	_, err := client.User.Tokens.Verify(context.TODO())
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
