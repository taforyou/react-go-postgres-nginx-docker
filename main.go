// ไม่ใช้แล้ว

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/", DoHealthCheck).Methods("GET")
// 	router.HandleFunc("/api/", DoHealthCheckApi).Methods("GET")
// 	log.Fatal(http.ListenAndServe(":8081", router))
// }

// func DoHealthCheck(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, i'm a golang microservice")
// }
// func DoHealthCheckApi(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, i'm a golang microservice API")
// }
