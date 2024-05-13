package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Carro struct {
	Nome   string
	Modelo string
	Ano    int
}

func (c Carro) Andar() {
	fmt.Println("O carro, " + c.Nome + " esta andando")
}

func (c Carro) Parar() {

	fmt.Println("O carro, " + c.Nome + " esta parando.")
}

func main() {
	carro1 := Carro{Nome: "Ford", Modelo: "qualquer", Ano: 2000}

	carro1.Andar()
	carro1.Parar()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ola, mundo")
		json.NewEncoder(w).Encode(carro1)
	})

	http.ListenAndServe(":8080", nil)
}
