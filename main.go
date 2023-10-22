package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

type Post struct {
	Title, Content string
}

const blog = `
	{{template post .}}
`

func main() {

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
	buff := bytes.Buffer{}
	if err := blogTemplate.ExecuteTemplate(&buff, "post", post); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", buff.String())
}
