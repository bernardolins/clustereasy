package templates

func ScopeTemplateContent() string {
	return `#cloud-config

{{ .Name}}:
  {{ range .Services }}{{ .GetName }}:
    {{ range $key, $value := .GetParameters }}{{ $key }}: {{ $value }}
    {{ end }}
  {{ end }}
`
}
