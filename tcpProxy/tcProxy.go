package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {

	dst, err := net.Dial("tcp", "OBJETIVO:PUERTO")
	if err != nil {
		log.Fatalln("No se puede conectar al host")
	}

	defer dst.Close()

	// Ahora una goroutine para prevenir io.Copy de bloquearse
	go func() {
		// Copiamos la salida de la fuente a nuestro destino
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	// Ahora copiemos la salida del destino de vuelta a la fuente
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {

	// Activamos escucha en el puerto 80
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("ERROR: NOT BINDING")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("ERROR: No he logrado aceptar la conexión")
		}
		go handle(conn)
	}

}
