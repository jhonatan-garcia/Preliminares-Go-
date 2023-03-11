package main

import (
	"fmt"
	"os"
)

func main() {
	
	file, err := os.Create("ejemplo.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()


	texto := "Para este proyecto se implementara el lenguaje Go.\n Este se empleara en el backend."
	_, err = file.WriteString(texto)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Println("Archivo creado y texto escrito con éxito.")
}
