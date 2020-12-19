package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(puertos, results chan int) {

	for p := range puertos {
		direccion := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", direccion)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	puertos := make(chan int, 100)
	resultados := make(chan int)
	var puertosAbiertos []int

	for i := 0; i < cap(puertos); i++ {
		go worker(puertos, resultados)
	}

	go func() {
		for i := 0; i <= 1024; i++ {
			puertos <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		puerto := <-resultados
		if puerto != 0 {
			puertosAbiertos = append(puertosAbiertos, puerto)
		}
	}

	close(puertos)
	close(resultados)
	sort.Ints(puertosAbiertos)

	for _, puerto := range puertosAbiertos {
		fmt.Printf("%d open\n", puerto)
	}
}
