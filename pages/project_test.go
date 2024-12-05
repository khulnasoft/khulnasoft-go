// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/pages"
)

func TestProjectNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pages.Projects.New(context.TODO(), pages.ProjectNewParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Project: pages.ProjectParam{
			BuildConfig: khulnasoft.F(pages.ProjectBuildConfigParam{
				BuildCaching:      khulnasoft.F(true),
				BuildCommand:      khulnasoft.F("npm run build"),
				DestinationDir:    khulnasoft.F("build"),
				RootDir:           khulnasoft.F("/"),
				WebAnalyticsTag:   khulnasoft.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
				WebAnalyticsToken: khulnasoft.F("021e1057c18547eca7b79f2516f06o7x"),
			}),
			DeploymentConfigs: khulnasoft.F(pages.ProjectDeploymentConfigsParam{
				Preview: khulnasoft.F(pages.ProjectDeploymentConfigsPreviewParam{
					AIBindings: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewAIBindingParam{
						"AI_BINDING": {
							ProjectID: khulnasoft.F("some-project-id"),
						},
					}),
					AnalyticsEngineDatasets: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetParam{
						"ANALYTICS_ENGINE_BINDING": {
							Dataset: khulnasoft.F("api_analytics"),
						},
					}),
					Browsers: khulnasoft.F(map[string]interface{}{
						"BROWSER": "bar",
					}),
					CompatibilityDate:  khulnasoft.F("2022-01-01"),
					CompatibilityFlags: khulnasoft.F([]string{"url_standard"}),
					D1Databases: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewD1DatabaseParam{
						"D1_BINDING": {
							ID: khulnasoft.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
						},
					}),
					DurableObjectNamespaces: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewDurableObjectNamespaceParam{
						"DO_BINDING": {
							NamespaceID: khulnasoft.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						},
					}),
					EnvVars: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewEnvVarParam{
						"foo": {
							Value: khulnasoft.F("hello world"),
							Type:  khulnasoft.F(pages.ProjectDeploymentConfigsPreviewEnvVarsTypePlainText),
						},
					}),
					HyperdriveBindings: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewHyperdriveBindingParam{
						"HYPERDRIVE": {
							ID: khulnasoft.F("a76a99bc342644deb02c38d66082262a"),
						},
					}),
					KVNamespaces: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewKVNamespaceParam{
						"KV_BINDING": {
							NamespaceID: khulnasoft.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						},
					}),
					MTLSCertificates: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewMTLSCertificateParam{
						"MTLS": {
							CertificateID: khulnasoft.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
						},
					}),
					Placement: khulnasoft.F(pages.ProjectDeploymentConfigsPreviewPlacementParam{
						Mode: khulnasoft.F("smart"),
					}),
					QueueProducers: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewQueueProducerParam{
						"QUEUE_PRODUCER_BINDING": {
							Name: khulnasoft.F("some-queue"),
						},
					}),
					R2Buckets: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewR2BucketParam{
						"R2_BINDING": {
							Jurisdiction: khulnasoft.F("eu"),
							Name:         khulnasoft.F("some-bucket"),
						},
					}),
					Services: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewServiceParam{
						"SERVICE_BINDING": {
							Entrypoint:  khulnasoft.F("MyHandler"),
							Environment: khulnasoft.F("production"),
							Service:     khulnasoft.F("example-worker"),
						},
					}),
					VectorizeBindings: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewVectorizeBindingParam{
						"VECTORIZE": {
							IndexName: khulnasoft.F("my_index"),
						},
					}),
				}),
				Production: khulnasoft.F(pages.ProjectDeploymentConfigsProductionParam{
					AIBindings: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionAIBindingParam{
						"AI_BINDING": {
							ProjectID: khulnasoft.F("some-project-id"),
						},
					}),
					AnalyticsEngineDatasets: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionAnalyticsEngineDatasetParam{
						"ANALYTICS_ENGINE_BINDING": {
							Dataset: khulnasoft.F("api_analytics"),
						},
					}),
					Browsers: khulnasoft.F(map[string]interface{}{
						"BROWSER": "bar",
					}),
					CompatibilityDate:  khulnasoft.F("2022-01-01"),
					CompatibilityFlags: khulnasoft.F([]string{"url_standard"}),
					D1Databases: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionD1DatabaseParam{
						"D1_BINDING": {
							ID: khulnasoft.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
						},
					}),
					DurableObjectNamespaces: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionDurableObjectNamespaceParam{
						"DO_BINDING": {
							NamespaceID: khulnasoft.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						},
					}),
					EnvVars: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionEnvVarParam{
						"foo": {
							Value: khulnasoft.F("hello world"),
							Type:  khulnasoft.F(pages.ProjectDeploymentConfigsProductionEnvVarsTypePlainText),
						},
					}),
					HyperdriveBindings: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionHyperdriveBindingParam{
						"HYPERDRIVE": {
							ID: khulnasoft.F("a76a99bc342644deb02c38d66082262a"),
						},
					}),
					KVNamespaces: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionKVNamespaceParam{
						"KV_BINDING": {
							NamespaceID: khulnasoft.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						},
					}),
					MTLSCertificates: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionMTLSCertificateParam{
						"MTLS": {
							CertificateID: khulnasoft.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
						},
					}),
					Placement: khulnasoft.F(pages.ProjectDeploymentConfigsProductionPlacementParam{
						Mode: khulnasoft.F("smart"),
					}),
					QueueProducers: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionQueueProducerParam{
						"QUEUE_PRODUCER_BINDING": {
							Name: khulnasoft.F("some-queue"),
						},
					}),
					R2Buckets: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionR2BucketParam{
						"R2_BINDING": {
							Jurisdiction: khulnasoft.F("eu"),
							Name:         khulnasoft.F("some-bucket"),
						},
					}),
					Services: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionServiceParam{
						"SERVICE_BINDING": {
							Entrypoint:  khulnasoft.F("MyHandler"),
							Environment: khulnasoft.F("production"),
							Service:     khulnasoft.F("example-worker"),
						},
					}),
					VectorizeBindings: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionVectorizeBindingParam{
						"VECTORIZE": {
							IndexName: khulnasoft.F("my_index"),
						},
					}),
				}),
			}),
			Name:             khulnasoft.F("NextJS Blog"),
			ProductionBranch: khulnasoft.F("main"),
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

func TestProjectList(t *testing.T) {
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
	_, err := client.Pages.Projects.List(context.TODO(), pages.ProjectListParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectDelete(t *testing.T) {
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
	_, err := client.Pages.Projects.Delete(
		context.TODO(),
		"this-is-my-project-01",
		pages.ProjectDeleteParams{
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

func TestProjectEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Pages.Projects.Edit(
		context.TODO(),
		"this-is-my-project-01",
		pages.ProjectEditParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Project: pages.ProjectParam{
				BuildConfig: khulnasoft.F(pages.ProjectBuildConfigParam{
					BuildCaching:      khulnasoft.F(true),
					BuildCommand:      khulnasoft.F("npm run build"),
					DestinationDir:    khulnasoft.F("build"),
					RootDir:           khulnasoft.F("/"),
					WebAnalyticsTag:   khulnasoft.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
					WebAnalyticsToken: khulnasoft.F("021e1057c18547eca7b79f2516f06o7x"),
				}),
				DeploymentConfigs: khulnasoft.F(pages.ProjectDeploymentConfigsParam{
					Preview: khulnasoft.F(pages.ProjectDeploymentConfigsPreviewParam{
						AIBindings: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewAIBindingParam{
							"AI_BINDING": {
								ProjectID: khulnasoft.F("some-project-id"),
							},
						}),
						AnalyticsEngineDatasets: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetParam{
							"ANALYTICS_ENGINE_BINDING": {
								Dataset: khulnasoft.F("api_analytics"),
							},
						}),
						Browsers: khulnasoft.F(map[string]interface{}{
							"BROWSER": "bar",
						}),
						CompatibilityDate:  khulnasoft.F("2022-01-01"),
						CompatibilityFlags: khulnasoft.F([]string{"url_standard"}),
						D1Databases: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewD1DatabaseParam{
							"D1_BINDING": {
								ID: khulnasoft.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
							},
						}),
						DurableObjectNamespaces: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewDurableObjectNamespaceParam{
							"DO_BINDING": {
								NamespaceID: khulnasoft.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
							},
						}),
						EnvVars: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewEnvVarParam{
							"foo": {
								Value: khulnasoft.F("hello world"),
								Type:  khulnasoft.F(pages.ProjectDeploymentConfigsPreviewEnvVarsTypePlainText),
							},
						}),
						HyperdriveBindings: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewHyperdriveBindingParam{
							"HYPERDRIVE": {
								ID: khulnasoft.F("a76a99bc342644deb02c38d66082262a"),
							},
						}),
						KVNamespaces: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewKVNamespaceParam{
							"KV_BINDING": {
								NamespaceID: khulnasoft.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
							},
						}),
						MTLSCertificates: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewMTLSCertificateParam{
							"MTLS": {
								CertificateID: khulnasoft.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
							},
						}),
						Placement: khulnasoft.F(pages.ProjectDeploymentConfigsPreviewPlacementParam{
							Mode: khulnasoft.F("smart"),
						}),
						QueueProducers: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewQueueProducerParam{
							"QUEUE_PRODUCER_BINDING": {
								Name: khulnasoft.F("some-queue"),
							},
						}),
						R2Buckets: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewR2BucketParam{
							"R2_BINDING": {
								Jurisdiction: khulnasoft.F("eu"),
								Name:         khulnasoft.F("some-bucket"),
							},
						}),
						Services: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewServiceParam{
							"SERVICE_BINDING": {
								Entrypoint:  khulnasoft.F("MyHandler"),
								Environment: khulnasoft.F("production"),
								Service:     khulnasoft.F("example-worker"),
							},
						}),
						VectorizeBindings: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsPreviewVectorizeBindingParam{
							"VECTORIZE": {
								IndexName: khulnasoft.F("my_index"),
							},
						}),
					}),
					Production: khulnasoft.F(pages.ProjectDeploymentConfigsProductionParam{
						AIBindings: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionAIBindingParam{
							"AI_BINDING": {
								ProjectID: khulnasoft.F("some-project-id"),
							},
						}),
						AnalyticsEngineDatasets: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionAnalyticsEngineDatasetParam{
							"ANALYTICS_ENGINE_BINDING": {
								Dataset: khulnasoft.F("api_analytics"),
							},
						}),
						Browsers: khulnasoft.F(map[string]interface{}{
							"BROWSER": "bar",
						}),
						CompatibilityDate:  khulnasoft.F("2022-01-01"),
						CompatibilityFlags: khulnasoft.F([]string{"url_standard"}),
						D1Databases: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionD1DatabaseParam{
							"D1_BINDING": {
								ID: khulnasoft.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
							},
						}),
						DurableObjectNamespaces: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionDurableObjectNamespaceParam{
							"DO_BINDING": {
								NamespaceID: khulnasoft.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
							},
						}),
						EnvVars: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionEnvVarParam{
							"BUILD_VERSION": {
								Value: khulnasoft.F("3.3"),
								Type:  khulnasoft.F(pages.ProjectDeploymentConfigsProductionEnvVarsTypePlainText),
							},
							"delete_this_env_var": {
								Value: khulnasoft.F("value"),
								Type:  khulnasoft.F(pages.ProjectDeploymentConfigsProductionEnvVarsTypePlainText),
							},
							"secret_var": {
								Value: khulnasoft.F("A_CMS_API_TOKEN"),
								Type:  khulnasoft.F(pages.ProjectDeploymentConfigsProductionEnvVarsTypePlainText),
							},
						}),
						HyperdriveBindings: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionHyperdriveBindingParam{
							"HYPERDRIVE": {
								ID: khulnasoft.F("a76a99bc342644deb02c38d66082262a"),
							},
						}),
						KVNamespaces: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionKVNamespaceParam{
							"KV_BINDING": {
								NamespaceID: khulnasoft.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
							},
						}),
						MTLSCertificates: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionMTLSCertificateParam{
							"MTLS": {
								CertificateID: khulnasoft.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
							},
						}),
						Placement: khulnasoft.F(pages.ProjectDeploymentConfigsProductionPlacementParam{
							Mode: khulnasoft.F("smart"),
						}),
						QueueProducers: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionQueueProducerParam{
							"QUEUE_PRODUCER_BINDING": {
								Name: khulnasoft.F("some-queue"),
							},
						}),
						R2Buckets: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionR2BucketParam{
							"R2_BINDING": {
								Jurisdiction: khulnasoft.F("eu"),
								Name:         khulnasoft.F("some-bucket"),
							},
						}),
						Services: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionServiceParam{
							"SERVICE_BINDING": {
								Entrypoint:  khulnasoft.F("MyHandler"),
								Environment: khulnasoft.F("production"),
								Service:     khulnasoft.F("example-worker"),
							},
						}),
						VectorizeBindings: khulnasoft.F(map[string]pages.ProjectDeploymentConfigsProductionVectorizeBindingParam{
							"VECTORIZE": {
								IndexName: khulnasoft.F("my_index"),
							},
						}),
					}),
				}),
				Name:             khulnasoft.F("NextJS Blog"),
				ProductionBranch: khulnasoft.F("main"),
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

func TestProjectGet(t *testing.T) {
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
	_, err := client.Pages.Projects.Get(
		context.TODO(),
		"this-is-my-project-01",
		pages.ProjectGetParams{
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

func TestProjectPurgeBuildCache(t *testing.T) {
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
	_, err := client.Pages.Projects.PurgeBuildCache(
		context.TODO(),
		"this-is-my-project-01",
		pages.ProjectPurgeBuildCacheParams{
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
