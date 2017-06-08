//go:generate goversioninfo -icon=pikachu.ico
package main

import (
	"flag"
	"net/http"
)

func main(){
	path := flag.String("path", ".\\", "Путь до папки")
	port := flag.String("port", "8000", "Порт")
	flag.Parse()
	println("Запустил сервер")
	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*path)))
}
