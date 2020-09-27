package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl05 := `
Variables, así se crea una variable en un template
{{- $mynumber := 41 }}
El valor de la variable es:
{{ $mynumber }}
Reasignando el valor a la variable:
{{- $mynumber = "hola mundo" }}
El nuevo valor es:
{{ $mynumber }}
`
	t, err := template.New("ejercicio05").Parse(tpl05)
	if err != nil {
		log.Fatalf("Error al hacer parse del template %v", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("Error al ejecutar el template %v", err)
	}
}
