// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_authorities_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/certificate_authorities"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestHostnameAssociationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CertificateAuthorities.HostnameAssociations.Update(context.TODO(), certificate_authorities.HostnameAssociationUpdateParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		TLSHostnameAssociation: certificate_authorities.TLSHostnameAssociationParam{
			Hostnames:         khulnasoft.F([]certificate_authorities.HostnameAssociationParam{"api.example.com", "api.example.com", "api.example.com"}),
			MTLSCertificateID: khulnasoft.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		},
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHostnameAssociationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.CertificateAuthorities.HostnameAssociations.Get(context.TODO(), certificate_authorities.HostnameAssociationGetParams{
		ZoneID:            khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		MTLSCertificateID: khulnasoft.F("b2134436-2555-4acf-be5b-26c48136575e"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
