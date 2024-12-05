// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/zero_trust"
)

func TestDLPProfileCustomNew(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Custom.New(context.TODO(), zero_trust.DLPProfileCustomNewParams{
		AccountID: khulnasoft.F("account_id"),
		Profiles: khulnasoft.F([]zero_trust.DLPProfileCustomNewParamsProfile{{
			Entries: khulnasoft.F([]zero_trust.DLPProfileCustomNewParamsProfilesEntryUnion{zero_trust.DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry{
				Enabled: khulnasoft.F(true),
				Name:    khulnasoft.F("name"),
				Pattern: khulnasoft.F(zero_trust.PatternParam{
					Regex:      khulnasoft.F("regex"),
					Validation: khulnasoft.F(zero_trust.PatternValidationLuhn),
				}),
			}, zero_trust.DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry{
				Enabled: khulnasoft.F(true),
				Name:    khulnasoft.F("name"),
				Pattern: khulnasoft.F(zero_trust.PatternParam{
					Regex:      khulnasoft.F("regex"),
					Validation: khulnasoft.F(zero_trust.PatternValidationLuhn),
				}),
			}, zero_trust.DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry{
				Enabled: khulnasoft.F(true),
				Name:    khulnasoft.F("name"),
				Pattern: khulnasoft.F(zero_trust.PatternParam{
					Regex:      khulnasoft.F("regex"),
					Validation: khulnasoft.F(zero_trust.PatternValidationLuhn),
				}),
			}}),
			Name:              khulnasoft.F("name"),
			AllowedMatchCount: khulnasoft.F(int64(5)),
			ContextAwareness: khulnasoft.F(zero_trust.ContextAwarenessParam{
				Enabled: khulnasoft.F(true),
				Skip: khulnasoft.F(zero_trust.SkipConfigurationParam{
					Files: khulnasoft.F(true),
				}),
			}),
			Description: khulnasoft.F("description"),
			OCREnabled:  khulnasoft.F(true),
			SharedEntries: khulnasoft.F([]zero_trust.DLPProfileCustomNewParamsProfilesSharedEntryUnion{zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustom{
				Enabled:   khulnasoft.F(true),
				EntryID:   khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: khulnasoft.F(zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryTypeCustom),
			}, zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustom{
				Enabled:   khulnasoft.F(true),
				EntryID:   khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: khulnasoft.F(zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryTypeCustom),
			}, zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustom{
				Enabled:   khulnasoft.F(true),
				EntryID:   khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: khulnasoft.F(zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryTypeCustom),
			}}),
		}, {
			Entries: khulnasoft.F([]zero_trust.DLPProfileCustomNewParamsProfilesEntryUnion{zero_trust.DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry{
				Enabled: khulnasoft.F(true),
				Name:    khulnasoft.F("name"),
				Pattern: khulnasoft.F(zero_trust.PatternParam{
					Regex:      khulnasoft.F("regex"),
					Validation: khulnasoft.F(zero_trust.PatternValidationLuhn),
				}),
			}, zero_trust.DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry{
				Enabled: khulnasoft.F(true),
				Name:    khulnasoft.F("name"),
				Pattern: khulnasoft.F(zero_trust.PatternParam{
					Regex:      khulnasoft.F("regex"),
					Validation: khulnasoft.F(zero_trust.PatternValidationLuhn),
				}),
			}, zero_trust.DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry{
				Enabled: khulnasoft.F(true),
				Name:    khulnasoft.F("name"),
				Pattern: khulnasoft.F(zero_trust.PatternParam{
					Regex:      khulnasoft.F("regex"),
					Validation: khulnasoft.F(zero_trust.PatternValidationLuhn),
				}),
			}}),
			Name:              khulnasoft.F("name"),
			AllowedMatchCount: khulnasoft.F(int64(5)),
			ContextAwareness: khulnasoft.F(zero_trust.ContextAwarenessParam{
				Enabled: khulnasoft.F(true),
				Skip: khulnasoft.F(zero_trust.SkipConfigurationParam{
					Files: khulnasoft.F(true),
				}),
			}),
			Description: khulnasoft.F("description"),
			OCREnabled:  khulnasoft.F(true),
			SharedEntries: khulnasoft.F([]zero_trust.DLPProfileCustomNewParamsProfilesSharedEntryUnion{zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustom{
				Enabled:   khulnasoft.F(true),
				EntryID:   khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: khulnasoft.F(zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryTypeCustom),
			}, zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustom{
				Enabled:   khulnasoft.F(true),
				EntryID:   khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: khulnasoft.F(zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryTypeCustom),
			}, zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustom{
				Enabled:   khulnasoft.F(true),
				EntryID:   khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: khulnasoft.F(zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryTypeCustom),
			}}),
		}, {
			Entries: khulnasoft.F([]zero_trust.DLPProfileCustomNewParamsProfilesEntryUnion{zero_trust.DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry{
				Enabled: khulnasoft.F(true),
				Name:    khulnasoft.F("name"),
				Pattern: khulnasoft.F(zero_trust.PatternParam{
					Regex:      khulnasoft.F("regex"),
					Validation: khulnasoft.F(zero_trust.PatternValidationLuhn),
				}),
			}, zero_trust.DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry{
				Enabled: khulnasoft.F(true),
				Name:    khulnasoft.F("name"),
				Pattern: khulnasoft.F(zero_trust.PatternParam{
					Regex:      khulnasoft.F("regex"),
					Validation: khulnasoft.F(zero_trust.PatternValidationLuhn),
				}),
			}, zero_trust.DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry{
				Enabled: khulnasoft.F(true),
				Name:    khulnasoft.F("name"),
				Pattern: khulnasoft.F(zero_trust.PatternParam{
					Regex:      khulnasoft.F("regex"),
					Validation: khulnasoft.F(zero_trust.PatternValidationLuhn),
				}),
			}}),
			Name:              khulnasoft.F("name"),
			AllowedMatchCount: khulnasoft.F(int64(5)),
			ContextAwareness: khulnasoft.F(zero_trust.ContextAwarenessParam{
				Enabled: khulnasoft.F(true),
				Skip: khulnasoft.F(zero_trust.SkipConfigurationParam{
					Files: khulnasoft.F(true),
				}),
			}),
			Description: khulnasoft.F("description"),
			OCREnabled:  khulnasoft.F(true),
			SharedEntries: khulnasoft.F([]zero_trust.DLPProfileCustomNewParamsProfilesSharedEntryUnion{zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustom{
				Enabled:   khulnasoft.F(true),
				EntryID:   khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: khulnasoft.F(zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryTypeCustom),
			}, zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustom{
				Enabled:   khulnasoft.F(true),
				EntryID:   khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: khulnasoft.F(zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryTypeCustom),
			}, zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustom{
				Enabled:   khulnasoft.F(true),
				EntryID:   khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: khulnasoft.F(zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryTypeCustom),
			}}),
		}}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDLPProfileCustomUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Custom.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.DLPProfileCustomUpdateParams{
			AccountID: khulnasoft.F("account_id"),
			Entries: khulnasoft.F([]zero_trust.DLPProfileCustomUpdateParamsEntryUnion{zero_trust.DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID{
				Enabled: khulnasoft.F(true),
				EntryID: khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Name:    khulnasoft.F("name"),
				Pattern: khulnasoft.F(zero_trust.PatternParam{
					Regex:      khulnasoft.F("regex"),
					Validation: khulnasoft.F(zero_trust.PatternValidationLuhn),
				}),
			}, zero_trust.DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID{
				Enabled: khulnasoft.F(true),
				EntryID: khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Name:    khulnasoft.F("name"),
				Pattern: khulnasoft.F(zero_trust.PatternParam{
					Regex:      khulnasoft.F("regex"),
					Validation: khulnasoft.F(zero_trust.PatternValidationLuhn),
				}),
			}, zero_trust.DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID{
				Enabled: khulnasoft.F(true),
				EntryID: khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Name:    khulnasoft.F("name"),
				Pattern: khulnasoft.F(zero_trust.PatternParam{
					Regex:      khulnasoft.F("regex"),
					Validation: khulnasoft.F(zero_trust.PatternValidationLuhn),
				}),
			}}),
			Name:              khulnasoft.F("name"),
			AllowedMatchCount: khulnasoft.F(int64(0)),
			ContextAwareness: khulnasoft.F(zero_trust.ContextAwarenessParam{
				Enabled: khulnasoft.F(true),
				Skip: khulnasoft.F(zero_trust.SkipConfigurationParam{
					Files: khulnasoft.F(true),
				}),
			}),
			Description: khulnasoft.F("description"),
			OCREnabled:  khulnasoft.F(true),
			SharedEntries: khulnasoft.F([]zero_trust.DLPProfileCustomUpdateParamsSharedEntryUnion{zero_trust.DLPProfileCustomUpdateParamsSharedEntriesPredefined{
				Enabled:   khulnasoft.F(true),
				EntryID:   khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: khulnasoft.F(zero_trust.DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryTypePredefined),
			}, zero_trust.DLPProfileCustomUpdateParamsSharedEntriesPredefined{
				Enabled:   khulnasoft.F(true),
				EntryID:   khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: khulnasoft.F(zero_trust.DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryTypePredefined),
			}, zero_trust.DLPProfileCustomUpdateParamsSharedEntriesPredefined{
				Enabled:   khulnasoft.F(true),
				EntryID:   khulnasoft.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: khulnasoft.F(zero_trust.DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryTypePredefined),
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

func TestDLPProfileCustomDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Custom.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.DLPProfileCustomDeleteParams{
			AccountID: khulnasoft.F("account_id"),
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

func TestDLPProfileCustomGet(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Custom.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.DLPProfileCustomGetParams{
			AccountID: khulnasoft.F("account_id"),
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
