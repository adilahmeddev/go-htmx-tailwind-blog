package posts

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) GetAll(res http.ResponseWriter, req *http.Request) {
	blogTemplate, err := template.ParseFiles("./pkg/templates/posts.tmpl", "./pkg/templates/post.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	allPosts, err := h.storage.GetAll()
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

func (h *Handler) PostLorem(res http.ResponseWriter, req *http.Request) {
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

	err := h.storage.Add(post)
	if err != nil {
		log.Fatal(err)
	}

}