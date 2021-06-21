// SPDX-FileCopyrightText: Pachyderm, Inc. <info@pachyderm.com>
// SPDX-License-Identifier: Apache-2.0

package helmtest

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	v1 "k8s.io/api/core/v1"
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
		//TODO Fill in rest of values
	}

	helmValues := map[string]string{
		"pachd.storage.backend": "AMAZON",
	}
	for _, tc := range testCases {
		helmValues[tc.helmKey] = tc.value
	}

	render := helm.RenderTemplate(t,
		&helm.Options{
			SetStrValues: helmValues,
		}, "../pachyderm", "release-name", []string{"templates/pachd/storage-secret.yaml"})

	var secret *v1.Secret

	helm.UnmarshalK8SYaml(t, render, &secret)

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s equals %s", tc.envVar, tc.value), func(t *testing.T) {
			if got := string(secret.Data[tc.envVar]); got != tc.value {
				t.Errorf("got %s; want %s", got, tc.value)
			}
		})
	}
}
