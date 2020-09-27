package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl07 := `
Condicionales
{{- if ge . 18 }}
Bienvenido al Bar!!! Disfruta!!!
{{ else if ge . 12 }}
Aún eres muy joven, espera un poco
{{- else }}
Lo lamento, aún no puedes ingresar :(
{{ end }}
`
	t, err := template.New("ejercicio07").Parse(tpl07)
	if err != nil {
		log.Fatalf("Error al hacer parse del template %v", err)
	}

	err = t.Execute(os.Stdout, 8)
	if err != nil {
		log.Fatalf("Error al ejecutar el template %v", err)
	}
}
