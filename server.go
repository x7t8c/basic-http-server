package main

import (
	"net/http"
)

type HTTPFuncHandler struct {
	Name     string
	Function func(rw http.ResponseWriter, rq *http.Request)
}

type Plugin struct {
	Name         string
	Version      string
	FuncHandlers []HTTPFuncHandler
}

var Plugins = []Plugin{}

func main() {
	for x := 0; x < len(Plugins); x++ {
		for y := 0; y < len(Plugins[x].FuncHandlers); y++ {
			http.HandleFunc(Plugins[x].FuncHandlers[y].Name, Plugins[x].FuncHandlers[y].Function)
		}
	}
	http.ListenAndServe("127.0.0.1:8124", nil)
}
