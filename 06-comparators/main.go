package main

import (
	"log"
	"os"
	"text/template"
)

type numbers struct {
	X int
	Y int
}

func main() {
	tpl06 := `
Comparadores
eq: Es igual a. {{ .X }} es igual a {{ .Y }}? = {{ eq .X .Y }}
ne: {{.X}} es diferente a "!=" {{.Y}}? {{ne .X .Y}}
lt: {{.X}} es menor a "<" {{.Y}}? {{lt .X .Y}}
le: {{.X}} es menor o igual a "<=" {{.Y}}? {{le .X .Y}}
gt: {{.X}} es mayor a ">" {{.Y}}? {{gt .X .Y}}
ge: {{.X}} es mayor o igual a ">=" {{.Y}}? {{ge .X .Y}}
`
	t, err := template.New("ejercicio06").Parse(tpl06)
	if err != nil {
		log.Fatalf("Error al hacer parse del template %v", err)
	}

	n := numbers{5, 6}
	err = t.Execute(os.Stdout, &n)
	if err != nil {
		log.Fatalf("Error al ejecutar el template %v", err)
	}
}
