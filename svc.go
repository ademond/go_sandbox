package main
import (
  "net/http"
  "fmt"
  "log"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  r.HandleFunc("/price/{cusip}", Pricer)
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe(":4000", nil))
}

func HomeHandler(w http.ResponseWriter,r *http.Request) {
  fmt.Fprintf(w, "Hi!")
}

func Pricer(w http.ResponseWriter,r *http.Request) {
  vars := mux.Vars(r)
  cusip := vars["cusip"]
  fmt.Fprintf(w, "%s is worth 100.3", cusip)
}

