// SPDX-FileCopyrightText: Pachyderm, Inc. <info@pachyderm.com>
// SPDX-License-Identifier: Apache-2.0

package helmtest

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestMicrosoft(t *testing.T) {
	var (
		storageBackend = "MICROSOFT"
		container      = "foo-container"
		id             = "ms-id"
		secret         = "ms-secret"
		objects, err   = manifestToObjects(helm.RenderTemplate(t,
			&helm.Options{
				SetStrValues: map[string]string{
					"pachd.storage.backend":             storageBackend,
					"pachd.storage.microsoft.container": container,
					"pachd.storage.microsoft.id":        id,
					"pachd.storage.microsoft.secret":    secret,
				}},
			"../pachyderm/", "release-name", nil))
	)
	if err != nil {
		t.Fatalf("could not render templates to objects: %v", err)
	}

	for _, object := range objects {

		switch object := object.(type) {
		case *v1.Secret:
			// TODO Doesn't check if secret not found
			if object.Name != "pachyderm-storage-secret" {
				continue
			}
			testCases := map[string]string{
				"microsoft-container": container,
				"microsoft-id":        id,
				"microsoft-secret":    "blah",
			}

			for k, v := range testCases {
				t.Run(fmt.Sprintf("%s equals %s", k, v), func(t *testing.T) {
					if got := string(object.Data[k]); got != v {
						t.Errorf("got %s; want %s", got, v)
					}
				})
			}
		case *appsV1.Deployment:
			if object.Name != "pachd" {
				continue
			}
			t.Run("Deployment Env Vars", func(t *testing.T) {
				err, c := GetContainerByName(object.Spec.Template.Spec.Containers, "pachd")
				if err != nil {
					t.Error(err)
				}
				//TODO STORAGE_BACKEND check
				testCases := map[string]string{
					"MICROSOFT_CONTAINER": container,
					"MICROSOFT_ID":        id,
					"MICROSOFT_SECRET":    secret,
				}

				for k, v := range testCases {
					t.Run(fmt.Sprintf("%s equals %s", k, v), func(t *testing.T) {
						if got := e.ValueFrom.SecretKeyRef.Key; got != v {
							t.Errorf("got %s; want %s", got, v)
						}
					})
				}
				for _, e := range c.Env {
					switch e.Name {
					case "STORAGE_BACKEND":
						if e.Value != storageBackend {
							t.Errorf("expected STORAGE_BACKEND to be %q, not %q", storageBackend, e.Value)
						}
					case "MICROSOFT_CONTAINER":
						if e.ValueFrom.SecretKeyRef.Key != container {
							t.Errorf("expected MICROSOFT_CONTAINER to be %q, not %q", container, e.Value)
						}
					case "MICROSOFT_ID":
						//TODO (envFrom)
					case "MICROSOFT_SECRET":
						//TODO (envFrom)
					}
				}
			})
		case *appsV1.StatefulSet:
			if object.Name != "etcd" {
				continue
			}
			if *object.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests.Storage() != resource.MustParse("256Gi") {
				t.Errorf("expected storage size to be %q, not %q", "256Gi", object.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests.Storage())
			}

		}
	}
}
