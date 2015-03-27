package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/elazarl/go-bindata-assetfs"
)

var rootDir = ""

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	}))
	mux.Get("/ui/", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir}))

	http.Handle("/", mux)
	println("Listening on :9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
