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

func TestDevicePostureNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Posture.New(context.TODO(), zero_trust.DevicePostureNewParams{
		AccountID:   khulnasoft.F("699d98642c564d2e855e9661899b7252"),
		Name:        khulnasoft.F("Admin Serial Numbers"),
		Type:        khulnasoft.F(zero_trust.DevicePostureNewParamsTypeFile),
		Description: khulnasoft.F("The rule for admin serial numbers"),
		Expiration:  khulnasoft.F("1h"),
		Input: khulnasoft.F[zero_trust.DeviceInputUnionParam](zero_trust.FileInputParam{
			OperatingSystem: khulnasoft.F(zero_trust.FileInputOperatingSystemWindows),
			Path:            khulnasoft.F("/bin/cat"),
			Exists:          khulnasoft.F(true),
			Sha256:          khulnasoft.F("https://api.us-2.crowdstrike.com"),
			Thumbprint:      khulnasoft.F("0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"),
		}),
		Match: khulnasoft.F([]zero_trust.DeviceMatchParam{{
			Platform: khulnasoft.F(zero_trust.DeviceMatchPlatformWindows),
		}, {
			Platform: khulnasoft.F(zero_trust.DeviceMatchPlatformWindows),
		}, {
			Platform: khulnasoft.F(zero_trust.DeviceMatchPlatformWindows),
		}}),
		Schedule: khulnasoft.F("1h"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDevicePostureUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Posture.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePostureUpdateParams{
			AccountID:   khulnasoft.F("699d98642c564d2e855e9661899b7252"),
			Name:        khulnasoft.F("Admin Serial Numbers"),
			Type:        khulnasoft.F(zero_trust.DevicePostureUpdateParamsTypeFile),
			Description: khulnasoft.F("The rule for admin serial numbers"),
			Expiration:  khulnasoft.F("1h"),
			Input: khulnasoft.F[zero_trust.DeviceInputUnionParam](zero_trust.FileInputParam{
				OperatingSystem: khulnasoft.F(zero_trust.FileInputOperatingSystemWindows),
				Path:            khulnasoft.F("/bin/cat"),
				Exists:          khulnasoft.F(true),
				Sha256:          khulnasoft.F("https://api.us-2.crowdstrike.com"),
				Thumbprint:      khulnasoft.F("0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"),
			}),
			Match: khulnasoft.F([]zero_trust.DeviceMatchParam{{
				Platform: khulnasoft.F(zero_trust.DeviceMatchPlatformWindows),
			}, {
				Platform: khulnasoft.F(zero_trust.DeviceMatchPlatformWindows),
			}, {
				Platform: khulnasoft.F(zero_trust.DeviceMatchPlatformWindows),
			}}),
			Schedule: khulnasoft.F("1h"),
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

func TestDevicePostureList(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Posture.List(context.TODO(), zero_trust.DevicePostureListParams{
		AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDevicePostureDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Posture.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePostureDeleteParams{
			AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
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

func TestDevicePostureGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Posture.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePostureGetParams{
			AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
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
