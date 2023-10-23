package posts_test

import (
	"goblog/pkg/posts"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAll(t *testing.T) {
	blogTemplate, err := template.ParseFiles("../templates/posts.tmpl", "../templates/post.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	storage := posts.MemoryStorage{}
	handler := posts.NewHandler(&storage, blogTemplate)

	storage.Add(posts.Post{
		Title: "test post",
		Content: `test content
		hello
		content`,
	})

	storage.Add(posts.Post{
		Title: "test post 2",
		Content: `test content
		hello
		content`,
	})

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/placeholder", nil)

	handler.GetAll(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expected %v, got %v", http.StatusOK, res.Code)
	}

	recievedContent, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(recievedContent), "test post") {
		t.Error("does not contain post title")
	}

	if !strings.Contains(string(recievedContent), "test content") {
		t.Error("does not contain post content")
	}

	if !strings.Contains(string(recievedContent), "test post 2") {
		t.Error("does not contain post 2 title")
	}

}

func TestPostLorem(t *testing.T) {
	blogTemplate, err := template.ParseFiles("../templates/posts.tmpl", "../templates/post.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	store := posts.MemoryStorage{}
	handler := posts.NewHandler(&store, blogTemplate)

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/placeholder", nil)

	handler.PostLorem(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("expected %v, got %v", http.StatusOK, res.Code)
	}

	allposts, err := store.GetAll()
	if err != nil {
		t.Error(err)
	}

	if len(allposts) == 0 {
		t.Error("should have posts")
	}
}
