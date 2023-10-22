package main

import (
	_ "embed"
	"html/template"
	"io"
	"log"
	"mime"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Post struct {
	Title, Content string
}

const blog = `
	{{template post .}}
`

func mm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		FixMimeTypes()

		next.ServeHTTP(w, r)
	})
}

func main() {
	file, err := os.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}

	m := mux.NewRouter()
	fs := http.FileServer(http.Dir("./src/"))

	m.PathPrefix("/src/").Handler(http.StripPrefix("/src/", fs))

	m.Use(loggingMiddleware)
	homepage, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	m.HandleFunc("/home", (func(w http.ResponseWriter, r *http.Request) {
		w.Write(homepage)
	}))
	m.HandleFunc("/posts", handler)
	http.Handle("/", m)
	http.ListenAndServe(":7000", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {

	blogTemplate, err := template.ParseFiles("post.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	post := Post{
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
	if err := blogTemplate.ExecuteTemplate(res, "post", post); err != nil {
		log.Fatal(err)
	}

}

func FixMimeTypes() {
	err1 := mime.AddExtensionType(".js", "text/javascript")
	if err1 != nil {
		log.Printf("Error in mime js %s", err1.Error())
	}

	err2 := mime.AddExtensionType(".css", "text/css")
	if err2 != nil {
		log.Printf("Error in mime js %s", err2.Error())
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
