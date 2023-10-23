# go-htmx-tailwind-blog

## Requirements
- [tailwind](https://tailwindcss.com/docs/installation) 
- [go](https://go.dev/doc/install) 

----------
## Running

### Start the tailwind build process
`tailwind -i ./pkg/input.css -o ./static/output.css --watch`

### Start the go server
`go run ./pkg/cmd/main.go`

The server will start listening on the port `:7000`

View the homepage at `http://localhost:7000/home`

The posts are served on the `/posts/{id}` endpoint
