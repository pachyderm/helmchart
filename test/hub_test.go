package helmtest

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/gruntwork-io/terratest/modules/helm"
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/networking/v1beta1"
)

type pach struct {
	DashURL        string
	StorageSize    int
	ClusterIP      string
	PachVersion    string
	BucketName     string
	ServiceAccount string
}

type config struct {
	HubServerHostname string
}

type templateValues struct {
	Pach   pach
	Config config
}

func TestHub(t *testing.T) {
	var (
		p = pach{
			DashURL:        "http://foo.test/",
			StorageSize:    6,
			ClusterIP:      "::1",
			PachVersion:    "v1.12.6",
			BucketName:     "fake-bucket",
			ServiceAccount: "test-service-account",
		}
		c              = config{"hub-server.test"}
		valuesTemplate = template.Must(template.New("hub-values").Funcs(sprig.TxtFuncMap()).ParseFiles("../examples/hub-values.yaml"))
		objects        []interface{}
		checks         = map[string]bool{
			"ingress":            false,
			"metricsEndpoint":    false,
			"dash limits":        false,
			"etcd limits":        false,
			"loki logging":       false,
			"docker socket":      false,
			"pachd service type": false,
		}
		f, err = ioutil.TempFile("", "values.yaml")
	)
	if err != nil {
		t.Fatalf("couldn’t open temporary values file: %v", err)
	}
	defer os.Remove(f.Name())

	if err = valuesTemplate.Lookup("hub-values.yaml").Execute(f, templateValues{p, c}); err != nil {
		t.Fatalf("couldn’t execute template: %v", err)
	}
	f.Close()

	if objects, err = manifestToObjects(helm.RenderTemplate(t,
		&helm.Options{
			ValuesFiles: []string{f.Name()},
		},
		"../pachyderm/", "pachd", nil)); err != nil {
		t.Fatalf("could not render templates to objects: %v", err)
	}
	for _, object := range objects {
		switch object := object.(type) {
		case *v1beta1.Ingress:
			for _, rule := range object.Spec.Rules {
				if rule.Host == p.DashURL {
					checks["ingress"] = true
				}
			}
		case *appsV1.Deployment:
			switch object.Name {
			case "pachd":
				for _, cc := range object.Spec.Template.Spec.Containers {
					if cc.Name != "pachd" {
						continue
					}
					for _, v := range cc.Env {
						switch v.Name {
						case "METRICS_ENDPOINT":
							expected := fmt.Sprintf("https://%s/api/v1/metrics", c.HubServerHostname)
							if v.Value != expected {
								t.Errorf("metrics endpoint %q ≠ %q", v.Value, expected)
							}
							checks["metricsEndpoint"] = true
						case "LOKI_LOGGING":
							if v.Value != "true" {
								t.Error("Loki logging should be enabled")
							}
							checks["loki logging"] = true
						case "NO_EXPOSE_DOCKER_SOCKET":
							if v.Value != "false" {
								t.Error("Docker socket should be exposed")
							}
							checks["docker socket"] = true
						}
					}
				}
			case "dash":
				for _, cc := range object.Spec.Template.Spec.Containers {
					if cc.Name != "dash" {
						continue
					}
					if len(cc.Resources.Limits) > 0 {
						t.Errorf("dash should have no resource limits")
					}
					checks["dash limits"] = true
				}
			}
		case *v1.Secret:
			if object.Name != "dash-tls" {
				continue
			}
			t.Errorf("there should be no dash-tls secret")
		case *v1.Service:
			if object.Name != "pachd" {
				continue
			}
			if object.Spec.Type != "ClusterIP" {
				t.Errorf("pachd service type should be \"ClusterIP\", not %q", object.Spec.Type)
			}
			checks["pachd service type"] = true
		case *appsV1.StatefulSet:
			if object.Name != "etcd" {
				continue
			}
			for _, cc := range object.Spec.Template.Spec.Containers {
				if cc.Name != "etcd" {
					continue
				}
				if len(cc.Resources.Limits) > 0 {
					t.Errorf("etcd should have no resource limits")
				}
				checks["etcd limits"] = true
			}
		}
	}

	for check := range checks {
		if !checks[check] {
			t.Errorf("%q incomplete", check)
		}
	}
}
