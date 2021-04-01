package helmtest

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/gruntwork-io/terratest/modules/helm"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
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
		f, err         = ioutil.TempFile("", "values.yaml")
		valuesTemplate = template.Must(template.New("hub-values").Funcs(sprig.TxtFuncMap()).ParseFiles("../examples/hub-values.yaml"))
		objects        []interface{}
		checks         = map[string]bool{}
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
		case *networkingv1beta1.Ingress:
			log.Println(object, checks)
		default:
		}
	}
}
