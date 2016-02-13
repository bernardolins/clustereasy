package templates

func Header() string {
	return `
{{ .Content }}
`
}

func ScopeTemplateContent() string {
	//	return `#cloud-config
	//
	//{{ .Name}}:
	//  {{ range .Services }}{{ .GetName }}:
	//    {{ range $key, $value := .GetParameters }}{{ $key }}: {{ $value }}
	//    {{ end }}
	//  {{ end }}units:
	//    {{ range .Units }}- name: {{ .GetName }}
	//      command: {{ .GetCommand }}
	//      {{ if .GetContent }}content: |{{ end }}
	//        {{ range $key, $line := .ContentLines }}{{ $line }}
	//        {{ end }}
	//    {{ end }}
	//`

	return "{{ . }}"
}

func CoreOS() string {
	return `coreos:
  etcd2:
    {{ range $parameter, $value := .Etcd2.Parameters }}{{ $parameter }}: {{ $value }}
    {{end}}
  fleet:
    {{ range $parameter, $value := .Fleet.Parameters }}{{ $parameter }}: {{ $value }}
	  {{ end }}
  flannel:
    {{ range $parameter, $value := .Flannel.Parameters }}{{ $parameter }}: {{ $value }}
	  {{ end }}
  units:
    {{ range $name, $unit := .Units}}- {{ $name }}
		{{ $unit }}
      {{ range $param, $value := .Parameters }}{{ if $value }}{{ $param }}: {{ $value }}{{ end }}
    {{ end }}
      {{ if .Content }}content: |
      {{ range $key, $line := .ContentLines }}{{ $line }}{{ end }}
			{{ end }}
    {{ end }}
	`
}
