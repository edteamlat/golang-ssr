package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
	password string
}

func (p *person) Saludar() string {
	return "Hola, soy " + p.Name
}

func (p *person) Adicion2(a int) int {
	return a + 2
}

func (p *person) Multiplicar(a, b int) int {
	return a * b
}

func main() {
	tpl03 := `Hola {{ .Name }}, su saludo es: {{ .Saludar }}
	Adicionando 2 a {{ .Age }}, da como resultado {{ .Adicion2 .Age }}
	Multiplicando {{ .Age }} * 6 da como resultado {{ .Multiplicar .Age 6 }}
`
	t, err := template.New("ejercicio03").Parse(tpl03)
	if err != nil {
		log.Fatalf("Error al hacer parse del template %v", err)
	}

	p := person{"EDteam", 5, "123456"}
	err = t.Execute(os.Stdout, &p)
	if err != nil {
		log.Fatalf("Error al ejecutar el template %v", err)
	}
}
