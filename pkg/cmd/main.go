package main

import (
	_ "embed"
	"fmt"
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
	router.HandleFunc("/home", homeHandler(homepage))

	router.HandleFunc("/posts", GetPostsHandler(&storage)).Methods("GET")

	router.HandleFunc("/posts/lorem", AddLoremPostsHandler(&storage)).Methods("POST")

	http.Handle("/", router)

	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}

func AddLoremPostsHandler(storage posts.Storage) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {

		post := posts.Post{
			Title: "First post pog",
			Content: `        aliquet tempor justo. In hac
        habitasse platea dictumst. Interdum et malesuada fames ac ante ipsum primis in faucibus. In vehicula augue non
        ante finibus, a tincidunt enim elementum. Fusce ac justo diam. Suspendisse condimentum consectetur laoreet.
        Proin vel pellentesque est, non hendrerit lorem. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
        Suspendisse convallis augue turpis. Sed nec fermentum eros. Nam vel orci sagittis, malesuada magna non,
        rhoncus
        est. Integer sem ipsum, lobortis quis nunc sit amet, elementum malesuada felis. Vestibulum ante ipsum primis
        in
        faucibus orci luctus et ultrices posuere cubilia curae;

        Sed suscipit tempus nibh sed eleifend. Quisque volutpat felis quis ex vehicula dictum. Etiam ornare ex at
        sapien
        rutrum accumsan. Sed ut odio non purus hendrerit porta. Integer nisi enim, porta at vestibulum vitae, tempor
        ut
        enim. Sed bibendum justo ac vehicula dapibus. Praesent at vestibulum diam. Aliquam auctor porttitor lorem, a
        euismod nisi sollicitudin at. Sed quis ipsum ut massa fermentum vulputate nec a justo. In vitae ligula leo. In
`,
		}

		err := storage.Add(post)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func GetPostsHandler(storage posts.Storage) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {

		blogTemplate, err := template.ParseFiles("./pkg/templates/posts.tmpl", "./pkg/templates/post.tmpl")
		if err != nil {
			log.Fatal(err)
		}
		allPosts, err := storage.GetAll()
		if err != nil {
			res.WriteHeader(500)
			res.Write([]byte(err.Error()))
		}
		for _, post := range allPosts {
			fmt.Println(post.Title)
		}

		if err := blogTemplate.ExecuteTemplate(res, "posts", map[string]interface{}{"Posts": allPosts}); err != nil {
			log.Fatal(err)
		}
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
