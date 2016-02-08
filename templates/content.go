package templates

func ScopeTemplateContent() string {
	return `{{ .Name}}:
  {{ range .Services }}{{ .GetName }}:
    {{ range $key, $value := .GetParameters }}{{ $key }}: {{ $value }}
    {{ end }}{{ end }}
`
}
