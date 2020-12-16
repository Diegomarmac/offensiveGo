package main

import (
	"fmt"
	"net"
)

func main() {
	// solo vamos a escanear 1024 puertos, pero recuerda que hay 65535 puertos en un server
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		// si la conexión es satisfactoria, err será igual a nil
		if err != nil {
			// el puerto está cerrado o filtrado
			continue
		}
		conn.Close()
		fmt.Printf("%d abierto\n", i)

	}
}
