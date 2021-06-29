{{- /*
SPDX-FileCopyrightText: Pachyderm, Inc. <info@pachyderm.com>
SPDX-License-Identifier: Apache-2.0
*/ -}}
{{- /* vim: set filetype=mustache: */ -}}
{{- define "imagePullSecret" }}
{{- with .Values.imageCredentials }}
{{- printf "{\"auths\":{\"%s\":{\"username\":\"%s\",\"password\":\"%s\",\"email\":\"%s\",\"auth\":\"%s\"}}}" .registry .username .password .email (printf "%s:%s" .username .password | b64enc) | b64enc }}
{{ end -}}
{{ end -}}

{{- define "pachyderm.storageBackend" -}}
{{- if .Values.pachd.storage.backend -}}
{{ .Values.pachd.storage.backend }}
{{- else if eq .Values.deployTarget "AMAZON" -}}
AMAZON
{{- else if eq .Values.deployTarget "GOOGLE" -}}
GOOGLE
{{- else if eq .Values.deployTarget "MICROSOFT" -}}
MICROSOFT
{{- else if eq .Values.deployTarget "LOCAL" -}}
LOCAL
{{- else -}}
{{ fail "pachd.storage.backend required when no matching deploy target found" }}
{{- end -}}
{{- end -}}

{{- define "pachyderm.etcd.storageParameters" -}}
{{- if .Values.etcd.storageParameters -}}
{{- range $k, $v := .Values.etcd.storageParameters }}
{{ $k }}:
  {{ $v }}
{{- end -}}
{{- else if eq (include "pachyderm.storageBackend" .) "AMAZON" -}}
type: gp3
{{- else if eq (include "pachyderm.storageBackend" .) "GOOGLE" -}}
type: pd-ssd
{{- else if eq (include "pachyderm.storageBackend" .) "MICROSOFT" -}}
storageaccounttype: Premium_LRS
kind: Managed
{{- end -}}
{{- end -}}

{{- define "pachyderm.etcd.storageProvisioner" -}}
{{- if .Values.etcd.storageProvisioner -}}
{{ .Values.etc.storageProvisioner }}
{{- else if eq (include "pachyderm.storageBackend" .) "AMAZON" -}}
kubernetes.io/aws-ebs
{{- else if eq (include "pachyderm.storageBackend" .) "GOOGLE" -}}
kubernetes.io/gce-pd
{{- else if eq (include "pachyderm.storageBackend" .) "MICROSOFT" -}}
kubernetes.io/azure-disk
{{- end -}}
{{- end -}}

{{- define "pachyderm.postgresql.storageParameters" -}}
{{- if .Values.postgresql.storageParameters -}}
{{- range $k, $v := .Values.postgresql.storageParameters }}
{{ $k }}:
  {{ $v }}
{{- end -}}
{{- else if eq (include "pachyderm.storageBackend" .) "AMAZON" -}}
type: gp3
{{- else if eq (include "pachyderm.storageBackend" .) "GOOGLE" -}}
type: pd-ssd
{{- else if eq (include "pachyderm.storageBackend" .) "MICROSOFT" -}}
storageaccounttype: Premium_LRS
kind: Managed
{{- end -}}
{{- end -}}

{{- define "pachyderm.postgresql.storageProvisioner" -}}
{{- if .Values.postgresql.storageProvisioner -}}
{{ .Values.etc.storageProvisioner }}
{{- else if eq (include "pachyderm.storageBackend" .) "AMAZON" -}}
kubernetes.io/aws-ebs
{{- else if eq (include "pachyderm.storageBackend" .) "GOOGLE" -}}
kubernetes.io/gce-pd
{{- else if eq (include "pachyderm.storageBackend" .) "MICROSOFT" -}}
kubernetes.io/azure-disk
{{- end -}}
{{- end -}}
