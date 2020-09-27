Hola {{ .Name }}, espero que estés bien.

{{ .Name }}, recuerda que este sábado tenemos una fiesta de despedida.

En vista que tu tienes {{ .Age }} años,
{{- if lt .Age 18 }} no puedes ingresar, pero te enviaremos un detalle que esperemos disfrutes mucho.
{{- else }} te esperamos a las 8pm en el salón comunal del barrio.
{{ end -}}

Nos vemos pronto!
