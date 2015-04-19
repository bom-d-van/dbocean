package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/bmizerany/pat"
	"github.com/elazarl/go-bindata-assetfs"
)

var (
	rootDir = ""
	tmpl    = loadTmpl()
	devMode = true
)

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(home))
	mux.Get("/ui/", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir}))

	mux.Get("/dbs", http.HandlerFunc(getDBs))
	mux.Post("/db", http.HandlerFunc(addDB))
	mux.Del("/db", http.HandlerFunc(removeDB))
	mux.Post("/db/exec", http.HandlerFunc(execCmd))

	http.Handle("/", mux)
	println("Listening on :9000")
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
