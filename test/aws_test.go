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

func TestAWS(t *testing.T) {

	type envVarMap struct {
		helmKey string
		envVar  string
		value   string
	}

	testCases := []envVarMap{
		{
			helmKey: "pachd.storage.amazon.bucket",
			envVar:  "AMAZON_BUCKET",
			value:   "my-bucket",
		},
		{
			helmKey: "pachd.storage.amazon.cloudFrontDistribution",
			envVar:  "AMAZON_DISTRIBUTION",
			value:   "test-cf",
		},
		{
			helmKey: "pachd.storage.amazon.id",
			envVar:  "AMAZON_ID",
			value:   "testazid123",
		},
		{
			helmKey: "pachd.storage.amazon.region",
			envVar:  "AMAZON_REGION",
			value:   "", //TODO
		},
		{
			helmKey: "pachd.storage.amazon.secret",
			envVar:  "AMAZON_SECRET",
			value:   "", //TODO
		},
		{
			helmKey: "pachd.storage.amazon.token",
			envVar:  "AMAZON_TOKEN",
			value:   "", //TODO
		},
		{
			helmKey: "pachd.storage.amazon.customEndpoint",
			envVar:  "CUSTOM_ENDPOINT",
			value:   "", //TODO
		},
		//TODO Fill in rest of values
		{
			helmKey: "pachd.storage.amazon.disableSSL",
			envVar:  "DISABLE_SSL",
			value:   "false",
		},
		{
			helmKey: "pachd.storage.amazon.logOptions",
			envVar:  "OBJ_LOG_OPTS",
			value:   "", //TODO
		},
		/*{
			helmKey: "pachd.storage.amazon.maxUploadParts",
			envVar:  "MAX_UPLOAD_PARTS",
			value:   "", //TODO
		},*/
		/*{
			helmKey: "pachd.storage.amazon.verifySSL",
			envVar:  "NO_VERIFY_SSL",
			value:   "", //TODO
		},*/
		{
			helmKey: "pachd.storage.amazon.partSize",
			envVar:  "PART_SIZE",
			value:   "", //TODO
		},
		/*{
			helmKey: "pachd.storage.amazon.retries",
			envVar:  "RETRIES",
			value:   "", //TODO
		},*/
		/*{
			helmKey: "pachd.storage.amazon.reverse",
			envVar:  "REVERSE",
			value:   "", //TODO
		},*/
		{
			helmKey: "pachd.storage.amazon.timeout",
			envVar:  "TIMEOUT",
			value:   "", //TODO
		},
		{
			helmKey: "pachd.storage.amazon.uploadACL",
			envVar:  "UPLOAD_ACL",
			value:   "", //TODO
		},
	}
	var (
		//expectedServiceAccount = "my-fine-sa"
		expectedProvisioner = "ebs.csi.aws.com"
		//storageBackendEnvVar   = "STORAGE_BACKEND"
		//expectedStorageBackend = "AMAZON"
	)

	helmValues := map[string]string{
		"pachd.storage.backend": "AMAZON",
	}
	for _, tc := range testCases {
		helmValues[tc.helmKey] = tc.value
	}

	templatesToRender := []string{
		"templates/pachd/storage-secret.yaml",
		"templates/pachd/deployment.yaml",
		"templates/pachd/rbac/serviceaccount.yaml",
		"templates/pachd/rbac/worker-serviceaccount.yaml",
		"templates/etcd/statefulset.yaml",
		"templates/etcd/storageclass-aws.yaml",
		"templates/postgresql/statefulset.yaml",
		"templates/postgresql/storageclass-aws.yaml",
	}

	output := helm.RenderTemplate(t, &helm.Options{
		SetValues: helmValues,
	}, "../pachyderm", "blah", templatesToRender)

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
				//TODO Check default storage size for amazon
			}

		}
	}
}
