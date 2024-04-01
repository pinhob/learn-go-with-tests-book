package blogposts

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i alaways fail")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
		Description: Description 1`
		secondBody = `Title: Post 2
		Description: Description 2`
	)

	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := Post{Title: "Post 1", Description: "Description 1"}

	assertPost(t, got, want)
}

func assertPost(t *testing.T, got, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
