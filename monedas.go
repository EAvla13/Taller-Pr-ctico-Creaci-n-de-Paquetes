package main

import "fmt"

func main() {
	dinero := 0.0
	moneda := 0
	fmt.Println("Bienvenido, ingrese el valor en dolares que desea convertir: ")
	fmt.Scan(&dinero)
	fmt.Println("A que desea transformar este valor en dolares?\n1) Euros\n2) LB(Libras Esterlinas)\n3) Won(Sur Koreano)\n4) BTC")
	fmt.Scan(&moneda)

	switch moneda {
	case 1:
		fmt.Println(dinero, "dolares, a Euros son", (dinero * 0.86))
	case 2:
		fmt.Println(dinero, "dolares, a LB son", (dinero * 0.75))
	case 3:
		fmt.Println(dinero, "dolares, a Won son", (dinero * 1412.76))
	case 4:
		fmt.Println(dinero, "dolares, a BTC son", (dinero * 0.0000091))
	default:
		fmt.Println("Ingrese una opcion valida!")
	}
}
