package louvois

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func webServices() {
	r := mux.NewRouter()

	// liste des soldats
	r.HandleFunc("/soldier/", soldierListHandler)

	// recherche d'un soldat
	r.HandleFunc("/soldier/search", soldierSearchHandler)

	// edition d'un soldat
	r.HandleFunc("/soldier/{:id}", soldierHandler)

	//gestion des fichiers statiques HTML & CSS & JS
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", r)

	http.ListenAndServe(":3000", nil)
}

func StartServer() {
	webServices()
	log.Println("Listening web serveur on port : 3000 activated !")
}

func StopServer() {
	os.Exit(0)
}

func soldierListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "soldier list")
}

func soldierSearchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "soldier search")
}

func soldierHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "soldier")
}
