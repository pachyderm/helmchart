// SPDX-FileCopyrightText: Pachyderm, Inc. <info@pachyderm.com>
// SPDX-License-Identifier: Apache-2.0

package helmtest

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	v1 "k8s.io/api/core/v1"
	storageV1 "k8s.io/api/storage/v1"
)

func TestMicrosoft(t *testing.T) {
	helmChartPath := "../pachyderm"

	type envVarMap struct {
		helmKey string
		envVar  string
		value   string
	}
	testCases := []envVarMap{
		{
			helmKey: "pachd.storage.microsoft.container",
			value:   "fake-bucket",
			envVar:  "MICROSOFT_CONTAINER",
		},
		{
			helmKey: "pachd.storage.microsoft.id",
			value:   "a-fine-id",
			envVar:  "MICROSOFT_ID",
		},
		{
			helmKey: "pachd.storage.microsoft.secret",
			value:   "super-secret-something",
			envVar:  "MICROSOFT_SECRET",
		},
	}
	var (
		expectedProvisioner = "kubernetes.io/azure-disk"
		//storageBackendEnvVar   = "STORAGE_BACKEND"
		expectedStorageBackend = "MICROSOFT"
	)
	helmValues := map[string]string{
		"pachd.storage.backend": expectedStorageBackend,
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
		"templates/etcd/storageclass-azure.yaml",
		"templates/postgresql/statefulset.yaml",
		"templates/postgresql/storageclass-azure.yaml",
	}
	output := helm.RenderTemplate(t, options, helmChartPath, "blah", templatesToRender)

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
			//TODO checks["secret"] = true
			for _, tc := range testCases {
				t.Run(fmt.Sprintf("%s equals %s", tc.envVar, tc.value), func(t *testing.T) {
					if got := string(resource.Data[tc.envVar]); got != tc.value {
						t.Errorf("got %s; want %s", got, tc.value)
					}
				})
			}
		case *storageV1.StorageClass:
			if resource.Name == "postgresql-storage-class" || resource.Name == "etcd-storage-class" {

				//checks["storage class"] = true TODO
				t.Run(fmt.Sprintf("%s storage class annotation equals %s", resource.Name, expectedProvisioner), func(t *testing.T) {
					if resource.Provisioner != expectedProvisioner {
						t.Errorf("expected storageclass provisioner to be %q but it was %q", expectedProvisioner, resource.Provisioner)
					}
				})
				//TODO Check default storage size for microsoft
				//Check all params in storage class set for microsoft
			}

		} // TODO: Deployment check
	}
}

/*
func TestMicrosoft(t *testing.T) {
	var (
		container    = "foo-container"
		id           = "ms-id"
		secret       = "ms-secret"
		objects, err = manifestToObjects(helm.RenderTemplate(t,
			&helm.Options{
				SetStrValues: map[string]string{
					"pachd.storage.backend":             "MICROSOFT",
					"pachd.storage.microsoft.container": container,
					"pachd.storage.microsoft.id":        id,
					"pachd.storage.microsoft.secret":    secret,
				}},
			"../pachyderm/", "release-name", nil))
		checks = map[string]bool{
			"STORAGE_BACKEND":     false,
			"MICROSOFT_CONTAINER": false,
			"MICROSOFT_ID":        false,
			"MICROSOFT_SECRET":    false,
			"container":           false,
			"id":                  false,
			"secret":              false,
			"etcd-volume-size":    false,
		}
	)
	if err != nil {
		t.Fatalf("could not render templates to objects: %v", err)
	}
	for _, object := range objects {
		switch object := object.(type) {
		case *v1.Secret:
			if object.Name != "pachyderm-storage-secret" {
				continue
			}
			if s := string(object.Data["microsoft-container"]); s != container {
				t.Errorf("expected container to be %q but it is %q", container, s)
			}
			checks["container"] = true
			if s := string(object.Data["microsoft-id"]); s != id {
				t.Errorf("expected id to be %q but it is %q", id, s)
			}
			checks["id"] = true
			if s := string(object.Data["microsoft-secret"]); s != secret {
				t.Errorf("expected secret to be %q but it is %q", secret, s)
			}
			checks["secret"] = true
		case *appsV1.Deployment:
			if object.Name != "pachd" {
				continue
			}
			for _, c := range object.Spec.Template.Spec.Containers {
				if c.Name != "pachd" {
					continue
				}
				for _, e := range c.Env {
					switch e.Name {
					case "STORAGE_BACKEND":
						if e.Value != "MICROSOFT" {
							t.Errorf("expected STORAGE_BACKEND to be %q, not %q", "GOOGLE", e.Value)
						}
						checks["STORAGE_BACKEND"] = true
					case "MICROSOFT_CONTAINER":
						checks["MICROSOFT_CONTAINER"] = true
					case "MICROSOFT_ID":
						checks["MICROSOFT_ID"] = true
					case "MICROSOFT_SECRET":
						checks["MICROSOFT_SECRET"] = true
					}
				}
			}
		case *appsV1.StatefulSet:
			if object.Name != "etcd" {
				continue
			}
			if *object.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests.Storage() != resource.MustParse("256Gi") {
				t.Errorf("expected storage size to be %q, not %q", "256Gi", object.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests.Storage())
			}
			checks["etcd-volume-size"] = true
		}
	}
	for check := range checks {
		if !checks[check] {
			t.Errorf("%s unchecked", check)
		}
	}
}
*/
