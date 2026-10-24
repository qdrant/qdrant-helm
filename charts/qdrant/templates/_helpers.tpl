{{/*
Expand the name of the chart.
*/}}
{{- define "qdrant.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "qdrant.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "qdrant.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "qdrant.labels" -}}
helm.sh/chart: {{ include "qdrant.chart" . }}
{{ include "qdrant.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "qdrant.selectorLabels" -}}
app: {{ include "qdrant.name" . }}
app.kubernetes.io/name: {{ include "qdrant.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Additional (global) labels
*/}}
{{- define "qdrant.additionalLabels" -}}
{{- with .Values.additionalLabels }}
{{- range $key, $value := . }}
{{ $key }}: {{ $value | quote }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "qdrant.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "qdrant.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Api key given as a plain value, as a string.
Numbers are converted: --set turns 12345 into an integer, and a values file or --reuse-values into a float.
Whole numbers are written without decimals; anything else can't be converted back to the original key.
Note: an unquoted number in a values file is already changed by YAML (e.g. 0123 becomes 83), it has to be quoted.
*/}}
{{- define "qdrant.apiKeyString" -}}
{{- if kindIs "string" .value -}}
{{- .value -}}
{{- else if kindIs "int64" .value -}}
{{- .value | toString -}}
{{- else if and (kindIs "float64" .value) (eq .value (floor .value)) (lt .value 9007199254740992.0) (gt .value -9007199254740992.0) -}}
{{- .value | int64 | toString -}}
{{- else -}}
{{- fail (printf "%s must be a string, quote it in your values file: %s: \"<key>\"" .name .name) -}}
{{- end -}}
{{- end -}}

{{/*
Create secret
*/}}
{{- define "qdrant.secret" -}}
{{- $readOnlyApiKey := false }}
{{- $apiKey := false }}
{{- if kindIs "map" .Values.apiKey -}}
{{- if .Values.apiKey.valueFrom -}}
{{- /* Retrieve the value from the secret as specified in valueFrom */ -}}
{{- $secretName := .Values.apiKey.valueFrom.secretKeyRef.name -}}
{{- $secretKey := .Values.apiKey.valueFrom.secretKeyRef.key -}}
{{- $secretObj := (lookup "v1" "Secret" .Release.Namespace $secretName) | default dict -}}
{{- $secretData := (get $secretObj "data") | default dict -}}
{{- $apiKey = (get $secretData $secretKey | b64dec) -}}
{{- end -}}
{{- else if .Values.apiKey | toJson | eq "true" -}}
{{- /* Retrieve existing randomly generated api key or create a new one */ -}}
{{- $secretObj := (lookup "v1" "Secret" .Release.Namespace (printf "%s-apikey" (include "qdrant.fullname" . ))) | default dict -}}
{{- $secretData := (get $secretObj "data") | default dict -}}
{{- $apiKey = (get $secretData "api-key" | b64dec) | default (randAlphaNum 32) -}}
{{- else if .Values.apiKey -}}
{{- $apiKey = include "qdrant.apiKeyString" (dict "name" "apiKey" "value" .Values.apiKey) -}}
{{- end -}}
{{- if kindIs "map" .Values.readOnlyApiKey -}}
{{- if .Values.readOnlyApiKey.valueFrom -}}
{{- /* Retrieve the value from the secret as specified in valueFrom */ -}}
{{- $secretName := .Values.readOnlyApiKey.valueFrom.secretKeyRef.name -}}
{{- $secretKey := .Values.readOnlyApiKey.valueFrom.secretKeyRef.key -}}
{{- $secretObj := (lookup "v1" "Secret" .Release.Namespace $secretName) | default dict -}}
{{- $secretData := (get $secretObj "data") | default dict -}}
{{- $readOnlyApiKey = (get $secretData $secretKey | b64dec) -}}
{{- end -}}
{{- else if eq (.Values.readOnlyApiKey | toJson) "true" -}}
{{- /* retrieve existing randomly generated api key or create new one */ -}}
{{- $secretObj := (lookup "v1" "Secret" .Release.Namespace (printf "%s-apikey" (include "qdrant.fullname" . ))) | default dict -}}
{{- $secretData := (get $secretObj "data") | default dict -}}
{{- $readOnlyApiKey = (get $secretData "read-only-api-key" | b64dec) | default (randAlphaNum 32) -}}
{{- else if .Values.readOnlyApiKey -}}
{{- $readOnlyApiKey = include "qdrant.apiKeyString" (dict "name" "readOnlyApiKey" "value" .Values.readOnlyApiKey) -}}
{{- end -}}
{{- if and $apiKey $readOnlyApiKey -}}
api-key: {{ $apiKey | b64enc }}
read-only-api-key: {{ $readOnlyApiKey | b64enc }}
local.yaml: {{ printf "service:\n  api_key: %s\n  read_only_api_key: %s" $apiKey $readOnlyApiKey | b64enc }}
{{- else if $apiKey -}}
api-key: {{ $apiKey | b64enc }}
local.yaml: {{ printf "service:\n  api_key: %s" $apiKey | b64enc }}
{{- else if $readOnlyApiKey -}}
read-only-api-key: {{ $readOnlyApiKey | b64enc }}
local.yaml: {{ printf "service:\n  read_only_api_key: %s" $readOnlyApiKey | b64enc }}
{{- end -}}
{{- end -}}

{{/*
Protocol to use for inter cluster communication
*/}}
{{- define "qdrant.p2p.protocol" -}}
{{ if eq (.Values.config.cluster.p2p.enable_tls | toJson) "true" -}}
https
{{- else -}}
http
{{- end -}}
{{- end -}}

{{/*
Port to use for inter cluster communication
*/}}
{{- define "qdrant.p2p.port" -}}
{{- default 6335 .Values.config.cluster.p2p.port -}}
{{- end -}}
