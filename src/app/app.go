package app

import (
	"net/http"
	"reflect"
	"fmt"
)

type application struct {
	routes map[string]interface{}
}

func New() *application {
	return &application{
		routes: make(map[string]interface{}),
	}
}

func (p *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	controllerName := r.URL.Query().Get("c")
	actionName := r.URL.Query().Get("a")
	if controllerName == "" || actionName == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	route, ok := p.routes[controllerName]
	if !ok {
		http.Error(w, "Controller Not Found", http.StatusNotFound)
		return
	}
	ele := reflect.ValueOf(route).Elem()
	ele.FieldByName("Request").Set(reflect.ValueOf(r))
	ele.FieldByName("Response").Set(reflect.ValueOf(w))
	ele.MethodByName(actionName).Call([]reflect.Value{})
}

func (p *application) printRoutes() {
	for route, controller := range p.routes {
		ele := reflect.ValueOf(controller).Type().String()
		fmt.Printf("%s %s\n", route, ele)
	}
}

func (p *application) Get(route string, controller interface{}) {
	p.routes[route] = controller
}

func (p *application) Run(addr string) error {
	p.printRoutes()
	fmt.Printf("listen on %s\n", addr)
	return http.ListenAndServe(addr, p)
}
