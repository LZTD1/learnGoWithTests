package files

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1\nDescription: hola hola\nTags: row, click\n---\nBeautiful\nBody")},
		"hello-world2.md": {Data: []byte("Title: Post 2\nDescription: awaga awaga\nTags: row2, click2\n---\nBody2")},
	}

	posts, err := newPostsFromFs(fs)

	if err != nil {
		t.Errorf("got %v; want nil", err)
	}
	if len(posts) != len(fs) {
		t.Errorf("got %v; want %v", len(posts), len(fs))
	}
	assertPost(t, posts[0], Post{Title: "Post 1", Description: "hola hola", Tags: []string{"row", "click"}, Body: "Beautiful\nBody"})
}

func TestNotMDFile(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.txt": {Data: []byte("Title: Post 1\nDescription: hola hola\nTags: row, click\n---\nBeautiful\nBody")},
	}
	posts, _ := newPostsFromFs(fs)

	if len(posts) != 0 {
		t.Errorf("got %v; want empty", len(posts))
	}
}

func TestIncorrectFile(t *testing.T) {
	var cases = []struct {
		Name       string
		Filesystem fstest.MapFS
		Want       string
	}{
		{
			Name: "Wanted Title",
			Filesystem: fstest.MapFS{
				"test.md": {
					Data: []byte("T1tle: Post 1\nDescription: hola hola\nTags: row, click\n---\nBeautiful\nBody"),
				},
			},
			Want: "Wanted Title, but got T1tle",
		},
		{
			Name: "Wanted Description",
			Filesystem: fstest.MapFS{
				"test.md": {
					Data: []byte("Title: Post 1\nD3scription: hola hola\nTags: row, click\n---\nBeautiful\nBody"),
				},
			},
			Want: "Wanted Description, but got D3scription",
		},
	}
	for _, s := range cases {
		t.Run(s.Name, func(t *testing.T) {
			fs := s.Filesystem
			_, err := newPostsFromFs(fs)
			assertErrors(t, err, s.Want)
		})
	}
}

func assertErrors(t *testing.T, err error, want string) {
	t.Helper()

	if err == nil {
		t.Fatalf("got nil; want: %v", want)
	}
	if err.Error() != want {
		t.Fatalf("got %v; want: %v", err.Error(), want)
	}
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
