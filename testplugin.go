package main

import "net/http"

var Start = func() *bool {
	// run plugin
	Plugins = append(Plugins, Plugin{
		Name:    "testPlugin",
		Version: "1.0",
		FuncHandlers: []HTTPFuncHandler{
			HTTPFuncHandler{
				Name: "/",
				Function: func(rw http.ResponseWriter, rq *http.Request) {
					if rq.URL.Path != "/" {
						rw.Write([]byte("404!"))
						rq.Response.StatusCode = 404
					}
				},
			},
		},
	})
	return nil
}()
