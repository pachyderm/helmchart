# Pachyderm Helm Chart

## ** Please submit all PRs to the main pachyderm repo! **

The repo contains the pachyderm helm chart.

## Usage

Create a `values.yaml` file with your storage provider of choice, and other options, then run

```shell
$ helm repo add pach https://pachyderm.github.io/helmchart
$ helm install pachd pach/pachyderm -f values.yaml
```



<!-- SPDX-FileCopyrightText: Pachyderm, Inc. <info@pachyderm.com>
SPDX-License-Identifier: Apache-2.0 -->
