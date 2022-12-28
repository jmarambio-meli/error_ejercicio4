package main

import (
	"errors"
	"fmt"
)

var (
	Error4 = errors.New("Error: El salario es menor a 150000")
)

func main() {
	var salary int = 180000
	err := validarSalario(salary)
	if err != nil {
		if errors.Is(err, Error4) {
			fmt.Println(fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary))
		}
	} else {
		fmt.Println("Debe pagar Impuesto")
	}

}

func validarSalario(salary int) error {
	if salary < 150000 {
		return Error4
	}
	return nil
}
