package main

import (
  "database/sql"
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  _ "github.com/lib/pq"
)

type App struct {
  Router  *mux.Router
  DB      *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
  connectionString :=
    fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

  var err error
  a.DB, err = sql.Open("postgres", connectionString)
  if err != nil {
    log.Fatal(err)
  }

  a.Router = mux.NewRouter()
  a.InitializeRoutes()
}

func (a *App) Run(addr string) {
  log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) InitializeRoutes() {
  a.Router.HandleFunc("/", a.getIndex).Methods("GET")
}

func (a *App) getIndex(w http.ResponseWriter, r *http.Request) {
  data := "Example App"
  w.Write([]byte(data))
}
