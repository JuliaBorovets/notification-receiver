{{- define "notification-receiver.selectorLabels" -}}
app.kubernetes.io/name: 'notification-receiver'
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}