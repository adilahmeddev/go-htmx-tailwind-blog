package main

import (
	_ "embed"
	"goblog/pkg/posts"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	file, err := os.Open("./static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static/"))

	router.PathPrefix("/src/").Handler(http.StripPrefix("/src/", fs))

	router.Use(loggingMiddleware)
	homepage, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	storage := posts.MemoryStorage{}

	blogTemplate, err := template.ParseFiles("./pkg/templates/posts.tmpl", "./pkg/templates/post.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	postsHandler := posts.NewHandler(&storage, blogTemplate)
	router.HandleFunc("/home", homeHandler(homepage))

	router.HandleFunc("/posts", postsHandler.GetAll).Methods("GET")

	router.HandleFunc("/post/{id}", postsHandler.Get).Methods("GET")

	router.HandleFunc("/posts/lorem", postsHandler.PostLorem).Methods("POST")

	http.Handle("/", router)

	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(homepage []byte) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(homepage)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
