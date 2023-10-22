package posts

import "fmt"

type Storage interface {
	GetAll() ([]Post, error)
	Add(Post) error
}

type MemoryStorage struct {
	posts []Post
}

func (m *MemoryStorage) Add(post Post) error {
	m.posts = append(m.posts, post)

	return nil
}

func (m *MemoryStorage) GetAll() ([]Post, error) {
	if len(m.posts) == 0 {
		return nil, fmt.Errorf("no posts in storage")
	}

	return m.posts, nil
}

