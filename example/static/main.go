package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/go-on/routerstatic.v0/example/static/site"

	"gopkg.in/go-on/routerstatic.v0"
	"gopkg.in/go-on/router.v2/route"
)

type resolver struct{}

func (rs resolver) Params(rt *route.Route) []map[string]string {
	switch rt {
	case site.DRoute:
		return []map[string]string{
			map[string]string{"a": "a0", "b": "b0", "d": "d0.html"},
			map[string]string{"a": "a1", "b": "b1", "d": "d1.html"},
			map[string]string{"a": "a2", "b": "b2", "d": "d2.html"},
		}
	default:
		panic("unhandled route: " + rt.DefinitionPath)
	}
}

func main() {
	site.Router.Mount("/", nil)

	gopath := os.Getenv("GOPATH")
	dir := filepath.Join(gopath, "src", "github.com", "go-on", "routerstatic", "example", "static", "result")

	os.RemoveAll(dir)
	os.Mkdir(dir, os.FileMode(0755))

	fmt.Println("dump paths")

	routerstatic.MustSavePages(site.Router, resolver{}, site.App, dir)

	fmt.Println("running static fileserver at localhost:8080")

	err := http.ListenAndServe(":8080", http.FileServer(http.Dir(dir)))

	if err != nil {
		panic(err.Error())
	}
}
