// SPDX-FileCopyrightText: Pachyderm, Inc. <info@pachyderm.com>
// SPDX-License-Identifier: Apache-2.0

package helmtest

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	storageV1 "k8s.io/api/storage/v1"
)

func TestGoogleServiceAccount(t *testing.T) {
	helmChartPath := "../pachyderm"

	expectedServiceAccount := "my-fine-sa"
	options := &helm.Options{
		SetValues: map[string]string{
			"pachd.image.tag":                         "1.12.3",
			"pachd.storage.backend":                   "GOOGLE",
			"pachd.storage.google.bucket":             "fake-bucket",
			"pachd.storage.google.serviceAccountName": expectedServiceAccount,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, "blah", []string{"templates/pachd/rbac/serviceaccount.yaml"})

	var serviceAccount v1.ServiceAccount

	helm.UnmarshalK8SYaml(t, output, &serviceAccount)

	manifestServiceAccount := serviceAccount.Annotations["iam.gke.io/gcp-service-account"]
	if manifestServiceAccount != expectedServiceAccount {
		t.Fatalf("Google service account expected (%s) actual (%s) ", expectedServiceAccount, manifestServiceAccount)
	}
}

func TestGoogleWorkerServiceAccount(t *testing.T) {
	helmChartPath := "../pachyderm"

	expectedServiceAccount := "my-fine-sa"
	options := &helm.Options{
		SetValues: map[string]string{
			"pachd.image.tag":                         "1.12.3",
			"pachd.storage.backend":                   "GOOGLE",
			"pachd.storage.google.bucket":             "fake-bucket",
			"pachd.storage.google.serviceAccountName": expectedServiceAccount,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, "blah", []string{"templates/pachd/rbac/worker-serviceaccount.yaml"})

	var serviceAccount v1.ServiceAccount

	helm.UnmarshalK8SYaml(t, output, &serviceAccount)

	manifestServiceAccount := serviceAccount.Annotations["iam.gke.io/gcp-service-account"]
	if manifestServiceAccount != expectedServiceAccount {
		t.Fatalf("Google service account expected (%s) actual (%s) ", expectedServiceAccount, manifestServiceAccount)
	}
}

func TestGoogleValues(t *testing.T) {
	type envVarMap struct {
		helmKey string
		envVar  string
		value   string
	}
	testCases := []envVarMap{
		{
			helmKey: "pachd.storage.google.bucket",
		},
		{
			helmKey: "pachd.storage.google.cred",
		}, {
			helmKey: "pachd.storage.google.serviceAccountName",
		},
	}
	var (
		bucket               = "fake-bucket"
		cred                 = `INSERT JSON HERE`
		pachdServiceAccount  = "128"
		serviceAccount       = "a-service-account"
		helmChartPath        = "../pachyderm"
		provisioner          = "kubernetes.io/gce-pd"
		expectedStorageClass = "etcd-storage-class"
		checks               = map[string]bool{
			"bucket":                false,
			"cred":                  false,
			"service account":       false,
			"STORAGE_BACKEND":       false,
			"GOOGLE_BUCKET":         false,
			"GOOGLE_CRED":           false,
			"volume claim template": false,
			"storage class":         false,
		}

		options = &helm.Options{
			SetStrValues: map[string]string{
				"pachd.serviceAccount.name":               pachdServiceAccount,
				"pachd.storage.backend":                   "GOOGLE",
				"pachd.storage.google.bucket":             bucket,
				"pachd.storage.google.cred":               cred,
				"pachd.storage.google.serviceAccountName": serviceAccount,
			},
		}
		output = helm.RenderTemplate(t, options, helmChartPath, "release-name", nil)
	)
	objects, err := manifestToObjects(output)
	if err != nil {
		t.Fatal(err)
	}

	for _, object := range objects {
		switch resource := object.(type) {
		case *v1.Secret:
			if resource.Name != "pachyderm-storage-secret" {
				continue
			}
			if b := string(resource.Data["google-bucket"]); b != bucket {
				t.Errorf("expected bucket to be %q but was %q", bucket, b)
			}
			checks["bucket"] = true
			if c := string(resource.Data["google-cred"]); c != cred {
				t.Errorf("expected cred to be %q but was %q", cred, c)
			}
			checks["cred"] = true
		case *v1.ServiceAccount:
			if resource.Name != pachdServiceAccount {
				continue
			}
			if sa := resource.Annotations["iam.gke.io/gcp-service-account"]; sa != serviceAccount {
				t.Errorf("expected service account to be %q but was %q", serviceAccount, sa)
			}
			checks["service account"] = true
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
						if e.Value != "GOOGLE" {
							t.Errorf("expected STORAGE_BACKEND to be %q, not %q", "GOOGLE", e.Value)
						}
						checks["STORAGE_BACKEND"] = true
					case "GOOGLE_BUCKET":
						checks["GOOGLE_BUCKET"] = true
					case "GOOGLE_CRED":
						checks["GOOGLE_CRED"] = true
					}
				}
			}
		case *appsV1.StatefulSet:
			if resource.Name != "etcd" {
				continue
			}
			for _, v := range resource.Spec.VolumeClaimTemplates {
				if *v.Spec.StorageClassName != expectedStorageClass {
					continue
				}
				checks["volume claim template"] = true
			}
		case *storageV1.StorageClass:
			if resource.Name != expectedStorageClass {
				continue
			}
			if resource.Provisioner != provisioner {
				t.Errorf("expected storageclass provisioner to be %q but it was %q", provisioner, resource.Provisioner)
			}
			checks["storage class"] = true
		}
	}
	for check := range checks {
		if !checks[check] {
			t.Errorf("check %q not performed", check)
		}
	}
}
