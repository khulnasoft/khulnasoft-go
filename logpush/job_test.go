// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/logpush"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestJobNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Logpush.Jobs.New(context.TODO(), logpush.JobNewParams{
		DestinationConf:          khulnasoft.F("s3://mybucket/logs?region=us-west-2"),
		AccountID:                khulnasoft.F("account_id"),
		Dataset:                  khulnasoft.F("http_requests"),
		Enabled:                  khulnasoft.F(false),
		Frequency:                khulnasoft.F(logpush.JobNewParamsFrequencyHigh),
		Kind:                     khulnasoft.F(logpush.JobNewParamsKindEdge),
		LogpullOptions:           khulnasoft.F("fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"),
		MaxUploadBytes:           khulnasoft.F(int64(5000000)),
		MaxUploadIntervalSeconds: khulnasoft.F(int64(30)),
		MaxUploadRecords:         khulnasoft.F(int64(1000)),
		Name:                     khulnasoft.F("example.com"),
		OutputOptions: khulnasoft.F(logpush.OutputOptionsParam{
			BatchPrefix:     khulnasoft.F("batch_prefix"),
			BatchSuffix:     khulnasoft.F("batch_suffix"),
			Cve2021_4428:    khulnasoft.F(true),
			FieldDelimiter:  khulnasoft.F("field_delimiter"),
			FieldNames:      khulnasoft.F([]string{"ClientIP", "EdgeStartTimestamp", "RayID"}),
			OutputType:      khulnasoft.F(logpush.OutputOptionsOutputTypeNdjson),
			RecordDelimiter: khulnasoft.F("record_delimiter"),
			RecordPrefix:    khulnasoft.F("record_prefix"),
			RecordSuffix:    khulnasoft.F("record_suffix"),
			RecordTemplate:  khulnasoft.F("record_template"),
			SampleRate:      khulnasoft.F(0.000000),
			TimestampFormat: khulnasoft.F(logpush.OutputOptionsTimestampFormatUnixnano),
		}),
		OwnershipChallenge: khulnasoft.F("00000000000000000000"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Logpush.Jobs.Update(
		context.TODO(),
		int64(1),
		logpush.JobUpdateParams{
			AccountID:                khulnasoft.F("account_id"),
			DestinationConf:          khulnasoft.F("s3://mybucket/logs?region=us-west-2"),
			Enabled:                  khulnasoft.F(false),
			Frequency:                khulnasoft.F(logpush.JobUpdateParamsFrequencyHigh),
			Kind:                     khulnasoft.F(logpush.JobUpdateParamsKindEdge),
			LogpullOptions:           khulnasoft.F("fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"),
			MaxUploadBytes:           khulnasoft.F(int64(5000000)),
			MaxUploadIntervalSeconds: khulnasoft.F(int64(30)),
			MaxUploadRecords:         khulnasoft.F(int64(1000)),
			OutputOptions: khulnasoft.F(logpush.OutputOptionsParam{
				BatchPrefix:     khulnasoft.F("batch_prefix"),
				BatchSuffix:     khulnasoft.F("batch_suffix"),
				Cve2021_4428:    khulnasoft.F(true),
				FieldDelimiter:  khulnasoft.F("field_delimiter"),
				FieldNames:      khulnasoft.F([]string{"ClientIP", "EdgeStartTimestamp", "RayID"}),
				OutputType:      khulnasoft.F(logpush.OutputOptionsOutputTypeNdjson),
				RecordDelimiter: khulnasoft.F("record_delimiter"),
				RecordPrefix:    khulnasoft.F("record_prefix"),
				RecordSuffix:    khulnasoft.F("record_suffix"),
				RecordTemplate:  khulnasoft.F("record_template"),
				SampleRate:      khulnasoft.F(0.000000),
				TimestampFormat: khulnasoft.F(logpush.OutputOptionsTimestampFormatUnixnano),
			}),
			OwnershipChallenge: khulnasoft.F("00000000000000000000"),
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

func TestJobListWithOptionalParams(t *testing.T) {
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
	_, err := client.Logpush.Jobs.List(context.TODO(), logpush.JobListParams{
		AccountID: khulnasoft.F("account_id"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Logpush.Jobs.Delete(
		context.TODO(),
		int64(1),
		logpush.JobDeleteParams{
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

func TestJobGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Logpush.Jobs.Get(
		context.TODO(),
		int64(1),
		logpush.JobGetParams{
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
