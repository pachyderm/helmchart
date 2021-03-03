# Pachyderm Helm Chart

The repo contains the pachyderm helm chart.

Status: **Experimental**

## Folder Structure

```
pachyderm - the helm chart itself
test - Go based tests for the helm chart
```

## Developer Guide
To run the tests for the helm chart, run the following:

```
make test
```

# Diff against `pachctl`

The scheme here is to take a recent `pachctl` and diff its manifest &
the Helm-generated manifest.

 1. Generate a `pachctl` manifest with `pachctl deploy google blah 10
    --dynamic-etcd-nodes 1 -o yaml --dry-run > pachmanifest.yaml`
 2. Generate a Helm manifest with `helm template -f
    examples/gcp-values.yaml ./pachyderm --generate-name >
    helmmanifest.yaml`
 3. Visually diff the two.  I did this with the currently checked-in
    version, with my comments in `pachctlmanifest.yaml`.

There are a bunch of fiddly little bits to do, but the major portions
appear to be there.

N.b.: neither of the manifests is checked in, as they are temporary
files.

# Validate Helm manifest
 1. `go install github.com/instrumenta/kubeval`
 2. `kubeval helmmanifest.yaml`
