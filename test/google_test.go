// SPDX-FileCopyrightText: Pachyderm, Inc. <info@pachyderm.com>
// SPDX-License-Identifier: Apache-2.0

package helmtest

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	storageV1 "k8s.io/api/storage/v1"
)

//Service Account Test
//manifestServiceAccount := serviceAccount.Annotations["iam.gke.io/gcp-service-account"]
//if manifestServiceAccount != expectedServiceAccount {
//	t.Fatalf("Google service account expected (%s) actual (%s) ", expectedServiceAccount, manifestServiceAccount)
//}

//Worker Service Account Test (Same as Service Account test)

//Storage Secret Test

//Pachd Deployment - Storage backend

//Etcd / Pachd Storage Class - Should test  storage class name elsewhere
//Service Account name - Should test  service account name elsewhere

func TestGoogle(t *testing.T) {
	helmChartPath := "../pachyderm"

	type envVarMap struct {
		helmKey string
		envVar  string
		value   string
	}
	testCases := []envVarMap{
		{
			helmKey: "pachd.storage.google.bucket",
			value:   "fake-bucket",
			envVar:  "GOOGLE_BUCKET",
		},
		{
			helmKey: "pachd.storage.google.cred",
			value:   `INSERT JSON HERE`,
			envVar:  "GOOGLE_CRED",
		},
	}
	var (
		expectedServiceAccount = "my-fine-sa"
		expectedProvisioner    = "kubernetes.io/gce-pd"
		storageBackendEnvVar   = "STORAGE_BACKEND"
		expectedStorageBackend = "GOOGLE"
	)
	helmValues := map[string]string{
		"pachd.storage.backend":                   expectedStorageBackend,
		"pachd.storage.google.serviceAccountName": expectedServiceAccount,
	}
	for _, tc := range testCases {
		helmValues[tc.helmKey] = tc.value
	}

	options := &helm.Options{
		SetValues: helmValues,
	}

	templatesToRender := []string{
		"templates/pachd/storage-secret.yaml",
		"templates/pachd/deployment.yaml",
		"templates/pachd/rbac/serviceaccount.yaml",
		"templates/pachd/rbac/worker-serviceaccount.yaml",
		"templates/etcd/statefulset.yaml",
		"templates/etcd/storageclass-gcp.yaml",
		"templates/postgresql/statefulset.yaml",
		"templates/postgresql/storageclass-gcp.yaml",
	}
	output := helm.RenderTemplate(t, options, helmChartPath, "blah", templatesToRender)

	objects, err := manifestToObjects(output)
	if err != nil {
		t.Fatal(err)
	}

	/*checked := map[string]bool{
		"pachyderm-storage-secret": false,
		"pachyderm-worker":         false,
		"pachyderm":                false,
		"pachd":                    false,
		"postgresql-storage-class": false,
		"etcd-storage-class":       false,
	}*/

	//NOTE: If adding a new check to this for loop, be sure to add to it the checked map above to ensure it's found
	for _, object := range objects {
		switch resource := object.(type) {
		case *v1.Secret:
			if resource.Name != "pachyderm-storage-secret" {
				continue
			}
			//TODO checks["secret"] = true
			for _, tc := range testCases {
				t.Run(fmt.Sprintf("%s equals %s", tc.envVar, tc.value), func(t *testing.T) {
					if got := string(resource.Data[tc.envVar]); got != tc.value {
						t.Errorf("got %s; want %s", got, tc.value)
					}
				})
			}

		case *v1.ServiceAccount:
			if resource.Name == "pachyderm-worker" || resource.Name == "pachyderm" {

				t.Run(fmt.Sprintf("%s service account annotation equals %s", resource.Name, expectedServiceAccount), func(t *testing.T) {
					if sa := resource.Annotations["iam.gke.io/gcp-service-account"]; sa != expectedServiceAccount {
						t.Errorf("expected service account to be %q but was %q", expectedServiceAccount, sa)
					}
				})
				//TODO checks["service account"] = true
			}
		case *appsV1.Deployment:
			if resource.Name != "pachd" {
				continue
			}
			//checks["storage class"] = true TODO
			t.Run("pachd deployment env vars", func(t *testing.T) {
				c := GetContainerByName("pachd", resource.Spec.Template.Spec.Containers)
				if c == nil {
					t.Errorf("pachd container not found in pachd deployment")
				}
				if _, got := GetEnvVarByName(c.Env, storageBackendEnvVar); got != "GOOGLE" {
					t.Errorf("expected %s to be %q, not %q", storageBackendEnvVar, expectedStorageBackend, got)
				}

			})
		case *storageV1.StorageClass:
			if resource.Name == "postgresql-storage-class" || resource.Name == "etcd-storage-class" {

				//checks["storage class"] = true TODO
				t.Run(fmt.Sprintf("%s storage class annotation equals %s", resource.Name, expectedProvisioner), func(t *testing.T) {
					if resource.Provisioner != expectedProvisioner {
						t.Errorf("expected storageclass provisioner to be %q but it was %q", expectedProvisioner, resource.Provisioner)
					}
				})
				//TODO Check default storage size for google
			}

		}
	}
}
