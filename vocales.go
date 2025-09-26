package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	vocales := "aeiouAEIOU"
	fmt.Println("Bienvenido, ingrese una frase para poder contar las vocales que contiene: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	fraseclean := strings.ReplaceAll(frase, " ", "")

	contador := 0
	for _, r := range fraseclean {
		if strings.ContainsRune(vocales, r) {
			contador++
		}
	}
	fmt.Printf("La frase contiene %d vocales\n", contador)
}
