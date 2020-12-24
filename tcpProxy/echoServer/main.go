package main

import (
	"io"
	"log"
	"net"
)

// La función echo solamente hara eco de los datos recibidos
func echo(conn net.Conn) {

	defer conn.Close()

	// Creamos un buffer para almacenar los datos recibidos
	b := make([]byte, 512)
	for {
		// Recibimos los datos via conn.Read
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Cliente desconectado")
			break
		}
		if err != nil {
			log.Println("Error inesperado")
			break
		}
		log.Printf("Recibidos %d bytes : %s \n", size, string(b))

		// enviamos los datos via conn.Write
		log.Println("Escribiendo información")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("NO ha sido posible escribir los datos")
		}
	}
}

func main() {

	// Bind to TCP port 20080 en todas las inbterfaces de red
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("No ha sido posible hacer bind al puerto 20080")
	}
	log.Println("Escuchando en 0.0.0.0:20080")
	for {
		// esperamos conexión, creamos net.Conn sobre la conezión establecida
		conn, err := listener.Accept()
		log.Println("Conexión recibida")
		if err != nil {
			log.Fatalln("No ha sido posible aceptar la comunicación")
		}
		// usamos una goroutine para la conexión
		go echo(conn)
	}
}
