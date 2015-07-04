package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/bmizerany/pat"
	"github.com/elazarl/go-bindata-assetfs"
)

var (
	rootDir = ""
	tmpl    = loadTmpl()
	devMode = os.Getenv("ENV") != "prod"
)

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(home))
	mux.Get("/ui/", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir}))

	mux.Get("/dbs", handlerWrap(listDBs))
	mux.Post("/dbs", handlerWrap(addDB))
	mux.Del("/dbs/:id", handlerWrap(removeDB))

	mux.Get("/dbs/:id", handlerWrap(getDB))
	// mux.Get("/dbs/:id/tables", handlerWrap(getDBTables))
	mux.Get("/dbs/:id/tables/:name", handlerWrap(getDBTable))
	mux.Post("/dbs/:id/exec", handlerWrap(query))

	http.Handle("/", mux)
	log.Println("Listening on :9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func home(rw http.ResponseWriter, req *http.Request) {
	execTmpl(rw, "app", nil)
}

func execTmpl(rw io.Writer, name string, data interface{}) {
	if devMode {
		tmpl = loadTmpl()
	}
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, name, data); err != nil {
		log.Printf("%s\n%s", err, debug.Stack())
		rw.Write([]byte("<div><pre>"))
		rw.Write([]byte(err.Error() + "\n"))
		rw.Write(debug.Stack())
		rw.Write([]byte("</pre></div>"))
		return
	}

	io.Copy(rw, &buf)
}

func loadTmpl() *template.Template {
	t := template.New("").Funcs(template.FuncMap{
		"now": func() time.Time { return time.Now() },
	})
	var err error
	for _, name := range AssetNames() {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		t, err = t.New(name).Parse(string(MustAsset(name)))
		if err != nil {
			panic(err)
		}
	}
	return t
}

func handlerWrap(h func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() { log.Printf("%s %s %s", r.Method, r.URL.Path, time.Now().Sub(start)) }()
		defer func() {
			if r := recover(); r != nil {
				err := fmt.Errorf("%+v", r)
				log.Println(err)
				debug.PrintStack()
				w.Write([]byte(err.Error() + "\n"))
				w.Write(debug.Stack())
			}
		}()

		h(w, r)
	})
}
