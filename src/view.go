package main

import (
  "net/http"
  "html/template"
)

type View struct {
}

func getView(r *http.Request) (*View, error) {
  return &View {}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  //fmt.Fprintf(w, "%s", r.URL.Path[0:])
  t, _ := template.ParseFiles("./webroot/guest/index.html")
  v, _ := getView(r)
  t.Execute(w, v)
}
