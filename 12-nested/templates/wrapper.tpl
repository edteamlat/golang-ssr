{{ define "wrapper" }}
Este es un encabezado
{{- template "body" . -}}
Este sería el pie de página
{{ end }}