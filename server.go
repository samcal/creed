package main

import (
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "html/template"
)

type Assassin struct {
  Name string
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", index).Methods("GET")

  http.Handle("/", r)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  a := &Assassin{Name: "Brutus"}
  t, _ := template.ParseFiles("tmpl/index.html")
  t.Execute(w, a)
}

