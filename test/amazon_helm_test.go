package helmtest

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	logger "github.com/gruntwork-io/terratest/modules/logger"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

func init() {
	logger.Default = logger.Discard
}

/*
	amazonRegion:
	amazonBucket:
	amazonID:
	amazonSecret:
	amazonToken:
	amazonDistribution:
	customEndpoint:
	retries:
	timeout:
	uploadACL:
	reverse:
	partSize:
	maxUploadParts:
	disableSSL:
	noVerifySSL:
	#TODO Vault env vars
	#TODO iamRole - Check all places IAM role rendered
*/

type TestValues struct {
	ValuePath string //pachd.storage.amazon.region
	ValueKey  string //amazonRegion
	Value     string //us-east-1a
}

func ValuesToHelmValues(t []TestValues) map[string]string {
	helmValues := make(map[string]string)
	for _, tv := range t {
		helmValues[tv.ValuePath] = tv.Value
	}
	return helmValues
}

func TestAmazonStorageSecrets(t *testing.T) {
	helmChartPath := "../pachyderm"

	//expectedBucket := "mybucket123"

	flagtests := []TestValues{
		{
			ValuePath: "pachd.storage.amazon.bucket",
			ValueKey:  "amazon-bucket",
			Value:     "blahblah12",
		}, {
			ValuePath: "pachd.storage.amazon.id",
			ValueKey:  "amazon-id",
			Value:     "testkey",
		},
	}

	helmValues := ValuesToHelmValues(flagtests)

	//Set additional values
	helmValues["pachd.storage.backend"] = "AMAZON"

	options := &helm.Options{
		SetValues: helmValues,
	}

	templates := []string{
		"templates/pachd/storage-secret.yaml",
		//"templates/pachd/deployment.yaml",
	}
	output := helm.RenderTemplate(t, options, helmChartPath, "blah", templates)

	var secret v1.Secret
	helm.UnmarshalK8SYaml(t, output, &secret)

	for _, tt := range flagtests {
		t.Run(tt.ValueKey, func(t *testing.T) {
			v := string(secret.Data[tt.ValueKey])
			if v != tt.Value {
				t.Errorf("got %q, want %q", v, tt.Value)
			}
		})
	}
	//if manifestServiceAccount != expectedSeexpectedBucketrviceAccount {
	//	t.Fatalf("Amazon bucket expected (%s) actual (%s) ", expectedBucket, secret.)
	//}

}

type DeploymentTestValues struct {
	ValuePath string
	Value     string
	ValueKey  string
	EnvVar    string
}

func CheckSecretEnvVar(envVarSource *v1.EnvVarSource) {

}

func DeploymentValuesToHelmValues(t []DeploymentTestValues) map[string]string {
	helmValues := make(map[string]string)
	for _, tv := range t {
		helmValues[tv.ValuePath] = tv.Value
	}
	return helmValues
}

func TestAmazonStoragePachDeploymentEnvVars(t *testing.T) {
	helmChartPath := "../pachyderm"

	//expectedBucket := "mybucket123"

	flagtests := []DeploymentTestValues{
		{
			ValuePath: "pachd.storage.amazon.bucket",
			ValueKey:  "amazon-bucket",
			Value:     "blahblah12",
			EnvVar:    "AMAZON_BUCKET",
		}, {
			ValuePath: "pachd.storage.amazon.id",
			ValueKey:  "amazon-id",
			Value:     "testkey",
			EnvVar:    "AMAZON_ID",
		},
	}

	helmValues := DeploymentValuesToHelmValues(flagtests)

	//Set additional values
	helmValues["pachd.storage.backend"] = "AMAZON"

	options := &helm.Options{
		SetValues: helmValues,
	}

	templates := []string{
		//"templates/pachd/storage-secret.yaml",
		"templates/pachd/deployment.yaml",
	}
	output := helm.RenderTemplate(t, options, helmChartPath, "blah", templates)

	var deployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, output, &deployment)

	container := deployment.Spec.Template.Spec.Containers[0]

	secretName := "pachyderm-storage-secret" //TODO check if it matches the name of the secret

	for _, tt := range flagtests {
		t.Run(tt.ValueKey, func(t *testing.T) {
			err, v := GetSecretEnvVarByName(container.Env, tt.EnvVar)
			if err != nil {
				t.Errorf(err.Error())
			}
			secretKeyRef := v.SecretKeyRef
			if secretKeyRef == nil {
				t.Errorf("No secret key ref found")
			}

			//fmt.Printf("%#v\n\n", secretKeyRef.Name)
			fmt.Printf("%T\n\n", secretKeyRef.Name)

			if secretKeyRef.Key != tt.ValueKey {
				t.Errorf("got %q, want %q", v, tt.Value)
			}
			if secretKeyRef.Name != secretName {
				t.Errorf("got %q, want %q", secretKeyRef.Name, secretName)
			}
			boolRef := true
			if *secretKeyRef.Optional != boolRef {
				t.Errorf("got %t, want %t", *secretKeyRef.Optional, boolRef)
			}

		})
	}
}

/*
   - name: AMAZON_REGION
     valueFrom:
       secretKeyRef:
         key: amazon-region
         name: pachyderm-storage-secret
         optional: true
*/

/*
Multi file handling
files, err := splitYAML(output)
if err != nil {
	t.Fatal(err)
}
decodedKubeObjects := []interface{}{}
decode := scheme.Codecs.UniversalDeserializer().Decode
for _, f := range files {
	obj, _, err := decode([]byte(f), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	decodedKubeObjects = append(decodedKubeObjects, obj)

}

for _, f := range decodedKubeObjects {
	fmt.Printf("%#v\n", f)
}


func splitYAML(manifest string) ([]string, error) {
	dec := goyaml.NewDecoder(bytes.NewReader([]byte(manifest)))
	var res []string
	for {
		var value interface{}
		if err := dec.Decode(&value); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		b, err := goyaml.Marshal(value)
		if err != nil {
			return nil, err
		}
		res = append(res, string(b))
	}
	return res, nil
}
*/
