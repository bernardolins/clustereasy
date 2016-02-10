package templates

func ScopeTemplateContent() string {
	return `#cloud-config

{{ .Name}}:
  {{ range .Services }}{{ .GetName }}:
    {{ range $key, $value := .GetParameters }}{{ $key }}: {{ $value }}
    {{ end }}
  {{ end }}units:
    {{ range .Units }}- name: {{ .GetName }}
      command: {{ .GetCommand }}
      {{ if .GetContent }}content: |{{ end }}
        {{ range $key, $line := .ContentLines }}{{ $line }}
        {{ end }}
    {{ end }}
`
}
