package main

import (
	"fmt"
	"log"
	"os"
)

// FooReader definido como un lector desde stdin
type FooReader struct{}

// Read lee datos desde stdin
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// FooWriter definido como un escritor a stdout
type FooWriter struct{}

//Write escribe en stdout
func (fooWrtier *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	// instanciamos lectura y escritura
	var (
		reader FooReader
		writer FooWriter
	)

	// Creamos un buffer para mantener input/output
	input := make([]byte, 4096)

	// Usamos read para leer input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("No he sido capaz de leer los datos")
	}
	fmt.Printf("Leyendo %d bytes desde standard input\n", s)

	// Usamos writer para escribir outpu
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("No he sido capaz de escribir datos")
	}
	fmt.Printf("He escrito %d bytes a salida standard\n", s)
}
