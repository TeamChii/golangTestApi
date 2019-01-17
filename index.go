package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//port := os.Getenv("PORT")
	router := mux.NewRouter()
	userDB := map[string]int{
		"team":  23,
		"park":  21,
		"poom":  30,
		"first": 44,
	}

	log.Print("Starting the service")

	////// ส่งค่าเป็น json
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-type", "application/json")
		fmt.Fprintf(w, `{"id":"aa"}`)
	})

	/// router  ธรรมดา
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "home.html")
	})
	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "signup.html")
	})
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Login.html")
	})
	router.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "welcome.html")
	})
	router.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "name.txt")
	})

	// ส่งพารามิเตอร์
	router.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := userDB[name]
		fmt.Fprintf(w, "My name is %s %d", name, age)
	}).Methods("GET")

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
