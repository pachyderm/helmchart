// SPDX-FileCopyrightText: Pachyderm, Inc. <info@pachyderm.com>
// SPDX-License-Identifier: Apache-2.0

package helmtest

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

func TestMinio(t *testing.T) {
	type envVarMap struct {
		helmKey string
		envVar  string
		value   string
	}

	testCases := []envVarMap{
		{
			helmKey: "pachd.storage.minio.bucket",
			envVar:  "MINIO_BUCKET",
			value:   "my-bucket",
		},
		{
			helmKey: "pachd.storage.minio.endpoint",
			envVar:  "MINIO_ENDPOINT",
			value:   "https://endpoint.test/",
		},
		{
			helmKey: "pachd.storage.minio.id",
			envVar:  "MINIO_ID",
			value:   "my-fine-id",
		},
		{
			helmKey: "pachd.storage.minio.secret",
			envVar:  "MINIO_SECRET",
			value:   "my-fine-secret",
		},
		{
			helmKey: "pachd.storage.minio.secure",
			envVar:  "MINIO_SECURE",
			value:   "", //TODO
		},
		{
			helmKey: "pachd.storage.minio.signature",
			envVar:  "MINIO_SIGNATURE",
			value:   "", //TODO
		},
	}

	helmValues := map[string]string{
		"pachd.storage.backend": "MINIO",
	}
	for _, tc := range testCases {
		helmValues[tc.helmKey] = tc.value
	}

	templatesToRender := []string{
		"templates/pachd/storage-secret.yaml",
		"templates/pachd/deployment.yaml",
	}

	objects, err := manifestToObjects(helm.RenderTemplate(t,
		&helm.Options{
			SetStrValues: helmValues,
		}, "../pachyderm/", "release-name", templatesToRender),
	)

	if err != nil {
		t.Fatalf("could not render templates to objects: %v", err)
	}
	for _, object := range objects {
		switch resource := object.(type) {
		case *v1.Secret:
			if resource.Name != "pachyderm-storage-secret" {
				continue
			}
			for _, tc := range testCases {
				t.Run(fmt.Sprintf("%s equals %s", tc.envVar, tc.value), func(t *testing.T) {
					if got := string(resource.Data[tc.envVar]); got != tc.value {
						t.Errorf("got %s; want %s", got, tc.value)
					}
				})
			}
		case *appsV1.Deployment:
			if resource.Name != "pachd" {
				continue
			}
			for _, c := range resource.Spec.Template.Spec.Containers {
				if c.Name != "pachd" {
					continue
				}
				for _, e := range c.Env {
					switch e.Name {
					case "STORAGE_BACKEND":
						if e.Value != "MINIO" {
							t.Errorf("expected STORAGE_BACKEND to be %q, not %q", "GOOGLE", e.Value)
						}
						//checks["STORAGE_BACKEND"] = true

					}
				}
			}
		}
	}
	/*for check := range checks {
		if !checks[check] {
			t.Errorf("%q incomplete", check)
		}
	}*/ //TODO
}
